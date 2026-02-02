package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = BusinessFlow{}

type BusinessFlow struct {
	CustomData       nullable.Type[string] `json:"customData,omitempty"`
	DeDuplicationId  nullable.Type[string] `json:"deDuplicationId,omitempty"`
	Description      nullable.Type[string] `json:"description,omitempty"`
	DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
	Policy           *GovernancePolicy     `json:"policy,omitempty"`
	PolicyTemplateId nullable.Type[string] `json:"policyTemplateId,omitempty"`
	RecordVersion    nullable.Type[string] `json:"recordVersion,omitempty"`
	SchemaId         nullable.Type[string] `json:"schemaId,omitempty"`
	Settings         *BusinessFlowSettings `json:"settings,omitempty"`

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

func (s BusinessFlow) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BusinessFlow{}

func (s BusinessFlow) MarshalJSON() ([]byte, error) {
	type wrapper BusinessFlow
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BusinessFlow: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessFlow: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.businessFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BusinessFlow: %+v", err)
	}

	return encoded, nil
}
