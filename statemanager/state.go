package statemanager

import (
	"staterecovery/common/entity"
	"staterecovery/common/util"
	"sync"
)

var (
	StorageNodeHashCalculateParam = 8
	BalanceNodeHashCalculateParam = 5
	AccountNodeHashCalculateParam = 12
	AccountAssetNodeHashCalculateParam = 6
	TreeNodeCalculateParam = 5

	accountDepth = 16
	balanceDepth = 16
	storageDepth = 7

	treeNodeChildrenNum = 4
)

func init() {
	defaultBalanceTree, _ = calculateDefaultBalanceTree()
	defaultStorageTree, _ = calculateDefaultStorageTree()
}

func CreateEmptyState() (state *entity.State) {
	state = &entity.State{}
	state.BlockId = 0
	// // default accounts_tree and root
	// calculateDefaultAccountTree(state)
	// default accounts_asset_tree and root
	// calculateDefaultAccountAssetTree(state)
	// default account value
	state.AccountsValues = make(map[string]*entity.AccountValue)
	state.ExchangeID = 0
	return state
}

func GetAccount(state *entity.State, accountID string) (account *entity.AccountUpdate) {
	accountValue := state.AccountsValues[accountID]
	if accountValue == nil {
		// create empty account
		accountValue = getDefaultAccountValue()
		state.AccountsValues[accountID] = accountValue
	}
	return &entity.AccountUpdate{
		AccountID: accountID,
		AccountValue: accountValue,
	}
}

func GetAccountByOwnerAndNonce(state *entity.State, owner string, nonce uint64) (account *entity.AccountUpdate) {
	for accountID, accountValue := range state.AccountsValues {
		if accountValue.Owner == owner && accountValue.Nonce == nonce {
			return &entity.AccountUpdate{
				AccountID: accountID,
				AccountValue: accountValue,
			}
		}
	}
	return nil
}

func AddBalance(accountUpdate *entity.AccountUpdate, amount, tokenID string) {
	accountUpdate.AccountValue.UpdateAccount = true
	updateLeaf := getAndNewBalanceLeaf(accountUpdate.AccountValue, tokenID)
	amountBefore := util.NewDecimalFromString(updateLeaf.Balance)
	amountAdd := util.NewDecimalFromString(amount)
	amountAfter := amountBefore.Add(amountAdd)
	updateLeaf.Balance = amountAfter.String()
}

func SubBalance(accountUpdate *entity.AccountUpdate, amount, tokenID string) {
	accountUpdate.AccountValue.UpdateAccount = true
	updateLeaf := getAndNewBalanceLeaf(accountUpdate.AccountValue, tokenID)
	amountBefore := util.NewDecimalFromString(updateLeaf.Balance)
	amountSub := util.NewDecimalFromString(amount)
	amountAfter := amountBefore.Sub(amountSub)
	updateLeaf.Balance = amountAfter.String()
}

func UpdateAccountAsset(accountUpdate *entity.AccountUpdate) {
	accountUpdate.AccountValue.UpdateAccount = true
}

func UpdateState(accountUpdateList []*entity.AccountUpdate, state *entity.State) {
	// calculate balance and storage
	for _, accountUpdate := range accountUpdateList {
		updateAccount(accountUpdate.AccountID, accountUpdate.AccountValue, state)
	}
}

func CalculateStateTree(state *entity.State) {
	calculateDefaultAccountAssetTree(state)
	// for accountID, accountValue := range state.AccountsValues {
	// 	// calculate balance
	// 	accountValue.BalancesTree = getDefaultChildTree(defaultBalanceTree, balanceDepth, treeNodeChildrenNum, defaultBalanceRoot.String())
	// 	for key, leaf := range accountValue.BalancesLeafs {
	// 		leafHash := poseidonHash(entity.BalanceNodeToPoseidonParam(leaf), BalanceNodeHashCalculateParam)
	// 		accountValue.BalancesTree.Root = updateTree(accountValue.BalancesTree.Root, accountValue.BalancesTree.Db.Kv, 16, key, leafHash.String())
	// 	}
	// 	// calculate account
	// 	leafAccountAssetHash := poseidonHash(entity.AccountAssetNodeToPoseidonParam(accountValue), AccountAssetNodeHashCalculateParam)
	// 	state.AccountsAssetRoot = updateTree(state.AccountsAssetRoot, state.AccountsAssetTree, accountDepth, accountID, leafAccountAssetHash.String())
	// }

	var wg sync.WaitGroup
	for _, accountValue := range state.AccountsValues {
		// calculate balance
		wg.Add(1)
		go func (accountValue *entity.AccountValue)  {
			accountValue.BalancesTree = getDefaultChildTree(defaultBalanceTree, balanceDepth, treeNodeChildrenNum, defaultBalanceRoot.String())
			for key, leaf := range accountValue.BalancesLeafs {
				leafHash := poseidonHash(entity.BalanceNodeToPoseidonParam(leaf), BalanceNodeHashCalculateParam)
				accountValue.BalancesTree.Root = updateTree(accountValue.BalancesTree.Root, accountValue.BalancesTree.Db.Kv, 16, key, leafHash.String())
			}
			accountValue.LeafHash = poseidonHash(entity.AccountAssetNodeToPoseidonParam(accountValue), AccountAssetNodeHashCalculateParam)
			wg.Done()
		} (accountValue)
	}
	wg.Wait()
	for accountID, accountValue := range state.AccountsValues {
		// calculate account
		state.AccountsAssetRoot = updateTree(state.AccountsAssetRoot, state.AccountsAssetTree, accountDepth, accountID, accountValue.LeafHash.String())
	}
}

func GetAllAccountIDByOwner(state *entity.State, owner string) (accountIDList []string) {
	accountIDList = make([]string, 0)
	for accountID, accountValue := range state.AccountsValues {
		if accountValue.Owner != "0" && accountValue.Owner == owner {
			accountIDList = append(accountIDList, accountID)
		}
	}
	return accountIDList
}

func CreateWithdrawModeMerkleProof(state *entity.State, accountID, tokenID string) (withdrawModeMerkleProof *entity.WithdrawModeMerkleProof) {
	accountValue := state.AccountsValues[accountID]
	if accountValue == nil {
		return nil
	}
	withdrawModeMerkleProof = &entity.WithdrawModeMerkleProof {
		AccountLeaf: entity.GetAccountLeafForWithdrawMode(accountValue, accountID),
		BalanceLeaf: entity.GetBalanceLeafForWithdrawMode(accountValue.BalancesLeafs[tokenID], tokenID),
		AccountMerkleProof: CreateProof(state.AccountsAssetRoot, state.AccountsAssetTree, accountDepth, accountID),
		BalanceMerkleProof: CreateProof(accountValue.BalancesTree.Root, accountValue.BalancesTree.Db.Kv, balanceDepth, tokenID),
	}
	return withdrawModeMerkleProof
}

func updateAccount(accountID string, updateAccount *entity.AccountValue, state *entity.State) {
	// balance update
	// log.Println("updateAccount.UpdateBalancesLeafs length:", strconv.Itoa(len(updateAccount.UpdateBalancesLeafs)))
	// if updateAccount.UpdateBalancesLeafs != nil && len(updateAccount.UpdateBalancesLeafs) > 0 {
	// 	for key, leaf := range updateAccount.UpdateBalancesLeafs {
	// 		log.Println("in loop leaf:", leaf.Balance)
	// 		leafHash := poseidonHash(entity.BalanceNodeToPoseidonParam(leaf), BalanceNodeHashCalculateParam)
	// 		updateAccount.BalancesTree.Root = updateTree(updateAccount.BalancesTree.Root, updateAccount.BalancesTree.Db.Kv, balanceDepth, key, leafHash.String())
	// 		log.Println("in updateAccount updateAccount.BalancesTree.Root:", updateAccount.BalancesTree.Root)
	// 	}
	// 	updateAccount.UpdateBalancesLeafs = make(map[string]*entity.BalanceNode)
	// }
	// storage update
	// if updateAccount.UpdateStorageLeafs != nil && len(updateAccount.UpdateStorageLeafs) > 0 {
	// 	for key, leaf := range updateAccount.UpdateStorageLeafs {
	// 		leafHash := poseidonHash(entity.StorageNodeToPoseidonParam(leaf), StorageNodeHashCalculateParam)
	// 		updateAccount.StorageTree.Root = updateTree(updateAccount.StorageTree.Root, updateAccount.StorageTree.Db.Kv, storageDepth, key, leafHash.String())
	// 	}
	// 	updateAccount.UpdateStorageLeafs = make(map[string]*entity.StorageNode)
	// }
	// account update
	// leafAccountHash := poseidonHash(entity.AccountNodeToPoseidonParam(updateAccount), AccountNodeHashCalculateParam)
	// state.AccountsRoot = updateTree(state.AccountsRoot, state.AccountsTree, accountDepth, accountID, leafAccountHash.String())

	// leafAccountAssetHash := poseidonHash(entity.AccountAssetNodeToPoseidonParam(updateAccount), AccountAssetNodeHashCalculateParam)
	// log.Println("state.AccountsAssetRoot:", state.AccountsAssetRoot)
	// state.AccountsAssetRoot = updateTree(state.AccountsAssetRoot, state.AccountsAssetTree, accountDepth, accountID, leafAccountAssetHash.String())

	updateAccount.UpdateBalancesLeafs = make(map[string]*entity.BalanceNode)
}

func getAndNewBalanceLeaf(accountValue *entity.AccountValue, tokenID string) (updateLeaf *entity.BalanceNode) {
	updateLeaf = getBalanceLeaf(accountValue.UpdateBalancesLeafs, tokenID)
	if updateLeaf == nil {
		// get updateLeaf from balanceLeafs
		updateLeaf = getBalanceLeaf(accountValue.BalancesLeafs, tokenID)
		if updateLeaf == nil {
			// create empty leaf
			updateLeaf = getDefaultBalanceLeaf()
			accountValue.BalancesLeafs[tokenID] = updateLeaf
			accountValue.UpdateBalancesLeafs[tokenID] = updateLeaf
		} else {
			accountValue.UpdateBalancesLeafs[tokenID] = updateLeaf
		}
	}
	return updateLeaf
}

func getAndNewStorageLeaf(accountValue *entity.AccountValue, storageID string) (updateLeaf *entity.StorageNode) {
	updateLeaf = getStorageLeaf(accountValue.UpdateStorageLeafs, storageID)
	if updateLeaf == nil {
		// get updateLeaf from balanceLeafs
		updateLeaf = getStorageLeaf(accountValue.StorageLeafs, storageID)
		if updateLeaf == nil {
			// create empty leaf
			updateLeaf = getDefaultStorageLeaf()
			accountValue.StorageLeafs[storageID] = updateLeaf
			accountValue.UpdateStorageLeafs[storageID] = updateLeaf
		}
	}
	return updateLeaf
}

func getBalanceLeaf(leafs map[string]*entity.BalanceNode, tokenId string) (leaf *entity.BalanceNode) {
	if leafs == nil || len(leafs) == 0 {
		return nil
	}
	for tempKey, tempLeaf := range leafs {
		if tempKey == tokenId {
			return tempLeaf
		}
	}
	return nil
}

func getStorageLeaf(leafs map[string]*entity.StorageNode, storageID string) (leaf *entity.StorageNode) {
	if leafs == nil || len(leafs) == 0 {
		return nil
	}
	for tempKey, tempLeaf := range leafs {
		if tempKey == storageID {
			return tempLeaf
		}
	}
	return nil
}