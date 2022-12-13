package entity

import (
	"math/big"
	"staterecovery/common/util"

)

type State struct {
	BlockId int `json:"blockId"`
	AccountsAssetRoot string `json:"accounts_asset_root"`
	AccountsAssetTree map[string][]string `json:"accounts_asset_tree"`
	AccountsRoot string `json:"accounts_root"`
	AccountsTree map[string][]string `json:"accounts_tree"`
	AccountsValues map[string]*AccountValue `json:"accounts_values"`
	ExchangeID int `json:"exchangeID"`
	ProtocolAccount string `json:"protocolAccount"`
	WhiteAccounts []*WhiteAccount `json:"whiteAccounts"`
}

type WhiteAccount struct {
	AccountID string `json:"accountID"`
	Owner string `json:"owner"`
}

type AccountValue struct {
	BalancesLeafs map[string]*BalanceNode `json:"_balancesLeafs"`
	BalancesTree *ChildTree `json:"_balancesTree"`
	StorageLeafs map[string]*StorageNode `json:"_storageLeafs"`
	StorageTree *ChildTree `json:"_storageTree"`
	AppKeyPublicKeyX string `json:"appKeyPublicKeyX"`
	AppKeyPublicKeyY string `json:"appKeyPublicKeyY"`
	DisableAppKeySpotTrade int `json:"disableAppKeySpotTrade"`
	DisableAppKeyTransferToOther int `json:"disableAppKeyTransferToOther"`
	DisableAppKeyWithdraw int `json:"disableAppKeyWithdraw"`
	Nonce uint64 `json:"nonce"`
	Owner string `json:"owner"`
	PublicKeyX string `json:"publicKeyX"`
	PublicKeyY string `json:"publicKeyY"`

	UpdateBalancesLeafs map[string]*BalanceNode
	UpdateStorageLeafs map[string]*StorageNode
	UpdateAccount bool
	LeafHash *big.Int
}

type StorageNode struct {
	Cancelled string `json:"cancelled"`
	Data string `json:"data"`
	Forward int `json:"forward"`
	GasFee string `json:"gasFee"`
	StorageID string `json:"storageID"`
	TokenBID string `json:"tokenBID"`
	TokenSID string `json:"tokenSID"`
}

type BalanceNode struct {
	Balance string `json:"balance"`
}

type ChildTree struct {
	Db struct{
		Kv map[string][]string `json:"kv"`
	} `json:"_db"`
	Depth int `json:"_depth"`
	Hasher struct {NumChildren int `json:"_num_children"`} `json:"_hasher"`
	NumChildren int `json:"_num_children"`
	Root string `json:"_root"`
}

func BalanceNodeToPoseidonParam(node *BalanceNode) (poseidonParam []*big.Int) {
	poseidonParam = make([]*big.Int, 0)
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.Balance))
	return poseidonParam
}

func StorageNodeToPoseidonParam(node *StorageNode) (poseidonParam []*big.Int) {
	poseidonParam = make([]*big.Int, 0)
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.TokenSID))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.TokenBID))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.Data))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.StorageID))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.GasFee))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.Cancelled))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.Forward)))
	return poseidonParam
}

func AccountNodeToPoseidonParam(node *AccountValue) (poseidonParam []*big.Int) {
	poseidonParam = make([]*big.Int, 0)
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.Owner))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.PublicKeyX))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.PublicKeyY))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.AppKeyPublicKeyX))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.AppKeyPublicKeyY))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.Nonce)))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.DisableAppKeySpotTrade)))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.DisableAppKeyWithdraw)))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.DisableAppKeyTransferToOther)))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.BalancesTree.Root))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.StorageTree.Root))
	return poseidonParam
}

func AccountAssetNodeToPoseidonParam(node *AccountValue) (poseidonParam []*big.Int) {
	poseidonParam = make([]*big.Int, 0)
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.Owner))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.PublicKeyX))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.PublicKeyY))
	poseidonParam = append(poseidonParam, big.NewInt(int64(node.Nonce)))
	poseidonParam = append(poseidonParam, util.NewBigIntFromStr(node.BalancesTree.Root))
	return poseidonParam
}
