package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimConditionBase = CustomClaimCondition{}

type CustomClaimCondition struct {
	// A list of groups (GUIDs) to which the user/application must be a member for this condition to be applied.
	MemberOf *[]string `json:"memberOf,omitempty"`

	// The type of user this condition applies to. The possible values are: any, members, allGuests, aadGuests,
	// externalGuests, unknownFutureValue.
	UserType *ClaimConditionUserType `json:"userType,omitempty"`

	// Fields inherited from CustomClaimConditionBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomClaimCondition) CustomClaimConditionBase() BaseCustomClaimConditionBaseImpl {
	return BaseCustomClaimConditionBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomClaimCondition{}

func (s CustomClaimCondition) MarshalJSON() ([]byte, error) {
	type wrapper CustomClaimCondition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomClaimCondition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimCondition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customClaimCondition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomClaimCondition: %+v", err)
	}

	return encoded, nil
}
