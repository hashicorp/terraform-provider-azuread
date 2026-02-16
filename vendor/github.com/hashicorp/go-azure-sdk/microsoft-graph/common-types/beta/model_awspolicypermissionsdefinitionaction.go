package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AwsPermissionsDefinitionAction = AwsPolicyPermissionsDefinitionAction{}

type AwsPolicyPermissionsDefinitionAction struct {
	// ID for the role.
	AssignToRoleId nullable.Type[string] `json:"assignToRoleId,omitempty"`

	Policies *[]PermissionsDefinitionAwsPolicy `json:"policies,omitempty"`

	// Fields inherited from PermissionsDefinitionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AwsPolicyPermissionsDefinitionAction) AwsPermissionsDefinitionAction() BaseAwsPermissionsDefinitionActionImpl {
	return BaseAwsPermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AwsPolicyPermissionsDefinitionAction) PermissionsDefinitionAction() BasePermissionsDefinitionActionImpl {
	return BasePermissionsDefinitionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsPolicyPermissionsDefinitionAction{}

func (s AwsPolicyPermissionsDefinitionAction) MarshalJSON() ([]byte, error) {
	type wrapper AwsPolicyPermissionsDefinitionAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsPolicyPermissionsDefinitionAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsPolicyPermissionsDefinitionAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsPolicyPermissionsDefinitionAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsPolicyPermissionsDefinitionAction: %+v", err)
	}

	return encoded, nil
}
