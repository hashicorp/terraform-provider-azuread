package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExternalConnectorsExternalActivity = ExternalConnectorsExternalActivityResult{}

type ExternalConnectorsExternalActivityResult struct {
	// Error information that explains the failure to process an external activity.
	Error *PublicError `json:"error,omitempty"`

	// Fields inherited from ExternalConnectorsExternalActivity

	// Represents an identity used to identify who is responsible for the activity.
	PerformedBy *ExternalConnectorsIdentity `json:"performedBy,omitempty"`

	// The date and time when the particular activity occurred. The DateTimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime *string `json:"startDateTime,omitempty"`

	Type *ExternalConnectorsExternalActivityType `json:"type,omitempty"`

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

func (s ExternalConnectorsExternalActivityResult) ExternalConnectorsExternalActivity() BaseExternalConnectorsExternalActivityImpl {
	return BaseExternalConnectorsExternalActivityImpl{
		PerformedBy:   s.PerformedBy,
		StartDateTime: s.StartDateTime,
		Type:          s.Type,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s ExternalConnectorsExternalActivityResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnectorsExternalActivityResult{}

func (s ExternalConnectorsExternalActivityResult) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsExternalActivityResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsExternalActivityResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsExternalActivityResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.externalActivityResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsExternalActivityResult: %+v", err)
	}

	return encoded, nil
}
