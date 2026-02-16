package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemTypeAction interface {
	Entity
	AuthorizationSystemTypeAction() BaseAuthorizationSystemTypeActionImpl
}

var _ AuthorizationSystemTypeAction = BaseAuthorizationSystemTypeActionImpl{}

type BaseAuthorizationSystemTypeActionImpl struct {
	// The type of action allowed in the authorization system's service. The possible values are: delete, read,
	// unknownFutureValue. Supports $filter and (eq).
	ActionType *AuthorizationSystemActionType `json:"actionType,omitempty"`

	// The display name of an action. Read-only. Supports $filter and (eq).
	ExternalId *string `json:"externalId,omitempty"`

	// The resource types in the authorization system's service where the action can be performed. Supports $filter and
	// (eq).
	ResourceTypes *[]string `json:"resourceTypes,omitempty"`

	Severity *AuthorizationSystemActionSeverity `json:"severity,omitempty"`

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

func (s BaseAuthorizationSystemTypeActionImpl) AuthorizationSystemTypeAction() BaseAuthorizationSystemTypeActionImpl {
	return s
}

func (s BaseAuthorizationSystemTypeActionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthorizationSystemTypeAction = RawAuthorizationSystemTypeActionImpl{}

// RawAuthorizationSystemTypeActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthorizationSystemTypeActionImpl struct {
	authorizationSystemTypeAction BaseAuthorizationSystemTypeActionImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawAuthorizationSystemTypeActionImpl) AuthorizationSystemTypeAction() BaseAuthorizationSystemTypeActionImpl {
	return s.authorizationSystemTypeAction
}

func (s RawAuthorizationSystemTypeActionImpl) Entity() BaseEntityImpl {
	return s.authorizationSystemTypeAction.Entity()
}

var _ json.Marshaler = BaseAuthorizationSystemTypeActionImpl{}

func (s BaseAuthorizationSystemTypeActionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthorizationSystemTypeActionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthorizationSystemTypeActionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthorizationSystemTypeActionImpl: %+v", err)
	}

	delete(decoded, "externalId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authorizationSystemTypeAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthorizationSystemTypeActionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthorizationSystemTypeActionImplementation(input []byte) (AuthorizationSystemTypeAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthorizationSystemTypeAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.awsAuthorizationSystemTypeAction") {
		var out AwsAuthorizationSystemTypeAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AwsAuthorizationSystemTypeAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAuthorizationSystemTypeAction") {
		var out AzureAuthorizationSystemTypeAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAuthorizationSystemTypeAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.gcpAuthorizationSystemTypeAction") {
		var out GcpAuthorizationSystemTypeAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GcpAuthorizationSystemTypeAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthorizationSystemTypeActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthorizationSystemTypeActionImpl: %+v", err)
	}

	return RawAuthorizationSystemTypeActionImpl{
		authorizationSystemTypeAction: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
