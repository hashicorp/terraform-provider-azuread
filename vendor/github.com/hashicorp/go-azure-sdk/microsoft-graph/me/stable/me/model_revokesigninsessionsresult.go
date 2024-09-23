package me

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RevokeSignInSessionsResult struct {
	Value nullable.Type[bool] `json:"value,omitempty"`
}
