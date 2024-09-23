package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityImpactedAsset = SecurityImpactedMailboxAsset{}

type SecurityImpactedMailboxAsset struct {
	Identifier *SecurityMailboxAssetIdentifier `json:"identifier,omitempty"`

	// Fields inherited from SecurityImpactedAsset

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityImpactedMailboxAsset) SecurityImpactedAsset() BaseSecurityImpactedAssetImpl {
	return BaseSecurityImpactedAssetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityImpactedMailboxAsset{}

func (s SecurityImpactedMailboxAsset) MarshalJSON() ([]byte, error) {
	type wrapper SecurityImpactedMailboxAsset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityImpactedMailboxAsset: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityImpactedMailboxAsset: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.impactedMailboxAsset"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityImpactedMailboxAsset: %+v", err)
	}

	return encoded, nil
}
