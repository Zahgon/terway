package main

import (
	"github.com/AliyunContainerService/terway/rpc"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func runList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// list types

// list resources

func runShow(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// todo: 去锁
	return nil
}

// only type, select the first resource returned

func runMapping(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func renderPrefixSection(list pterm.LeveledList, title string, prefixes []*rpc.PrefixInfo) pterm.LeveledList {
	_ = "STUB: not implemented"
	return *new(pterm.LeveledList)
}

func runExecute(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// <type> <resource> <command> [args...]
	return nil
}

// print message

const (
	metadataErrorStringIPV6 = "can't get ipv6 info for eni %s, skipping. (%s)"

	metadataLevelVSwitch       = 0
	metadataLevelENI           = 1
	metadataLevelAttribute     = 2
	metadataLevelAttributeItem = 3
)

func runMetadata(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// make primary to the first

// network interface

// ipv4

// when len(ipv4) == 1, only primary ipv4 exists

// ipv6

func printKV(key, value string) string { _ = "STUB: not implemented"; return "" }

// getFirstNameWithType finds the first resource in the given type
func getFirstNameWithType(typ string) (string, error) { _ = "STUB: not implemented"; return "", nil }
