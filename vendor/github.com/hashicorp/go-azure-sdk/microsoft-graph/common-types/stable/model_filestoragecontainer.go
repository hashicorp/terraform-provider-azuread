package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FileStorageContainer{}

type FileStorageContainer struct {
	// Container type ID of the fileStorageContainer. For details about container types, see Container Types. Each container
	// must have only one container type. Read-only.
	ContainerTypeId *string `json:"containerTypeId,omitempty"`

	// Date and time of the fileStorageContainer creation. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Custom property collection for the fileStorageContainer. Read-write.
	CustomProperties *FileStorageContainerCustomPropertyDictionary `json:"customProperties,omitempty"`

	// Provides a user-visible description of the fileStorageContainer. Read-write.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the fileStorageContainer. Read-write.
	DisplayName *string `json:"displayName,omitempty"`

	// The drive of the resource fileStorageContainer. Read-only.
	Drive *Drive `json:"drive,omitempty"`

	// The set of permissions for users in the fileStorageContainer. Permission for each user is set by the roles property.
	// The possible values are: reader, writer, manager, and owner. Read-write.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Status of the fileStorageContainer. Containers are created as inactive and require activation. Inactive containers
	// are subjected to automatic deletion in 24 hours. The possible values are: inactive, active. Read-only.
	Status *FileStorageContainerStatus `json:"status,omitempty"`

	// Data specific to the current user. Read-only.
	Viewpoint *FileStorageContainerViewpoint `json:"viewpoint,omitempty"`

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

func (s FileStorageContainer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FileStorageContainer{}

func (s FileStorageContainer) MarshalJSON() ([]byte, error) {
	type wrapper FileStorageContainer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileStorageContainer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileStorageContainer: %+v", err)
	}

	delete(decoded, "containerTypeId")
	delete(decoded, "createdDateTime")
	delete(decoded, "drive")
	delete(decoded, "status")
	delete(decoded, "viewpoint")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fileStorageContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileStorageContainer: %+v", err)
	}

	return encoded, nil
}
