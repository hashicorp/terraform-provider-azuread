package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ContactFolder{}

type ContactFolder struct {
	// The collection of child folders in the folder. Navigation property. Read-only. Nullable.
	ChildFolders *[]ContactFolder `json:"childFolders,omitempty"`

	// The contacts in the folder. Navigation property. Read-only. Nullable.
	Contacts *[]Contact `json:"contacts,omitempty"`

	// The folder's display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The ID of the folder's parent folder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The name of the folder if the folder is a recognized folder. Currently contacts is the only recognized contacts
	// folder.
	WellKnownName nullable.Type[string] `json:"wellKnownName,omitempty"`

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

func (s ContactFolder) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ContactFolder{}

func (s ContactFolder) MarshalJSON() ([]byte, error) {
	type wrapper ContactFolder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ContactFolder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ContactFolder: %+v", err)
	}

	delete(decoded, "childFolders")
	delete(decoded, "contacts")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.contactFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ContactFolder: %+v", err)
	}

	return encoded, nil
}
