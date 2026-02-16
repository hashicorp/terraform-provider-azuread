package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageResourceScope{}

type AccessPackageResourceScope struct {
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`

	// The description of the scope.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the scope.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// True if the scopes are arranged in a hierarchy and this is the top or root scope of the resource.
	IsRootScope nullable.Type[bool] `json:"isRootScope,omitempty"`

	// The unique identifier for the scope in the resource as defined in the origin system.
	OriginId nullable.Type[string] `json:"originId,omitempty"`

	// The origin system for the scope.
	OriginSystem nullable.Type[string] `json:"originSystem,omitempty"`

	// The origin system for the role, if different.
	RoleOriginId nullable.Type[string] `json:"roleOriginId,omitempty"`

	// A resource locator for the scope.
	Url nullable.Type[string] `json:"url,omitempty"`

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

func (s AccessPackageResourceScope) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageResourceScope{}

func (s AccessPackageResourceScope) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageResourceScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageResourceScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageResourceScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageResourceScope: %+v", err)
	}

	return encoded, nil
}
