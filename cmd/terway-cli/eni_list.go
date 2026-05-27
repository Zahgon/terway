package main

import (
	"github.com/spf13/cobra"
)

var (
	listStatus  string
	listNodeID  string
	listVpcID   string
	listShowAll bool
)

// eniListCmd represents the list command for ENI
var eniListCmd = &cobra.Command{
	Use:   "list",
	Short: "List ENI network interfaces",
	Long:  `Query and display ENI network interfaces. By default, only abnormal interfaces are shown.`,
	RunE:  runEniList,
}

func init() {
	eniListCmd.Flags().StringVar(&listStatus, "status", "", "Filter by status (optional)")
	eniListCmd.Flags().StringVar(&listNodeID, "node-id", "", "Filter by node ID (optional)")
	eniListCmd.Flags().StringVar(&listVpcID, "vpc-id", "", "Filter by VPC ID (optional)")
	eniListCmd.Flags().BoolVar(&listShowAll, "show-all", false, "Show all network interfaces (default: only show abnormal)")

	eniCmd.AddCommand(eniListCmd)
}

func runEniList(cmd *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

// Build query options

// Use ECS client for regular ENIs

// Use EFLO client for LENI/HDENI (default)

// Always use raw status for accurate filtering for EFLO

// Filter to only abnormal statuses if not showing all

// Display results in a table

// Format tags as key=value pairs separated by commas

// Colorize abnormal statuses

// Green for normal/healthy status
