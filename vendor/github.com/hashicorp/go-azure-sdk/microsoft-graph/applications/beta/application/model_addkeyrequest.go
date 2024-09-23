package application

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddKeyRequest struct {
	KeyCredential      *beta.KeyCredential      `json:"keyCredential,omitempty"`
	PasswordCredential *beta.PasswordCredential `json:"passwordCredential,omitempty"`
	Proof              *string                  `json:"proof,omitempty"`
}
