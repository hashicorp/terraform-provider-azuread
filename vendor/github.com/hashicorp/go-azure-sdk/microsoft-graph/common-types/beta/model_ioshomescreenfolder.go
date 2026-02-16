package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosHomeScreenItem = IosHomeScreenFolder{}

type IosHomeScreenFolder struct {
	// Pages of Home Screen Layout Icons which must be applications or web clips. This collection can contain a maximum of
	// 500 elements.
	Pages *[]IosHomeScreenFolderPage `json:"pages,omitempty"`

	// Fields inherited from IosHomeScreenItem

	// Name of the app
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosHomeScreenFolder) IosHomeScreenItem() BaseIosHomeScreenItemImpl {
	return BaseIosHomeScreenItemImpl{
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = IosHomeScreenFolder{}

func (s IosHomeScreenFolder) MarshalJSON() ([]byte, error) {
	type wrapper IosHomeScreenFolder
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosHomeScreenFolder: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosHomeScreenFolder: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosHomeScreenFolder"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosHomeScreenFolder: %+v", err)
	}

	return encoded, nil
}
