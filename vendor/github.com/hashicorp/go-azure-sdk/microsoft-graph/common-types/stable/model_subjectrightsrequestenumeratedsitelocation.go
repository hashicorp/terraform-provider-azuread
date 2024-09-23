package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectRightsRequestSiteLocation = SubjectRightsRequestEnumeratedSiteLocation{}

type SubjectRightsRequestEnumeratedSiteLocation struct {
	// Collection of site URLs that should be included. Includes the URL of each site, for example,
	// https://www.contoso.com/site1.
	Urls *[]string `json:"urls,omitempty"`

	// Fields inherited from SubjectRightsRequestSiteLocation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SubjectRightsRequestEnumeratedSiteLocation) SubjectRightsRequestSiteLocation() BaseSubjectRightsRequestSiteLocationImpl {
	return BaseSubjectRightsRequestSiteLocationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubjectRightsRequestEnumeratedSiteLocation{}

func (s SubjectRightsRequestEnumeratedSiteLocation) MarshalJSON() ([]byte, error) {
	type wrapper SubjectRightsRequestEnumeratedSiteLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubjectRightsRequestEnumeratedSiteLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequestEnumeratedSiteLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subjectRightsRequestEnumeratedSiteLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubjectRightsRequestEnumeratedSiteLocation: %+v", err)
	}

	return encoded, nil
}
