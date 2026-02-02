package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = ResourceSpecificPermissionGrant{}

type ResourceSpecificPermissionGrant struct {
	// ID of the service principal of the Microsoft Entra app that has been granted access. Read-only.
	ClientAppId nullable.Type[string] `json:"clientAppId,omitempty"`

	// ID of the Microsoft Entra app that has been granted access. Read-only.
	ClientId nullable.Type[string] `json:"clientId,omitempty"`

	// The name of the resource-specific permission. Read-only.
	Permission nullable.Type[string] `json:"permission,omitempty"`

	// The type of permission. Possible values are: Application, Delegated. Read-only.
	PermissionType nullable.Type[string] `json:"permissionType,omitempty"`

	// ID of the Microsoft Entra app that is hosting the resource. Read-only.
	ResourceAppId nullable.Type[string] `json:"resourceAppId,omitempty"`

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

func (s ResourceSpecificPermissionGrant) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ResourceSpecificPermissionGrant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ResourceSpecificPermissionGrant{}

func (s ResourceSpecificPermissionGrant) MarshalJSON() ([]byte, error) {
	type wrapper ResourceSpecificPermissionGrant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceSpecificPermissionGrant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceSpecificPermissionGrant: %+v", err)
	}

	delete(decoded, "clientAppId")
	delete(decoded, "clientId")
	delete(decoded, "permission")
	delete(decoded, "permissionType")
	delete(decoded, "resourceAppId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.resourceSpecificPermissionGrant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceSpecificPermissionGrant: %+v", err)
	}

	return encoded, nil
}
