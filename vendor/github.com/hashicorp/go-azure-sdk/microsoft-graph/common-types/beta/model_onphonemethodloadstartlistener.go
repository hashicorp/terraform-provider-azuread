package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationEventListener = OnPhoneMethodLoadStartListener{}

type OnPhoneMethodLoadStartListener struct {
	Handler OnPhoneMethodLoadStartHandler `json:"handler"`

	// Fields inherited from AuthenticationEventListener

	// The identifier of the authenticationEventsFlow object.
	AuthenticationEventsFlowId nullable.Type[string] `json:"authenticationEventsFlowId,omitempty"`

	// The conditions on which this authenticationEventListener should trigger.
	Conditions *AuthenticationConditions `json:"conditions,omitempty"`

	// The priority of this handler. Between 0 (lower priority) and 1000 (higher priority).
	Priority *int64 `json:"priority,omitempty"`

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

func (s OnPhoneMethodLoadStartListener) AuthenticationEventListener() BaseAuthenticationEventListenerImpl {
	return BaseAuthenticationEventListenerImpl{
		AuthenticationEventsFlowId: s.AuthenticationEventsFlowId,
		Conditions:                 s.Conditions,
		Priority:                   s.Priority,
		Id:                         s.Id,
		ODataId:                    s.ODataId,
		ODataType:                  s.ODataType,
	}
}

func (s OnPhoneMethodLoadStartListener) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OnPhoneMethodLoadStartListener{}

func (s OnPhoneMethodLoadStartListener) MarshalJSON() ([]byte, error) {
	type wrapper OnPhoneMethodLoadStartListener
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OnPhoneMethodLoadStartListener: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OnPhoneMethodLoadStartListener: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.onPhoneMethodLoadStartListener"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OnPhoneMethodLoadStartListener: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OnPhoneMethodLoadStartListener{}

func (s *OnPhoneMethodLoadStartListener) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationEventsFlowId nullable.Type[string]     `json:"authenticationEventsFlowId,omitempty"`
		Conditions                 *AuthenticationConditions `json:"conditions,omitempty"`
		Priority                   *int64                    `json:"priority,omitempty"`
		Id                         *string                   `json:"id,omitempty"`
		ODataId                    *string                   `json:"@odata.id,omitempty"`
		ODataType                  *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationEventsFlowId = decoded.AuthenticationEventsFlowId
	s.Conditions = decoded.Conditions
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Priority = decoded.Priority

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OnPhoneMethodLoadStartListener into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["handler"]; ok {
		impl, err := UnmarshalOnPhoneMethodLoadStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Handler' for 'OnPhoneMethodLoadStartListener': %+v", err)
		}
		s.Handler = impl
	}

	return nil
}
