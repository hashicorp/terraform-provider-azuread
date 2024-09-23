package me

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangePasswordRequest struct {
	CurrentPassword nullable.Type[string] `json:"currentPassword,omitempty"`
	NewPassword     nullable.Type[string] `json:"newPassword,omitempty"`
}
