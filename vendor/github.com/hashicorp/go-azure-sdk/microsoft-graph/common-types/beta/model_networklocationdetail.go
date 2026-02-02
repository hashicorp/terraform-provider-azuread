package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkLocationDetail struct {
	// Provides the name of the network used when signing in.
	NetworkNames *[]string `json:"networkNames,omitempty"`

	// Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted,
	// unknownFutureValue.
	NetworkType *NetworkType `json:"networkType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
