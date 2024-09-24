package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookItem interface {
	Entity
	OutlookItem() BaseOutlookItemImpl
}

var _ OutlookItem = BaseOutlookItemImpl{}

type BaseOutlookItemImpl struct {
	// The categories associated with the item
	Categories *[]string `json:"categories,omitempty"`

	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange
	// to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseOutlookItemImpl) OutlookItem() BaseOutlookItemImpl {
	return s
}

func (s BaseOutlookItemImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OutlookItem = RawOutlookItemImpl{}

// RawOutlookItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOutlookItemImpl struct {
	outlookItem BaseOutlookItemImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawOutlookItemImpl) OutlookItem() BaseOutlookItemImpl {
	return s.outlookItem
}

func (s RawOutlookItemImpl) Entity() BaseEntityImpl {
	return s.outlookItem.Entity()
}

var _ json.Marshaler = BaseOutlookItemImpl{}

func (s BaseOutlookItemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOutlookItemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOutlookItemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOutlookItemImpl: %+v", err)
	}

	delete(decoded, "changeKey")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outlookItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOutlookItemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOutlookItemImplementation(input []byte) (OutlookItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OutlookItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.contact") {
		var out Contact
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Contact: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.event") {
		var out Event
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Event: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.message") {
		var out Message
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Message: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.post") {
		var out Post
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Post: %+v", err)
		}
		return out, nil
	}

	var parent BaseOutlookItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOutlookItemImpl: %+v", err)
	}

	return RawOutlookItemImpl{
		outlookItem: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
