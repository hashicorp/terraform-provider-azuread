package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataReferenceDefinition{}

type IndustryDataReferenceDefinition struct {
	// The code value for the definition that must be unique within the referenceType.
	Code *string `json:"code,omitempty"`

	// The date and time when the definition was created. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A human-readable representation of the reference code value for display in a user interface.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the definition is disabled.
	IsDisabled *bool `json:"isDisabled,omitempty"`

	// The date and time when the definition was most recently changed. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The categorical type for a collection of enumerated values.
	ReferenceType *string `json:"referenceType,omitempty"`

	// The index that specifies the order in which to present the definition to the user. Must be unique within the
	// referenceType.
	SortIndex *int64 `json:"sortIndex,omitempty"`

	// The standards body or organization source which defined the code.
	Source *string `json:"source,omitempty"`

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

func (s IndustryDataReferenceDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataReferenceDefinition{}

func (s IndustryDataReferenceDefinition) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataReferenceDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataReferenceDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataReferenceDefinition: %+v", err)
	}

	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "source")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.referenceDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataReferenceDefinition: %+v", err)
	}

	return encoded, nil
}
