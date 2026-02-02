package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessSettings{}

type NetworkaccessSettings struct {
	ConditionalAccess *NetworkaccessConditionalAccessSettings `json:"conditionalAccess,omitempty"`
	CrossTenantAccess *NetworkaccessCrossTenantAccessSettings `json:"crossTenantAccess,omitempty"`
	EnrichedAuditLogs *NetworkaccessEnrichedAuditLogs         `json:"enrichedAuditLogs,omitempty"`
	ForwardingOptions *NetworkaccessForwardingOptions         `json:"forwardingOptions,omitempty"`

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

func (s NetworkaccessSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessSettings{}

func (s NetworkaccessSettings) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.settings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessSettings: %+v", err)
	}

	return encoded, nil
}
