package directoryrole

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMemberObjectsRequest struct {
	SecurityEnabledOnly nullable.Type[bool] `json:"securityEnabledOnly,omitempty"`
}
