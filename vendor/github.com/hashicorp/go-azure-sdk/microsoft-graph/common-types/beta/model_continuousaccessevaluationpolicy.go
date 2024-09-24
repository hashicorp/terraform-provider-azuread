package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ContinuousAccessEvaluationPolicy{}

type ContinuousAccessEvaluationPolicy struct {
	// Continuous access evaluation automatically blocks access to resources and applications in near real time when a
	// user's access is removed or a client IP address changes. Read-only.
	Description *string `json:"description,omitempty"`

	// The value is always Continuous Access Evaluation. Read-only.
	DisplayName *string `json:"displayName,omitempty"`

	// The collection of group identifiers in scope for evaluation. All groups are in scope when the collection is empty.
	// Read-only.
	Groups *[]string `json:"groups,omitempty"`

	// true to indicate whether continuous access evaluation should be performed; otherwise false. Read-only.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// true to indicate that the continuous access evaluation policy settings should be or has been migrated to the
	// conditional access policy.
	Migrate *bool `json:"migrate,omitempty"`

	// The collection of user identifiers in scope for evaluation. All users are in scope when the collection is empty.
	// Read-only.
	Users *[]string `json:"users,omitempty"`

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

func (s ContinuousAccessEvaluationPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ContinuousAccessEvaluationPolicy{}

func (s ContinuousAccessEvaluationPolicy) MarshalJSON() ([]byte, error) {
	type wrapper ContinuousAccessEvaluationPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContinuousAccessEvaluationPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContinuousAccessEvaluationPolicy: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "groups")
	delete(decoded, "isEnabled")
	delete(decoded, "users")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.continuousAccessEvaluationPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContinuousAccessEvaluationPolicy: %+v", err)
	}

	return encoded, nil
}
