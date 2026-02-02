package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalActivity interface {
	Entity
	ExternalConnectorsExternalActivity() BaseExternalConnectorsExternalActivityImpl
}

var _ ExternalConnectorsExternalActivity = BaseExternalConnectorsExternalActivityImpl{}

type BaseExternalConnectorsExternalActivityImpl struct {
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

func (s BaseExternalConnectorsExternalActivityImpl) ExternalConnectorsExternalActivity() BaseExternalConnectorsExternalActivityImpl {
	return s
}

func (s BaseExternalConnectorsExternalActivityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ExternalConnectorsExternalActivity = RawExternalConnectorsExternalActivityImpl{}

// RawExternalConnectorsExternalActivityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExternalConnectorsExternalActivityImpl struct {
	externalConnectorsExternalActivity BaseExternalConnectorsExternalActivityImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawExternalConnectorsExternalActivityImpl) ExternalConnectorsExternalActivity() BaseExternalConnectorsExternalActivityImpl {
	return s.externalConnectorsExternalActivity
}

func (s RawExternalConnectorsExternalActivityImpl) Entity() BaseEntityImpl {
	return s.externalConnectorsExternalActivity.Entity()
}

var _ json.Marshaler = BaseExternalConnectorsExternalActivityImpl{}

func (s BaseExternalConnectorsExternalActivityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExternalConnectorsExternalActivityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExternalConnectorsExternalActivityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExternalConnectorsExternalActivityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.externalActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExternalConnectorsExternalActivityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExternalConnectorsExternalActivityImplementation(input []byte) (ExternalConnectorsExternalActivity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsExternalActivity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalConnectors.externalActivityResult") {
		var out ExternalConnectorsExternalActivityResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalConnectorsExternalActivityResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseExternalConnectorsExternalActivityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExternalConnectorsExternalActivityImpl: %+v", err)
	}

	return RawExternalConnectorsExternalActivityImpl{
		externalConnectorsExternalActivity: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
