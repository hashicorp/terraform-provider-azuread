package msgraph

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func WaitForCreationReplication(ctx context.Context, timeout time.Duration, f func() (interface{}, int, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"NotFound", "BadCast"},
		Target:                    []string{"Found"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			i, status, err := f()

			switch {
			case status >= 200 && status < 300:
				return i, "Found", nil
			case status == 404:
				return i, "NotFound", nil
			case i == nil:
				return nil, "BadCast", nil
			case err != nil:
				return i, "Error", fmt.Errorf("unable to retrieve object, received response with status %d: %v", status, err)
			}

			return i, "Error", fmt.Errorf("unrecognised response with status %d", status)
		},
	}).WaitForStateContext(ctx)
}

func WaitForListAdd(ctx context.Context, item string, f func() ([]string, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"NotFound"},
		Target:                    []string{"Found"},
		Timeout:                   5 * time.Minute,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			listItems, err := f()

			if err != nil {
				return listItems, "Error", err
			}

			for _, v := range listItems {
				if v == item {
					return listItems, "Found", nil
				}
			}

			return listItems, "NotFound", nil
		},
	}).WaitForStateContext(ctx)
}

func WaitForListRemove(ctx context.Context, item string, f func() ([]string, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"Found"},
		Target:                    []string{"NotFound"},
		Timeout:                   5 * time.Minute,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 2,
		Refresh: func() (interface{}, string, error) {
			listItems, err := f()

			if err != nil {
				return listItems, "Error", err
			}

			for _, v := range listItems {
				if v == item {
					return listItems, "Found", nil
				}
			}

			return listItems, "NotFound", nil
		},
	}).WaitForStateContext(ctx)
}
