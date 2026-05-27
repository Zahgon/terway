package main

import (
	"os"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/credentials/provider"
	"github.com/spf13/cobra"

	aliClient "github.com/AliyunContainerService/terway/pkg/aliyun/client"
)

var (
	regionID string
	linkType string
)

const (
	LinkTypeEFLO = "eflo"
	LinkTypeECS  = "ecs"
)

// eniCmd is the root command for ENI management (supports both EFLO/LENI/HDENI and ECS ENIs)
var eniCmd = &cobra.Command{
	Use:   "eni",
	Short: "Manage ENI network interfaces (EFLO/LENI/HDENI and ECS)",
	Long:  `Query and clean up ENI network interfaces. Supports both EFLO (LENI/HDENI) and regular ECS ENIs.`,
}

func init() {
	eniCmd.PersistentFlags().StringVar(&regionID, "region", os.Getenv("REGION_ID"), "Alibaba Cloud region ID (can also be set via REGION_ID env)")
	eniCmd.PersistentFlags().StringVar(&linkType, "link-type", LinkTypeECS, "Link type: 'eflo' for LENI/HDENI, 'ecs' for regular ECS ENIs")
}

// validateLinkType validates the link type flag
func validateLinkType() error { _ = "STUB: not implemented"; return nil }

// isValidEFLOENIID checks if an ENI ID has valid EFLO prefix (hdeni- or leni-)
func isValidEFLOENIID(eniID string) bool { _ = "STUB: not implemented"; return false }

// validateENIIDsForLinkType validates ENI IDs based on link type
func validateENIIDsForLinkType(eniIDs []string, lType string) error {
	_ = "STUB: not implemented"
	return nil
}

// getCredentialProvider creates a credential provider chain for Alibaba Cloud
func getCredentialProvider() provider.CredentialsProvider {
	_ = "STUB: not implemented"
	return *new(provider.CredentialsProvider)
}

// getEFLOClient initializes the EFLO client with the given region
func getEFLOClient(region string) (aliClient.EFLO, error) {
	_ = "STUB: not implemented"
	return *new(aliClient.EFLO), nil
}

// Use default rate limiter configuration

// Use noop tracer for CLI

// getECSClient initializes the ECS client with the given region
func getECSClient(region string) (aliClient.ECS, error) {
	_ = "STUB: not implemented"
	return *new(aliClient.ECS), nil
}

// Use default rate limiter configuration

// Use noop tracer for CLI

// getEFLOControlClient initializes the EFLO Control client with the given region
func getEFLOControlClient(region string) (aliClient.EFLOControl, error) {
	_ = "STUB: not implemented"
	return *new(aliClient.EFLOControl), nil
}

// isAbnormalStatus checks if an ENI status indicates an abnormal state
// For ENI (ECS): Available status means not attached, which is abnormal
// For ENO (LENI): Unattached status means not attached, which is abnormal
func isAbnormalStatus(status string) bool { _ = "STUB: not implemented"; return false }

// isAbnormalStatusForLinkType checks if an ENI status is abnormal based on link type
func isAbnormalStatusForLinkType(status, lType string) bool {
	_ = "STUB: not implemented"
	return false

	// For ECS ENIs, "Available" means not attached to any instance
}

// Any status other than InUse is considered abnormal for ENI

// For EFLO (LENI/HDENI), use the standard abnormal status check

// parseENIIDs parses a comma-separated list of ENI IDs
func parseENIIDs(ids string) []string { _ = "STUB: not implemented"; return nil }
