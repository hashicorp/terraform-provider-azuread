package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationListener = InvokeUserFlowListener{}

type InvokeUserFlowListener struct {
	// The user flow that is invoked when this action executes.
	UserFlow *B2xIdentityUserFlow `json:"userFlow,omitempty"`

	// Fields inherited from AuthenticationListener

	// The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority
	// is evaluated from low to high.
	Priority *int64 `json:"priority,omitempty"`

	// Filter based on the source of the authentication that is used to determine whether the listener is evaluated, and is
	// currently limited to evaluations based on application the user is authenticating to.
	SourceFilter *AuthenticationSourceFilter `json:"sourceFilter,omitempty"`

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

func (s InvokeUserFlowListener) AuthenticationListener() BaseAuthenticationListenerImpl {
	return BaseAuthenticationListenerImpl{
		Priority:     s.Priority,
		SourceFilter: s.SourceFilter,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s InvokeUserFlowListener) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InvokeUserFlowListener{}

func (s InvokeUserFlowListener) MarshalJSON() ([]byte, error) {
	type wrapper InvokeUserFlowListener
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InvokeUserFlowListener: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InvokeUserFlowListener: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.invokeUserFlowListener"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InvokeUserFlowListener: %+v", err)
	}

	return encoded, nil
}
