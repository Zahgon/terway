package main

import (
	"github.com/AliyunContainerService/terway/rpc"

	"github.com/pterm/pterm"
)

type tree struct {
	Key    string
	Value  string // value only exists in leaf
	IsLeaf bool
	Leaves []*tree
}

// add a leaf to the tree
func (t *tree) addLeaf(path []string, value string) {
	_ = "STUB: not implemented"
	// noop
	return
}

// is leaf

func (t *tree) leveledList(list pterm.LeveledList, level int) pterm.LeveledList {
	_ = "STUB: not implemented"
	return *new(pterm.LeveledList)
}

func printPTermTree(m []*rpc.MapKeyValueEntry) error {
	_ = "STUB: not implemented"
	// build a tree
	return nil
}
