package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = List{}

type List struct {
	// The recent activities that took place within this list.
	Activities *[]ItemActivityOLD `json:"activities,omitempty"`

	// The collection of field definitions for this list.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`

	// The collection of content types present in this list.
	ContentTypes *[]ContentType `json:"contentTypes,omitempty"`

	// The displayable title of the list.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Allows access to the list as a drive resource with driveItems. Only present on document libraries.
	Drive *Drive `json:"drive,omitempty"`

	// All items contained in the list.
	Items *[]ListItem `json:"items,omitempty"`

	// Contains more details about the list.
	List *ListInfo `json:"list,omitempty"`

	// The collection of long-running operations on the list.
	Operations *[]RichLongRunningOperation `json:"operations,omitempty"`

	// The set of permissions for the item. Read-only. Nullable.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// The set of subscriptions on the list.
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`

	// If present, indicates that the list is system-managed. Read-only.
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

func (s List) BaseItem() BaseBaseItemImpl {
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

func (s List) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = List{}

func (s List) MarshalJSON() ([]byte, error) {
	type wrapper List
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling List: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling List: %+v", err)
	}

	delete(decoded, "permissions")
	delete(decoded, "sharepointIds")
	delete(decoded, "system")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.list"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling List: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &List{}

func (s *List) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Activities           *[]ItemActivityOLD          `json:"activities,omitempty"`
		Columns              *[]ColumnDefinition         `json:"columns,omitempty"`
		ContentTypes         *[]ContentType              `json:"contentTypes,omitempty"`
		DisplayName          nullable.Type[string]       `json:"displayName,omitempty"`
		Drive                *Drive                      `json:"drive,omitempty"`
		Items                *[]ListItem                 `json:"items,omitempty"`
		List                 *ListInfo                   `json:"list,omitempty"`
		Operations           *[]RichLongRunningOperation `json:"operations,omitempty"`
		Permissions          *[]Permission               `json:"permissions,omitempty"`
		SharepointIds        *SharepointIds              `json:"sharepointIds,omitempty"`
		Subscriptions        *[]Subscription             `json:"subscriptions,omitempty"`
		System               *SystemFacet                `json:"system,omitempty"`
		CreatedByUser        *User                       `json:"createdByUser,omitempty"`
		CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
		Description          nullable.Type[string]       `json:"description,omitempty"`
		ETag                 nullable.Type[string]       `json:"eTag,omitempty"`
		LastModifiedByUser   *User                       `json:"lastModifiedByUser,omitempty"`
		LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
		Name                 nullable.Type[string]       `json:"name,omitempty"`
		ParentReference      *ItemReference              `json:"parentReference,omitempty"`
		WebUrl               nullable.Type[string]       `json:"webUrl,omitempty"`
		Id                   *string                     `json:"id,omitempty"`
		ODataId              *string                     `json:"@odata.id,omitempty"`
		ODataType            *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Activities = decoded.Activities
	s.Columns = decoded.Columns
	s.ContentTypes = decoded.ContentTypes
	s.DisplayName = decoded.DisplayName
	s.Drive = decoded.Drive
	s.Items = decoded.Items
	s.List = decoded.List
	s.Operations = decoded.Operations
	s.Permissions = decoded.Permissions
	s.SharepointIds = decoded.SharepointIds
	s.Subscriptions = decoded.Subscriptions
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
		return fmt.Errorf("unmarshaling List into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'List': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'List': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
