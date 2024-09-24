package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppListItem interface {
	AppListItem() BaseAppListItemImpl
}

var _ AppListItem = BaseAppListItemImpl{}

type BaseAppListItemImpl struct {
	// The application or bundle identifier of the application
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The Store URL of the application
	AppStoreUrl nullable.Type[string] `json:"appStoreUrl,omitempty"`

	// The application name
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The publisher of the application
	Publisher nullable.Type[string] `json:"publisher,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAppListItemImpl) AppListItem() BaseAppListItemImpl {
	return s
}

var _ AppListItem = RawAppListItemImpl{}

// RawAppListItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAppListItemImpl struct {
	appListItem BaseAppListItemImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawAppListItemImpl) AppListItem() BaseAppListItemImpl {
	return s.appListItem
}

func UnmarshalAppListItemImplementation(input []byte) (AppListItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AppListItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appleAppListItem") {
		var out AppleAppListItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppleAppListItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseAppListItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAppListItemImpl: %+v", err)
	}

	return RawAppListItemImpl{
		appListItem: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
