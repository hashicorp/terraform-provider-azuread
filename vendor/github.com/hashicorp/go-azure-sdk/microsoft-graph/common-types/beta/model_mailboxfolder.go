package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MailboxFolder{}

type MailboxFolder struct {
	// The number of immediate child folders in the current folder.
	ChildFolderCount nullable.Type[int64] `json:"childFolderCount,omitempty"`

	// The collection of child folders in this folder.
	ChildFolders *[]MailboxFolder `json:"childFolders,omitempty"`

	// The display name of the folder.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The collection of items in this folder.
	Items *[]MailboxItem `json:"items,omitempty"`

	// The collection of multi-value extended properties defined for the mailboxFolder.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The unique identifier for the parent folder of this folder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The routing link to the actual underlying mailbox where the folder physically resides. The folder can be accessed
	// using GET {parentMailboxUrl}/folders/{id}, which treats the entire URL as an opaque string. This method is especially
	// important when auto-expanding archiving is enabled for a user's in-place archive mailbox. The user's archive content
	// can span across multiple mailboxes in such scenarios.
	ParentMailboxUrl nullable.Type[string] `json:"parentMailboxUrl,omitempty"`

	// The collection of single-value extended properties defined for the mailboxFolder.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The number of items in the folder.
	TotalItemCount nullable.Type[int64] `json:"totalItemCount,omitempty"`

	// Describes the folder class type.
	Type nullable.Type[string] `json:"type,omitempty"`

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

func (s MailboxFolder) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MailboxFolder{}

func (s MailboxFolder) MarshalJSON() ([]byte, error) {
	type wrapper MailboxFolder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MailboxFolder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MailboxFolder: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailboxFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MailboxFolder: %+v", err)
	}

	return encoded, nil
}
