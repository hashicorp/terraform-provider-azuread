package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppInstallTimeSettings struct {
	// The time at which the app should be installed.
	DeadlineDateTime *string `json:"deadlineDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Whether the local device time or UTC time should be used when determining the deadline times.
	UseLocalTime *bool `json:"useLocalTime,omitempty"`
}
