package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedAsset interface {
	SecurityImpactedAsset() BaseSecurityImpactedAssetImpl
}

var _ SecurityImpactedAsset = BaseSecurityImpactedAssetImpl{}

type BaseSecurityImpactedAssetImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityImpactedAssetImpl) SecurityImpactedAsset() BaseSecurityImpactedAssetImpl {
	return s
}

var _ SecurityImpactedAsset = RawSecurityImpactedAssetImpl{}

// RawSecurityImpactedAssetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityImpactedAssetImpl struct {
	securityImpactedAsset BaseSecurityImpactedAssetImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawSecurityImpactedAssetImpl) SecurityImpactedAsset() BaseSecurityImpactedAssetImpl {
	return s.securityImpactedAsset
}

func UnmarshalSecurityImpactedAssetImplementation(input []byte) (SecurityImpactedAsset, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityImpactedAsset into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.impactedDeviceAsset") {
		var out SecurityImpactedDeviceAsset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityImpactedDeviceAsset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.impactedMailboxAsset") {
		var out SecurityImpactedMailboxAsset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityImpactedMailboxAsset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.impactedUserAsset") {
		var out SecurityImpactedUserAsset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityImpactedUserAsset: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityImpactedAssetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityImpactedAssetImpl: %+v", err)
	}

	return RawSecurityImpactedAssetImpl{
		securityImpactedAsset: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
