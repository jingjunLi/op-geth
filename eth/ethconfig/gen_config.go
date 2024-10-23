// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ethconfig

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool/blobpool"
	"github.com/ethereum/go-ethereum/core/txpool/bundlepool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/triedb/pathdb"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                                 *core.Genesis `toml:",omitempty"`
		NetworkId                               uint64
		SyncMode                                downloader.SyncMode
		EthDiscoveryURLs                        []string
		SnapDiscoveryURLs                       []string
		NoPruning                               bool
		NoPrefetch                              bool
		TxLookupLimit                           uint64                 `toml:",omitempty"`
		TransactionHistory                      uint64                 `toml:",omitempty"`
		StateHistory                            uint64                 `toml:",omitempty"`
		StateScheme                             string                 `toml:",omitempty"`
		PathNodeBuffer                          pathdb.NodeBufferType  `toml:",omitempty"`
		ProposeBlockInterval                    uint64                 `toml:",omitempty"`
		EnableProofKeeper                       bool                   `toml:",omitempty"`
		KeepProofBlockSpan                      uint64                 `toml:",omitempty"`
		JournalFileEnabled                      bool                   `toml:",omitempty"`
		RequiredBlocks                          map[uint64]common.Hash `toml:"-"`
		LightServ                               int                    `toml:",omitempty"`
		LightIngress                            int                    `toml:",omitempty"`
		LightEgress                             int                    `toml:",omitempty"`
		LightPeers                              int                    `toml:",omitempty"`
		LightNoPrune                            bool                   `toml:",omitempty"`
		LightNoSyncServe                        bool                   `toml:",omitempty"`
		SkipBcVersionCheck                      bool                   `toml:"-"`
		DatabaseHandles                         int                    `toml:"-"`
		DatabaseCache                           int
		BlockDbCachePer                         int `toml:",omitempty"`
		StateDbCachePer                         int `toml:",omitempty"`
		ChaindbCachePer                         int `toml:",omitempty"`
		DatabaseFreezer                         string
		TrieCleanCache                          int
		TrieDirtyCache                          int
		TrieTimeout                             time.Duration
		TrieCommitInterval                      uint64
		SnapshotCache                           int
		Preimages                               bool
		NoTries                                 bool
		FilterLogCacheSize                      int
		Miner                                   miner.Config
		TxPool                                  legacypool.Config
		BlobPool                                blobpool.Config
		BundlePool                              bundlepool.Config
		GPO                                     gasprice.Config
		EnablePreimageRecording                 bool
		DocRoot                                 string `toml:"-"`
		RPCGasCap                               uint64
		RPCEVMTimeout                           time.Duration
		RPCTxFeeCap                             float64
		OverrideCancun                          *uint64 `toml:",omitempty"`
		OverrideVerkle                          *uint64 `toml:",omitempty"`
		OverrideOptimismCanyon                  *uint64 `toml:",omitempty"`
		OverrideOptimismEcotone                 *uint64 `toml:",omitempty"`
		OverrideOptimismFjord                   *uint64 `toml:",omitempty"`
		OverrideOptimismInterop                 *uint64 `toml:",omitempty"`
		ApplySuperchainUpgrades                 bool    `toml:",omitempty"`
		RollupSequencerHTTP                     string
		RollupHistoricalRPC                     string
		RollupHistoricalRPCTimeout              time.Duration
		RollupDisableTxPoolGossip               bool
		RollupDisableTxPoolAdmission            bool
		RollupHaltOnIncompatibleProtocolVersion string
		EnableOpcodeOptimizing                  bool
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.EthDiscoveryURLs = c.EthDiscoveryURLs
	enc.SnapDiscoveryURLs = c.SnapDiscoveryURLs
	enc.NoPruning = c.NoPruning
	enc.NoPrefetch = c.NoPrefetch
	enc.TxLookupLimit = c.TxLookupLimit
	enc.TransactionHistory = c.TransactionHistory
	enc.StateHistory = c.StateHistory
	enc.StateScheme = c.StateScheme
	enc.PathNodeBuffer = c.PathNodeBuffer
	enc.ProposeBlockInterval = c.ProposeBlockInterval
	enc.EnableProofKeeper = c.EnableProofKeeper
	enc.KeepProofBlockSpan = c.KeepProofBlockSpan
	enc.JournalFileEnabled = c.JournalFileEnabled
	enc.RequiredBlocks = c.RequiredBlocks
	enc.LightServ = c.LightServ
	enc.LightIngress = c.LightIngress
	enc.LightEgress = c.LightEgress
	enc.LightPeers = c.LightPeers
	enc.LightNoPrune = c.LightNoPrune
	enc.LightNoSyncServe = c.LightNoSyncServe
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.BlockDbCachePer = c.BlockDbCachePer
	enc.StateDbCachePer = c.StateDbCachePer
	enc.ChaindbCachePer = c.ChaindbCachePer
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.TrieCommitInterval = c.TrieCommitInterval
	enc.SnapshotCache = c.SnapshotCache
	enc.Preimages = c.Preimages
	enc.NoTries = c.NoTries
	enc.FilterLogCacheSize = c.FilterLogCacheSize
	enc.Miner = c.Miner
	enc.TxPool = c.TxPool
	enc.BlobPool = c.BlobPool
	enc.BundlePool = c.BundlePool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCEVMTimeout = c.RPCEVMTimeout
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	enc.OverrideCancun = c.OverrideCancun
	enc.OverrideVerkle = c.OverrideVerkle
	enc.OverrideOptimismCanyon = c.OverrideOptimismCanyon
	enc.OverrideOptimismEcotone = c.OverrideOptimismEcotone
	enc.OverrideOptimismFjord = c.OverrideOptimismFjord
	enc.OverrideOptimismInterop = c.OverrideOptimismInterop
	enc.ApplySuperchainUpgrades = c.ApplySuperchainUpgrades
	enc.RollupSequencerHTTP = c.RollupSequencerHTTP
	enc.RollupHistoricalRPC = c.RollupHistoricalRPC
	enc.RollupHistoricalRPCTimeout = c.RollupHistoricalRPCTimeout
	enc.RollupDisableTxPoolGossip = c.RollupDisableTxPoolGossip
	enc.RollupDisableTxPoolAdmission = c.RollupDisableTxPoolAdmission
	enc.RollupHaltOnIncompatibleProtocolVersion = c.RollupHaltOnIncompatibleProtocolVersion
	enc.EnableOpcodeOptimizing = c.EnableOpcodeOptimizing
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                                 *core.Genesis `toml:",omitempty"`
		NetworkId                               *uint64
		SyncMode                                *downloader.SyncMode
		EthDiscoveryURLs                        []string
		SnapDiscoveryURLs                       []string
		NoPruning                               *bool
		NoPrefetch                              *bool
		TxLookupLimit                           *uint64                `toml:",omitempty"`
		TransactionHistory                      *uint64                `toml:",omitempty"`
		StateHistory                            *uint64                `toml:",omitempty"`
		StateScheme                             *string                `toml:",omitempty"`
		PathNodeBuffer                          *pathdb.NodeBufferType `toml:",omitempty"`
		ProposeBlockInterval                    *uint64                `toml:",omitempty"`
		EnableProofKeeper                       *bool                  `toml:",omitempty"`
		KeepProofBlockSpan                      *uint64                `toml:",omitempty"`
		JournalFileEnabled                      *bool                  `toml:",omitempty"`
		RequiredBlocks                          map[uint64]common.Hash `toml:"-"`
		LightServ                               *int                   `toml:",omitempty"`
		LightIngress                            *int                   `toml:",omitempty"`
		LightEgress                             *int                   `toml:",omitempty"`
		LightPeers                              *int                   `toml:",omitempty"`
		LightNoPrune                            *bool                  `toml:",omitempty"`
		LightNoSyncServe                        *bool                  `toml:",omitempty"`
		SkipBcVersionCheck                      *bool                  `toml:"-"`
		DatabaseHandles                         *int                   `toml:"-"`
		DatabaseCache                           *int
		BlockDbCachePer                         *int `toml:",omitempty"`
		StateDbCachePer                         *int `toml:",omitempty"`
		ChaindbCachePer                         *int `toml:",omitempty"`
		DatabaseFreezer                         *string
		TrieCleanCache                          *int
		TrieDirtyCache                          *int
		TrieTimeout                             *time.Duration
		TrieCommitInterval                      *uint64
		SnapshotCache                           *int
		Preimages                               *bool
		NoTries                                 *bool
		FilterLogCacheSize                      *int
		Miner                                   *miner.Config
		TxPool                                  *legacypool.Config
		BlobPool                                *blobpool.Config
		BundlePool                              *bundlepool.Config
		GPO                                     *gasprice.Config
		EnablePreimageRecording                 *bool
		DocRoot                                 *string `toml:"-"`
		RPCGasCap                               *uint64
		RPCEVMTimeout                           *time.Duration
		RPCTxFeeCap                             *float64
		OverrideCancun                          *uint64 `toml:",omitempty"`
		OverrideVerkle                          *uint64 `toml:",omitempty"`
		OverrideOptimismCanyon                  *uint64 `toml:",omitempty"`
		OverrideOptimismEcotone                 *uint64 `toml:",omitempty"`
		OverrideOptimismFjord                   *uint64 `toml:",omitempty"`
		OverrideOptimismInterop                 *uint64 `toml:",omitempty"`
		ApplySuperchainUpgrades                 *bool   `toml:",omitempty"`
		RollupSequencerHTTP                     *string
		RollupHistoricalRPC                     *string
		RollupHistoricalRPCTimeout              *time.Duration
		RollupDisableTxPoolGossip               *bool
		RollupDisableTxPoolAdmission            *bool
		RollupHaltOnIncompatibleProtocolVersion *string
		EnableOpcodeOptimizing                  *bool
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.EthDiscoveryURLs != nil {
		c.EthDiscoveryURLs = dec.EthDiscoveryURLs
	}
	if dec.SnapDiscoveryURLs != nil {
		c.SnapDiscoveryURLs = dec.SnapDiscoveryURLs
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.NoPrefetch != nil {
		c.NoPrefetch = *dec.NoPrefetch
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.TransactionHistory != nil {
		c.TransactionHistory = *dec.TransactionHistory
	}
	if dec.StateHistory != nil {
		c.StateHistory = *dec.StateHistory
	}
	if dec.StateScheme != nil {
		c.StateScheme = *dec.StateScheme
	}
	if dec.PathNodeBuffer != nil {
		c.PathNodeBuffer = *dec.PathNodeBuffer
	}
	if dec.ProposeBlockInterval != nil {
		c.ProposeBlockInterval = *dec.ProposeBlockInterval
	}
	if dec.EnableProofKeeper != nil {
		c.EnableProofKeeper = *dec.EnableProofKeeper
	}
	if dec.KeepProofBlockSpan != nil {
		c.KeepProofBlockSpan = *dec.KeepProofBlockSpan
	}
	if dec.JournalFileEnabled != nil {
		c.JournalFileEnabled = *dec.JournalFileEnabled
	}
	if dec.RequiredBlocks != nil {
		c.RequiredBlocks = dec.RequiredBlocks
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightIngress != nil {
		c.LightIngress = *dec.LightIngress
	}
	if dec.LightEgress != nil {
		c.LightEgress = *dec.LightEgress
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.LightNoPrune != nil {
		c.LightNoPrune = *dec.LightNoPrune
	}
	if dec.LightNoSyncServe != nil {
		c.LightNoSyncServe = *dec.LightNoSyncServe
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.BlockDbCachePer != nil {
		c.BlockDbCachePer = *dec.BlockDbCachePer
	}
	if dec.StateDbCachePer != nil {
		c.StateDbCachePer = *dec.StateDbCachePer
	}
	if dec.ChaindbCachePer != nil {
		c.ChaindbCachePer = *dec.ChaindbCachePer
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.TrieCommitInterval != nil {
		c.TrieCommitInterval = *dec.TrieCommitInterval
	}
	if dec.SnapshotCache != nil {
		c.SnapshotCache = *dec.SnapshotCache
	}
	if dec.Preimages != nil {
		c.Preimages = *dec.Preimages
	}
	if dec.NoTries != nil {
		c.NoTries = *dec.NoTries
	}
	if dec.FilterLogCacheSize != nil {
		c.FilterLogCacheSize = *dec.FilterLogCacheSize
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.BlobPool != nil {
		c.BlobPool = *dec.BlobPool
	}
	if dec.BundlePool != nil {
		c.BundlePool = *dec.BundlePool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = *dec.RPCGasCap
	}
	if dec.RPCEVMTimeout != nil {
		c.RPCEVMTimeout = *dec.RPCEVMTimeout
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	if dec.OverrideCancun != nil {
		c.OverrideCancun = dec.OverrideCancun
	}
	if dec.OverrideVerkle != nil {
		c.OverrideVerkle = dec.OverrideVerkle
	}
	if dec.OverrideOptimismCanyon != nil {
		c.OverrideOptimismCanyon = dec.OverrideOptimismCanyon
	}
	if dec.OverrideOptimismEcotone != nil {
		c.OverrideOptimismEcotone = dec.OverrideOptimismEcotone
	}
	if dec.OverrideOptimismFjord != nil {
		c.OverrideOptimismFjord = dec.OverrideOptimismFjord
	}
	if dec.OverrideOptimismInterop != nil {
		c.OverrideOptimismInterop = dec.OverrideOptimismInterop
	}
	if dec.ApplySuperchainUpgrades != nil {
		c.ApplySuperchainUpgrades = *dec.ApplySuperchainUpgrades
	}
	if dec.RollupSequencerHTTP != nil {
		c.RollupSequencerHTTP = *dec.RollupSequencerHTTP
	}
	if dec.RollupHistoricalRPC != nil {
		c.RollupHistoricalRPC = *dec.RollupHistoricalRPC
	}
	if dec.RollupHistoricalRPCTimeout != nil {
		c.RollupHistoricalRPCTimeout = *dec.RollupHistoricalRPCTimeout
	}
	if dec.RollupDisableTxPoolGossip != nil {
		c.RollupDisableTxPoolGossip = *dec.RollupDisableTxPoolGossip
	}
	if dec.RollupDisableTxPoolAdmission != nil {
		c.RollupDisableTxPoolAdmission = *dec.RollupDisableTxPoolAdmission
	}
	if dec.RollupHaltOnIncompatibleProtocolVersion != nil {
		c.RollupHaltOnIncompatibleProtocolVersion = *dec.RollupHaltOnIncompatibleProtocolVersion
	}
	if dec.EnableOpcodeOptimizing != nil {
		c.EnableOpcodeOptimizing = *dec.EnableOpcodeOptimizing
	}
	return nil
}
