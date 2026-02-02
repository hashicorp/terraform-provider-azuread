package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemReference struct {
	// Unique identifier of the drive instance that contains the driveItem. Only returned if the item is located in a drive.
	// Read-only.
	DriveId nullable.Type[string] `json:"driveId,omitempty"`

	// Identifies the type of drive. Only returned if the item is located in a drive. See drive resource for values.
	DriveType nullable.Type[string] `json:"driveType,omitempty"`

	// Unique identifier of the driveItem in the drive or a listItem in a list. Read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The name of the item being referenced. Read-only.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Percent-encoded path that can be used to navigate to the item. Read-only.
	Path nullable.Type[string] `json:"path,omitempty"`

	// A unique identifier for a shared resource that can be accessed via the Shares API.
	ShareId nullable.Type[string] `json:"shareId,omitempty"`

	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds *SharepointIds `json:"sharepointIds,omitempty"`

	// For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent
	// document library of the driveItem resource or the parent list of the listItem resource. The value is the same as the
	// id property of that site resource. It's an opaque string that consists of three identifiers of the site. For
	// OneDrive, this property isn't populated.
	SiteId nullable.Type[string] `json:"siteId,omitempty"`
}

var _ json.Marshaler = ItemReference{}

func (s ItemReference) MarshalJSON() ([]byte, error) {
	type wrapper ItemReference
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemReference: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemReference: %+v", err)
	}

	delete(decoded, "driveId")
	delete(decoded, "id")
	delete(decoded, "name")
	delete(decoded, "path")
	delete(decoded, "sharepointIds")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemReference: %+v", err)
	}

	return encoded, nil
}
