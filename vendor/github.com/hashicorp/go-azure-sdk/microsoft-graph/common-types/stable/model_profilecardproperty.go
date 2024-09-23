package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProfileCardProperty{}

type ProfileCardProperty struct {
	// Allows an administrator to set a custom display label for the directory property and localize it for the users in
	// their tenant.
	Annotations *[]ProfileCardAnnotation `json:"annotations,omitempty"`

	// Identifies a profileCardProperty resource in Get, Update, or Delete operations. Allows an administrator to surface
	// hidden Microsoft Entra ID properties on the Microsoft 365 profile card within their tenant. When present, the
	// Microsoft Entra ID field referenced in this property is visible to all users in your tenant on the contact pane of
	// the profile card. Allowed values for this field are: UserPrincipalName, Fax, StreetAddress, PostalCode,
	// StateOrProvince, Alias, CustomAttribute1, CustomAttribute2, CustomAttribute3, CustomAttribute4, CustomAttribute5,
	// CustomAttribute6, CustomAttribute7, CustomAttribute8, CustomAttribute9, CustomAttribute10, CustomAttribute11,
	// CustomAttribute12, CustomAttribute13, CustomAttribute14, CustomAttribute15.
	DirectoryPropertyName nullable.Type[string] `json:"directoryPropertyName,omitempty"`

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

func (s ProfileCardProperty) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProfileCardProperty{}

func (s ProfileCardProperty) MarshalJSON() ([]byte, error) {
	type wrapper ProfileCardProperty
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProfileCardProperty: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProfileCardProperty: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.profileCardProperty"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProfileCardProperty: %+v", err)
	}

	return encoded, nil
}
