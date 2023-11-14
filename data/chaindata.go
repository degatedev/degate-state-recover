package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"staterecovery/abi/exchangeV3"
	"staterecovery/conf"
	"strconv"
	"strings"
	"time"
)

var (
	SubmitBlocksWithCallbacksOpCode = crypto.Keccak256Hash([]byte("BlockSubmitted(uint256,bytes32,bytes32)"))
	blockFileParentPath             string
	blockNumFilePath                string
	zkpBlockFilePath                string
	blockNumFileContent             = "blockNum=%v&zkpBlockId=%v"
)

type ChainDataResource struct {
	exchangeAddress common.Address
	client          *ethclient.Client
}

func newChainDataResource(conf *conf.Config) *ChainDataResource {
	var err error
	if len(conf.ExchangeContract) == 0 {
		panic("not find exchangeContract")
	}
	dr := &ChainDataResource{
		exchangeAddress: common.HexToAddress(conf.ExchangeContract),
	}
	dr.client, err = ethclient.Dial(conf.ChainNode)
	if err != nil {
		panic(err)
	}
	if conf.StateBlockID > 0 {
		conf.LastL2BlockID = conf.StateBlockID
		return dr
	}
	if conf.FirstL1BlockID <= 0 {
		panic("not find FirstL1BlockID")
	}
	err = InitBlockPath()
	if err != nil {
		panic(err)
	}
	dr.Scanner()
	return dr
}

func (dr *ChainDataResource) GetDataByBlockId(blockId int) (data *BlockData, err error) {
	fileData, err := ioutil.ReadFile(fmt.Sprintf(zkpBlockFilePath, blockId))
	if err != nil {
		return
	}
	if len(fileData) == 0 {
		err = errors.New("block data is empty")
		return
	}
	data = &BlockData{
		BlockId: int64(blockId),
		Data:    fileData,
	}
	return
}

func (dr *ChainDataResource) GetTokenIDByAddress(tokenAddr string) (tokenID string, err error) {
	exchangeContract, err := exchangeV3.NewExchangeV3Caller(common.HexToAddress(conf.Conf.ExchangeContract), dr.client)
	if err != nil {
		return
	}

	tokenIDUint, err := exchangeContract.GetTokenID(nil, common.HexToAddress(tokenAddr))
	if err != nil {
		return
	}
	return strconv.FormatUint(uint64(tokenIDUint), 10), nil
}

func (dr *ChainDataResource) GetTokenAddressByID(tokenId uint32) (tokenAddr common.Address, err error) {
	exchangeContract, err := exchangeV3.NewExchangeV3Caller(common.HexToAddress(conf.Conf.ExchangeContract), dr.client)
	if err != nil {
		return
	}

	tokenAddr, err = exchangeContract.GetTokenAddress(nil, tokenId)
	if err != nil {
		return
	}
	return tokenAddr, nil
}

func (dr *ChainDataResource) Scanner() {
	var (
		zkpBlockId  uint64
		blockNum    uint64
		maxBlockNum uint64
		err         error
	)

	lastBlockId, err := dr.GetLastBlockId()
	if err != nil {
		panic(err)
	}
	conf.Conf.LastL2BlockID = int(lastBlockId)

	zkpBlockId, blockNum, err = dr.GetScannerBlockNum()
	if err != nil {
		panic(err)
	}
	if zkpBlockId >= lastBlockId {
		log.Info("all blocks scanner finish")
		return
	}
	blockNum++
	for {
		var (
			logs       []types.Log
			toBlockNum uint64
		)

		maxBlockNum, err = dr.client.BlockNumber(context.Background())
		if err != nil {
			log.Error("BlockNumber error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		if maxBlockNum < blockNum {
			time.Sleep(time.Second * 5)
			continue
		}
		if maxBlockNum-blockNum > 1000 {
			toBlockNum = blockNum + 1000
		} else {
			toBlockNum = maxBlockNum
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(blockNum)),
			ToBlock:   big.NewInt(int64(toBlockNum)),
			Addresses: []common.Address{
				dr.exchangeAddress,
			},
		}
		logs, err = dr.client.FilterLogs(context.Background(), query)
		if err != nil {
			log.Error("FilterLogs error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		savedTxHash := ""
		for _, event := range logs {
			var (
				tx *types.Transaction
			)
			if !isSubmitBlocks(&event) {
				continue
			}
			zkpBlockId, err = strconv.ParseUint(event.Topics[1].Hex()[2:], 16, 0)
			if err != nil {
				log.Error("GetSubmitBlock tx:%v parse blockId error: %v", tx.Hash().String(), err)
				panic(err)
			}
			for {
				tx, _, err = dr.client.TransactionByHash(context.Background(), event.TxHash)
				if err != nil {
					log.Error("TransactionByHash tx:%v error: %v", event.TxHash, err)
					time.Sleep(time.Second * 5)
					continue
				}
				break
			}
			if savedTxHash != event.TxHash.String() {
				savedTxHash = event.TxHash.String()
				dr.SaveBlockData(tx.Data(), zkpBlockId, event.BlockNumber)
			}
			log.Info("zkp block %v scanner finish", event.BlockNumber)
		}
		if int(zkpBlockId) >= conf.Conf.LastL2BlockID {
			log.Info("all block scanner finish")
			return
		}
		blockNum = toBlockNum + 1
	}
}

func (dr *ChainDataResource) GetScannerBlockNum() (zkpBlockId uint64, blockNum uint64, err error) {
	_, err = os.Stat(blockNumFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		var blockNumFile *os.File
		blockNumFile, err = os.Create(blockNumFilePath)
		if err != nil {
			return
		}
		_ = blockNumFile.Close()
	}
	fileData, err := ioutil.ReadFile(blockNumFilePath)
	if err != nil {
		return
	}
	if len(fileData) > 0 {
		var (
			mBlockNum   int
			mZkpBlockId int
		)
		bns := string(fileData)
		bnA := strings.Split(bns, "&")
		if len(bnA) != 2 {
			err = errors.New("error block_num file content")
			return
		}
		mBlockNum, err = strconv.Atoi(strings.Split(bnA[0], "=")[1])
		if err != nil {
			return
		}
		mZkpBlockId, err = strconv.Atoi(strings.Split(bnA[1], "=")[1])
		if err != nil {
			return
		}
		blockNum = uint64(int64(mBlockNum))
		zkpBlockId = uint64(mZkpBlockId)
	} else {
		blockNum = conf.Conf.FirstL1BlockID - 1
		zkpBlockId = 0
	}
	log.Info("scanner from block:%v", blockNum)
	return
}

func (dr *ChainDataResource) SaveBlockData(data []byte, zkpBlockId uint64, blockNum uint64) {
	var (
		zkpBlockNumFile *os.File
		blockNumFile    *os.File
		err             error
	)

	for {
		zkpBlockNumFile, err = os.OpenFile(fmt.Sprintf(zkpBlockFilePath, zkpBlockId), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			log.Error("OpenFile error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	for {
		n, _ := zkpBlockNumFile.Seek(io.SeekStart, io.SeekEnd)
		_, err = zkpBlockNumFile.WriteAt(data, n)
		if err != nil {
			log.Error("write zkp block data blockId:%v error: %v", zkpBlockId, err)
			time.Sleep(time.Second * 5)
			continue
		}
		_ = zkpBlockNumFile.Close()
		break
	}

	for {
		blockNumFile, err = os.OpenFile(blockNumFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			log.Error("OpenFile error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
	blockNumCont := fmt.Sprintf(blockNumFileContent, blockNum, zkpBlockId)
	for {
		n, _ := blockNumFile.Seek(io.SeekStart, io.SeekEnd)
		_, err = blockNumFile.WriteAt([]byte(blockNumCont), n)
		if err != nil {
			log.Error("write blockNum error: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}
		_ = blockNumFile.Close()
		break
	}
}

func (dr *ChainDataResource) GetLastBlockId() (blockId uint64, err error) {
	exchangeContract, err := exchangeV3.NewExchangeV3Caller(common.HexToAddress(conf.Conf.ExchangeContract), dr.client)
	if err != nil {
		return
	}

	blockState, err := exchangeContract.State(nil)
	if err != nil {
		return
	}
	blockId = blockState.NumBlocks.Uint64() - 1
	return
}

func InitBlockPath() (err error) {
	blockFileParentPath = conf.Conf.BlockFilePath + "/blockfile"
	blockNumFilePath = blockFileParentPath + "/block_num"
	zkpBlockFilePath = blockFileParentPath + "/zkp_block_%v"
	_, err = os.Stat(blockFileParentPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		err = os.Mkdir(blockFileParentPath, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

func isSubmitBlocks(event *types.Log) bool {
	return len(event.Topics) > 0 && event.Topics[0] == SubmitBlocksWithCallbacksOpCode
}
