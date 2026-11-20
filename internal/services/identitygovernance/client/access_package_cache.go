// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"sync"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// AccessPackageCache keeps one read of each access package with its role scopes, so refreshing all the resource
// package associations of an access package makes one Graph request instead of 2 for each association. It lives as
// long as the Client, which the provider builds once for each plan or apply. It does not keep errors.
type AccessPackageCache struct {
	entries sync.Map // access package ID: *accessPackageCacheEntry
}

type accessPackageCacheEntry struct {
	once          sync.Once
	accessPackage *beta.AccessPackage
	err           error
}

// Get returns the cached access package, or calls read once for all callers that ask for it at the same time.
func (c *AccessPackageCache) Get(accessPackageId string, read func() (*beta.AccessPackage, error)) (*beta.AccessPackage, error) {
	v, _ := c.entries.LoadOrStore(accessPackageId, &accessPackageCacheEntry{})
	entry := v.(*accessPackageCacheEntry)
	entry.once.Do(func() {
		entry.accessPackage, entry.err = read()
	})
	if entry.err != nil {
		c.entries.CompareAndDelete(accessPackageId, entry)
		return nil, entry.err
	}

	return entry.accessPackage, nil
}

// Evict removes the access package, so the next Get reads it again.
func (c *AccessPackageCache) Evict(accessPackageId string) {
	c.entries.Delete(accessPackageId)
}
