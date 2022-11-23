package statemanager

import (
	"math"
	"math/big"
	"staterecovery/common/util"
	"staterecovery/crypto/poseidon"
	"strconv"

)

func poseidonHash(children []*big.Int, poseidonParam int) (hash *big.Int) {
	hash, err := poseidon.HashAccurate(children, poseidonParam)
	if err != nil {
		panic("in calculateHash HashAccurate tree error:" + err.Error())
	}
	return hash
}

// key: leaf index
// value: leaf hash
func updateTree(rootBefore string, tree map[string][]string, depth int, key, value string) (rootAfter string) {
	parentRoot := rootBefore
	branchNodes := make([][]string, depth)
	leafsNum := int64(math.Pow(float64(treeNodeChildrenNum), float64(depth - 1)))
	path, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		panic("in updateTree strconv.ParseInt key error:" + err.Error())
	}
	// record origin branch nodes, from root to leaf
	for i := 0; i < depth; i++ {
		children := tree[parentRoot]
		branchNodes[i] = children
		branchIndex := int((path / leafsNum) % int64(treeNodeChildrenNum))
		parentRoot = children[branchIndex]
		path *= int64(treeNodeChildrenNum)
	}
	// calculate new branch nodes, from lead to root
	path, _ = strconv.ParseInt(key, 10, 64)
	calculateHash := value
	for i := 0; i < depth; i++ {
		branchIndex := int(path % int64(treeNodeChildrenNum))
		children := make([]string, 4)
		for j := 0; j < treeNodeChildrenNum; j++ {
			if j == branchIndex {
				children[j] = calculateHash
			} else {
				children[j] = branchNodes[depth - 1 - i][j]
			}
		}
		newCalculateHash := poseidonHash(util.ConvertStringArrayToBigIntArray(children), TreeNodeCalculateParam)
		tree[newCalculateHash.String()] = children
		path = path / int64(treeNodeChildrenNum)
		calculateHash = newCalculateHash.String()
	}
	rootAfter = calculateHash
	return rootAfter
}

func CreateProof(root string, tree map[string][]string, depth int, key string) (branchNodes []*big.Int) {
	parentRoot := root
	branchNodesTemp := make([][]*big.Int, depth)
	leafsNum := int64(math.Pow(float64(treeNodeChildrenNum), float64(depth - 1)))
	path, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		panic("in CreateProof strconv.ParseInt key error:" + err.Error())
	}
	for i := 0; i < depth; i++ {
		children := tree[parentRoot]
		branchIndex := int((path / leafsNum) % int64(treeNodeChildrenNum))
		parentRoot = children[branchIndex]
		for index, child := range children {
			if index != branchIndex {
				branchNodesTemp[depth - i - 1] = append(branchNodesTemp[depth - i - 1], util.NewBigIntFromStr(child))
			}
		}
		path *= int64(treeNodeChildrenNum)
	}
	branchNodes = make([]*big.Int, 0)
	for _, tmp := range branchNodesTemp {
		branchNodes = append(branchNodes, tmp...)
	}
	return branchNodes
}
