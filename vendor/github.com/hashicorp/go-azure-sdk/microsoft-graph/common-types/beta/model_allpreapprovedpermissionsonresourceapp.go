package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PreApprovedPermissions = AllPreApprovedPermissionsOnResourceApp{}

type AllPreApprovedPermissionsOnResourceApp struct {
	// The appId of the resource application (the API). Required.
	ResourceApplicationId nullable.Type[string] `json:"resourceApplicationId,omitempty"`

	// Fields inherited from PreApprovedPermissions

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the scope of permissions that are included in this condition set. Possible values: all for all permissions,
	// enumerated for a given list of permissions, or allPermissionsOnResourceApp for all permissions from a given API.
	// Required.
	PermissionKind PermissionKind `json:"permissionKind"`

	// The type of permission being granted. Possible values: application for application permissions, or delegated for
	// delegated permissions. Required.
	PermissionType PermissionType `json:"permissionType"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AllPreApprovedPermissionsOnResourceApp) PreApprovedPermissions() BasePreApprovedPermissionsImpl {
	return BasePreApprovedPermissionsImpl{
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		PermissionKind: s.PermissionKind,
		PermissionType: s.PermissionType,
	}
}

var _ json.Marshaler = AllPreApprovedPermissionsOnResourceApp{}

func (s AllPreApprovedPermissionsOnResourceApp) MarshalJSON() ([]byte, error) {
	type wrapper AllPreApprovedPermissionsOnResourceApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AllPreApprovedPermissionsOnResourceApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AllPreApprovedPermissionsOnResourceApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.allPreApprovedPermissionsOnResourceApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AllPreApprovedPermissionsOnResourceApp: %+v", err)
	}

	return encoded, nil
}
