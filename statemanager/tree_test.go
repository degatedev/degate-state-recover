package statemanager

import (
	"log"
	"staterecovery/common/entity"
	"testing"
)

func TestBalanceTree(t *testing.T) {
	defaultBalanceTree, _ = calculateDefaultBalanceTree()
	leafHash := poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "3963532413760000000"}), BalanceNodeHashCalculateParam)
	rootAfter := updateTree(defaultBalanceRoot.String(), defaultBalanceTree, 16, "0", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "87351461824260000000000"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "2", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "6188834770000000000000"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "3", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "299640000000000000000"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "32", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "8021935040000000000"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "4", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "29512504500000000000000"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "5", leafHash.String())
	leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "0"}), BalanceNodeHashCalculateParam)
	rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "1", leafHash.String())
	log.Println("rootAfter:", rootAfter)

}

func TestBalanceTree2(t *testing.T) {
	defaultBalanceTree, _ = calculateDefaultBalanceTree()
	leafHash := poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "100000000000000000"}), BalanceNodeHashCalculateParam)
	rootAfter := updateTree(defaultBalanceRoot.String(), defaultBalanceTree, 16, "0", leafHash.String())
	// leafHash = poseidonHash(entity.BalanceNodeToPoseidonParam(&entity.BalanceNode{Balance: "1000000000000000000000000"}), BalanceNodeHashCalculateParam)
	// rootAfter = updateTree(rootAfter, defaultBalanceTree, 16, "2", leafHash.String())
	log.Println("rootAfter:", rootAfter)

}
