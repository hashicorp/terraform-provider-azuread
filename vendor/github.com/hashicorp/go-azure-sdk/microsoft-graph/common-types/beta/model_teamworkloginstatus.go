package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkLoginStatus struct {
	// Information about the Exchange connection.
	ExchangeConnection *TeamworkConnection `json:"exchangeConnection,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about the Skype for Business connection.
	SkypeConnection *TeamworkConnection `json:"skypeConnection,omitempty"`

	// Information about the Teams connection.
	TeamsConnection *TeamworkConnection `json:"teamsConnection,omitempty"`
}
