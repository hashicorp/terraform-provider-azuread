package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCveSeverityInformation struct {
	ExploitedCves *[]WindowsUpdatesCveInformation `json:"exploitedCves,omitempty"`
	MaxSeverity   *WindowsUpdatesCveSeverityLevel `json:"maxSeverity,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
