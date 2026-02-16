package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessEnrichedAuditLogs{}

type NetworkaccessEnrichedAuditLogs struct {
	// Exchange Online enriched audit logs settings.
	Exchange *NetworkaccessEnrichedAuditLogsSettings `json:"exchange,omitempty"`

	// SharePoint Online enriched audit logs settings.
	Sharepoint *NetworkaccessEnrichedAuditLogsSettings `json:"sharepoint,omitempty"`

	// Teams enriched audit logs settings.
	Teams *NetworkaccessEnrichedAuditLogsSettings `json:"teams,omitempty"`

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

func (s NetworkaccessEnrichedAuditLogs) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessEnrichedAuditLogs{}

func (s NetworkaccessEnrichedAuditLogs) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessEnrichedAuditLogs
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessEnrichedAuditLogs: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessEnrichedAuditLogs: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.enrichedAuditLogs"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessEnrichedAuditLogs: %+v", err)
	}

	return encoded, nil
}
