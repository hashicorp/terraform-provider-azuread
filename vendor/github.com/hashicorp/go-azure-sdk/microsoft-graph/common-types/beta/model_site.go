package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = Site{}

type Site struct {
	// Analytics about the view activities that took place on this site.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`

	// The collection of column definitions reusable across lists under this site.
	Columns *[]ColumnDefinition `json:"columns,omitempty"`

	// The collection of content models applied to this site.
	ContentModels *[]ContentModel `json:"contentModels,omitempty"`

	// The collection of content types defined for this site.
	ContentTypes *[]ContentType `json:"contentTypes,omitempty"`

	Deleted *Deleted `json:"deleted,omitempty"`

	// The full title for the site. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The document processing jobs running on this site.
	DocumentProcessingJobs *[]DocumentProcessingJob `json:"documentProcessingJobs,omitempty"`

	// The default drive (document library) for this site.
	Drive *Drive `json:"drive,omitempty"`

	// The collection of drives (document libraries) under this site.
	Drives *[]Drive `json:"drives,omitempty"`

	// The collection of column definitions available in the site that is referenced from the sites in the parent hierarchy
	// of the current site.
	ExternalColumns *[]ColumnDefinition `json:"externalColumns,omitempty"`

	InformationProtection *InformationProtection `json:"informationProtection,omitempty"`
	IsPersonalSite        nullable.Type[bool]    `json:"isPersonalSite,omitempty"`

	// Used to address any item contained in this site. This collection can't be enumerated.
	Items *[]BaseItem `json:"items,omitempty"`

	// The collection of lists under this site.
	Lists *[]List `json:"lists,omitempty"`

	Onenote *Onenote `json:"onenote,omitempty"`

	// The collection of long running operations for the site.
	Operations *[]RichLongRunningOperation `json:"operations,omitempty"`

	// The collection of page templates on this site.
	PageTemplates *[]PageTemplate `json:"pageTemplates,omitempty"`

	// The collection of pages in the baseSitePages list on this site.
	Pages *[]BaseSitePage `json:"pages,omitempty"`

	// The permissions associated with the site. Nullable.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// A container for a collection of recycleBinItem resources in this site.
	RecycleBin *RecycleBin `json:"recycleBin,omitempty"`

	// If present, provides the root site in the site collection. Read-only.
	Root *Root `json:"root,omitempty"`

	// The settings on this site. Read-only.
	Settings *SiteSettings `json:"settings,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Provides details about the site's site collection. Available only on the root site. Read-only.
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`

	// The collection of the sub-sites under this site.
	Sites *[]Site `json:"sites,omitempty"`

	// The termStore under this site.
	TermStore *TermStoreStore `json:"termStore,omitempty"`

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
	delete(decoded, "root")
	delete(decoded, "settings")
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
		Analytics              *ItemAnalytics              `json:"analytics,omitempty"`
		Columns                *[]ColumnDefinition         `json:"columns,omitempty"`
		ContentModels          *[]ContentModel             `json:"contentModels,omitempty"`
		ContentTypes           *[]ContentType              `json:"contentTypes,omitempty"`
		Deleted                *Deleted                    `json:"deleted,omitempty"`
		DisplayName            nullable.Type[string]       `json:"displayName,omitempty"`
		DocumentProcessingJobs *[]DocumentProcessingJob    `json:"documentProcessingJobs,omitempty"`
		Drive                  *Drive                      `json:"drive,omitempty"`
		Drives                 *[]Drive                    `json:"drives,omitempty"`
		ExternalColumns        *[]ColumnDefinition         `json:"externalColumns,omitempty"`
		InformationProtection  *InformationProtection      `json:"informationProtection,omitempty"`
		IsPersonalSite         nullable.Type[bool]         `json:"isPersonalSite,omitempty"`
		Lists                  *[]List                     `json:"lists,omitempty"`
		Onenote                *Onenote                    `json:"onenote,omitempty"`
		Operations             *[]RichLongRunningOperation `json:"operations,omitempty"`
		PageTemplates          *[]PageTemplate             `json:"pageTemplates,omitempty"`
		Permissions            *[]Permission               `json:"permissions,omitempty"`
		RecycleBin             *RecycleBin                 `json:"recycleBin,omitempty"`
		Root                   *Root                       `json:"root,omitempty"`
		Settings               *SiteSettings               `json:"settings,omitempty"`
		SharepointIds          *SharepointIds              `json:"sharepointIds,omitempty"`
		SiteCollection         *SiteCollection             `json:"siteCollection,omitempty"`
		Sites                  *[]Site                     `json:"sites,omitempty"`
		TermStore              *TermStoreStore             `json:"termStore,omitempty"`
		CreatedByUser          *User                       `json:"createdByUser,omitempty"`
		CreatedDateTime        *string                     `json:"createdDateTime,omitempty"`
		Description            nullable.Type[string]       `json:"description,omitempty"`
		ETag                   nullable.Type[string]       `json:"eTag,omitempty"`
		LastModifiedByUser     *User                       `json:"lastModifiedByUser,omitempty"`
		LastModifiedDateTime   *string                     `json:"lastModifiedDateTime,omitempty"`
		Name                   nullable.Type[string]       `json:"name,omitempty"`
		ParentReference        *ItemReference              `json:"parentReference,omitempty"`
		WebUrl                 nullable.Type[string]       `json:"webUrl,omitempty"`
		Id                     *string                     `json:"id,omitempty"`
		ODataId                *string                     `json:"@odata.id,omitempty"`
		ODataType              *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Analytics = decoded.Analytics
	s.Columns = decoded.Columns
	s.ContentModels = decoded.ContentModels
	s.ContentTypes = decoded.ContentTypes
	s.Deleted = decoded.Deleted
	s.DisplayName = decoded.DisplayName
	s.DocumentProcessingJobs = decoded.DocumentProcessingJobs
	s.Drive = decoded.Drive
	s.Drives = decoded.Drives
	s.ExternalColumns = decoded.ExternalColumns
	s.InformationProtection = decoded.InformationProtection
	s.IsPersonalSite = decoded.IsPersonalSite
	s.Lists = decoded.Lists
	s.Onenote = decoded.Onenote
	s.Operations = decoded.Operations
	s.PageTemplates = decoded.PageTemplates
	s.Permissions = decoded.Permissions
	s.RecycleBin = decoded.RecycleBin
	s.Root = decoded.Root
	s.Settings = decoded.Settings
	s.SharepointIds = decoded.SharepointIds
	s.SiteCollection = decoded.SiteCollection
	s.Sites = decoded.Sites
	s.TermStore = decoded.TermStore
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
