package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosHomeScreenItem = IosHomeScreenApp{}

type IosHomeScreenApp struct {
	// BundleID of the app if isWebClip is false or the URL of a web clip if isWebClip is true.
	BundleId *string `json:"bundleID,omitempty"`

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

func (s IosHomeScreenApp) IosHomeScreenItem() BaseIosHomeScreenItemImpl {
	return BaseIosHomeScreenItemImpl{
		DisplayName: s.DisplayName,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = IosHomeScreenApp{}

func (s IosHomeScreenApp) MarshalJSON() ([]byte, error) {
	type wrapper IosHomeScreenApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosHomeScreenApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosHomeScreenApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosHomeScreenApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosHomeScreenApp: %+v", err)
	}

	return encoded, nil
}
