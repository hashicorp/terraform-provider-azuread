package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdleSessionSignOut struct {
	// Indicates whether the idle session sign-out policy is enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Number of seconds of inactivity after which a user is signed out.
	SignOutAfterInSeconds nullable.Type[int64] `json:"signOutAfterInSeconds,omitempty"`

	// Number of seconds of inactivity after which a user is notified that they'll be signed out.
	WarnAfterInSeconds nullable.Type[int64] `json:"warnAfterInSeconds,omitempty"`
}
