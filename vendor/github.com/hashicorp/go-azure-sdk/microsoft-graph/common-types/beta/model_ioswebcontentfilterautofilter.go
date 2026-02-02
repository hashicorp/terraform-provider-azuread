package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosWebContentFilterBase = IosWebContentFilterAutoFilter{}

type IosWebContentFilterAutoFilter struct {
	// Additional URLs allowed for access
	AllowedUrls *[]string `json:"allowedUrls,omitempty"`

	// Additional URLs blocked for access
	BlockedUrls *[]string `json:"blockedUrls,omitempty"`

	// Fields inherited from IosWebContentFilterBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosWebContentFilterAutoFilter) IosWebContentFilterBase() BaseIosWebContentFilterBaseImpl {
	return BaseIosWebContentFilterBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosWebContentFilterAutoFilter{}

func (s IosWebContentFilterAutoFilter) MarshalJSON() ([]byte, error) {
	type wrapper IosWebContentFilterAutoFilter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosWebContentFilterAutoFilter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosWebContentFilterAutoFilter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosWebContentFilterAutoFilter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosWebContentFilterAutoFilter: %+v", err)
	}

	return encoded, nil
}
