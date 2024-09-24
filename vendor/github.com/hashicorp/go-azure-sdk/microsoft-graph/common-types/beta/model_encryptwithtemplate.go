package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EncryptContent = EncryptWithTemplate{}

type EncryptWithTemplate struct {
	AvailableForEncryption nullable.Type[bool]   `json:"availableForEncryption,omitempty"`
	TemplateId             nullable.Type[string] `json:"templateId,omitempty"`

	// Fields inherited from EncryptContent

	EncryptWith *EncryptWith `json:"encryptWith,omitempty"`

	// Fields inherited from LabelActionBase

	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EncryptWithTemplate) EncryptContent() BaseEncryptContentImpl {
	return BaseEncryptContentImpl{
		EncryptWith: s.EncryptWith,
		Name:        s.Name,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s EncryptWithTemplate) LabelActionBase() BaseLabelActionBaseImpl {
	return BaseLabelActionBaseImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EncryptWithTemplate{}

func (s EncryptWithTemplate) MarshalJSON() ([]byte, error) {
	type wrapper EncryptWithTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EncryptWithTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EncryptWithTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.encryptWithTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EncryptWithTemplate: %+v", err)
	}

	return encoded, nil
}
