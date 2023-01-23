// Copyright 2022 bnb-chain. All Rights Reserved.
//
// Distributed under MIT license.
// See file LICENSE for detail or copy at https://opensource.org/licenses/MIT

package bsmt

type (
	// Version is the number represents version.
	Version uint64

	// Item is key, value pair.
	Item struct {
		Key uint64
		Val []byte
	}

	// SparseMerkleTree is implemented by types that represents a Sparse Merkle Tree and its operations.
	SparseMerkleTree interface {
		Size() uint64
		Get(key uint64, version *Version) ([]byte, error)
		Set(key uint64, val []byte) error
		SetWithVersion(key uint64, val []byte, newVersion Version) error
		MultiSet(items []Item) error
		MultiSetWithVersion(items []Item, newVersion Version) error
		IsEmpty() bool
		Root() []byte
		GetProof(key uint64) (Proof, error)
		VerifyProof(key uint64, proof Proof) bool
		LatestVersion() Version
		RecentVersion() Version
		Reset()
		Commit(recentVersion *Version) (Version, error)
		CommitWithNewVersion(recentVersion *Version, newVersion *Version) (Version, error)
		Rollback(version Version) error
		Versions() []Version
	}
)
