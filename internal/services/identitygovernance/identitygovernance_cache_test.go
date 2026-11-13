// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygovernance

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

func TestAccessPackageCache_get(t *testing.T) {
	var cache accessPackageCache
	var reads atomic.Int32
	reader := func(err error) func() (*beta.AccessPackage, error) {
		return func() (*beta.AccessPackage, error) {
			reads.Add(1)
			if err != nil {
				return nil, err
			}
			return &beta.AccessPackage{}, nil
		}
	}
	read := reader(nil)

	var wg sync.WaitGroup
	for range 20 {
		wg.Go(func() {
			if _, err := cache.get("package", read); err != nil {
				t.Errorf("get: %v", err)
			}
		})
	}
	wg.Wait()
	if n := reads.Load(); n != 1 {
		t.Fatalf("20 parallel gets made %d reads, want 1", n)
	}

	cache.evict("package")
	if _, err := cache.get("package", read); err != nil {
		t.Fatalf("get after evict: %v", err)
	}
	if n := reads.Load(); n != 2 {
		t.Fatalf("get after evict made %d reads in total, want 2", n)
	}

	if _, err := cache.get("other", reader(errors.New("throttled"))); err == nil {
		t.Fatal("get returned no error from a failing read")
	}
	if _, err := cache.get("other", read); err != nil {
		t.Fatalf("get after a failed read: %v", err)
	}
	if n := reads.Load(); n != 4 {
		t.Fatalf("a failed read was kept: %d reads in total, want 4", n)
	}
}
