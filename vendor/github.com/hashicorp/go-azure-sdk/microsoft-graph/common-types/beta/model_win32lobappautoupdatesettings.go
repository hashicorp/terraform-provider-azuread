package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAutoUpdateSettings struct {
	// Contains value for auto-update superseded apps.
	AutoUpdateSupersededAppsState *Win32LobAutoUpdateSupersededAppsState `json:"autoUpdateSupersededAppsState,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
