package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PreApprovedPermissions = EnumeratedPreApprovedPermissions{}

type EnumeratedPreApprovedPermissions struct {
	// The list of id values for the specific resource-specific application permissions to match with. See the
	// resourceSpecificApplicationPermissions property of the servicePrincipal object for the list of permissions.
	PermissionIds *[]string `json:"permissionIds,omitempty"`

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

func (s EnumeratedPreApprovedPermissions) PreApprovedPermissions() BasePreApprovedPermissionsImpl {
	return BasePreApprovedPermissionsImpl{
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		PermissionKind: s.PermissionKind,
		PermissionType: s.PermissionType,
	}
}

var _ json.Marshaler = EnumeratedPreApprovedPermissions{}

func (s EnumeratedPreApprovedPermissions) MarshalJSON() ([]byte, error) {
	type wrapper EnumeratedPreApprovedPermissions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnumeratedPreApprovedPermissions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnumeratedPreApprovedPermissions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enumeratedPreApprovedPermissions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnumeratedPreApprovedPermissions: %+v", err)
	}

	return encoded, nil
}
