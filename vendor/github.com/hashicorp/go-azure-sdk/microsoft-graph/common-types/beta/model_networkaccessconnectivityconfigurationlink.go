package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessConnectivityConfigurationLink{}

type NetworkaccessConnectivityConfigurationLink struct {
	// Specifies the name of the link.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies Microsoft's end of the tunnel configuration for a device link.
	LocalConfigurations *[]NetworkaccessLocalConnectivityConfiguration `json:"localConfigurations,omitempty"`

	PeerConfiguration *NetworkaccessPeerConnectivityConfiguration `json:"peerConfiguration,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessConnectivityConfigurationLink) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessConnectivityConfigurationLink{}

func (s NetworkaccessConnectivityConfigurationLink) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessConnectivityConfigurationLink
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessConnectivityConfigurationLink: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessConnectivityConfigurationLink: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.connectivityConfigurationLink"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessConnectivityConfigurationLink: %+v", err)
	}

	return encoded, nil
}
