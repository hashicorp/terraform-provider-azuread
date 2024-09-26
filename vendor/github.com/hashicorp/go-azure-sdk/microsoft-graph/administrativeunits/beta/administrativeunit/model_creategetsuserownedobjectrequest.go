package administrativeunit

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGetsUserOwnedObjectRequest struct {
	Type   nullable.Type[string] `json:"type,omitempty"`
	UserId nullable.Type[string] `json:"userId,omitempty"`
}
