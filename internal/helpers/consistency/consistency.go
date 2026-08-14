// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package consistency

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type ChangeFunc func(ctx context.Context) (*bool, error)

func WaitForDeletion(ctx context.Context, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	// Disable SDK-level 404 retries for deletion checks
	// Since we are explicitly waiting for a 404, we don't want the SDK to endlessly retry it.
	ctx = client.WithDisable404Retry(ctx)

	timeout := time.Until(deadline)
	_, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Deleted"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
		Refresh: func() (interface{}, string, error) {
			exists, err := f(ctx)
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving resource: %+v", err)
			}
			if exists == nil {
				return nil, "Error", fmt.Errorf("retrieving resource: exists was nil")
			}
			if *exists {
				return "stub", "Waiting", nil
			}
			return "stub", "Deleted", nil
		},
	}).WaitForStateContext(ctx)

	return err
}

type continuousTargetOccurrenceProvider interface {
	GetContinuousTargetOccurrence() int
}

func getCTO(meta interface{}) int {
	if p, ok := meta.(continuousTargetOccurrenceProvider); ok {
		return p.GetContinuousTargetOccurrence()
	}
	return 2
}

func WaitForUpdate(ctx context.Context, meta interface{}, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	_, err := WaitForUpdateWithTimeout(ctx, time.Until(deadline), meta, f)
	return err
}

func WaitForUpdateWithTimeout(ctx context.Context, timeout time.Duration, meta interface{}, f ChangeFunc) (bool, error) {
	res, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: getCTO(meta),
		Refresh: func() (interface{}, string, error) {
			updated, err := f(ctx)
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving resource: %+v", err)
			}
			if updated == nil {
				return nil, "Error", fmt.Errorf("retrieving resource: updated was nil")
			}
			if *updated {
				return true, "Done", nil
			}
			return false, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if res == nil {
		return false, err
	}
	return res.(bool), err
}

func WaitForUpdateDelayStart(ctx context.Context, delay time.Duration, meta interface{}, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	_, err := WaitForUpdateWithTimeoutDelayStart(ctx, time.Until(deadline), delay, meta, f)
	return err
}

func WaitForUpdateWithTimeoutDelayStart(ctx context.Context, timeout, delay time.Duration, meta interface{}, f ChangeFunc) (bool, error) {
	res, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Delay:                     delay,
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: getCTO(meta),
		Refresh: func() (interface{}, string, error) {
			updated, err := f(ctx)
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving resource: %+v", err)
			}
			if updated == nil {
				return nil, "Error", fmt.Errorf("retrieving resource: updated was nil")
			}
			if *updated {
				return true, "Done", nil
			}
			return false, "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if res == nil {
		return false, err
	}
	return res.(bool), err
}
