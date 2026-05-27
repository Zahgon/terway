package main

import (
	"context"

	"github.com/spf13/cobra"

	aliClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
)

var (
	deleteIDs    string
	deleteBatch  bool
	deleteStatus string
	deleteNodeID string
	deleteYes    bool
	deleteLimit  int
)

// eniDeleteCmd represents the delete command for ENI
var eniDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ENI network interfaces",
	Long: `Delete specified ENI network interfaces.

Examples:
  # Delete specific EFLO ENIs by ID
  terway-cli eni delete --ids leni-xxx,leni-yyy --region cn-hangzhou --yes

  # Delete specific ECS ENIs by ID
  terway-cli eni delete --link-type ecs --ids eni-xxx,eni-yyy --region cn-hangzhou --yes

  # Batch delete ECS ENIs with specific status on a node
  terway-cli eni delete --link-type ecs --batch --status Available --node-id node-xxx --region cn-hangzhou --yes`,
	RunE: runEniDelete,
}

func init() {
	eniDeleteCmd.Flags().StringVar(&deleteIDs, "ids", "", "Comma-separated list of ENI IDs to delete (mutually exclusive with --batch)")
	eniDeleteCmd.Flags().BoolVar(&deleteBatch, "batch", false, "Batch delete mode - delete all matching abnormal ENIs")
	eniDeleteCmd.Flags().StringVar(&deleteStatus, "status", "", "Filter by status (only used with --batch)")
	eniDeleteCmd.Flags().StringVar(&deleteNodeID, "node-id", "", "Filter by node ID (only used with --batch)")
	eniDeleteCmd.Flags().BoolVar(&deleteYes, "yes", false, "Skip confirmation prompt")
	eniDeleteCmd.Flags().IntVar(&deleteLimit, "limit", 50, "Maximum number of ENIs to delete per batch (max 50)")

	eniCmd.AddCommand(eniDeleteCmd)
}

func runEniDelete(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	// Validate arguments
	return nil
}

// In batch mode, must specify either --status or --node-id as filter

// Get the list of ENI IDs to delete

// Validate ENI IDs for EFLO mode

// Batch mode: query ENIs matching criteria

// Apply limit to the number of ENIs to delete

// Show the ENIs to be deleted

// Confirm deletion unless --yes is specified

// Delete ENIs in batches using appropriate client

func queryENIsForBatchDelete(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use ECS client for regular ENIs

// Use EFLO client for LENI/HDENI

// Always use raw status for accurate filtering for EFLO

// If no specific status is provided, filter for abnormal ones

func showDeleteConfirmation(eniIDs []string) error { _ = "STUB: not implemented"; return nil }

// Show first 10 ENIs

// eniDeleter interface for deleting ENIs
type eniDeleter interface {
	deleteENI(ctx context.Context, eniID string) error
}

// efloDeleter wraps EFLO client to implement eniDeleter
type efloDeleter struct {
	client aliClient.EFLO
}

func (e *efloDeleter) deleteENI(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

// ecsDeleter wraps ECS client to implement eniDeleter
type ecsDeleter struct {
	client aliClient.ECS
}

func (e *ecsDeleter) deleteENI(ctx context.Context, eniID string) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteENIsInBatches(ctx context.Context, deleter eniDeleter, eniIDs []string, batchSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// Process in batches

// Delete each ENI in the batch concurrently

// Continue processing other ENIs even if one fails

// Update spinner text

// Show final results
