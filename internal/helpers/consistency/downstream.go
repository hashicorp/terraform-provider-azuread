package consistency

import (
	"context"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// RetryOnSubjectNotFoundConsistencyFailureFunc creates a retry function that catches SubjectNotFound
// errors (400 Bad Request) and deterministically checks if the missing IDs exist in the main directory.
// If the missing IDs exist in the directory, the function returns true to retry the request (indicating a downstream
// consistency delay).
// If any missing ID does not exist in the directory (404 Not Found), the function returns false to fail the request
// immediately (indicating a user or config error).
func RetryOnSubjectNotFoundConsistencyFailureFunc(ctx context.Context, dirClient *directoryobject.DirectoryObjectClient, subjectIds ...string) func(resp *http.Response, o *odata.OData) (bool, error) {
	return func(resp *http.Response, o *odata.OData) (bool, error) {
		if response.WasBadRequest(resp) && o != nil && o.Error != nil && (o.Error.Match("SubjectNotFound") || o.Error.Match("are not found in your directory")) {
			// Respect context bypass
			if resp.Request != nil && resp.Request.Context() != nil {
				if bypass, ok := resp.Request.Context().Value(client.Disable404RetryContextKey).(bool); ok && bypass {
					return false, nil
				}
			}

			var idsToCheck []string

			if len(subjectIds) > 0 {
				idsToCheck = subjectIds
			} else {
				// Extract UUIDs from error message
				// E.g., "SpecificAllowedTargets contains Groups, Users or Service Principals: 96221727-c1df-4caa-bee3-b6f33291d3f7 that are not found in your directory."
				re := regexp.MustCompile(`(?i)([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`)
				for _, match := range re.FindAllStringSubmatch(*o.Error.Message, -1) {
					idsToCheck = append(idsToCheck, match[1])
				}
			}

			if len(idsToCheck) == 0 {
				return false, nil // Cannot verify, fail safe
			}

			// Deterministically check each ID against the main directory
			for _, id := range idsToCheck {
				getResp, err := dirClient.GetDirectoryObject(ctx, stable.NewDirectoryObjectID(id), directoryobject.DefaultGetDirectoryObjectOperationOptions())
				if err != nil {
					if response.WasNotFound(getResp.HttpResponse) {
						return false, nil // Deterministic proof of bad user input - FAIL FAST
					}
				}
			}

			// All missing IDs were found in the main directory!
			// This is confirmed downstream eventual consistency.
			return true, nil // RETRY
		}
		return false, nil
	}
}
