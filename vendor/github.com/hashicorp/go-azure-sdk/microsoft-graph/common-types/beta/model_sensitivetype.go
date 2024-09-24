package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SensitiveType{}

type SensitiveType struct {
	ClassificationMethod *ClassificationMethod `json:"classificationMethod,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	Name                 nullable.Type[string] `json:"name,omitempty"`
	PublisherName        nullable.Type[string] `json:"publisherName,omitempty"`
	RulePackageId        nullable.Type[string] `json:"rulePackageId,omitempty"`
	RulePackageType      nullable.Type[string] `json:"rulePackageType,omitempty"`
	Scope                *SensitiveTypeScope   `json:"scope,omitempty"`
	SensitiveTypeSource  *SensitiveTypeSource  `json:"sensitiveTypeSource,omitempty"`
	State                nullable.Type[string] `json:"state,omitempty"`

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

func (s SensitiveType) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SensitiveType{}

func (s SensitiveType) MarshalJSON() ([]byte, error) {
	type wrapper SensitiveType
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SensitiveType: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SensitiveType: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sensitiveType"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SensitiveType: %+v", err)
	}

	return encoded, nil
}
