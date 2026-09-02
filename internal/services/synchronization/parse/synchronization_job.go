// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type SynchronizationJobId struct {
	ServicePrincipalId string
	JobId              string
}

func NewSynchronizationJobID(servicePrincipalId string, jobId string) SynchronizationJobId {
	return SynchronizationJobId{
		ServicePrincipalId: servicePrincipalId,
		JobId:              jobId,
	}
}

func (id SynchronizationJobId) String() string {
	return id.ServicePrincipalId + "/job/" + id.JobId
}

func SynchronizationJobID(idString string) (*SynchronizationJobId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("object Resource ID should be in the format {servicePrincipalId}/{type}/{jobId} - but got %q", idString)
	}

	id := SynchronizationJobId{
		ServicePrincipalId: parts[0],
		JobId:              parts[2],
	}

	if _, err := uuid.ParseUUID(id.ServicePrincipalId); err != nil {
		return nil, fmt.Errorf("servicePrincipalId isn't a valid UUID (%q): %+v", id.ServicePrincipalId, err)
	}

	if parts[1] == "" {
		return nil, errors.New("type in {servicePrincipalId}/{type}/{subID} should not be empty")
	}

	return &id, nil
}
