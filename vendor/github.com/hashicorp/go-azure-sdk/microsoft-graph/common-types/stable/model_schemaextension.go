package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SchemaExtension{}

type SchemaExtension struct {
	// Description for the schema extension. Supports $filter (eq).
	Description nullable.Type[string] `json:"description,omitempty"`

	// The appId of the application that is the owner of the schema extension. The owner of the schema definition must be
	// explicitly specified during the Create and Update operations, or it will be implied and auto-assigned by Microsoft
	// Entra ID as follows: In delegated access: The signed-in user must be the owner of the app that calls Microsoft Graph
	// to create the schema extension definition. If the signed-in user isn't the owner of the calling app, they must
	// explicitly specify the owner property, and assign it the appId of an app that they own. In app-only access: The owner
	// property isn't required in the request body. Instead, the calling app is assigned ownership of the schema extension.
	// So, for example, if creating a new schema extension definition using Graph Explorer, you must supply the owner
	// property. Once set, this property is read-only and cannot be changed. Supports $filter (eq).
	Owner *string `json:"owner,omitempty"`

	// The collection of property names and types that make up the schema extension definition.
	Properties *[]ExtensionSchemaProperty `json:"properties,omitempty"`

	// The lifecycle state of the schema extension. Possible states are InDevelopment, Available, and Deprecated.
	// Automatically set to InDevelopment on creation. For more information about the possible state transitions and
	// behaviors, see Schema extensions lifecycle. Supports $filter (eq).
	Status *string `json:"status,omitempty"`

	// Set of Microsoft Graph types (that can support extensions) that the schema extension can be applied to. Select from
	// administrativeUnit, contact, device, event, group, message, organization, post, todoTask, todoTaskList, or user.
	TargetTypes *[]string `json:"targetTypes,omitempty"`

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

func (s SchemaExtension) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SchemaExtension{}

func (s SchemaExtension) MarshalJSON() ([]byte, error) {
	type wrapper SchemaExtension
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SchemaExtension: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SchemaExtension: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.schemaExtension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SchemaExtension: %+v", err)
	}

	return encoded, nil
}
