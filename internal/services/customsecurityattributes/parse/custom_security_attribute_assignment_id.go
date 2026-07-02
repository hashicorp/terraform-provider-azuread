// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

func ParseServicePrincipalID(id string) (*beta.ServicePrincipalId, error) {
	spId, err := beta.ParseServicePrincipalID(id)
	if err != nil {
		return nil, fmt.Errorf("parsing %q as a service principal ID: %w", id, err)
	}
	return spId, nil
}
