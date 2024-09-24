package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteItem struct {
	// Identity of the user, device, and application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// Date and time of item creation. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates that the remote item is a file. Read-only.
	File *File `json:"file,omitempty"`

	// Information about the remote item from the local file system. Read-only.
	FileSystemInfo *FileSystemInfo `json:"fileSystemInfo,omitempty"`

	// Indicates that the remote item is a folder. Read-only.
	Folder *Folder `json:"folder,omitempty"`

	// Unique identifier for the remote item in its drive. Read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Image metadata, if the item is an image. Read-only.
	Image *Image `json:"image,omitempty"`

	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Optional. Filename of the remote item. Read-only.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If present, indicates that this item is a package instead of a folder or file. Packages are treated like files in
	// some contexts and folders in others. Read-only.
	Package *Package `json:"package,omitempty"`

	// Properties of the parent of the remote item. Read-only.
	ParentReference *ItemReference `json:"parentReference,omitempty"`

	// Indicates that the item has been shared with others and provides information about the shared state of the item.
	// Read-only.
	Shared *Shared `json:"shared,omitempty"`

	// Provides interop between items in OneDrive for Business and SharePoint with the full set of item identifiers.
	// Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// Size of the remote item. Read-only.
	Size nullable.Type[int64] `json:"size,omitempty"`

	// If the current item is also available as a special folder, this facet is returned. Read-only.
	SpecialFolder *SpecialFolder `json:"specialFolder,omitempty"`

	// Video metadata, if the item is a video. Read-only.
	Video *Video `json:"video,omitempty"`

	// DAV compatible URL for the item.
	WebDavUrl nullable.Type[string] `json:"webDavUrl,omitempty"`

	// URL that displays the resource in the browser. Read-only.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}

var _ json.Marshaler = RemoteItem{}

func (s RemoteItem) MarshalJSON() ([]byte, error) {
	type wrapper RemoteItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RemoteItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RemoteItem: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "file")
	delete(decoded, "fileSystemInfo")
	delete(decoded, "folder")
	delete(decoded, "id")
	delete(decoded, "image")
	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "name")
	delete(decoded, "package")
	delete(decoded, "parentReference")
	delete(decoded, "shared")
	delete(decoded, "sharepointIds")
	delete(decoded, "size")
	delete(decoded, "specialFolder")
	delete(decoded, "video")
	delete(decoded, "webUrl")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RemoteItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &RemoteItem{}

func (s *RemoteItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		File                 *File                 `json:"file,omitempty"`
		FileSystemInfo       *FileSystemInfo       `json:"fileSystemInfo,omitempty"`
		Folder               *Folder               `json:"folder,omitempty"`
		Id                   nullable.Type[string] `json:"id,omitempty"`
		Image                *Image                `json:"image,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Name                 nullable.Type[string] `json:"name,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
		Package              *Package              `json:"package,omitempty"`
		ParentReference      *ItemReference        `json:"parentReference,omitempty"`
		Shared               *Shared               `json:"shared,omitempty"`
		SharepointIds        *SharepointIds        `json:"sharepointIds,omitempty"`
		Size                 nullable.Type[int64]  `json:"size,omitempty"`
		SpecialFolder        *SpecialFolder        `json:"specialFolder,omitempty"`
		Video                *Video                `json:"video,omitempty"`
		WebDavUrl            nullable.Type[string] `json:"webDavUrl,omitempty"`
		WebUrl               nullable.Type[string] `json:"webUrl,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.File = decoded.File
	s.FileSystemInfo = decoded.FileSystemInfo
	s.Folder = decoded.Folder
	s.Id = decoded.Id
	s.Image = decoded.Image
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Package = decoded.Package
	s.ParentReference = decoded.ParentReference
	s.Shared = decoded.Shared
	s.SharepointIds = decoded.SharepointIds
	s.Size = decoded.Size
	s.SpecialFolder = decoded.SpecialFolder
	s.Video = decoded.Video
	s.WebDavUrl = decoded.WebDavUrl
	s.WebUrl = decoded.WebUrl

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RemoteItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'RemoteItem': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'RemoteItem': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
