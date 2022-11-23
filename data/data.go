package data

import "staterecovery/conf"

type BlockData struct {
	BlockId int64  `json:"block_id"`
	Data    []byte `json:"data"`
}

type DataResource interface {
	GetDataByBlockId(blockId int) (data *BlockData, err error)
	GetTokenIDByAddress(tokenAddr string) (tokenID string, err error)
}

func NewDataResource(conf *conf.Config) DataResource {
	if conf.DataFrom == 0 {
		return newDBDataResource(conf)
	} else if conf.DataFrom == 1 {
		return newChainDataResource(conf)
	}
	return nil
}
