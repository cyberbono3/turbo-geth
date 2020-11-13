package cli

import (
	"github.com/ledgerwatch/turbo-geth/cmd/utils"

	"github.com/urfave/cli"
)

// DefaultFlags contains all flags that are used and supported by turbo-geth binary.
var DefaultFlags = []cli.Flag{
	utils.DataDirFlag,
	utils.KeyStoreDirFlag,
	utils.EthashDatasetDirFlag,
	utils.TxPoolLocalsFlag,
	utils.TxPoolNoLocalsFlag,
	utils.TxPoolJournalFlag,
	utils.TxPoolRejournalFlag,
	utils.TxPoolPriceLimitFlag,
	utils.TxPoolPriceBumpFlag,
	utils.TxPoolAccountSlotsFlag,
	utils.TxPoolGlobalSlotsFlag,
	utils.TxPoolAccountQueueFlag,
	utils.TxPoolGlobalQueueFlag,
	utils.TxPoolLifetimeFlag,
	utils.TxLookupLimitFlag,
	StorageModeFlag,
	SnapshotModeFlag,
	SeedSnapshotsFlag,
	ExternalSnapshotDownloaderAddrFlag,
	BatchSizeFlag,
	DatabaseFlag,
	PrivateApiAddr,
	EtlBufferSizeFlag,
	LMDBMapSizeFlag,
	LMDBMaxFreelistReuseFlag,
	TLSFlag,
	TLSCertFlag,
	TLSKeyFlag,
	TLSCACertFlag,
	utils.ListenPortFlag,
	utils.NATFlag,
	utils.NoDiscoverFlag,
	utils.DiscoveryV5Flag,
	utils.NetrestrictFlag,
	utils.NodeKeyFileFlag,
	utils.NodeKeyHexFlag,
	utils.DNSDiscoveryFlag,
	utils.RopstenFlag,
	utils.RinkebyFlag,
	utils.GoerliFlag,
	utils.YoloV1Flag,
	utils.VMEnableDebugFlag,
	utils.NetworkIdFlag,
	utils.FakePoWFlag,
	utils.GpoBlocksFlag,
	utils.GpoPercentileFlag,
	utils.EWASMInterpreterFlag,
	utils.EVMInterpreterFlag,
	utils.InsecureUnlockAllowedFlag,
	utils.MetricsEnabledFlag,
	utils.MetricsEnabledExpensiveFlag,
	utils.MetricsHTTPFlag,
	utils.MetricsPortFlag,
	utils.IdentityFlag,
	SilkwormFlag,
}
