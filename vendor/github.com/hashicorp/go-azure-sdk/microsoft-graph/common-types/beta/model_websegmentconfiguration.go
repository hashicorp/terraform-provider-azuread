package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SegmentConfiguration = WebSegmentConfiguration{}

type WebSegmentConfiguration struct {
	ApplicationSegments *[]WebApplicationSegment `json:"applicationSegments,omitempty"`

	// Fields inherited from SegmentConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WebSegmentConfiguration) SegmentConfiguration() BaseSegmentConfigurationImpl {
	return BaseSegmentConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WebSegmentConfiguration{}

func (s WebSegmentConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WebSegmentConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebSegmentConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebSegmentConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.webSegmentConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebSegmentConfiguration: %+v", err)
	}

	return encoded, nil
}
