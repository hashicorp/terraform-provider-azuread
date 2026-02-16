package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Dictionary = FileStorageContainerCustomPropertyDictionary{}

type FileStorageContainerCustomPropertyDictionary struct {

	// Fields inherited from Dictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s FileStorageContainerCustomPropertyDictionary) Dictionary() BaseDictionaryImpl {
	return BaseDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FileStorageContainerCustomPropertyDictionary{}

func (s FileStorageContainerCustomPropertyDictionary) MarshalJSON() ([]byte, error) {
	type wrapper FileStorageContainerCustomPropertyDictionary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FileStorageContainerCustomPropertyDictionary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FileStorageContainerCustomPropertyDictionary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.fileStorageContainerCustomPropertyDictionary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FileStorageContainerCustomPropertyDictionary: %+v", err)
	}

	return encoded, nil
}
