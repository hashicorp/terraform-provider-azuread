package serviceprincipal

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddTokenSigningCertificateRequest struct {
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`
}
