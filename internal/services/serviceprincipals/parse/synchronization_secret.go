// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-uuid"
)

type SynchronizationSecretId struct {
	ServicePrincipalId string
}

func NewSynchronizationSecretID(servicePrincipalId string) SynchronizationSecretId {
	return SynchronizationSecretId{
		ServicePrincipalId: servicePrincipalId,
	}
}

func (id SynchronizationSecretId) String() string {
	return id.ServicePrincipalId + "/secrets"
}

func SynchronizationSecretID(idString string) (*SynchronizationSecretId, error) {
	parts := strings.Split(idString, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Object Resource ID should be in the format {servicePrincipalId}/{type} - but got %q", idString)
	}

	id := SynchronizationSecretId{
		ServicePrincipalId: parts[0],
	}

	if _, err := uuid.ParseUUID(id.ServicePrincipalId); err != nil {
		return nil, fmt.Errorf("ServicePrincipalId isn't a valid UUID (%q): %+v", id.ServicePrincipalId, err)
	}

	if parts[1] == "" {
		return nil, fmt.Errorf("Type in {servicePrincipalId}/{type} should not be empty")
	}

	return &id, nil
}
