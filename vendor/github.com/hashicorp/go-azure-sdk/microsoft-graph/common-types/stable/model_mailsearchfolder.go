package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MailFolder = MailSearchFolder{}

type MailSearchFolder struct {
	// The OData query to filter the messages.
	FilterQuery nullable.Type[string] `json:"filterQuery,omitempty"`

	// Indicates how the mailbox folder hierarchy should be traversed in the search. true means that a deep search should be
	// done to include child folders in the hierarchy of each folder explicitly specified in sourceFolderIds. false means a
	// shallow search of only each of the folders explicitly specified in sourceFolderIds.
	IncludeNestedFolders nullable.Type[bool] `json:"includeNestedFolders,omitempty"`

	// Indicates whether a search folder is editable using REST APIs.
	IsSupported nullable.Type[bool] `json:"isSupported,omitempty"`

	// The mailbox folders that should be mined.
	SourceFolderIds *[]string `json:"sourceFolderIds,omitempty"`

	// Fields inherited from MailFolder

	// The number of immediate child mailFolders in the current mailFolder.
	ChildFolderCount nullable.Type[int64] `json:"childFolderCount,omitempty"`

	// The collection of child folders in the mailFolder.
	ChildFolders *[]MailFolder `json:"childFolders,omitempty"`

	// The mailFolder's display name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more
	// information in Hidden mail folders.
	IsHidden nullable.Type[bool] `json:"isHidden,omitempty"`

	// The collection of rules that apply to the user's Inbox folder.
	MessageRules *[]MessageRule `json:"messageRules,omitempty"`

	// The collection of messages in the mailFolder.
	Messages *[]Message `json:"messages,omitempty"`

	// The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The unique identifier for the mailFolder's parent mailFolder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The number of items in the mailFolder.
	TotalItemCount nullable.Type[int64] `json:"totalItemCount,omitempty"`

	// The number of items in the mailFolder marked as unread.
	UnreadItemCount nullable.Type[int64] `json:"unreadItemCount,omitempty"`

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

func (s MailSearchFolder) MailFolder() BaseMailFolderImpl {
	return BaseMailFolderImpl{
		ChildFolderCount:              s.ChildFolderCount,
		ChildFolders:                  s.ChildFolders,
		DisplayName:                   s.DisplayName,
		IsHidden:                      s.IsHidden,
		MessageRules:                  s.MessageRules,
		Messages:                      s.Messages,
		MultiValueExtendedProperties:  s.MultiValueExtendedProperties,
		ParentFolderId:                s.ParentFolderId,
		SingleValueExtendedProperties: s.SingleValueExtendedProperties,
		TotalItemCount:                s.TotalItemCount,
		UnreadItemCount:               s.UnreadItemCount,
		Id:                            s.Id,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
	}
}

func (s MailSearchFolder) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MailSearchFolder{}

func (s MailSearchFolder) MarshalJSON() ([]byte, error) {
	type wrapper MailSearchFolder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MailSearchFolder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MailSearchFolder: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailSearchFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MailSearchFolder: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MailSearchFolder{}

func (s *MailSearchFolder) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		FilterQuery                   nullable.Type[string]                `json:"filterQuery,omitempty"`
		IncludeNestedFolders          nullable.Type[bool]                  `json:"includeNestedFolders,omitempty"`
		IsSupported                   nullable.Type[bool]                  `json:"isSupported,omitempty"`
		SourceFolderIds               *[]string                            `json:"sourceFolderIds,omitempty"`
		ChildFolderCount              nullable.Type[int64]                 `json:"childFolderCount,omitempty"`
		DisplayName                   nullable.Type[string]                `json:"displayName,omitempty"`
		IsHidden                      nullable.Type[bool]                  `json:"isHidden,omitempty"`
		MessageRules                  *[]MessageRule                       `json:"messageRules,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		ParentFolderId                nullable.Type[string]                `json:"parentFolderId,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		TotalItemCount                nullable.Type[int64]                 `json:"totalItemCount,omitempty"`
		UnreadItemCount               nullable.Type[int64]                 `json:"unreadItemCount,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FilterQuery = decoded.FilterQuery
	s.IncludeNestedFolders = decoded.IncludeNestedFolders
	s.IsSupported = decoded.IsSupported
	s.SourceFolderIds = decoded.SourceFolderIds
	s.ChildFolderCount = decoded.ChildFolderCount
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.IsHidden = decoded.IsHidden
	s.MessageRules = decoded.MessageRules
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ParentFolderId = decoded.ParentFolderId
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.TotalItemCount = decoded.TotalItemCount
	s.UnreadItemCount = decoded.UnreadItemCount

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MailSearchFolder into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["childFolders"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ChildFolders into list []json.RawMessage: %+v", err)
		}

		output := make([]MailFolder, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMailFolderImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ChildFolders' for 'MailSearchFolder': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ChildFolders = &output
	}

	if v, ok := temp["messages"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Messages into list []json.RawMessage: %+v", err)
		}

		output := make([]Message, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMessageImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Messages' for 'MailSearchFolder': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Messages = &output
	}

	return nil
}
