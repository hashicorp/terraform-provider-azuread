package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesResourceConnection interface {
	Entity
	WindowsUpdatesResourceConnection() BaseWindowsUpdatesResourceConnectionImpl
}

var _ WindowsUpdatesResourceConnection = BaseWindowsUpdatesResourceConnectionImpl{}

type BaseWindowsUpdatesResourceConnectionImpl struct {
	// The state of the connection. The possible values are: connected, notAuthorized, notFound, unknownFutureValue.
	State *WindowsUpdatesResourceConnectionState `json:"state,omitempty"`

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

func (s BaseWindowsUpdatesResourceConnectionImpl) WindowsUpdatesResourceConnection() BaseWindowsUpdatesResourceConnectionImpl {
	return s
}

func (s BaseWindowsUpdatesResourceConnectionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesResourceConnection = RawWindowsUpdatesResourceConnectionImpl{}

// RawWindowsUpdatesResourceConnectionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesResourceConnectionImpl struct {
	windowsUpdatesResourceConnection BaseWindowsUpdatesResourceConnectionImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawWindowsUpdatesResourceConnectionImpl) WindowsUpdatesResourceConnection() BaseWindowsUpdatesResourceConnectionImpl {
	return s.windowsUpdatesResourceConnection
}

func (s RawWindowsUpdatesResourceConnectionImpl) Entity() BaseEntityImpl {
	return s.windowsUpdatesResourceConnection.Entity()
}

var _ json.Marshaler = BaseWindowsUpdatesResourceConnectionImpl{}

func (s BaseWindowsUpdatesResourceConnectionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesResourceConnectionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesResourceConnectionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesResourceConnectionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.resourceConnection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesResourceConnectionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesResourceConnectionImplementation(input []byte) (WindowsUpdatesResourceConnection, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesResourceConnection into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.operationalInsightsConnection") {
		var out WindowsUpdatesOperationalInsightsConnection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesOperationalInsightsConnection: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesResourceConnectionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesResourceConnectionImpl: %+v", err)
	}

	return RawWindowsUpdatesResourceConnectionImpl{
		windowsUpdatesResourceConnection: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
