package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityInformationProtectionAction = SecurityMetadataAction{}

type SecurityMetadataAction struct {
	// A collection of key-value pairs that should be added to the file.
	MetadataToAdd *[]SecurityKeyValuePair `json:"metadataToAdd,omitempty"`

	// A collection of strings that indicate which keys to remove from the file metadata.
	MetadataToRemove *[]string `json:"metadataToRemove,omitempty"`

	// Fields inherited from SecurityInformationProtectionAction

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityMetadataAction) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return BaseSecurityInformationProtectionActionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityMetadataAction{}

func (s SecurityMetadataAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityMetadataAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityMetadataAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMetadataAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.metadataAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityMetadataAction: %+v", err)
	}

	return encoded, nil
}
