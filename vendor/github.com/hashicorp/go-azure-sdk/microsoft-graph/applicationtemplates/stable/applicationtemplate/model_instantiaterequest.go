package applicationtemplate

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstantiateRequest struct {
	DisplayName                nullable.Type[string] `json:"displayName,omitempty"`
	ServiceManagementReference nullable.Type[string] `json:"serviceManagementReference,omitempty"`
}
