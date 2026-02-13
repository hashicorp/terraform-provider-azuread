// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package consistency

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type ChangeFunc func(ctx context.Context) (*bool, error)

func WaitForDeletion(ctx context.Context, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

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

func WaitForUpdate(ctx context.Context, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	_, err := WaitForUpdateWithTimeout(ctx, time.Until(deadline), f)
	return err
}

func WaitForUpdateWithTimeout(ctx context.Context, timeout time.Duration, f ChangeFunc) (bool, error) {
	res, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
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

func WaitForUpdateDelayStart(ctx context.Context, delay time.Duration, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	_, err := WaitForUpdateWithTimeoutDelayStart(ctx, time.Until(deadline), delay, f)
	return err
}

func WaitForUpdateWithTimeoutDelayStart(ctx context.Context, timeout, delay time.Duration, f ChangeFunc) (bool, error) {
	res, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Delay:                     delay,
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Done"},
		Timeout:                   timeout,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
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

func WaitForCreation(ctx context.Context, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	timeout := time.Until(deadline)
	_, err := WaitForCreationWithTimeout(ctx, timeout, f)

	return err
}

func WaitForCreationWithTimeout(ctx context.Context, timeout time.Duration, f ChangeFunc) (bool, error) {
	res, err := (&pluginsdk.StateChangeConf{
		Pending:                   []string{"Waiting"},
		Target:                    []string{"Created"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 3, // Require 3 successful checks
		Refresh: func() (interface{}, string, error) {
			exists, err := f(ctx)
			if err != nil {
				return nil, "Error", fmt.Errorf("checking resource: %+v", err)
			}
			if exists == nil {
				return nil, "Error", fmt.Errorf("checking resource: exists was nil")
			}
			if *exists {
				return "stub", "Created", nil
			}
			return "stub", "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	if res == nil {
		return false, err
	}

	return true, err
}
