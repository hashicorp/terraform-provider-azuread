package aadgraph

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func WaitForCreationReplication(ctx context.Context, timeout time.Duration, f func() (interface{}, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"NotFound", "BadCast"},
		Target:                    []string{"Found"},
		Timeout:                   timeout,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
		Refresh: func() (interface{}, string, error) {
			i, err := f()
			if err == nil {
				return i, "Found", nil
			}

			r, ok := i.(autorest.Response)
			if !ok {
				return i, "BadCast", nil // sometimes the SDK bubbles up an entirely empty object
			}
			if utils.ResponseWasNotFound(r) {
				return i, "NotFound", nil
			}
			return i, "Error", fmt.Errorf("unable to retrieve object, received response with status %d: %v", r.StatusCode, err)
		},
	}).WaitForStateContext(ctx)
}

func WaitForListAdd(ctx context.Context, item string, f func() ([]string, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"404"},
		Target:                    []string{"Found"},
		Timeout:                   5 * time.Minute,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
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

			return listItems, "404", nil
		},
	}).WaitForStateContext(ctx)
}

func WaitForListRemove(ctx context.Context, item string, f func() ([]string, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"Found"},
		Target:                    []string{"NotFound"},
		Timeout:                   5 * time.Minute,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
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
