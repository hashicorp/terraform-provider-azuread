package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppListItem = AppleAppListItem{}

type AppleAppListItem struct {

	// Fields inherited from AppListItem

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

func (s AppleAppListItem) AppListItem() BaseAppListItemImpl {
	return BaseAppListItemImpl{
		AppId:       s.AppId,
		AppStoreUrl: s.AppStoreUrl,
		Name:        s.Name,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Publisher:   s.Publisher,
	}
}

var _ json.Marshaler = AppleAppListItem{}

func (s AppleAppListItem) MarshalJSON() ([]byte, error) {
	type wrapper AppleAppListItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppleAppListItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppleAppListItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appleAppListItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppleAppListItem: %+v", err)
	}

	return encoded, nil
}
