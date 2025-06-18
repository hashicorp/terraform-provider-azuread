package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationEventsFlow = ExternalUsersSelfServiceSignUpEventsFlow{}

type ExternalUsersSelfServiceSignUpEventsFlow struct {
	// The configuration for what to invoke when attributes are ready to be collected from the user.
	OnAttributeCollection OnAttributeCollectionHandler `json:"onAttributeCollection"`

	// The configuration for what to invoke when attribution collection starts.
	OnAttributeCollectionStart OnAttributeCollectionStartHandler `json:"onAttributeCollectionStart"`

	// The configuration for what to invoke when attributes are submitted at the end of attribution collection.
	OnAttributeCollectionSubmit OnAttributeCollectionSubmitHandler `json:"onAttributeCollectionSubmit"`

	// Required. The configuration for what to invoke when authentication methods are ready to be presented to the user.
	// Must have at least one identity provider linked. Supports $filter (eq). See support for filtering on user flows for
	// syntax information.
	OnAuthenticationMethodLoadStart OnAuthenticationMethodLoadStartHandler `json:"onAuthenticationMethodLoadStart"`

	// Required. The configuration for what to invoke when an authentication flow is ready to be initiated.
	OnInteractiveAuthFlowStart OnInteractiveAuthFlowStartHandler `json:"onInteractiveAuthFlowStart"`

	// The configuration for what to invoke during user creation.
	OnUserCreateStart OnUserCreateStartHandler `json:"onUserCreateStart"`

	// Fields inherited from AuthenticationEventsFlow

	// The conditions representing the context of the authentication request that's used to decide whether the events policy
	// is invoked. Supports $filter (eq). See support for filtering on user flows for syntax information.
	Conditions *AuthenticationConditions `json:"conditions,omitempty"`

	// The description of the events policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Required. The display name for the events policy.
	DisplayName string `json:"displayName"`

	// The priority to use for each individual event of the events policy. If multiple competing listeners for an event have
	// the same priority, one is chosen and an error is silently logged. Defaults to 500.
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

func (s ExternalUsersSelfServiceSignUpEventsFlow) AuthenticationEventsFlow() BaseAuthenticationEventsFlowImpl {
	return BaseAuthenticationEventsFlowImpl{
		Conditions:  s.Conditions,
		Description: s.Description,
		DisplayName: s.DisplayName,
		Priority:    s.Priority,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s ExternalUsersSelfServiceSignUpEventsFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalUsersSelfServiceSignUpEventsFlow{}

func (s ExternalUsersSelfServiceSignUpEventsFlow) MarshalJSON() ([]byte, error) {
	type wrapper ExternalUsersSelfServiceSignUpEventsFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalUsersSelfServiceSignUpEventsFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalUsersSelfServiceSignUpEventsFlow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalUsersSelfServiceSignUpEventsFlow: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExternalUsersSelfServiceSignUpEventsFlow{}

func (s *ExternalUsersSelfServiceSignUpEventsFlow) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Conditions  *AuthenticationConditions `json:"conditions,omitempty"`
		Description nullable.Type[string]     `json:"description,omitempty"`
		DisplayName string                    `json:"displayName"`
		Priority    *int64                    `json:"priority,omitempty"`
		Id          *string                   `json:"id,omitempty"`
		ODataId     *string                   `json:"@odata.id,omitempty"`
		ODataType   *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Conditions = decoded.Conditions
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Priority = decoded.Priority

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExternalUsersSelfServiceSignUpEventsFlow into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["onAttributeCollection"]; ok {
		impl, err := UnmarshalOnAttributeCollectionHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnAttributeCollection' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnAttributeCollection = impl
	}

	if v, ok := temp["onAttributeCollectionStart"]; ok {
		impl, err := UnmarshalOnAttributeCollectionStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnAttributeCollectionStart' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnAttributeCollectionStart = impl
	}

	if v, ok := temp["onAttributeCollectionSubmit"]; ok {
		impl, err := UnmarshalOnAttributeCollectionSubmitHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnAttributeCollectionSubmit' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnAttributeCollectionSubmit = impl
	}

	if v, ok := temp["onAuthenticationMethodLoadStart"]; ok {
		impl, err := UnmarshalOnAuthenticationMethodLoadStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnAuthenticationMethodLoadStart' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnAuthenticationMethodLoadStart = impl
	}

	if v, ok := temp["onInteractiveAuthFlowStart"]; ok {
		impl, err := UnmarshalOnInteractiveAuthFlowStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnInteractiveAuthFlowStart' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnInteractiveAuthFlowStart = impl
	}

	if v, ok := temp["onUserCreateStart"]; ok {
		impl, err := UnmarshalOnUserCreateStartHandlerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnUserCreateStart' for 'ExternalUsersSelfServiceSignUpEventsFlow': %+v", err)
		}
		s.OnUserCreateStart = impl
	}

	return nil
}
