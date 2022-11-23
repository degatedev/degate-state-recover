// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchangeV3

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExchangeDataAccountLeaf is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataAccountLeaf struct {
	AccountID uint32
	Owner     common.Address
	PubKeyX   *big.Int
	PubKeyY   *big.Int
	Nonce     uint32
}

// ExchangeDataBalanceLeaf is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBalanceLeaf struct {
	TokenID uint32
	Balance *big.Int
}

// ExchangeDataBlock is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBlock struct {
	BlockType             uint8
	BlockSize             uint16
	BlockVersion          uint8
	Data                  []byte
	Proof                 [8]*big.Int
	StoreBlockInfoOnchain bool
	AuxiliaryData         []byte
	OffchainData          []byte
}

// ExchangeDataBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataBlockInfo struct {
	Timestamp     uint32
	BlockDataHash [28]byte
}

// ExchangeDataConstants is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataConstants struct {
	SNARKSCALARFIELD                         *big.Int
	MAXOPENFORCEDREQUESTS                    *big.Int
	MAXAGEFORCEDREQUESTUNTILWITHDRAWMODE     *big.Int
	TIMESTAMPHALFWINDOWSIZEINSECONDS         *big.Int
	MAXNUMACCOUNTS                           *big.Int
	MAXNUMTOKENS                             *big.Int
	MINAGEPROTOCOLFEESUNTILUPDATED           *big.Int
	MINTIMEINSHUTDOWN                        *big.Int
	TXDATAAVAILABILITYSIZE                   *big.Int
	MAXAGEDEPOSITUNTILWITHDRAWABLEUPPERBOUND *big.Int
}

// ExchangeDataDepositState is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataDepositState struct {
	FreeDepositMax      *big.Int
	FreeDepositRemained *big.Int
	LastDepositBlockNum *big.Int
	FreeSlotPerBlock    *big.Int
	DepositFee          *big.Int
}

// ExchangeDataMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataMerkleProof struct {
	AccountLeaf        ExchangeDataAccountLeaf
	BalanceLeaf        ExchangeDataBalanceLeaf
	AccountMerkleProof [48]*big.Int
	BalanceMerkleProof [48]*big.Int
}

// ExchangeDataModeTime is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataModeTime struct {
	ShutdownModeStartTime   *big.Int
	WithdrawalModeStartTime *big.Int
}

// ExchangeDataProtocolFeeData is an auto generated low-level Go binding around an user-defined struct.
type ExchangeDataProtocolFeeData struct {
	SyncedAt                uint32
	ProtocolFeeBips         uint8
	PreviousProtocolFeeBips uint8
}

// ExchangeV3ABI is the input ABI used to generate the binding from.
const ExchangeV3ABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockIdx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicDataHash\",\"type\":\"bytes32\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint248\",\"name\":\"amount\",\"type\":\"uint248\"}],\"name\":\"DepositRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"genesisMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"ExchangeCloned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"}],\"name\":\"ForcedWithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"protocolFeeBips\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"previousProtocolFeeBips\",\"type\":\"uint8\"}],\"name\":\"ProtocolFeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenId\",\"type\":\"uint32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"TransactionApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawalModeActivated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowOnchainTransferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loopringAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxAgeDepositUntilWithdrawable\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"DOMAIN_SEPARATOR\",\"type\":\"bytes32\"},{\"internalType\":\"contractILoopringV3\",\"name\":\"loopring\",\"type\":\"address\"},{\"internalType\":\"contractIBlockVerifier\",\"name\":\"blockVerifier\",\"type\":\"address\"},{\"internalType\":\"contractIAgentRegistry\",\"name\":\"agentRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIDepositContract\",\"name\":\"depositContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"freeDepositMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeDepositRemained\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlockNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeSlotPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositFee\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.DepositState\",\"name\":\"depositState\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleAssetRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"numPendingForcedTransactions\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"syncedAt\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"protocolFeeBips\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"previousProtocolFeeBips\",\"type\":\"uint8\"}],\"internalType\":\"structExchangeData.ProtocolFeeData\",\"name\":\"protocolFeeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shutdownModeStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalModeStartTime\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.ModeTime\",\"name\":\"modeTime\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_loopring\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisMerkleAssetRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshBlockVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"name\":\"setDepositContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawExchangeFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"freeDepositMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeDepositRemained\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeSlotPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositFee\",\"type\":\"uint256\"}],\"name\":\"setDepositParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isUserOrAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConstants\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"SNARK_SCALAR_FIELD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_OPEN_FORCED_REQUESTS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_AGE_FORCED_REQUEST_UNTIL_WITHDRAW_MODE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TIMESTAMP_HALF_WINDOW_SIZE_IN_SECONDS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_NUM_ACCOUNTS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_NUM_TOKENS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MIN_AGE_PROTOCOL_FEES_UNTIL_UPDATED\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MIN_TIME_IN_SHUTDOWN\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TX_DATA_AVAILABILITY_SIZE\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"MAX_AGE_DEPOSIT_UNTIL_WITHDRAWABLE_UPPERBOUND\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeData.Constants\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInWithdrawalMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"}],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawExchangeStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getProtocolFeeLastWithdrawnTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnExchangeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleAssetRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockIdx\",\"type\":\"uint256\"}],\"name\":\"getBlockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes28\",\"name\":\"blockDataHash\",\"type\":\"bytes28\"}],\"internalType\":\"structExchangeData.BlockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"blockType\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"blockSize\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"blockVersion\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bool\",\"name\":\"storeBlockInfoOnchain\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"auxiliaryData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainData\",\"type\":\"bytes\"}],\"internalType\":\"structExchangeData.Block[]\",\"name\":\"blocks\",\"type\":\"tuple[]\"}],\"name\":\"submitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumAvailableForcedSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"amount\",\"type\":\"uint248\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getPendingDepositAmount\",\"outputs\":[{\"internalType\":\"uint248\",\"name\":\"\",\"type\":\"uint248\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"}],\"name\":\"forceWithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isForcedWithdrawalPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFees\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"internalType\":\"structExchangeData.AccountLeaf\",\"name\":\"accountLeaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"tokenID\",\"type\":\"uint32\"},{\"internalType\":\"uint248\",\"name\":\"balance\",\"type\":\"uint248\"}],\"internalType\":\"structExchangeData.BalanceLeaf\",\"name\":\"balanceLeaf\",\"type\":\"tuple\"},{\"internalType\":\"uint256[48]\",\"name\":\"accountMerkleProof\",\"type\":\"uint256[48]\"},{\"internalType\":\"uint256[48]\",\"name\":\"balanceMerkleProof\",\"type\":\"uint256[48]\"}],\"internalType\":\"structExchangeData.MerkleProof\",\"name\":\"merkleProof\",\"type\":\"tuple\"}],\"name\":\"withdrawFromMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isWithdrawnInWithdrawalMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawFromDepositRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"withdrawFromApprovedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAmountWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"accountID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"notifyForcedRequestTooOld\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"amount\",\"type\":\"uint248\"},{\"internalType\":\"uint32\",\"name\":\"storageID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"setWithdrawalRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"amount\",\"type\":\"uint248\"},{\"internalType\":\"uint32\",\"name\":\"storageID\",\"type\":\"uint32\"}],\"name\":\"getWithdrawalRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onchainTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"approveTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"transactionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"approveTransactions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"isTransactionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setMaxAgeDepositUntilWithdrawable\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxAgeDepositUntilWithdrawable\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeValues\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"syncedAt\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"protocolFeeBips\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"previousProtocolFeeBips\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowOnchainTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExchangeV3 is an auto generated Go binding around an Ethereum contract.
type ExchangeV3 struct {
	ExchangeV3Caller     // Read-only binding to the contract
	ExchangeV3Transactor // Write-only binding to the contract
	ExchangeV3Filterer   // Log filterer for contract events
}

// ExchangeV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeV3Session struct {
	Contract     *ExchangeV3       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeV3CallerSession struct {
	Contract *ExchangeV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExchangeV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeV3TransactorSession struct {
	Contract     *ExchangeV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExchangeV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeV3Raw struct {
	Contract *ExchangeV3 // Generic contract binding to access the raw methods on
}

// ExchangeV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeV3CallerRaw struct {
	Contract *ExchangeV3Caller // Generic read-only contract binding to access the raw methods on
}

// ExchangeV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeV3TransactorRaw struct {
	Contract *ExchangeV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeV3 creates a new instance of ExchangeV3, bound to a specific deployed contract.
func NewExchangeV3(address common.Address, backend bind.ContractBackend) (*ExchangeV3, error) {
	contract, err := bindExchangeV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3{ExchangeV3Caller: ExchangeV3Caller{contract: contract}, ExchangeV3Transactor: ExchangeV3Transactor{contract: contract}, ExchangeV3Filterer: ExchangeV3Filterer{contract: contract}}, nil
}

// NewExchangeV3Caller creates a new read-only instance of ExchangeV3, bound to a specific deployed contract.
func NewExchangeV3Caller(address common.Address, caller bind.ContractCaller) (*ExchangeV3Caller, error) {
	contract, err := bindExchangeV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3Caller{contract: contract}, nil
}

// NewExchangeV3Transactor creates a new write-only instance of ExchangeV3, bound to a specific deployed contract.
func NewExchangeV3Transactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeV3Transactor, error) {
	contract, err := bindExchangeV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3Transactor{contract: contract}, nil
}

// NewExchangeV3Filterer creates a new log filterer instance of ExchangeV3, bound to a specific deployed contract.
func NewExchangeV3Filterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeV3Filterer, error) {
	contract, err := bindExchangeV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3Filterer{contract: contract}, nil
}

// bindExchangeV3 binds a generic wrapper to an already deployed contract.
func bindExchangeV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV3 *ExchangeV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV3.Contract.ExchangeV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV3 *ExchangeV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ExchangeV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV3 *ExchangeV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ExchangeV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV3 *ExchangeV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV3 *ExchangeV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV3 *ExchangeV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV3.Contract.contract.Transact(opts, method, params...)
}

// AllowOnchainTransferFrom is a free data retrieval call binding the contract method 0x1de67966.
//
// Solidity: function allowOnchainTransferFrom() view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) AllowOnchainTransferFrom(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "allowOnchainTransferFrom")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowOnchainTransferFrom is a free data retrieval call binding the contract method 0x1de67966.
//
// Solidity: function allowOnchainTransferFrom() view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) AllowOnchainTransferFrom() (bool, error) {
	return _ExchangeV3.Contract.AllowOnchainTransferFrom(&_ExchangeV3.CallOpts)
}

// AllowOnchainTransferFrom is a free data retrieval call binding the contract method 0x1de67966.
//
// Solidity: function allowOnchainTransferFrom() view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) AllowOnchainTransferFrom() (bool, error) {
	return _ExchangeV3.Contract.AllowOnchainTransferFrom(&_ExchangeV3.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Session) DomainSeparator() ([32]byte, error) {
	return _ExchangeV3.Contract.DomainSeparator(&_ExchangeV3.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3CallerSession) DomainSeparator() ([32]byte, error) {
	return _ExchangeV3.Contract.DomainSeparator(&_ExchangeV3.CallOpts)
}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_ExchangeV3 *ExchangeV3Caller) GetAmountWithdrawable(opts *bind.CallOpts, owner common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getAmountWithdrawable", owner, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) GetAmountWithdrawable(owner common.Address, token common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetAmountWithdrawable(&_ExchangeV3.CallOpts, owner, token)
}

// GetAmountWithdrawable is a free data retrieval call binding the contract method 0x2fa5825f.
//
// Solidity: function getAmountWithdrawable(address owner, address token) view returns(uint256)
func (_ExchangeV3 *ExchangeV3CallerSession) GetAmountWithdrawable(owner common.Address, token common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetAmountWithdrawable(&_ExchangeV3.CallOpts, owner, token)
}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Caller) GetBlockHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getBlockHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) GetBlockHeight() (*big.Int, error) {
	return _ExchangeV3.Contract.GetBlockHeight(&_ExchangeV3.CallOpts)
}

// GetBlockHeight is a free data retrieval call binding the contract method 0x7bb96acb.
//
// Solidity: function getBlockHeight() view returns(uint256)
func (_ExchangeV3 *ExchangeV3CallerSession) GetBlockHeight() (*big.Int, error) {
	return _ExchangeV3.Contract.GetBlockHeight(&_ExchangeV3.CallOpts)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_ExchangeV3 *ExchangeV3Caller) GetBlockInfo(opts *bind.CallOpts, blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getBlockInfo", blockIdx)

	if err != nil {
		return *new(ExchangeDataBlockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeDataBlockInfo)).(*ExchangeDataBlockInfo)

	return out0, err

}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_ExchangeV3 *ExchangeV3Session) GetBlockInfo(blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	return _ExchangeV3.Contract.GetBlockInfo(&_ExchangeV3.CallOpts, blockIdx)
}

// GetBlockInfo is a free data retrieval call binding the contract method 0xbb141cf4.
//
// Solidity: function getBlockInfo(uint256 blockIdx) view returns((uint32,bytes28))
func (_ExchangeV3 *ExchangeV3CallerSession) GetBlockInfo(blockIdx *big.Int) (ExchangeDataBlockInfo, error) {
	return _ExchangeV3.Contract.GetBlockInfo(&_ExchangeV3.CallOpts, blockIdx)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ExchangeV3 *ExchangeV3Caller) GetConstants(opts *bind.CallOpts) (ExchangeDataConstants, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getConstants")

	if err != nil {
		return *new(ExchangeDataConstants), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeDataConstants)).(*ExchangeDataConstants)

	return out0, err

}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ExchangeV3 *ExchangeV3Session) GetConstants() (ExchangeDataConstants, error) {
	return _ExchangeV3.Contract.GetConstants(&_ExchangeV3.CallOpts)
}

// GetConstants is a free data retrieval call binding the contract method 0x9a295e73.
//
// Solidity: function getConstants() pure returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_ExchangeV3 *ExchangeV3CallerSession) GetConstants() (ExchangeDataConstants, error) {
	return _ExchangeV3.Contract.GetConstants(&_ExchangeV3.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) GetDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ExchangeV3 *ExchangeV3Session) GetDepositContract() (common.Address, error) {
	return _ExchangeV3.Contract.GetDepositContract(&_ExchangeV3.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) GetDepositContract() (common.Address, error) {
	return _ExchangeV3.Contract.GetDepositContract(&_ExchangeV3.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Caller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Session) GetDomainSeparator() ([32]byte, error) {
	return _ExchangeV3.Contract.GetDomainSeparator(&_ExchangeV3.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3CallerSession) GetDomainSeparator() ([32]byte, error) {
	return _ExchangeV3.Contract.GetDomainSeparator(&_ExchangeV3.CallOpts)
}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Caller) GetExchangeStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getExchangeStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) GetExchangeStake() (*big.Int, error) {
	return _ExchangeV3.Contract.GetExchangeStake(&_ExchangeV3.CallOpts)
}

// GetExchangeStake is a free data retrieval call binding the contract method 0xf732e021.
//
// Solidity: function getExchangeStake() view returns(uint256)
func (_ExchangeV3 *ExchangeV3CallerSession) GetExchangeStake() (*big.Int, error) {
	return _ExchangeV3.Contract.GetExchangeStake(&_ExchangeV3.CallOpts)
}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_ExchangeV3 *ExchangeV3Caller) GetMaxAgeDepositUntilWithdrawable(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getMaxAgeDepositUntilWithdrawable")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_ExchangeV3 *ExchangeV3Session) GetMaxAgeDepositUntilWithdrawable() (uint32, error) {
	return _ExchangeV3.Contract.GetMaxAgeDepositUntilWithdrawable(&_ExchangeV3.CallOpts)
}

// GetMaxAgeDepositUntilWithdrawable is a free data retrieval call binding the contract method 0x4a14cd84.
//
// Solidity: function getMaxAgeDepositUntilWithdrawable() view returns(uint32)
func (_ExchangeV3 *ExchangeV3CallerSession) GetMaxAgeDepositUntilWithdrawable() (uint32, error) {
	return _ExchangeV3.Contract.GetMaxAgeDepositUntilWithdrawable(&_ExchangeV3.CallOpts)
}

// GetMerkleAssetRoot is a free data retrieval call binding the contract method 0xa7bff050.
//
// Solidity: function getMerkleAssetRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Caller) GetMerkleAssetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getMerkleAssetRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleAssetRoot is a free data retrieval call binding the contract method 0xa7bff050.
//
// Solidity: function getMerkleAssetRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Session) GetMerkleAssetRoot() ([32]byte, error) {
	return _ExchangeV3.Contract.GetMerkleAssetRoot(&_ExchangeV3.CallOpts)
}

// GetMerkleAssetRoot is a free data retrieval call binding the contract method 0xa7bff050.
//
// Solidity: function getMerkleAssetRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3CallerSession) GetMerkleAssetRoot() ([32]byte, error) {
	return _ExchangeV3.Contract.GetMerkleAssetRoot(&_ExchangeV3.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Caller) GetMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3Session) GetMerkleRoot() ([32]byte, error) {
	return _ExchangeV3.Contract.GetMerkleRoot(&_ExchangeV3.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(bytes32)
func (_ExchangeV3 *ExchangeV3CallerSession) GetMerkleRoot() ([32]byte, error) {
	return _ExchangeV3.Contract.GetMerkleRoot(&_ExchangeV3.CallOpts)
}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Caller) GetNumAvailableForcedSlots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getNumAvailableForcedSlots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) GetNumAvailableForcedSlots() (*big.Int, error) {
	return _ExchangeV3.Contract.GetNumAvailableForcedSlots(&_ExchangeV3.CallOpts)
}

// GetNumAvailableForcedSlots is a free data retrieval call binding the contract method 0x69b91432.
//
// Solidity: function getNumAvailableForcedSlots() view returns(uint256)
func (_ExchangeV3 *ExchangeV3CallerSession) GetNumAvailableForcedSlots() (*big.Int, error) {
	return _ExchangeV3.Contract.GetNumAvailableForcedSlots(&_ExchangeV3.CallOpts)
}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint248)
func (_ExchangeV3 *ExchangeV3Caller) GetPendingDepositAmount(opts *bind.CallOpts, owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getPendingDepositAmount", owner, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint248)
func (_ExchangeV3 *ExchangeV3Session) GetPendingDepositAmount(owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetPendingDepositAmount(&_ExchangeV3.CallOpts, owner, tokenAddress)
}

// GetPendingDepositAmount is a free data retrieval call binding the contract method 0x438c2a42.
//
// Solidity: function getPendingDepositAmount(address owner, address tokenAddress) view returns(uint248)
func (_ExchangeV3 *ExchangeV3CallerSession) GetPendingDepositAmount(owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetPendingDepositAmount(&_ExchangeV3.CallOpts, owner, tokenAddress)
}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_ExchangeV3 *ExchangeV3Caller) GetProtocolFeeLastWithdrawnTime(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getProtocolFeeLastWithdrawnTime", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) GetProtocolFeeLastWithdrawnTime(tokenAddress common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetProtocolFeeLastWithdrawnTime(&_ExchangeV3.CallOpts, tokenAddress)
}

// GetProtocolFeeLastWithdrawnTime is a free data retrieval call binding the contract method 0xc8e26cae.
//
// Solidity: function getProtocolFeeLastWithdrawnTime(address tokenAddress) view returns(uint256)
func (_ExchangeV3 *ExchangeV3CallerSession) GetProtocolFeeLastWithdrawnTime(tokenAddress common.Address) (*big.Int, error) {
	return _ExchangeV3.Contract.GetProtocolFeeLastWithdrawnTime(&_ExchangeV3.CallOpts, tokenAddress)
}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3Caller) GetProtocolFeeValues(opts *bind.CallOpts) (struct {
	SyncedAt                uint32
	ProtocolFeeBips         uint8
	PreviousProtocolFeeBips uint8
}, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getProtocolFeeValues")

	outstruct := new(struct {
		SyncedAt                uint32
		ProtocolFeeBips         uint8
		PreviousProtocolFeeBips uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SyncedAt = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.ProtocolFeeBips = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.PreviousProtocolFeeBips = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3Session) GetProtocolFeeValues() (struct {
	SyncedAt                uint32
	ProtocolFeeBips         uint8
	PreviousProtocolFeeBips uint8
}, error) {
	return _ExchangeV3.Contract.GetProtocolFeeValues(&_ExchangeV3.CallOpts)
}

// GetProtocolFeeValues is a free data retrieval call binding the contract method 0x4597d3ce.
//
// Solidity: function getProtocolFeeValues() view returns(uint32 syncedAt, uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3CallerSession) GetProtocolFeeValues() (struct {
	SyncedAt                uint32
	ProtocolFeeBips         uint8
	PreviousProtocolFeeBips uint8
}, error) {
	return _ExchangeV3.Contract.GetProtocolFeeValues(&_ExchangeV3.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x543d7a19.
//
// Solidity: function getTokenAddress(uint32 tokenID) view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) GetTokenAddress(opts *bind.CallOpts, tokenID uint32) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getTokenAddress", tokenID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x543d7a19.
//
// Solidity: function getTokenAddress(uint32 tokenID) view returns(address)
func (_ExchangeV3 *ExchangeV3Session) GetTokenAddress(tokenID uint32) (common.Address, error) {
	return _ExchangeV3.Contract.GetTokenAddress(&_ExchangeV3.CallOpts, tokenID)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x543d7a19.
//
// Solidity: function getTokenAddress(uint32 tokenID) view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) GetTokenAddress(tokenID uint32) (common.Address, error) {
	return _ExchangeV3.Contract.GetTokenAddress(&_ExchangeV3.CallOpts, tokenID)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint32)
func (_ExchangeV3 *ExchangeV3Caller) GetTokenID(opts *bind.CallOpts, tokenAddress common.Address) (uint32, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getTokenID", tokenAddress)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint32)
func (_ExchangeV3 *ExchangeV3Session) GetTokenID(tokenAddress common.Address) (uint32, error) {
	return _ExchangeV3.Contract.GetTokenID(&_ExchangeV3.CallOpts, tokenAddress)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address tokenAddress) view returns(uint32)
func (_ExchangeV3 *ExchangeV3CallerSession) GetTokenID(tokenAddress common.Address) (uint32, error) {
	return _ExchangeV3.Contract.GetTokenID(&_ExchangeV3.CallOpts, tokenAddress)
}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0xcea69b86.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID) view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) GetWithdrawalRecipient(opts *bind.CallOpts, from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "getWithdrawalRecipient", from, to, token, amount, storageID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0xcea69b86.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID) view returns(address)
func (_ExchangeV3 *ExchangeV3Session) GetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	return _ExchangeV3.Contract.GetWithdrawalRecipient(&_ExchangeV3.CallOpts, from, to, token, amount, storageID)
}

// GetWithdrawalRecipient is a free data retrieval call binding the contract method 0xcea69b86.
//
// Solidity: function getWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID) view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) GetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32) (common.Address, error) {
	return _ExchangeV3.Contract.GetWithdrawalRecipient(&_ExchangeV3.CallOpts, from, to, token, amount, storageID)
}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsForcedWithdrawalPending(opts *bind.CallOpts, accountID uint32, token common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isForcedWithdrawalPending", accountID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsForcedWithdrawalPending(accountID uint32, token common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsForcedWithdrawalPending(&_ExchangeV3.CallOpts, accountID, token)
}

// IsForcedWithdrawalPending is a free data retrieval call binding the contract method 0x92f54afc.
//
// Solidity: function isForcedWithdrawalPending(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsForcedWithdrawalPending(accountID uint32, token common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsForcedWithdrawalPending(&_ExchangeV3.CallOpts, accountID, token)
}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsInWithdrawalMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isInWithdrawalMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsInWithdrawalMode() (bool, error) {
	return _ExchangeV3.Contract.IsInWithdrawalMode(&_ExchangeV3.CallOpts)
}

// IsInWithdrawalMode is a free data retrieval call binding the contract method 0x91b11ad4.
//
// Solidity: function isInWithdrawalMode() view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsInWithdrawalMode() (bool, error) {
	return _ExchangeV3.Contract.IsInWithdrawalMode(&_ExchangeV3.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsShutdown() (bool, error) {
	return _ExchangeV3.Contract.IsShutdown(&_ExchangeV3.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsShutdown() (bool, error) {
	return _ExchangeV3.Contract.IsShutdown(&_ExchangeV3.CallOpts)
}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsTransactionApproved(opts *bind.CallOpts, owner common.Address, transactionHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isTransactionApproved", owner, transactionHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsTransactionApproved(owner common.Address, transactionHash [32]byte) (bool, error) {
	return _ExchangeV3.Contract.IsTransactionApproved(&_ExchangeV3.CallOpts, owner, transactionHash)
}

// IsTransactionApproved is a free data retrieval call binding the contract method 0x6008cd1f.
//
// Solidity: function isTransactionApproved(address owner, bytes32 transactionHash) view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsTransactionApproved(owner common.Address, transactionHash [32]byte) (bool, error) {
	return _ExchangeV3.Contract.IsTransactionApproved(&_ExchangeV3.CallOpts, owner, transactionHash)
}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsUserOrAgent(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isUserOrAgent", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsUserOrAgent(owner common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsUserOrAgent(&_ExchangeV3.CallOpts, owner)
}

// IsUserOrAgent is a free data retrieval call binding the contract method 0x8a554abe.
//
// Solidity: function isUserOrAgent(address owner) view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsUserOrAgent(owner common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsUserOrAgent(&_ExchangeV3.CallOpts, owner)
}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3Caller) IsWithdrawnInWithdrawalMode(opts *bind.CallOpts, accountID uint32, token common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "isWithdrawnInWithdrawalMode", accountID, token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3Session) IsWithdrawnInWithdrawalMode(accountID uint32, token common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsWithdrawnInWithdrawalMode(&_ExchangeV3.CallOpts, accountID, token)
}

// IsWithdrawnInWithdrawalMode is a free data retrieval call binding the contract method 0x2c3d356b.
//
// Solidity: function isWithdrawnInWithdrawalMode(uint32 accountID, address token) view returns(bool)
func (_ExchangeV3 *ExchangeV3CallerSession) IsWithdrawnInWithdrawalMode(accountID uint32, token common.Address) (bool, error) {
	return _ExchangeV3.Contract.IsWithdrawnInWithdrawalMode(&_ExchangeV3.CallOpts, accountID, token)
}

// LoopringAddr is a free data retrieval call binding the contract method 0xfa3139e4.
//
// Solidity: function loopringAddr() view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) LoopringAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "loopringAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LoopringAddr is a free data retrieval call binding the contract method 0xfa3139e4.
//
// Solidity: function loopringAddr() view returns(address)
func (_ExchangeV3 *ExchangeV3Session) LoopringAddr() (common.Address, error) {
	return _ExchangeV3.Contract.LoopringAddr(&_ExchangeV3.CallOpts)
}

// LoopringAddr is a free data retrieval call binding the contract method 0xfa3139e4.
//
// Solidity: function loopringAddr() view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) LoopringAddr() (common.Address, error) {
	return _ExchangeV3.Contract.LoopringAddr(&_ExchangeV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV3 *ExchangeV3Session) Owner() (common.Address, error) {
	return _ExchangeV3.Contract.Owner(&_ExchangeV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) Owner() (common.Address, error) {
	return _ExchangeV3.Contract.Owner(&_ExchangeV3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ExchangeV3 *ExchangeV3Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ExchangeV3 *ExchangeV3Session) PendingOwner() (common.Address, error) {
	return _ExchangeV3.Contract.PendingOwner(&_ExchangeV3.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ExchangeV3 *ExchangeV3CallerSession) PendingOwner() (common.Address, error) {
	return _ExchangeV3.Contract.PendingOwner(&_ExchangeV3.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint32 maxAgeDepositUntilWithdrawable, bytes32 DOMAIN_SEPARATOR, address loopring, address blockVerifier, address agentRegistry, address depositContract, (uint256,uint256,uint256,uint256,uint256) depositState, bytes32 merkleRoot, bytes32 merkleAssetRoot, uint256 numBlocks, uint32 numPendingForcedTransactions, (uint32,uint8,uint8) protocolFeeData, (uint256,uint256) modeTime)
func (_ExchangeV3 *ExchangeV3Caller) State(opts *bind.CallOpts) (struct {
	MaxAgeDepositUntilWithdrawable uint32
	DOMAINSEPARATOR                [32]byte
	Loopring                       common.Address
	BlockVerifier                  common.Address
	AgentRegistry                  common.Address
	DepositContract                common.Address
	DepositState                   ExchangeDataDepositState
	MerkleRoot                     [32]byte
	MerkleAssetRoot                [32]byte
	NumBlocks                      *big.Int
	NumPendingForcedTransactions   uint32
	ProtocolFeeData                ExchangeDataProtocolFeeData
	ModeTime                       ExchangeDataModeTime
}, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		MaxAgeDepositUntilWithdrawable uint32
		DOMAINSEPARATOR                [32]byte
		Loopring                       common.Address
		BlockVerifier                  common.Address
		AgentRegistry                  common.Address
		DepositContract                common.Address
		DepositState                   ExchangeDataDepositState
		MerkleRoot                     [32]byte
		MerkleAssetRoot                [32]byte
		NumBlocks                      *big.Int
		NumPendingForcedTransactions   uint32
		ProtocolFeeData                ExchangeDataProtocolFeeData
		ModeTime                       ExchangeDataModeTime
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxAgeDepositUntilWithdrawable = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DOMAINSEPARATOR = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Loopring = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.BlockVerifier = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.AgentRegistry = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.DepositContract = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.DepositState = *abi.ConvertType(out[6], new(ExchangeDataDepositState)).(*ExchangeDataDepositState)
	outstruct.MerkleRoot = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)
	outstruct.MerkleAssetRoot = *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	outstruct.NumBlocks = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.NumPendingForcedTransactions = *abi.ConvertType(out[10], new(uint32)).(*uint32)
	outstruct.ProtocolFeeData = *abi.ConvertType(out[11], new(ExchangeDataProtocolFeeData)).(*ExchangeDataProtocolFeeData)
	outstruct.ModeTime = *abi.ConvertType(out[12], new(ExchangeDataModeTime)).(*ExchangeDataModeTime)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint32 maxAgeDepositUntilWithdrawable, bytes32 DOMAIN_SEPARATOR, address loopring, address blockVerifier, address agentRegistry, address depositContract, (uint256,uint256,uint256,uint256,uint256) depositState, bytes32 merkleRoot, bytes32 merkleAssetRoot, uint256 numBlocks, uint32 numPendingForcedTransactions, (uint32,uint8,uint8) protocolFeeData, (uint256,uint256) modeTime)
func (_ExchangeV3 *ExchangeV3Session) State() (struct {
	MaxAgeDepositUntilWithdrawable uint32
	DOMAINSEPARATOR                [32]byte
	Loopring                       common.Address
	BlockVerifier                  common.Address
	AgentRegistry                  common.Address
	DepositContract                common.Address
	DepositState                   ExchangeDataDepositState
	MerkleRoot                     [32]byte
	MerkleAssetRoot                [32]byte
	NumBlocks                      *big.Int
	NumPendingForcedTransactions   uint32
	ProtocolFeeData                ExchangeDataProtocolFeeData
	ModeTime                       ExchangeDataModeTime
}, error) {
	return _ExchangeV3.Contract.State(&_ExchangeV3.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint32 maxAgeDepositUntilWithdrawable, bytes32 DOMAIN_SEPARATOR, address loopring, address blockVerifier, address agentRegistry, address depositContract, (uint256,uint256,uint256,uint256,uint256) depositState, bytes32 merkleRoot, bytes32 merkleAssetRoot, uint256 numBlocks, uint32 numPendingForcedTransactions, (uint32,uint8,uint8) protocolFeeData, (uint256,uint256) modeTime)
func (_ExchangeV3 *ExchangeV3CallerSession) State() (struct {
	MaxAgeDepositUntilWithdrawable uint32
	DOMAINSEPARATOR                [32]byte
	Loopring                       common.Address
	BlockVerifier                  common.Address
	AgentRegistry                  common.Address
	DepositContract                common.Address
	DepositState                   ExchangeDataDepositState
	MerkleRoot                     [32]byte
	MerkleAssetRoot                [32]byte
	NumBlocks                      *big.Int
	NumPendingForcedTransactions   uint32
	ProtocolFeeData                ExchangeDataProtocolFeeData
	ModeTime                       ExchangeDataModeTime
}, error) {
	return _ExchangeV3.Contract.State(&_ExchangeV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ExchangeV3 *ExchangeV3Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExchangeV3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ExchangeV3 *ExchangeV3Session) Version() (string, error) {
	return _ExchangeV3.Contract.Version(&_ExchangeV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_ExchangeV3 *ExchangeV3CallerSession) Version() (string, error) {
	return _ExchangeV3.Contract.Version(&_ExchangeV3.CallOpts)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_ExchangeV3 *ExchangeV3Transactor) ApproveTransaction(opts *bind.TransactOpts, owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "approveTransaction", owner, transactionHash)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_ExchangeV3 *ExchangeV3Session) ApproveTransaction(owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ApproveTransaction(&_ExchangeV3.TransactOpts, owner, transactionHash)
}

// ApproveTransaction is a paid mutator transaction binding the contract method 0xd59acd25.
//
// Solidity: function approveTransaction(address owner, bytes32 transactionHash) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) ApproveTransaction(owner common.Address, transactionHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ApproveTransaction(&_ExchangeV3.TransactOpts, owner, transactionHash)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_ExchangeV3 *ExchangeV3Transactor) ApproveTransactions(opts *bind.TransactOpts, owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "approveTransactions", owners, transactionHashes)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_ExchangeV3 *ExchangeV3Session) ApproveTransactions(owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ApproveTransactions(&_ExchangeV3.TransactOpts, owners, transactionHashes)
}

// ApproveTransactions is a paid mutator transaction binding the contract method 0xcdb1b44b.
//
// Solidity: function approveTransactions(address[] owners, bytes32[] transactionHashes) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) ApproveTransactions(owners []common.Address, transactionHashes [][32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ApproveTransactions(&_ExchangeV3.TransactOpts, owners, transactionHashes)
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_ExchangeV3 *ExchangeV3Transactor) BurnExchangeStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "burnExchangeStake")
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_ExchangeV3 *ExchangeV3Session) BurnExchangeStake() (*types.Transaction, error) {
	return _ExchangeV3.Contract.BurnExchangeStake(&_ExchangeV3.TransactOpts)
}

// BurnExchangeStake is a paid mutator transaction binding the contract method 0x972f7565.
//
// Solidity: function burnExchangeStake() returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) BurnExchangeStake() (*types.Transaction, error) {
	return _ExchangeV3.Contract.BurnExchangeStake(&_ExchangeV3.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ExchangeV3 *ExchangeV3Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ExchangeV3 *ExchangeV3Session) ClaimOwnership() (*types.Transaction, error) {
	return _ExchangeV3.Contract.ClaimOwnership(&_ExchangeV3.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _ExchangeV3.Contract.ClaimOwnership(&_ExchangeV3.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8070f2ca.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint248 amount, bytes extraData) payable returns()
func (_ExchangeV3 *ExchangeV3Transactor) Deposit(opts *bind.TransactOpts, from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "deposit", from, to, tokenAddress, amount, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0x8070f2ca.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint248 amount, bytes extraData) payable returns()
func (_ExchangeV3 *ExchangeV3Session) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.Deposit(&_ExchangeV3.TransactOpts, from, to, tokenAddress, amount, extraData)
}

// Deposit is a paid mutator transaction binding the contract method 0x8070f2ca.
//
// Solidity: function deposit(address from, address to, address tokenAddress, uint248 amount, bytes extraData) payable returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) Deposit(from common.Address, to common.Address, tokenAddress common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.Deposit(&_ExchangeV3.TransactOpts, from, to, tokenAddress, amount, extraData)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_ExchangeV3 *ExchangeV3Transactor) ForceWithdraw(opts *bind.TransactOpts, owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "forceWithdraw", owner, token, accountID)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_ExchangeV3 *ExchangeV3Session) ForceWithdraw(owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ForceWithdraw(&_ExchangeV3.TransactOpts, owner, token, accountID)
}

// ForceWithdraw is a paid mutator transaction binding the contract method 0x47b67d05.
//
// Solidity: function forceWithdraw(address owner, address token, uint32 accountID) payable returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) ForceWithdraw(owner common.Address, token common.Address, accountID uint32) (*types.Transaction, error) {
	return _ExchangeV3.Contract.ForceWithdraw(&_ExchangeV3.TransactOpts, owner, token, accountID)
}

// Initialize is a paid mutator transaction binding the contract method 0x76cfcff6.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot, bytes32 _genesisMerkleAssetRoot) returns()
func (_ExchangeV3 *ExchangeV3Transactor) Initialize(opts *bind.TransactOpts, _loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte, _genesisMerkleAssetRoot [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "initialize", _loopring, _owner, _genesisMerkleRoot, _genesisMerkleAssetRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x76cfcff6.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot, bytes32 _genesisMerkleAssetRoot) returns()
func (_ExchangeV3 *ExchangeV3Session) Initialize(_loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte, _genesisMerkleAssetRoot [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.Initialize(&_ExchangeV3.TransactOpts, _loopring, _owner, _genesisMerkleRoot, _genesisMerkleAssetRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x76cfcff6.
//
// Solidity: function initialize(address _loopring, address _owner, bytes32 _genesisMerkleRoot, bytes32 _genesisMerkleAssetRoot) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) Initialize(_loopring common.Address, _owner common.Address, _genesisMerkleRoot [32]byte, _genesisMerkleAssetRoot [32]byte) (*types.Transaction, error) {
	return _ExchangeV3.Contract.Initialize(&_ExchangeV3.TransactOpts, _loopring, _owner, _genesisMerkleRoot, _genesisMerkleAssetRoot)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0x70f30de3.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, address token) returns()
func (_ExchangeV3 *ExchangeV3Transactor) NotifyForcedRequestTooOld(opts *bind.TransactOpts, accountID uint32, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "notifyForcedRequestTooOld", accountID, token)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0x70f30de3.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, address token) returns()
func (_ExchangeV3 *ExchangeV3Session) NotifyForcedRequestTooOld(accountID uint32, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.NotifyForcedRequestTooOld(&_ExchangeV3.TransactOpts, accountID, token)
}

// NotifyForcedRequestTooOld is a paid mutator transaction binding the contract method 0x70f30de3.
//
// Solidity: function notifyForcedRequestTooOld(uint32 accountID, address token) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) NotifyForcedRequestTooOld(accountID uint32, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.NotifyForcedRequestTooOld(&_ExchangeV3.TransactOpts, accountID, token)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_ExchangeV3 *ExchangeV3Transactor) OnchainTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "onchainTransferFrom", from, to, token, amount)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_ExchangeV3 *ExchangeV3Session) OnchainTransferFrom(from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.Contract.OnchainTransferFrom(&_ExchangeV3.TransactOpts, from, to, token, amount)
}

// OnchainTransferFrom is a paid mutator transaction binding the contract method 0x28068da3.
//
// Solidity: function onchainTransferFrom(address from, address to, address token, uint256 amount) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) OnchainTransferFrom(from common.Address, to common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.Contract.OnchainTransferFrom(&_ExchangeV3.TransactOpts, from, to, token, amount)
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_ExchangeV3 *ExchangeV3Transactor) RefreshBlockVerifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "refreshBlockVerifier")
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_ExchangeV3 *ExchangeV3Session) RefreshBlockVerifier() (*types.Transaction, error) {
	return _ExchangeV3.Contract.RefreshBlockVerifier(&_ExchangeV3.TransactOpts)
}

// RefreshBlockVerifier is a paid mutator transaction binding the contract method 0xc97890f1.
//
// Solidity: function refreshBlockVerifier() returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) RefreshBlockVerifier() (*types.Transaction, error) {
	return _ExchangeV3.Contract.RefreshBlockVerifier(&_ExchangeV3.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint32)
func (_ExchangeV3 *ExchangeV3Transactor) RegisterToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "registerToken", tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint32)
func (_ExchangeV3 *ExchangeV3Session) RegisterToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.RegisterToken(&_ExchangeV3.TransactOpts, tokenAddress)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tokenAddress) returns(uint32)
func (_ExchangeV3 *ExchangeV3TransactorSession) RegisterToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.RegisterToken(&_ExchangeV3.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV3 *ExchangeV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV3 *ExchangeV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeV3.Contract.RenounceOwnership(&_ExchangeV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExchangeV3.Contract.RenounceOwnership(&_ExchangeV3.TransactOpts)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_ExchangeV3 *ExchangeV3Transactor) SetAllowOnchainTransferFrom(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "setAllowOnchainTransferFrom", value)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_ExchangeV3 *ExchangeV3Session) SetAllowOnchainTransferFrom(value bool) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetAllowOnchainTransferFrom(&_ExchangeV3.TransactOpts, value)
}

// SetAllowOnchainTransferFrom is a paid mutator transaction binding the contract method 0x01b1eb07.
//
// Solidity: function setAllowOnchainTransferFrom(bool value) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) SetAllowOnchainTransferFrom(value bool) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetAllowOnchainTransferFrom(&_ExchangeV3.TransactOpts, value)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_ExchangeV3 *ExchangeV3Transactor) SetDepositContract(opts *bind.TransactOpts, _depositContract common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "setDepositContract", _depositContract)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_ExchangeV3 *ExchangeV3Session) SetDepositContract(_depositContract common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetDepositContract(&_ExchangeV3.TransactOpts, _depositContract)
}

// SetDepositContract is a paid mutator transaction binding the contract method 0x0ec2e821.
//
// Solidity: function setDepositContract(address _depositContract) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) SetDepositContract(_depositContract common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetDepositContract(&_ExchangeV3.TransactOpts, _depositContract)
}

// SetDepositParams is a paid mutator transaction binding the contract method 0xe7957992.
//
// Solidity: function setDepositParams(uint256 freeDepositMax, uint256 freeDepositRemained, uint256 freeSlotPerBlock, uint256 depositFee) returns()
func (_ExchangeV3 *ExchangeV3Transactor) SetDepositParams(opts *bind.TransactOpts, freeDepositMax *big.Int, freeDepositRemained *big.Int, freeSlotPerBlock *big.Int, depositFee *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "setDepositParams", freeDepositMax, freeDepositRemained, freeSlotPerBlock, depositFee)
}

// SetDepositParams is a paid mutator transaction binding the contract method 0xe7957992.
//
// Solidity: function setDepositParams(uint256 freeDepositMax, uint256 freeDepositRemained, uint256 freeSlotPerBlock, uint256 depositFee) returns()
func (_ExchangeV3 *ExchangeV3Session) SetDepositParams(freeDepositMax *big.Int, freeDepositRemained *big.Int, freeSlotPerBlock *big.Int, depositFee *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetDepositParams(&_ExchangeV3.TransactOpts, freeDepositMax, freeDepositRemained, freeSlotPerBlock, depositFee)
}

// SetDepositParams is a paid mutator transaction binding the contract method 0xe7957992.
//
// Solidity: function setDepositParams(uint256 freeDepositMax, uint256 freeDepositRemained, uint256 freeSlotPerBlock, uint256 depositFee) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) SetDepositParams(freeDepositMax *big.Int, freeDepositRemained *big.Int, freeSlotPerBlock *big.Int, depositFee *big.Int) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetDepositParams(&_ExchangeV3.TransactOpts, freeDepositMax, freeDepositRemained, freeSlotPerBlock, depositFee)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_ExchangeV3 *ExchangeV3Transactor) SetMaxAgeDepositUntilWithdrawable(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "setMaxAgeDepositUntilWithdrawable", newValue)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_ExchangeV3 *ExchangeV3Session) SetMaxAgeDepositUntilWithdrawable(newValue uint32) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetMaxAgeDepositUntilWithdrawable(&_ExchangeV3.TransactOpts, newValue)
}

// SetMaxAgeDepositUntilWithdrawable is a paid mutator transaction binding the contract method 0x960af2d9.
//
// Solidity: function setMaxAgeDepositUntilWithdrawable(uint32 newValue) returns(uint32)
func (_ExchangeV3 *ExchangeV3TransactorSession) SetMaxAgeDepositUntilWithdrawable(newValue uint32) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetMaxAgeDepositUntilWithdrawable(&_ExchangeV3.TransactOpts, newValue)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0x7907140e.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID, address newRecipient) returns()
func (_ExchangeV3 *ExchangeV3Transactor) SetWithdrawalRecipient(opts *bind.TransactOpts, from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "setWithdrawalRecipient", from, to, token, amount, storageID, newRecipient)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0x7907140e.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID, address newRecipient) returns()
func (_ExchangeV3 *ExchangeV3Session) SetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetWithdrawalRecipient(&_ExchangeV3.TransactOpts, from, to, token, amount, storageID, newRecipient)
}

// SetWithdrawalRecipient is a paid mutator transaction binding the contract method 0x7907140e.
//
// Solidity: function setWithdrawalRecipient(address from, address to, address token, uint248 amount, uint32 storageID, address newRecipient) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) SetWithdrawalRecipient(from common.Address, to common.Address, token common.Address, amount *big.Int, storageID uint32, newRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SetWithdrawalRecipient(&_ExchangeV3.TransactOpts, from, to, token, amount, storageID, newRecipient)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_ExchangeV3 *ExchangeV3Transactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_ExchangeV3 *ExchangeV3Session) Shutdown() (*types.Transaction, error) {
	return _ExchangeV3.Contract.Shutdown(&_ExchangeV3.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns(bool success)
func (_ExchangeV3 *ExchangeV3TransactorSession) Shutdown() (*types.Transaction, error) {
	return _ExchangeV3.Contract.Shutdown(&_ExchangeV3.TransactOpts)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_ExchangeV3 *ExchangeV3Transactor) SubmitBlocks(opts *bind.TransactOpts, blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "submitBlocks", blocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_ExchangeV3 *ExchangeV3Session) SubmitBlocks(blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SubmitBlocks(&_ExchangeV3.TransactOpts, blocks)
}

// SubmitBlocks is a paid mutator transaction binding the contract method 0x53228430.
//
// Solidity: function submitBlocks((uint8,uint16,uint8,bytes,uint256[8],bool,bytes,bytes)[] blocks) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) SubmitBlocks(blocks []ExchangeDataBlock) (*types.Transaction, error) {
	return _ExchangeV3.Contract.SubmitBlocks(&_ExchangeV3.TransactOpts, blocks)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV3 *ExchangeV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV3 *ExchangeV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.TransferOwnership(&_ExchangeV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.TransferOwnership(&_ExchangeV3.TransactOpts, newOwner)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawExchangeFees(opts *bind.TransactOpts, token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawExchangeFees", token, recipient)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_ExchangeV3 *ExchangeV3Session) WithdrawExchangeFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawExchangeFees(&_ExchangeV3.TransactOpts, token, recipient)
}

// WithdrawExchangeFees is a paid mutator transaction binding the contract method 0xa75f8a4e.
//
// Solidity: function withdrawExchangeFees(address token, address recipient) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawExchangeFees(token common.Address, recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawExchangeFees(&_ExchangeV3.TransactOpts, token, recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawExchangeStake(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawExchangeStake", recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_ExchangeV3 *ExchangeV3Session) WithdrawExchangeStake(recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawExchangeStake(&_ExchangeV3.TransactOpts, recipient)
}

// WithdrawExchangeStake is a paid mutator transaction binding the contract method 0xce2ec5de.
//
// Solidity: function withdrawExchangeStake(address recipient) returns(uint256)
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawExchangeStake(recipient common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawExchangeStake(&_ExchangeV3.TransactOpts, recipient)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawFromApprovedWithdrawals(opts *bind.TransactOpts, owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawFromApprovedWithdrawals", owners, tokens)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_ExchangeV3 *ExchangeV3Session) WithdrawFromApprovedWithdrawals(owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromApprovedWithdrawals(&_ExchangeV3.TransactOpts, owners, tokens)
}

// WithdrawFromApprovedWithdrawals is a paid mutator transaction binding the contract method 0xcd097b4f.
//
// Solidity: function withdrawFromApprovedWithdrawals(address[] owners, address[] tokens) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawFromApprovedWithdrawals(owners []common.Address, tokens []common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromApprovedWithdrawals(&_ExchangeV3.TransactOpts, owners, tokens)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawFromDepositRequest(opts *bind.TransactOpts, owner common.Address, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawFromDepositRequest", owner, token)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_ExchangeV3 *ExchangeV3Session) WithdrawFromDepositRequest(owner common.Address, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromDepositRequest(&_ExchangeV3.TransactOpts, owner, token)
}

// WithdrawFromDepositRequest is a paid mutator transaction binding the contract method 0x8d2a8888.
//
// Solidity: function withdrawFromDepositRequest(address owner, address token) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawFromDepositRequest(owner common.Address, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromDepositRequest(&_ExchangeV3.TransactOpts, owner, token)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x1280f724.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32),(uint32,uint248),uint256[48],uint256[48]) merkleProof) returns()
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawFromMerkleTree(opts *bind.TransactOpts, merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawFromMerkleTree", merkleProof)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x1280f724.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32),(uint32,uint248),uint256[48],uint256[48]) merkleProof) returns()
func (_ExchangeV3 *ExchangeV3Session) WithdrawFromMerkleTree(merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromMerkleTree(&_ExchangeV3.TransactOpts, merkleProof)
}

// WithdrawFromMerkleTree is a paid mutator transaction binding the contract method 0x1280f724.
//
// Solidity: function withdrawFromMerkleTree(((uint32,address,uint256,uint256,uint32),(uint32,uint248),uint256[48],uint256[48]) merkleProof) returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawFromMerkleTree(merkleProof ExchangeDataMerkleProof) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawFromMerkleTree(&_ExchangeV3.TransactOpts, merkleProof)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_ExchangeV3 *ExchangeV3Transactor) WithdrawProtocolFees(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.contract.Transact(opts, "withdrawProtocolFees", token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_ExchangeV3 *ExchangeV3Session) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawProtocolFees(&_ExchangeV3.TransactOpts, token)
}

// WithdrawProtocolFees is a paid mutator transaction binding the contract method 0x2d80caa5.
//
// Solidity: function withdrawProtocolFees(address token) payable returns()
func (_ExchangeV3 *ExchangeV3TransactorSession) WithdrawProtocolFees(token common.Address) (*types.Transaction, error) {
	return _ExchangeV3.Contract.WithdrawProtocolFees(&_ExchangeV3.TransactOpts, token)
}

// ExchangeV3BlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the ExchangeV3 contract.
type ExchangeV3BlockSubmittedIterator struct {
	Event *ExchangeV3BlockSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3BlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3BlockSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3BlockSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3BlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3BlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3BlockSubmitted represents a BlockSubmitted event raised by the ExchangeV3 contract.
type ExchangeV3BlockSubmitted struct {
	BlockIdx       *big.Int
	MerkleRoot     [32]byte
	PublicDataHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_ExchangeV3 *ExchangeV3Filterer) FilterBlockSubmitted(opts *bind.FilterOpts, blockIdx []*big.Int) (*ExchangeV3BlockSubmittedIterator, error) {

	var blockIdxRule []interface{}
	for _, blockIdxItem := range blockIdx {
		blockIdxRule = append(blockIdxRule, blockIdxItem)
	}

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "BlockSubmitted", blockIdxRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3BlockSubmittedIterator{contract: _ExchangeV3.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_ExchangeV3 *ExchangeV3Filterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *ExchangeV3BlockSubmitted, blockIdx []*big.Int) (event.Subscription, error) {

	var blockIdxRule []interface{}
	for _, blockIdxItem := range blockIdx {
		blockIdxRule = append(blockIdxRule, blockIdxItem)
	}

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "BlockSubmitted", blockIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3BlockSubmitted)
				if err := _ExchangeV3.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockSubmitted is a log parse operation binding the contract event 0xcc86d9ed29ebae540f9d25a4976d4da36ea4161b854b8ecf18f491cf6b0feb5c.
//
// Solidity: event BlockSubmitted(uint256 indexed blockIdx, bytes32 merkleRoot, bytes32 publicDataHash)
func (_ExchangeV3 *ExchangeV3Filterer) ParseBlockSubmitted(log types.Log) (*ExchangeV3BlockSubmitted, error) {
	event := new(ExchangeV3BlockSubmitted)
	if err := _ExchangeV3.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3DepositRequestedIterator is returned from FilterDepositRequested and is used to iterate over the raw logs and unpacked data for DepositRequested events raised by the ExchangeV3 contract.
type ExchangeV3DepositRequestedIterator struct {
	Event *ExchangeV3DepositRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3DepositRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3DepositRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3DepositRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3DepositRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3DepositRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3DepositRequested represents a DepositRequested event raised by the ExchangeV3 contract.
type ExchangeV3DepositRequested struct {
	From    common.Address
	To      common.Address
	Token   common.Address
	TokenId uint32
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositRequested is a free log retrieval operation binding the contract event 0x5cc43c6b47c350258362c540eb61d09c2060eaa8b760707b2892d08a01f6e9c3.
//
// Solidity: event DepositRequested(address from, address to, address token, uint32 tokenId, uint248 amount)
func (_ExchangeV3 *ExchangeV3Filterer) FilterDepositRequested(opts *bind.FilterOpts) (*ExchangeV3DepositRequestedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3DepositRequestedIterator{contract: _ExchangeV3.contract, event: "DepositRequested", logs: logs, sub: sub}, nil
}

// WatchDepositRequested is a free log subscription operation binding the contract event 0x5cc43c6b47c350258362c540eb61d09c2060eaa8b760707b2892d08a01f6e9c3.
//
// Solidity: event DepositRequested(address from, address to, address token, uint32 tokenId, uint248 amount)
func (_ExchangeV3 *ExchangeV3Filterer) WatchDepositRequested(opts *bind.WatchOpts, sink chan<- *ExchangeV3DepositRequested) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "DepositRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3DepositRequested)
				if err := _ExchangeV3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositRequested is a log parse operation binding the contract event 0x5cc43c6b47c350258362c540eb61d09c2060eaa8b760707b2892d08a01f6e9c3.
//
// Solidity: event DepositRequested(address from, address to, address token, uint32 tokenId, uint248 amount)
func (_ExchangeV3 *ExchangeV3Filterer) ParseDepositRequested(log types.Log) (*ExchangeV3DepositRequested, error) {
	event := new(ExchangeV3DepositRequested)
	if err := _ExchangeV3.contract.UnpackLog(event, "DepositRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3ExchangeClonedIterator is returned from FilterExchangeCloned and is used to iterate over the raw logs and unpacked data for ExchangeCloned events raised by the ExchangeV3 contract.
type ExchangeV3ExchangeClonedIterator struct {
	Event *ExchangeV3ExchangeCloned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3ExchangeClonedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3ExchangeCloned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3ExchangeCloned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3ExchangeClonedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3ExchangeClonedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3ExchangeCloned represents a ExchangeCloned event raised by the ExchangeV3 contract.
type ExchangeV3ExchangeCloned struct {
	ExchangeAddress   common.Address
	Owner             common.Address
	GenesisMerkleRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterExchangeCloned is a free log retrieval operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_ExchangeV3 *ExchangeV3Filterer) FilterExchangeCloned(opts *bind.FilterOpts) (*ExchangeV3ExchangeClonedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "ExchangeCloned")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3ExchangeClonedIterator{contract: _ExchangeV3.contract, event: "ExchangeCloned", logs: logs, sub: sub}, nil
}

// WatchExchangeCloned is a free log subscription operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_ExchangeV3 *ExchangeV3Filterer) WatchExchangeCloned(opts *bind.WatchOpts, sink chan<- *ExchangeV3ExchangeCloned) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "ExchangeCloned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3ExchangeCloned)
				if err := _ExchangeV3.contract.UnpackLog(event, "ExchangeCloned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExchangeCloned is a log parse operation binding the contract event 0x51499c854dc582463b59fee9dfc65a8ff68a6e4dbd4297d122f396b16950ddea.
//
// Solidity: event ExchangeCloned(address exchangeAddress, address owner, bytes32 genesisMerkleRoot)
func (_ExchangeV3 *ExchangeV3Filterer) ParseExchangeCloned(log types.Log) (*ExchangeV3ExchangeCloned, error) {
	event := new(ExchangeV3ExchangeCloned)
	if err := _ExchangeV3.contract.UnpackLog(event, "ExchangeCloned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3ForcedWithdrawalRequestedIterator is returned from FilterForcedWithdrawalRequested and is used to iterate over the raw logs and unpacked data for ForcedWithdrawalRequested events raised by the ExchangeV3 contract.
type ExchangeV3ForcedWithdrawalRequestedIterator struct {
	Event *ExchangeV3ForcedWithdrawalRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3ForcedWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3ForcedWithdrawalRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3ForcedWithdrawalRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3ForcedWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3ForcedWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3ForcedWithdrawalRequested represents a ForcedWithdrawalRequested event raised by the ExchangeV3 contract.
type ExchangeV3ForcedWithdrawalRequested struct {
	Owner     common.Address
	Token     common.Address
	AccountID uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterForcedWithdrawalRequested is a free log retrieval operation binding the contract event 0xd5e6217f4bf11b187f7e454fa39ac993048526aa50b7a4125cd3e02eea66e3dd.
//
// Solidity: event ForcedWithdrawalRequested(address owner, address token, uint32 accountID)
func (_ExchangeV3 *ExchangeV3Filterer) FilterForcedWithdrawalRequested(opts *bind.FilterOpts) (*ExchangeV3ForcedWithdrawalRequestedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "ForcedWithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3ForcedWithdrawalRequestedIterator{contract: _ExchangeV3.contract, event: "ForcedWithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchForcedWithdrawalRequested is a free log subscription operation binding the contract event 0xd5e6217f4bf11b187f7e454fa39ac993048526aa50b7a4125cd3e02eea66e3dd.
//
// Solidity: event ForcedWithdrawalRequested(address owner, address token, uint32 accountID)
func (_ExchangeV3 *ExchangeV3Filterer) WatchForcedWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *ExchangeV3ForcedWithdrawalRequested) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "ForcedWithdrawalRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3ForcedWithdrawalRequested)
				if err := _ExchangeV3.contract.UnpackLog(event, "ForcedWithdrawalRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForcedWithdrawalRequested is a log parse operation binding the contract event 0xd5e6217f4bf11b187f7e454fa39ac993048526aa50b7a4125cd3e02eea66e3dd.
//
// Solidity: event ForcedWithdrawalRequested(address owner, address token, uint32 accountID)
func (_ExchangeV3 *ExchangeV3Filterer) ParseForcedWithdrawalRequested(log types.Log) (*ExchangeV3ForcedWithdrawalRequested, error) {
	event := new(ExchangeV3ForcedWithdrawalRequested)
	if err := _ExchangeV3.contract.UnpackLog(event, "ForcedWithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExchangeV3 contract.
type ExchangeV3OwnershipTransferredIterator struct {
	Event *ExchangeV3OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3OwnershipTransferred represents a OwnershipTransferred event raised by the ExchangeV3 contract.
type ExchangeV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExchangeV3 *ExchangeV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeV3OwnershipTransferredIterator{contract: _ExchangeV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExchangeV3 *ExchangeV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3OwnershipTransferred)
				if err := _ExchangeV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExchangeV3 *ExchangeV3Filterer) ParseOwnershipTransferred(log types.Log) (*ExchangeV3OwnershipTransferred, error) {
	event := new(ExchangeV3OwnershipTransferred)
	if err := _ExchangeV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3ProtocolFeesUpdatedIterator is returned from FilterProtocolFeesUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeesUpdated events raised by the ExchangeV3 contract.
type ExchangeV3ProtocolFeesUpdatedIterator struct {
	Event *ExchangeV3ProtocolFeesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3ProtocolFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3ProtocolFeesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3ProtocolFeesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3ProtocolFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3ProtocolFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3ProtocolFeesUpdated represents a ProtocolFeesUpdated event raised by the ExchangeV3 contract.
type ExchangeV3ProtocolFeesUpdated struct {
	ProtocolFeeBips         uint8
	PreviousProtocolFeeBips uint8
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeesUpdated is a free log retrieval operation binding the contract event 0xe75a0680f39fd3ca6a9010d6e1d8633ab5bdcf5d2559bae1d55b936e73649db4.
//
// Solidity: event ProtocolFeesUpdated(uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3Filterer) FilterProtocolFeesUpdated(opts *bind.FilterOpts) (*ExchangeV3ProtocolFeesUpdatedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "ProtocolFeesUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3ProtocolFeesUpdatedIterator{contract: _ExchangeV3.contract, event: "ProtocolFeesUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeesUpdated is a free log subscription operation binding the contract event 0xe75a0680f39fd3ca6a9010d6e1d8633ab5bdcf5d2559bae1d55b936e73649db4.
//
// Solidity: event ProtocolFeesUpdated(uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3Filterer) WatchProtocolFeesUpdated(opts *bind.WatchOpts, sink chan<- *ExchangeV3ProtocolFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "ProtocolFeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3ProtocolFeesUpdated)
				if err := _ExchangeV3.contract.UnpackLog(event, "ProtocolFeesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProtocolFeesUpdated is a log parse operation binding the contract event 0xe75a0680f39fd3ca6a9010d6e1d8633ab5bdcf5d2559bae1d55b936e73649db4.
//
// Solidity: event ProtocolFeesUpdated(uint8 protocolFeeBips, uint8 previousProtocolFeeBips)
func (_ExchangeV3 *ExchangeV3Filterer) ParseProtocolFeesUpdated(log types.Log) (*ExchangeV3ProtocolFeesUpdated, error) {
	event := new(ExchangeV3ProtocolFeesUpdated)
	if err := _ExchangeV3.contract.UnpackLog(event, "ProtocolFeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3ShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the ExchangeV3 contract.
type ExchangeV3ShutdownIterator struct {
	Event *ExchangeV3Shutdown // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3ShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3Shutdown)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3Shutdown)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3ShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3ShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3Shutdown represents a Shutdown event raised by the ExchangeV3 contract.
type ExchangeV3Shutdown struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) FilterShutdown(opts *bind.FilterOpts) (*ExchangeV3ShutdownIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3ShutdownIterator{contract: _ExchangeV3.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *ExchangeV3Shutdown) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3Shutdown)
				if err := _ExchangeV3.contract.UnpackLog(event, "Shutdown", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseShutdown is a log parse operation binding the contract event 0x574214b195bf5273a95bb4498e35cf1fde0ce327c727a95ec2ab359f7ba4e11a.
//
// Solidity: event Shutdown(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) ParseShutdown(log types.Log) (*ExchangeV3Shutdown, error) {
	event := new(ExchangeV3Shutdown)
	if err := _ExchangeV3.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3TokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the ExchangeV3 contract.
type ExchangeV3TokenRegisteredIterator struct {
	Event *ExchangeV3TokenRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3TokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3TokenRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3TokenRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3TokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3TokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3TokenRegistered represents a TokenRegistered event raised by the ExchangeV3 contract.
type ExchangeV3TokenRegistered struct {
	Token   common.Address
	TokenId uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0x335b50bd5127c8bc5db255e26c4714815399743ddb497c0a3502f0e1e09c5a86.
//
// Solidity: event TokenRegistered(address token, uint32 tokenId)
func (_ExchangeV3 *ExchangeV3Filterer) FilterTokenRegistered(opts *bind.FilterOpts) (*ExchangeV3TokenRegisteredIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3TokenRegisteredIterator{contract: _ExchangeV3.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0x335b50bd5127c8bc5db255e26c4714815399743ddb497c0a3502f0e1e09c5a86.
//
// Solidity: event TokenRegistered(address token, uint32 tokenId)
func (_ExchangeV3 *ExchangeV3Filterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *ExchangeV3TokenRegistered) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3TokenRegistered)
				if err := _ExchangeV3.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenRegistered is a log parse operation binding the contract event 0x335b50bd5127c8bc5db255e26c4714815399743ddb497c0a3502f0e1e09c5a86.
//
// Solidity: event TokenRegistered(address token, uint32 tokenId)
func (_ExchangeV3 *ExchangeV3Filterer) ParseTokenRegistered(log types.Log) (*ExchangeV3TokenRegistered, error) {
	event := new(ExchangeV3TokenRegistered)
	if err := _ExchangeV3.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3TransactionApprovedIterator is returned from FilterTransactionApproved and is used to iterate over the raw logs and unpacked data for TransactionApproved events raised by the ExchangeV3 contract.
type ExchangeV3TransactionApprovedIterator struct {
	Event *ExchangeV3TransactionApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3TransactionApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3TransactionApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3TransactionApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3TransactionApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3TransactionApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3TransactionApproved represents a TransactionApproved event raised by the ExchangeV3 contract.
type ExchangeV3TransactionApproved struct {
	Owner           common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionApproved is a free log retrieval operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_ExchangeV3 *ExchangeV3Filterer) FilterTransactionApproved(opts *bind.FilterOpts) (*ExchangeV3TransactionApprovedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3TransactionApprovedIterator{contract: _ExchangeV3.contract, event: "TransactionApproved", logs: logs, sub: sub}, nil
}

// WatchTransactionApproved is a free log subscription operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_ExchangeV3 *ExchangeV3Filterer) WatchTransactionApproved(opts *bind.WatchOpts, sink chan<- *ExchangeV3TransactionApproved) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "TransactionApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3TransactionApproved)
				if err := _ExchangeV3.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionApproved is a log parse operation binding the contract event 0x7115b98931e7aca6d0aa11d10fe28877316a661a44c4bfc93c76b19dbbf5b107.
//
// Solidity: event TransactionApproved(address owner, bytes32 transactionHash)
func (_ExchangeV3 *ExchangeV3Filterer) ParseTransactionApproved(log types.Log) (*ExchangeV3TransactionApproved, error) {
	event := new(ExchangeV3TransactionApproved)
	if err := _ExchangeV3.contract.UnpackLog(event, "TransactionApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3WithdrawalCompletedIterator is returned from FilterWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for WithdrawalCompleted events raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalCompletedIterator struct {
	Event *ExchangeV3WithdrawalCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3WithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3WithdrawalCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3WithdrawalCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3WithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3WithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3WithdrawalCompleted represents a WithdrawalCompleted event raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalCompleted struct {
	Category uint8
	From     common.Address
	To       common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCompleted is a free log retrieval operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) FilterWithdrawalCompleted(opts *bind.FilterOpts) (*ExchangeV3WithdrawalCompletedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3WithdrawalCompletedIterator{contract: _ExchangeV3.contract, event: "WithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCompleted is a free log subscription operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) WatchWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ExchangeV3WithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "WithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3WithdrawalCompleted)
				if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalCompleted is a log parse operation binding the contract event 0x0d22d7344fc6871a839149fd89f9fd88a6c29cf797a67114772a9d4df5f8c96b.
//
// Solidity: event WithdrawalCompleted(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) ParseWithdrawalCompleted(log types.Log) (*ExchangeV3WithdrawalCompleted, error) {
	event := new(ExchangeV3WithdrawalCompleted)
	if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3WithdrawalFailedIterator is returned from FilterWithdrawalFailed and is used to iterate over the raw logs and unpacked data for WithdrawalFailed events raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalFailedIterator struct {
	Event *ExchangeV3WithdrawalFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3WithdrawalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3WithdrawalFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3WithdrawalFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3WithdrawalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3WithdrawalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3WithdrawalFailed represents a WithdrawalFailed event raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalFailed struct {
	Category uint8
	From     common.Address
	To       common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFailed is a free log retrieval operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) FilterWithdrawalFailed(opts *bind.FilterOpts) (*ExchangeV3WithdrawalFailedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "WithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3WithdrawalFailedIterator{contract: _ExchangeV3.contract, event: "WithdrawalFailed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFailed is a free log subscription operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) WatchWithdrawalFailed(opts *bind.WatchOpts, sink chan<- *ExchangeV3WithdrawalFailed) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "WithdrawalFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3WithdrawalFailed)
				if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalFailed is a log parse operation binding the contract event 0xc554475da3c47d29ea41e4f1ba76e92dd5b4b7764dbd96e1fb0d5d1e419bcf8b.
//
// Solidity: event WithdrawalFailed(uint8 category, address from, address to, address token, uint256 amount)
func (_ExchangeV3 *ExchangeV3Filterer) ParseWithdrawalFailed(log types.Log) (*ExchangeV3WithdrawalFailed, error) {
	event := new(ExchangeV3WithdrawalFailed)
	if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV3WithdrawalModeActivatedIterator is returned from FilterWithdrawalModeActivated and is used to iterate over the raw logs and unpacked data for WithdrawalModeActivated events raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalModeActivatedIterator struct {
	Event *ExchangeV3WithdrawalModeActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeV3WithdrawalModeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV3WithdrawalModeActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeV3WithdrawalModeActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeV3WithdrawalModeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV3WithdrawalModeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV3WithdrawalModeActivated represents a WithdrawalModeActivated event raised by the ExchangeV3 contract.
type ExchangeV3WithdrawalModeActivated struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalModeActivated is a free log retrieval operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) FilterWithdrawalModeActivated(opts *bind.FilterOpts) (*ExchangeV3WithdrawalModeActivatedIterator, error) {

	logs, sub, err := _ExchangeV3.contract.FilterLogs(opts, "WithdrawalModeActivated")
	if err != nil {
		return nil, err
	}
	return &ExchangeV3WithdrawalModeActivatedIterator{contract: _ExchangeV3.contract, event: "WithdrawalModeActivated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalModeActivated is a free log subscription operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) WatchWithdrawalModeActivated(opts *bind.WatchOpts, sink chan<- *ExchangeV3WithdrawalModeActivated) (event.Subscription, error) {

	logs, sub, err := _ExchangeV3.contract.WatchLogs(opts, "WithdrawalModeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV3WithdrawalModeActivated)
				if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalModeActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalModeActivated is a log parse operation binding the contract event 0x5b8f5ce93a49fc6eab534327f9c77fd2966e16278d6135cc0d99e6b6450c7963.
//
// Solidity: event WithdrawalModeActivated(uint256 timestamp)
func (_ExchangeV3 *ExchangeV3Filterer) ParseWithdrawalModeActivated(log types.Log) (*ExchangeV3WithdrawalModeActivated, error) {
	event := new(ExchangeV3WithdrawalModeActivated)
	if err := _ExchangeV3.contract.UnpackLog(event, "WithdrawalModeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
