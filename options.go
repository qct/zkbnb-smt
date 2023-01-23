// Copyright 2022 bnb-chain. All Rights Reserved.
//
// Distributed under MIT license.
// See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

package bsmt

import (
	"github.com/bnb-chain/zkbnb-smt/metrics"
	"github.com/panjf2000/ants/v2"
)

// Option is a function that configures SMT.
type Option func(*BNBSparseMerkleTree)

// InitializeVersion configures SMT with given version.
func InitializeVersion(version Version) Option {
	return func(smt *BNBSparseMerkleTree) {
		smt.version = version
	}
}

// BatchSizeLimit configures SMT with given batchSizeLimit.
func BatchSizeLimit(limit int) Option {
	return func(smt *BNBSparseMerkleTree) {
		smt.batchSizeLimit = limit
	}
}

// DBCacheSize configures SMT with given dbCacheSize.
func DBCacheSize(size int) Option {
	return func(smt *BNBSparseMerkleTree) {
		smt.dbCacheSize = size
	}
}

// GoRoutinePool configures SMT with given goroutinePool.
func GoRoutinePool(pool *ants.Pool) Option {
	return func(smt *BNBSparseMerkleTree) {
		smt.goroutinePool = pool
	}
}

// GCThreshold configures SMT with given gc threshold.
func GCThreshold(threshold uint64) Option {
	return func(smt *BNBSparseMerkleTree) {
		if smt.gcStatus != nil {
			smt.gcStatus.threshold = threshold
			smt.gcStatus.segment = threshold / 10
		}

	}
}

// EnableMetrics configures SMT with given metrics.
func EnableMetrics(metrics metrics.Metrics) Option {
	return func(smt *BNBSparseMerkleTree) {
		smt.metrics = metrics
	}
}
