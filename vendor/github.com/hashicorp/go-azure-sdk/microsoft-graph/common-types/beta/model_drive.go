package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = Drive{}

type Drive struct {
	// The list of recent activities that took place under this drive.
	Activities *[]ItemActivityOLD `json:"activities,omitempty"`

	// Collection of bundles (albums and multi-select-shared sets of items). Only in personal OneDrive.
	Bundles *[]DriveItem `json:"bundles,omitempty"`

	// Describes the type of drive represented by this resource. OneDrive personal drives return personal. OneDrive for
	// Business returns business. SharePoint document libraries return documentLibrary. Read-only.
	DriveType nullable.Type[string] `json:"driveType,omitempty"`

	// The list of items the user is following. Only in OneDrive for Business.
	Following *[]DriveItem `json:"following,omitempty"`

	// All items contained in the drive. Read-only. Nullable.
	Items *[]DriveItem `json:"items,omitempty"`

	// For drives in SharePoint, the underlying document library list. Read-only. Nullable.
	List *List `json:"list,omitempty"`

	// Optional. The user account that owns the drive. Read-only.
	Owner *IdentitySet `json:"owner,omitempty"`

	// Optional. Information about the drive's storage space quota. Read-only.
	Quota *Quota `json:"quota,omitempty"`

	// The root folder of the drive. Read-only.
	Root *DriveItem `json:"root,omitempty"`

	SharePointIds *SharepointIds `json:"sharePointIds,omitempty"`

	// Collection of common folders available in OneDrive. Read-only. Nullable.
	Special *[]DriveItem `json:"special,omitempty"`

	// If present, indicates that this is a system-managed drive. Read-only.
	System *SystemFacet `json:"system,omitempty"`

	// Fields inherited from BaseItem

	// Identity of the user, device, or application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	CreatedByUser *User `json:"createdByUser,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The description of the item.
	Description nullable.Type[string] `json:"description,omitempty"`

	// ETag for the item. Read-only.
	ETag nullable.Type[string] `json:"eTag,omitempty"`

	// Identity of the user, device, and application that last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The name of the item. Read-write.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// URL that either displays the resource in the browser (for Office file formats), or is a direct link to the file (for
	// other formats). Read-only.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`

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

func (s Drive) BaseItem() BaseBaseItemImpl {
	return BaseBaseItemImpl{
		CreatedBy:            s.CreatedBy,
		CreatedByUser:        s.CreatedByUser,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		ETag:                 s.ETag,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedByUser:   s.LastModifiedByUser,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Name:                 s.Name,
		ParentReference:      s.ParentReference,
		WebUrl:               s.WebUrl,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s Drive) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Drive{}

func (s Drive) MarshalJSON() ([]byte, error) {
	type wrapper Drive
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Drive: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Drive: %+v", err)
	}

	delete(decoded, "driveType")
	delete(decoded, "items")
	delete(decoded, "list")
	delete(decoded, "owner")
	delete(decoded, "quota")
	delete(decoded, "root")
	delete(decoded, "special")
	delete(decoded, "system")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.drive"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Drive: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Drive{}

func (s *Drive) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Activities           *[]ItemActivityOLD    `json:"activities,omitempty"`
		Bundles              *[]DriveItem          `json:"bundles,omitempty"`
		DriveType            nullable.Type[string] `json:"driveType,omitempty"`
		Following            *[]DriveItem          `json:"following,omitempty"`
		Items                *[]DriveItem          `json:"items,omitempty"`
		List                 *List                 `json:"list,omitempty"`
		Quota                *Quota                `json:"quota,omitempty"`
		Root                 *DriveItem            `json:"root,omitempty"`
		SharePointIds        *SharepointIds        `json:"sharePointIds,omitempty"`
		Special              *[]DriveItem          `json:"special,omitempty"`
		System               *SystemFacet          `json:"system,omitempty"`
		CreatedByUser        *User                 `json:"createdByUser,omitempty"`
		CreatedDateTime      *string               `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string] `json:"description,omitempty"`
		ETag                 nullable.Type[string] `json:"eTag,omitempty"`
		LastModifiedByUser   *User                 `json:"lastModifiedByUser,omitempty"`
		LastModifiedDateTime *string               `json:"lastModifiedDateTime,omitempty"`
		Name                 nullable.Type[string] `json:"name,omitempty"`
		ParentReference      *ItemReference        `json:"parentReference,omitempty"`
		WebUrl               nullable.Type[string] `json:"webUrl,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Activities = decoded.Activities
	s.Bundles = decoded.Bundles
	s.DriveType = decoded.DriveType
	s.Following = decoded.Following
	s.Items = decoded.Items
	s.List = decoded.List
	s.Quota = decoded.Quota
	s.Root = decoded.Root
	s.SharePointIds = decoded.SharePointIds
	s.Special = decoded.Special
	s.System = decoded.System
	s.CreatedByUser = decoded.CreatedByUser
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.ETag = decoded.ETag
	s.Id = decoded.Id
	s.LastModifiedByUser = decoded.LastModifiedByUser
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParentReference = decoded.ParentReference
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Drive into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'Drive': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'Drive': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	if v, ok := temp["owner"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Owner' for 'Drive': %+v", err)
		}
		s.Owner = &impl
	}

	return nil
}
