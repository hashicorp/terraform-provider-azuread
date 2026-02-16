package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkTeamsClientConfiguration struct {
	// The configuration of the Microsoft Teams client user account for a device.
	AccountConfiguration *TeamworkAccountConfiguration `json:"accountConfiguration,omitempty"`

	// The configuration of Microsoft Teams client features for a device.
	FeaturesConfiguration *TeamworkFeaturesConfiguration `json:"featuresConfiguration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
