package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolder interface {
	Entity
	MailFolder() BaseMailFolderImpl
}

var _ MailFolder = BaseMailFolderImpl{}

type BaseMailFolderImpl struct {
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

	// The collection of long-running operations in the mailFolder.
	Operations *[]MailFolderOperation `json:"operations,omitempty"`

	// The unique identifier for the mailFolder's parent mailFolder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The number of items in the mailFolder.
	TotalItemCount nullable.Type[int64] `json:"totalItemCount,omitempty"`

	// The number of items in the mailFolder marked as unread.
	UnreadItemCount nullable.Type[int64] `json:"unreadItemCount,omitempty"`

	UserConfigurations *[]UserConfiguration `json:"userConfigurations,omitempty"`

	// The well-known folder name for the folder. The possible values are listed above. This property is only set for
	// default folders created by Outlook. For other folders, this property is null.
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

func (s BaseMailFolderImpl) MailFolder() BaseMailFolderImpl {
	return s
}

func (s BaseMailFolderImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MailFolder = RawMailFolderImpl{}

// RawMailFolderImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMailFolderImpl struct {
	mailFolder BaseMailFolderImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawMailFolderImpl) MailFolder() BaseMailFolderImpl {
	return s.mailFolder
}

func (s RawMailFolderImpl) Entity() BaseEntityImpl {
	return s.mailFolder.Entity()
}

var _ json.Marshaler = BaseMailFolderImpl{}

func (s BaseMailFolderImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMailFolderImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMailFolderImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMailFolderImpl: %+v", err)
	}

	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMailFolderImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseMailFolderImpl{}

func (s *BaseMailFolderImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChildFolderCount              nullable.Type[int64]                 `json:"childFolderCount,omitempty"`
		DisplayName                   nullable.Type[string]                `json:"displayName,omitempty"`
		IsHidden                      nullable.Type[bool]                  `json:"isHidden,omitempty"`
		MessageRules                  *[]MessageRule                       `json:"messageRules,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		ParentFolderId                nullable.Type[string]                `json:"parentFolderId,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		TotalItemCount                nullable.Type[int64]                 `json:"totalItemCount,omitempty"`
		UnreadItemCount               nullable.Type[int64]                 `json:"unreadItemCount,omitempty"`
		UserConfigurations            *[]UserConfiguration                 `json:"userConfigurations,omitempty"`
		WellKnownName                 nullable.Type[string]                `json:"wellKnownName,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChildFolderCount = decoded.ChildFolderCount
	s.DisplayName = decoded.DisplayName
	s.IsHidden = decoded.IsHidden
	s.MessageRules = decoded.MessageRules
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.ParentFolderId = decoded.ParentFolderId
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.TotalItemCount = decoded.TotalItemCount
	s.UnreadItemCount = decoded.UnreadItemCount
	s.UserConfigurations = decoded.UserConfigurations
	s.WellKnownName = decoded.WellKnownName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseMailFolderImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'ChildFolders' for 'BaseMailFolderImpl': %+v", i, err)
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
				return fmt.Errorf("unmarshaling index %d field 'Messages' for 'BaseMailFolderImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Messages = &output
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]MailFolderOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMailFolderOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'BaseMailFolderImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}

func UnmarshalMailFolderImplementation(input []byte) (MailFolder, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MailFolder into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.mailSearchFolder") {
		var out MailSearchFolder
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailSearchFolder: %+v", err)
		}
		return out, nil
	}

	var parent BaseMailFolderImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMailFolderImpl: %+v", err)
	}

	return RawMailFolderImpl{
		mailFolder: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
