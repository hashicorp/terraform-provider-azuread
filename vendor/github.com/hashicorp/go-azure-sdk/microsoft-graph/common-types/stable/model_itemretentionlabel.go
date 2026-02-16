package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ItemRetentionLabel{}

type ItemRetentionLabel struct {
	// Specifies whether the label is applied explicitly on the item. True indicates that the label is applied explicitly;
	// otherwise, the label is inherited from its parent. Read-only.
	IsLabelAppliedExplicitly nullable.Type[bool] `json:"isLabelAppliedExplicitly,omitempty"`

	// Identity of the user who applied the label. Read-only.
	LabelAppliedBy *IdentitySet `json:"labelAppliedBy,omitempty"`

	// The date and time when the label was applied on the item. The timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only.
	LabelAppliedDateTime nullable.Type[string] `json:"labelAppliedDateTime,omitempty"`

	// The retention label on the document. Read-write.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The retention settings enforced on the item. Read-write.
	RetentionSettings *RetentionLabelSettings `json:"retentionSettings,omitempty"`

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

func (s ItemRetentionLabel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ItemRetentionLabel{}

func (s ItemRetentionLabel) MarshalJSON() ([]byte, error) {
	type wrapper ItemRetentionLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ItemRetentionLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ItemRetentionLabel: %+v", err)
	}

	delete(decoded, "isLabelAppliedExplicitly")
	delete(decoded, "labelAppliedBy")
	delete(decoded, "labelAppliedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.itemRetentionLabel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ItemRetentionLabel: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ItemRetentionLabel{}

func (s *ItemRetentionLabel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsLabelAppliedExplicitly nullable.Type[bool]     `json:"isLabelAppliedExplicitly,omitempty"`
		LabelAppliedDateTime     nullable.Type[string]   `json:"labelAppliedDateTime,omitempty"`
		Name                     nullable.Type[string]   `json:"name,omitempty"`
		RetentionSettings        *RetentionLabelSettings `json:"retentionSettings,omitempty"`
		Id                       *string                 `json:"id,omitempty"`
		ODataId                  *string                 `json:"@odata.id,omitempty"`
		ODataType                *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsLabelAppliedExplicitly = decoded.IsLabelAppliedExplicitly
	s.LabelAppliedDateTime = decoded.LabelAppliedDateTime
	s.Name = decoded.Name
	s.RetentionSettings = decoded.RetentionSettings
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ItemRetentionLabel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["labelAppliedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LabelAppliedBy' for 'ItemRetentionLabel': %+v", err)
		}
		s.LabelAppliedBy = &impl
	}

	return nil
}
