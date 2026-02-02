package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = ExtensionProperty{}

type ExtensionProperty struct {
	// Display name of the application object on which this extension property is defined. Read-only.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// Specifies the data type of the value the extension property can hold. Following values are supported. Binary - 256
	// bytes maximumBooleanDateTime - Must be specified in ISO 8601 format. Will be stored in UTC.Integer - 32-bit
	// value.LargeInteger - 64-bit value.String - 256 characters maximumNot nullable. For multivalued directory extensions,
	// these limits apply per value in the collection.
	DataType *string `json:"dataType,omitempty"`

	// Defines the directory extension as a multi-valued property. When true, the directory extension property can store a
	// collection of objects of the dataType; for example, a collection of string types such as
	// 'extensionb7b1c57b532f40b8b5ed4b7a7ba67401jobGroupTracker': ['String 1', 'String 2']. The default value is false.
	// Supports $filter (eq).
	IsMultiValued *bool `json:"isMultiValued,omitempty"`

	// Indicates if this extension property was synced from on-premises active directory using Microsoft Entra Connect.
	// Read-only.
	IsSyncedFromOnPremises nullable.Type[bool] `json:"isSyncedFromOnPremises,omitempty"`

	// Name of the extension property. Not nullable. Supports $filter (eq).
	Name *string `json:"name,omitempty"`

	// Following values are supported. Not nullable. UserGroupAdministrativeUnitApplicationDeviceOrganization
	TargetObjects *[]string `json:"targetObjects,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

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

func (s ExtensionProperty) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ExtensionProperty) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExtensionProperty{}

func (s ExtensionProperty) MarshalJSON() ([]byte, error) {
	type wrapper ExtensionProperty
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExtensionProperty: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExtensionProperty: %+v", err)
	}

	delete(decoded, "appDisplayName")
	delete(decoded, "isSyncedFromOnPremises")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.extensionProperty"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExtensionProperty: %+v", err)
	}

	return encoded, nil
}
