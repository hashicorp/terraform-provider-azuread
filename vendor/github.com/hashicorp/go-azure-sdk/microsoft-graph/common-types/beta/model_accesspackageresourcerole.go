package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResourceRole{}

type AccessPackageResourceRole struct {
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`

	// A description for the resource role.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the resource role such as the role defined by the application.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The unique identifier of the resource role in the origin system. For a SharePoint Online site, the originId is the
	// sequence number of the role in the site.
	OriginId nullable.Type[string] `json:"originId,omitempty"`

	// The type of the resource in the origin system, such as SharePointOnline, AadApplication or AadGroup.
	OriginSystem nullable.Type[string] `json:"originSystem,omitempty"`

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

func (s AccessPackageResourceRole) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResourceRole{}

func (s AccessPackageResourceRole) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceRole
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceRole: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceRole: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceRole"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceRole: %+v", err)
	}

	return encoded, nil
}
