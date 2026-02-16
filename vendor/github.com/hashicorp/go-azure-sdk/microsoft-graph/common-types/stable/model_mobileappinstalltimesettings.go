package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallTimeSettings struct {
	// The time at which the app should be installed.
	DeadlineDateTime nullable.Type[string] `json:"deadlineDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time at which the app should be available for installation.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Whether the local device time or UTC time should be used when determining the available and deadline times.
	UseLocalTime *bool `json:"useLocalTime,omitempty"`
}
