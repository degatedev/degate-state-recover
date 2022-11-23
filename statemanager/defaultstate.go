package statemanager

import (
	"log"
	"math/big"
	"staterecovery/common/entity"

	"github.com/shopspring/decimal"
)

var (
	defaultBalanceRoot, _ = decimal.NewFromString("3626386379762139238426088069940068312069344602207393459612601721558984385997")
	defaultStorageRoot, _ = decimal.NewFromString("17168846436385410234776549269474130900971613041027057153527920776001261983060")
	defaultAccountRoot, _ = decimal.NewFromString("1755311117727461112937066252003540264424472859778551426333315695520434999065")
	defaultAccountAssetRoot, _ = decimal.NewFromString("3216621562491977239625612062438587439774929574181340738308412865392963758824")

	defaultBalanceTree map[string][]string
	defaultStorageTree map[string][]string

	defaultAccountNode = []*big.Int {
		// poseidon([int(self.owner), int(self.publicKeyX), int(self.publicKeyY), int(self.appKeyPublicKeyX), int(self.appKeyPublicKeyY), int(self.nonce), int(self.disableAppKeySpotTrade), int(self.disableAppKeyWithdraw), int(self.disableAppKeyTransferToOther), int(self._balancesTree._root), int(self._storageTree._root)], poseidonParamsAccount)
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		defaultBalanceRoot.BigInt(),
		defaultStorageRoot.BigInt(),
	}
	defaultAccountAssetNode = []*big.Int {
		// poseidon([int(self.owner), int(self.publicKeyX), int(self.publicKeyY), int(self.nonce), int(self._balancesTree._root)], poseidonParamsAccountAsset)
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		defaultBalanceRoot.BigInt(),
	}
	defaultBalanceNode = []*big.Int {
		big.NewInt(0),
	}
	defaultStorageNode = []*big.Int {
		// poseidon([int(self.tokenSID), int(self.tokenBID), int(self.data), int(self.storageID), int(self.gasFee), int(self.cancelled), int(self.forward)], poseidonParamsStorage)
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	}

)

func GetDefaultBalanceRoot() string {
	return defaultBalanceRoot.String()
}

func calculateDefaultAccountTree(state *entity.State) {
	state.AccountsTree, state.AccountsRoot = calculateTree(defaultAccountNode, AccountNodeHashCalculateParam, accountDepth)
	log.Println("account root:", state.AccountsRoot)
}

func calculateDefaultAccountAssetTree(state *entity.State) {
	state.AccountsAssetTree, state.AccountsAssetRoot = calculateTree(defaultAccountAssetNode, AccountAssetNodeHashCalculateParam, accountDepth)
	log.Println("account asset root:", state.AccountsAssetRoot)
}

func calculateDefaultBalanceTree() (tree map[string][]string, root string) {
	tree, root = calculateTree(defaultBalanceNode, BalanceNodeHashCalculateParam, balanceDepth)
	if root != defaultBalanceRoot.String() {
		panic("in calculateDefaultBalanceTree balance root not match")
	}
	return tree, root
}

func calculateDefaultStorageTree() (tree map[string][]string, root string) {
	tree, root = calculateTree(defaultStorageNode, StorageNodeHashCalculateParam, storageDepth)
	if root != defaultStorageRoot.String() {
		panic("in calculateDefaultStorageTree storage root not match")
	}
	return tree, root
}

func getDefaultAccountValue() (account *entity.AccountValue) {
	account = &entity.AccountValue{}
	account.BalancesLeafs = make(map[string]*entity.BalanceNode)
	// account.BalancesTree = getDefaultChildTree(defaultBalanceTree, balanceDepth, treeNodeChildrenNum, defaultBalanceRoot.String())
	// account.StorageLeafs = make(map[string]*entity.StorageNode)
	// account.StorageTree = getDefaultChildTree(defaultStorageTree, storageDepth, treeNodeChildrenNum, defaultStorageRoot.String())
	// account.AppKeyPublicKeyX = "0"
	// account.AppKeyPublicKeyY = "0"
	account.PublicKeyX = "0"
	account.PublicKeyY = "0"
	// account.DisableAppKeySpotTrade = 0
	// account.DisableAppKeyTransferToOther = 0
	// account.DisableAppKeyWithdraw = 0
	account.Nonce = 0
	account.Owner = "0"

	account.UpdateBalancesLeafs = make(map[string]*entity.BalanceNode)
	account.UpdateStorageLeafs = make(map[string]*entity.StorageNode)
	account.UpdateAccount = false
	return account
}

func getDefaultChildTree(tree map[string][]string, depth, numChildren int, treeRoot string) (childTree *entity.ChildTree) {
	childTree = &entity.ChildTree{}
	childTree.Db.Kv = copyDefaultTree(tree)
	childTree.Depth = depth
	childTree.Hasher.NumChildren = numChildren
	childTree.NumChildren = numChildren
	childTree.Root = treeRoot
	return childTree
}

func getDefaultBalanceLeaf() (leaf *entity.BalanceNode) {
	leaf = &entity.BalanceNode{}
	leaf.Balance = "0"
	return leaf
}

func getDefaultStorageLeaf() (leaf *entity.StorageNode) {
	leaf = &entity.StorageNode{}
	leaf.TokenSID = "0"
	leaf.TokenBID = "0"
	leaf.StorageID = "0"
	leaf.Data = "0"
	leaf.GasFee = "0"
	leaf.Cancelled = "0"
	leaf.Forward = 1
	return leaf
}

func copyDefaultTree(tree map[string][]string) (copiedTree map[string][]string) {
	if tree == nil || len(tree) == 0 {
		panic("in copyDefaultTree tree is empty")
	}
	copiedTree = make(map[string][]string)
	for key, value := range tree {
		copiedTree[key] = value
	}
	return copiedTree
}

func calculateTree(nodeParam []*big.Int, poseidonParam, depth int) (tree map[string][]string, root string) {
	// calculate node
	hashValue := poseidonHash(nodeParam, poseidonParam)
	// calculate node
	var newHashValue *big.Int
	tree = make(map[string][]string)
	for i := 0; i < depth; i++ {
		children := []*big.Int {
			hashValue, 
			hashValue, 
			hashValue, 
			hashValue,
		}
		newHashValue = poseidonHash(children, TreeNodeCalculateParam)
		tree[newHashValue.String()] = [] string {hashValue.String(), hashValue.String(), hashValue.String(), hashValue.String()}
		hashValue = newHashValue
	}
	return tree, newHashValue.String()
}