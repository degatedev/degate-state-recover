package entity

import (
	"math/big"
	"staterecovery/common/util"

	"github.com/ethereum/go-ethereum/common"
)

var (
	WithdrawModeAbi = `[{"inputs":[{"components":[{"components":[{"internalType":"uint32","name":"accountID","type":"uint32"},{"internalType":"address","name":"owner","type":"address"},{"internalType":"uint256","name":"pubKeyX","type":"uint256"},{"internalType":"uint256","name":"pubKeyY","type":"uint256"},{"internalType":"uint32","name":"nonce","type":"uint32"}],"internalType":"struct ExchangeData.AccountLeaf","name":"accountLeaf","type":"tuple"},{"components":[{"internalType":"uint32","name":"tokenID","type":"uint32"},{"internalType":"uint248","name":"balance","type":"uint248"}],"internalType":"struct ExchangeData.BalanceLeaf","name":"balanceLeaf","type":"tuple"},{"internalType":"uint256[48]","name":"accountMerkleProof","type":"uint256[48]"},{"internalType":"uint256[48]","name":"balanceMerkleProof","type":"uint256[48]"}],"internalType":"struct ExchangeData.MerkleProof","name":"merkleProof","type":"tuple"}],"name":"withdrawFromMerkleTree","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	WithdrawModeMethod = "withdrawFromMerkleTree"
)

type AccountLeafWithdrawMode struct {
	AccountID uint32 `json:"accountID"`
	Nonce uint32 `json:"nonce"`
	Owner common.Address `json:"owner"`
	PubKeyX *big.Int `json:"pubKeyX"`
	PubKeyY *big.Int `json:"pubKeyY"`
}

type BalanceLeafWithdrawMode struct {
	TokenID uint32 `json:"tokenID"`
	Balance *big.Int `json:"balance"`
}

type WithdrawModeMerkleProof struct {
	AccountLeaf *AccountLeafWithdrawMode `json:"accountLeaf"`
	BalanceLeaf *BalanceLeafWithdrawMode `json:"balanceLeaf"`
	AccountMerkleProof []*big.Int `json:"accountMerkleProof"`
	BalanceMerkleProof []*big.Int `json:"balanceMerkleProof"`
}

func GetAccountLeafForWithdrawMode(accountValue *AccountValue, accountID string) *AccountLeafWithdrawMode {
	accountLeaf := &AccountLeafWithdrawMode{
		AccountID: uint32(util.NewDecimalFromString(accountID).IntPart()),
		Nonce: uint32(accountValue.Nonce),
		Owner: common.HexToAddress(util.IntStrToHex(accountValue.Owner)),
		PubKeyX: util.NewBigIntFromString(accountValue.PublicKeyX),
		PubKeyY: util.NewBigIntFromString(accountValue.PublicKeyY),
	}
	return accountLeaf
}

func GetBalanceLeafForWithdrawMode(balanceNode *BalanceNode, tokenID string) *BalanceLeafWithdrawMode {
	if balanceNode == nil {
		return &BalanceLeafWithdrawMode{
			TokenID: uint32(util.NewDecimalFromString(tokenID).IntPart()),
			Balance: util.NewBigIntFromString("0"),
		}
	}
	balanceLeaf := &BalanceLeafWithdrawMode{
		TokenID: uint32(util.NewDecimalFromString(tokenID).IntPart()),
		Balance: util.NewBigIntFromString(balanceNode.Balance),
	}
	return balanceLeaf
}