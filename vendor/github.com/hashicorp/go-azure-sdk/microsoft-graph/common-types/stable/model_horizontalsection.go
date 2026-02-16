package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = HorizontalSection{}

type HorizontalSection struct {
	// The set of vertical columns in this section.
	Columns *[]HorizontalSectionColumn `json:"columns,omitempty"`

	// Enumeration value that indicates the emphasis of the section background. The possible values are: none, netural,
	// soft, strong, unknownFutureValue.
	Emphasis *SectionEmphasisType `json:"emphasis,omitempty"`

	// Layout type of the section. The possible values are: none, oneColumn, twoColumns, threeColumns, oneThirdLeftColumn,
	// oneThirdRightColumn, fullWidth, unknownFutureValue.
	Layout *HorizontalSectionLayoutType `json:"layout,omitempty"`

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

func (s HorizontalSection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = HorizontalSection{}

func (s HorizontalSection) MarshalJSON() ([]byte, error) {
	type wrapper HorizontalSection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HorizontalSection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HorizontalSection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.horizontalSection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HorizontalSection: %+v", err)
	}

	return encoded, nil
}
