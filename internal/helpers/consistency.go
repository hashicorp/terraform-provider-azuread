package helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type ChangeFunc func(ctx context.Context) (*bool, error)

func WaitForDeletion(ctx context.Context, f ChangeFunc) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return errors.New("context has no deadline")
	}

	timeout := time.Until(deadline)
	_, err := (&resource.StateChangeConf{
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

	timeout := time.Until(deadline)
	_, err := (&resource.StateChangeConf{
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
				return "stub", "Done", nil
			}
			return "stub", "Waiting", nil
		},
	}).WaitForStateContext(ctx)

	return err
}
