package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BaseItem = DriveItem{}

type DriveItem struct {
	// The list of recent activities that took place on this item.
	Activities *[]ItemActivityOLD `json:"activities,omitempty"`

	// Analytics about the view activities that took place on this item.
	Analytics *ItemAnalytics `json:"analytics,omitempty"`

	// Audio metadata, if the item is an audio file. Read-only. Only on OneDrive Personal.
	Audio *Audio `json:"audio,omitempty"`

	// Bundle metadata, if the item is a bundle. Read-only.
	Bundle *Bundle `json:"bundle,omitempty"`

	// An eTag for the content of the item. This eTag isn't changed if only the metadata is changed. Note This property
	// isn't returned if the item is a folder. Read-only.
	CTag nullable.Type[string] `json:"cTag,omitempty"`

	// Collection containing Item objects for the immediate children of Item. Only items representing folders have children.
	// Read-only. Nullable.
	Children *[]DriveItem `json:"children,omitempty"`

	// The content stream, if the item represents a file. The content property will have a potentially breaking change in
	// behavior in the future. It will stream content directly instead of redirecting. To proactively opt in to the new
	// behavior ahead of time, use the contentStream property instead.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The content stream, if the item represents a file.
	ContentStream nullable.Type[string] `json:"contentStream,omitempty"`

	// Information about the deleted state of the item. Read-only.
	Deleted *Deleted `json:"deleted,omitempty"`

	// File metadata, if the item is a file. Read-only.
	File *File `json:"file,omitempty"`

	// File system information on client. Read-write.
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`

	// Folder metadata, if the item is a folder. Read-only.
	Folder *Folder `json:"folder,omitempty"`

	// Image metadata, if the item is an image. Read-only.
	Image *Image `json:"image,omitempty"`

	// For drives in SharePoint, the associated document library list item. Read-only. Nullable.
	ListItem *ListItem `json:"listItem,omitempty"`

	// Location metadata, if the item has location data. Read-only.
	Location *GeoCoordinates `json:"location,omitempty"`

	// Malware metadata, if the item was detected to contain malware. Read-only.
	Malware *Malware `json:"malware,omitempty"`

	// Information about the media (audio or video) item. Read-write. Only on OneDrive for Business and SharePoint.
	Media *Media `json:"media,omitempty"`

	// If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in
	// some contexts and folders in others. Read-only.
	Package *Package `json:"package,omitempty"`

	// If present, indicates that indicates that one or more operations that might affect the state of the driveItem are
	// pending completion. Read-only.
	PendingOperations *PendingOperations `json:"pendingOperations,omitempty"`

	// The set of permissions for the item. Read-only. Nullable.
	Permissions *[]Permission `json:"permissions,omitempty"`

	// Photo metadata, if the item is a photo. Read-only.
	Photo *Photo `json:"photo,omitempty"`

	// Provides information about the published or checked-out state of an item, in locations that support such actions.
	// This property isn't returned by default. Read-only.
	Publication *PublicationFacet `json:"publication,omitempty"`

	// Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.
	RemoteItem *RemoteItem `json:"remoteItem,omitempty"`

	// Information about retention label and settings enforced on the driveItem. Read-write.
	RetentionLabel *ItemRetentionLabel `json:"retentionLabel,omitempty"`

	// If this property is non-null, it indicates that the driveItem is the top-most driveItem in the drive.
	Root *Root `json:"root,omitempty"`

	// Search metadata, if the item is from a search result. Read-only.
	SearchResult *SearchResult `json:"searchResult,omitempty"`

	// Indicates that the item was shared with others and provides information about the shared state of the item.
	// Read-only.
	Shared *Shared `json:"shared,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Size of the item in bytes. Read-only.
	Size nullable.Type[int64] `json:"size,omitempty"`

	// Information about the drive item source. Read-only. Only on OneDrive for Business and SharePoint.
	Source *DriveItemSource `json:"source,omitempty"`

	// If the current item is also available as a special folder, this facet is returned. Read-only.
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`

	// The set of subscriptions on the item. Only supported on the root of a drive.
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`

	// Collection of thumbnailSet objects associated with the item. For more information, see getting thumbnails. Read-only.
	// Nullable.
	Thumbnails *[]ThumbnailSet `json:"thumbnails,omitempty"`

	// The list of previous versions of the item. For more info, see getting previous versions. Read-only. Nullable.
	Versions *[]DriveItemVersion `json:"versions,omitempty"`

	// Video metadata, if the item is a video. Read-only.
	Video *Video `json:"video,omitempty"`

	// Returns information specific to the calling user for this drive item. Read-only.
	Viewpoint *DriveItemViewpoint `json:"viewpoint,omitempty"`

	// WebDAV compatible URL for the item.
	WebDavUrl nullable.Type[string] `json:"webDavUrl,omitempty"`

	// For files that are Excel spreadsheets, access to the workbook API to work with the spreadsheet's contents. Nullable.
	Workbook *Workbook `json:"workbook,omitempty"`

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

func (s DriveItem) BaseItem() BaseBaseItemImpl {
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

func (s DriveItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DriveItem{}

func (s DriveItem) MarshalJSON() ([]byte, error) {
	type wrapper DriveItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DriveItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DriveItem: %+v", err)
	}

	delete(decoded, "audio")
	delete(decoded, "bundle")
	delete(decoded, "cTag")
	delete(decoded, "children")
	delete(decoded, "deleted")
	delete(decoded, "file")
	delete(decoded, "folder")
	delete(decoded, "image")
	delete(decoded, "listItem")
	delete(decoded, "location")
	delete(decoded, "malware")
	delete(decoded, "package")
	delete(decoded, "pendingOperations")
	delete(decoded, "permissions")
	delete(decoded, "photo")
	delete(decoded, "publication")
	delete(decoded, "remoteItem")
	delete(decoded, "searchResult")
	delete(decoded, "shared")
	delete(decoded, "sharepointIds")
	delete(decoded, "size")
	delete(decoded, "source")
	delete(decoded, "specialFolder")
	delete(decoded, "thumbnails")
	delete(decoded, "versions")
	delete(decoded, "video")
	delete(decoded, "viewpoint")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.driveItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DriveItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DriveItem{}

func (s *DriveItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Activities           *[]ItemActivityOLD    `json:"activities,omitempty"`
		Analytics            *ItemAnalytics        `json:"analytics,omitempty"`
		Audio                *Audio                `json:"audio,omitempty"`
		Bundle               *Bundle               `json:"bundle,omitempty"`
		CTag                 nullable.Type[string] `json:"cTag,omitempty"`
		Children             *[]DriveItem          `json:"children,omitempty"`
		Content              nullable.Type[string] `json:"content,omitempty"`
		ContentStream        nullable.Type[string] `json:"contentStream,omitempty"`
		Deleted              *Deleted              `json:"deleted,omitempty"`
		File                 *File                 `json:"file,omitempty"`
		FileSystemInfo       *FileSystemInfo       `json:"fileSystemInfo,omitempty"`
		Folder               *Folder               `json:"folder,omitempty"`
		Image                *Image                `json:"image,omitempty"`
		ListItem             *ListItem             `json:"listItem,omitempty"`
		Location             *GeoCoordinates       `json:"location,omitempty"`
		Malware              *Malware              `json:"malware,omitempty"`
		Media                *Media                `json:"media,omitempty"`
		Package              *Package              `json:"package,omitempty"`
		PendingOperations    *PendingOperations    `json:"pendingOperations,omitempty"`
		Permissions          *[]Permission         `json:"permissions,omitempty"`
		Photo                *Photo                `json:"photo,omitempty"`
		Publication          *PublicationFacet     `json:"publication,omitempty"`
		RemoteItem           *RemoteItem           `json:"remoteItem,omitempty"`
		RetentionLabel       *ItemRetentionLabel   `json:"retentionLabel,omitempty"`
		Root                 *Root                 `json:"root,omitempty"`
		SearchResult         *SearchResult         `json:"searchResult,omitempty"`
		Shared               *Shared               `json:"shared,omitempty"`
		SharepointIds        *SharepointIds        `json:"sharepointIds,omitempty"`
		Size                 nullable.Type[int64]  `json:"size,omitempty"`
		Source               *DriveItemSource      `json:"source,omitempty"`
		SpecialFolder        *SpecialFolder        `json:"specialFolder,omitempty"`
		Subscriptions        *[]Subscription       `json:"subscriptions,omitempty"`
		Thumbnails           *[]ThumbnailSet       `json:"thumbnails,omitempty"`
		Versions             *[]DriveItemVersion   `json:"versions,omitempty"`
		Video                *Video                `json:"video,omitempty"`
		Viewpoint            *DriveItemViewpoint   `json:"viewpoint,omitempty"`
		WebDavUrl            nullable.Type[string] `json:"webDavUrl,omitempty"`
		Workbook             *Workbook             `json:"workbook,omitempty"`
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
	s.Analytics = decoded.Analytics
	s.Audio = decoded.Audio
	s.Bundle = decoded.Bundle
	s.CTag = decoded.CTag
	s.Children = decoded.Children
	s.Content = decoded.Content
	s.ContentStream = decoded.ContentStream
	s.Deleted = decoded.Deleted
	s.File = decoded.File
	s.FileSystemInfo = decoded.FileSystemInfo
	s.Folder = decoded.Folder
	s.Image = decoded.Image
	s.ListItem = decoded.ListItem
	s.Location = decoded.Location
	s.Malware = decoded.Malware
	s.Media = decoded.Media
	s.Package = decoded.Package
	s.PendingOperations = decoded.PendingOperations
	s.Permissions = decoded.Permissions
	s.Photo = decoded.Photo
	s.Publication = decoded.Publication
	s.RemoteItem = decoded.RemoteItem
	s.RetentionLabel = decoded.RetentionLabel
	s.Root = decoded.Root
	s.SearchResult = decoded.SearchResult
	s.Shared = decoded.Shared
	s.SharepointIds = decoded.SharepointIds
	s.Size = decoded.Size
	s.Source = decoded.Source
	s.SpecialFolder = decoded.SpecialFolder
	s.Subscriptions = decoded.Subscriptions
	s.Thumbnails = decoded.Thumbnails
	s.Versions = decoded.Versions
	s.Video = decoded.Video
	s.Viewpoint = decoded.Viewpoint
	s.WebDavUrl = decoded.WebDavUrl
	s.Workbook = decoded.Workbook
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
		return fmt.Errorf("unmarshaling DriveItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'DriveItem': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'DriveItem': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
