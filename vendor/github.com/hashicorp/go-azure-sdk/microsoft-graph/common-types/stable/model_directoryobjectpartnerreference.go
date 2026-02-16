package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = DirectoryObjectPartnerReference{}

type DirectoryObjectPartnerReference struct {
	// Description of the object returned. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Name of directory object being returned, like group or application. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The tenant identifier for the partner tenant. Read-only.
	ExternalPartnerTenantId nullable.Type[string] `json:"externalPartnerTenantId,omitempty"`

	// The type of the referenced object in the partner tenant. Read-only.
	ObjectType nullable.Type[string] `json:"objectType,omitempty"`

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

func (s DirectoryObjectPartnerReference) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s DirectoryObjectPartnerReference) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectoryObjectPartnerReference{}

func (s DirectoryObjectPartnerReference) MarshalJSON() ([]byte, error) {
	type wrapper DirectoryObjectPartnerReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectoryObjectPartnerReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectoryObjectPartnerReference: %+v", err)
	}

	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "externalPartnerTenantId")
	delete(decoded, "objectType")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directoryObjectPartnerReference"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectoryObjectPartnerReference: %+v", err)
	}

	return encoded, nil
}
