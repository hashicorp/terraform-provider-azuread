package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TextClassificationRequest{}

type TextClassificationRequest struct {
	ContentMetaData          *ClassificationRequestContentMetaData `json:"contentMetaData,omitempty"`
	FileExtension            nullable.Type[string]                 `json:"fileExtension,omitempty"`
	MatchTolerancesToInclude *MlClassificationMatchTolerance       `json:"matchTolerancesToInclude,omitempty"`
	ScopesToRun              *SensitiveTypeScope                   `json:"scopesToRun,omitempty"`
	SensitiveTypeIds         *[]string                             `json:"sensitiveTypeIds,omitempty"`
	Text                     nullable.Type[string]                 `json:"text,omitempty"`

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

func (s TextClassificationRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TextClassificationRequest{}

func (s TextClassificationRequest) MarshalJSON() ([]byte, error) {
	type wrapper TextClassificationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TextClassificationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TextClassificationRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.textClassificationRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TextClassificationRequest: %+v", err)
	}

	return encoded, nil
}
