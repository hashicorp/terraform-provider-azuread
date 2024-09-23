package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = Site{}

type Site struct {
	// Analytics about the view activities that took place on this site.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`

	// The collection of column definitions reusable across lists under this site.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`

	// The collection of content types defined for this site.
	ContentTypes *[]ContentType `json:"contentTypes,omitempty"`

	// The full title for the site. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The default drive (document library) for this site.
	Drive *Drive `json:"drive,omitempty"`

	// The collection of drives (document libraries) under this site.
	Drives *[]Drive `json:"drives,omitempty"`

	Error           *PublicError        `json:"error,omitempty"`
	ExternalColumns *[]ColumnDefinition `json:"externalColumns,omitempty"`

	// Identifies whether the site is personal or not. Read-only.
	IsPersonalSite nullable.Type[bool] `json:"isPersonalSite,omitempty"`

	// Used to address any item contained in this site. This collection can't be enumerated.
	Items *[]BaseItem `json:"items,omitempty"`

	// The collection of lists under this site.
	Lists *[]List `json:"lists,omitempty"`

	// Calls the OneNote service for notebook related operations.
	Onenote *Onenote `json:"onenote,omitempty"`

	// The collection of long-running operations on the site.
	Operations *[]RichLongRunningOperation `json:"operations,omitempty"`

	// The collection of pages in the baseSitePages list in this site.
	Pages *[]BaseSitePage `json:"pages,omitempty"`

	// The permissions associated with the site. Nullable.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// If present, provides the root site in the site collection. Read-only.
	Root *Root `json:"root,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Provides details about the site's site collection. Available only on the root site. Read-only.
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`

	// The collection of the sub-sites under this site.
	Sites *[]Site `json:"sites,omitempty"`

	// The default termStore under this site.
	TermStore *TermStoreStore `json:"termStore,omitempty"`

	// The collection of termStores under this site.
	TermStores *[]TermStoreStore `json:"termStores,omitempty"`

	// Fields inherited from BaseItem

	// Identity of the user, device, or application that created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Provides a user-visible description of the item. Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// ETag for the item. Read-only.
	ETag nullable.Type[string] `json:"eTag,omitempty"`

	// Identity of the user, device, and application that last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Identity of the user who last modified the item. Read-only.
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

func (s Site) BaseItem() BaseBaseItemImpl {
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

func (s Site) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Site{}

func (s Site) MarshalJSON() ([]byte, error) {
	type wrapper Site
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Site: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Site: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "isPersonalSite")
	delete(decoded, "root")
	delete(decoded, "sharepointIds")
	delete(decoded, "siteCollection")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.site"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Site: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Site{}

func (s *Site) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Analytics            *ItemAnalytics              `json:"analytics,omitempty"`
		Columns              *[]ColumnDefinition         `json:"columns,omitempty"`
		ContentTypes         *[]ContentType              `json:"contentTypes,omitempty"`
		DisplayName          nullable.Type[string]       `json:"displayName,omitempty"`
		Drive                *Drive                      `json:"drive,omitempty"`
		Drives               *[]Drive                    `json:"drives,omitempty"`
		Error                *PublicError                `json:"error,omitempty"`
		ExternalColumns      *[]ColumnDefinition         `json:"externalColumns,omitempty"`
		IsPersonalSite       nullable.Type[bool]         `json:"isPersonalSite,omitempty"`
		Lists                *[]List                     `json:"lists,omitempty"`
		Onenote              *Onenote                    `json:"onenote,omitempty"`
		Operations           *[]RichLongRunningOperation `json:"operations,omitempty"`
		Permissions          *[]Permission               `json:"permissions,omitempty"`
		Root                 *Root                       `json:"root,omitempty"`
		SharepointIds        *SharepointIds              `json:"sharepointIds,omitempty"`
		SiteCollection       *SiteCollection             `json:"siteCollection,omitempty"`
		Sites                *[]Site                     `json:"sites,omitempty"`
		TermStore            *TermStoreStore             `json:"termStore,omitempty"`
		TermStores           *[]TermStoreStore           `json:"termStores,omitempty"`
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

	s.Analytics = decoded.Analytics
	s.Columns = decoded.Columns
	s.ContentTypes = decoded.ContentTypes
	s.DisplayName = decoded.DisplayName
	s.Drive = decoded.Drive
	s.Drives = decoded.Drives
	s.Error = decoded.Error
	s.ExternalColumns = decoded.ExternalColumns
	s.IsPersonalSite = decoded.IsPersonalSite
	s.Lists = decoded.Lists
	s.Onenote = decoded.Onenote
	s.Operations = decoded.Operations
	s.Permissions = decoded.Permissions
	s.Root = decoded.Root
	s.SharepointIds = decoded.SharepointIds
	s.SiteCollection = decoded.SiteCollection
	s.Sites = decoded.Sites
	s.TermStore = decoded.TermStore
	s.TermStores = decoded.TermStores
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
		return fmt.Errorf("unmarshaling Site into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'Site': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["items"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Items into list []json.RawMessage: %+v", err)
		}

		output := make([]BaseItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalBaseItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Items' for 'Site': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Items = &output
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'Site': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	if v, ok := temp["pages"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Pages into list []json.RawMessage: %+v", err)
		}

		output := make([]BaseSitePage, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalBaseSitePageImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Pages' for 'Site': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Pages = &output
	}

	return nil
}
