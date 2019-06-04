package graph

import (
	"fmt"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func WaitForReplication(f func() (interface{}, error)) (interface{}, error) {
	return (&resource.StateChangeConf{
		Pending:                   []string{"404"},
		Target:                    []string{"Found"},
		Timeout:                   5 * time.Minute,
		MinTimeout:                1 * time.Second,
		ContinuousTargetOccurence: 10,
		Refresh: func() (interface{}, string, error) {
			i, err := f()
			if err != nil {
				r, ok := i.(autorest.Response)
				if !ok {
					return i, "Error", fmt.Errorf("Unable to cast to response: %v", i)
				}
				if ar.ResponseWasNotFound(r) {
					return i, "404", nil
				}
				return i, "Error", fmt.Errorf("Error calling f, response was not 404 (%d): %v", r.StatusCode, err)
			}

			return i, "Found", nil
		},
	}).WaitForState()
}
