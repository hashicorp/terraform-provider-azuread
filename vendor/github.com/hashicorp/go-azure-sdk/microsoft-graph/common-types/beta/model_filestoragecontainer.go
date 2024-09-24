package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FileStorageContainer{}

type FileStorageContainer struct {
	// Sensitivity label assigned to the fileStorageContainer. Read-write.
	AssignedSensitivityLabel *AssignedLabel `json:"assignedSensitivityLabel,omitempty"`

	// The set of custom structured metadata supported by the fileStorageContainer. Read-write.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`

	// Container type ID of the fileStorageContainer. Each container must have only one container type. Read-only.
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

	ExternalGroupId nullable.Type[string] `json:"externalGroupId,omitempty"`

	// Indicates whether versioning is enabled for the fileStorageContainer. The setting is applicable to all items in the
	// fileStorageContainer. Read-Write.
	IsItemVersioningEnabled nullable.Type[bool] `json:"isItemVersioningEnabled,omitempty"`

	// Maximum number of major versions allowed for items in the fileStorageContainer. Read-write.
	ItemMajorVersionLimit nullable.Type[int64] `json:"itemMajorVersionLimit,omitempty"`

	// Indicates the lock state of the fileStorageContainer. The possible values are unlocked and lockedReadOnly. Read-only.
	LockState *SiteLockState `json:"lockState,omitempty"`

	// List of users who own the fileStorageContainer. Read-only.
	Owners *[]UserIdentity `json:"owners,omitempty"`

	// Ownership type of the fileStorageContainer.The possible values are: tenantOwned. Read-only.
	OwnershipType *FileStorageContainerOwnershipType `json:"ownershipType,omitempty"`

	// The set of permissions for users in the fileStorageContainer. The permission for each user is set by the roles
	// property. The possible values are 'reader', 'writer', 'manager', and 'owner'. Read-write.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Recycle bin of the fileStorageContainer. Read-only.
	RecycleBin *RecycleBin `json:"recycleBin,omitempty"`

	Settings *FileStorageContainerSettings `json:"settings,omitempty"`

	// Status of the fileStorageContainer. Containers are created as inactive and require activation. Inactive containers
	// are subjected to automatic deletion in 24 hours. The possible values are: inactive, active. Read-only.
	Status *FileStorageContainerStatus `json:"status,omitempty"`

	// Storage used in the fileStorageContainer, in bytes. Read-only.
	StorageUsedInBytes nullable.Type[int64] `json:"storageUsedInBytes,omitempty"`

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
	delete(decoded, "lockState")
	delete(decoded, "owners")
	delete(decoded, "ownershipType")
	delete(decoded, "recycleBin")
	delete(decoded, "status")
	delete(decoded, "storageUsedInBytes")
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

var _ json.Unmarshaler = &FileStorageContainer{}

func (s *FileStorageContainer) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedSensitivityLabel *AssignedLabel                                `json:"assignedSensitivityLabel,omitempty"`
		Columns                  *[]ColumnDefinition                           `json:"columns,omitempty"`
		ContainerTypeId          *string                                       `json:"containerTypeId,omitempty"`
		CreatedDateTime          *string                                       `json:"createdDateTime,omitempty"`
		CustomProperties         *FileStorageContainerCustomPropertyDictionary `json:"customProperties,omitempty"`
		Description              nullable.Type[string]                         `json:"description,omitempty"`
		DisplayName              *string                                       `json:"displayName,omitempty"`
		Drive                    *Drive                                        `json:"drive,omitempty"`
		ExternalGroupId          nullable.Type[string]                         `json:"externalGroupId,omitempty"`
		IsItemVersioningEnabled  nullable.Type[bool]                           `json:"isItemVersioningEnabled,omitempty"`
		ItemMajorVersionLimit    nullable.Type[int64]                          `json:"itemMajorVersionLimit,omitempty"`
		LockState                *SiteLockState                                `json:"lockState,omitempty"`
		OwnershipType            *FileStorageContainerOwnershipType            `json:"ownershipType,omitempty"`
		Permissions              *[]Permission                                 `json:"permissions,omitempty"`
		RecycleBin               *RecycleBin                                   `json:"recycleBin,omitempty"`
		Settings                 *FileStorageContainerSettings                 `json:"settings,omitempty"`
		Status                   *FileStorageContainerStatus                   `json:"status,omitempty"`
		StorageUsedInBytes       nullable.Type[int64]                          `json:"storageUsedInBytes,omitempty"`
		Viewpoint                *FileStorageContainerViewpoint                `json:"viewpoint,omitempty"`
		Id                       *string                                       `json:"id,omitempty"`
		ODataId                  *string                                       `json:"@odata.id,omitempty"`
		ODataType                *string                                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedSensitivityLabel = decoded.AssignedSensitivityLabel
	s.Columns = decoded.Columns
	s.ContainerTypeId = decoded.ContainerTypeId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomProperties = decoded.CustomProperties
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Drive = decoded.Drive
	s.ExternalGroupId = decoded.ExternalGroupId
	s.IsItemVersioningEnabled = decoded.IsItemVersioningEnabled
	s.ItemMajorVersionLimit = decoded.ItemMajorVersionLimit
	s.LockState = decoded.LockState
	s.OwnershipType = decoded.OwnershipType
	s.Permissions = decoded.Permissions
	s.RecycleBin = decoded.RecycleBin
	s.Settings = decoded.Settings
	s.Status = decoded.Status
	s.StorageUsedInBytes = decoded.StorageUsedInBytes
	s.Viewpoint = decoded.Viewpoint
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FileStorageContainer into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["owners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Owners into list []json.RawMessage: %+v", err)
		}

		output := make([]UserIdentity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Owners' for 'FileStorageContainer': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Owners = &output
	}

	return nil
}
