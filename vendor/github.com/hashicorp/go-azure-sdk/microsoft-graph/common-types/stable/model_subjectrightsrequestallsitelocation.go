package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectRightsRequestSiteLocation = SubjectRightsRequestAllSiteLocation{}

type SubjectRightsRequestAllSiteLocation struct {

	// Fields inherited from SubjectRightsRequestSiteLocation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SubjectRightsRequestAllSiteLocation) SubjectRightsRequestSiteLocation() BaseSubjectRightsRequestSiteLocationImpl {
	return BaseSubjectRightsRequestSiteLocationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubjectRightsRequestAllSiteLocation{}

func (s SubjectRightsRequestAllSiteLocation) MarshalJSON() ([]byte, error) {
	type wrapper SubjectRightsRequestAllSiteLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubjectRightsRequestAllSiteLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestAllSiteLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subjectRightsRequestAllSiteLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubjectRightsRequestAllSiteLocation: %+v", err)
	}

	return encoded, nil
}
