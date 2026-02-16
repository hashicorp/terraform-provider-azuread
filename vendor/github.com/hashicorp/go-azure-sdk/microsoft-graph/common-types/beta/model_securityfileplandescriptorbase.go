package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFilePlanDescriptorBase interface {
	SecurityFilePlanDescriptorBase() BaseSecurityFilePlanDescriptorBaseImpl
}

var _ SecurityFilePlanDescriptorBase = BaseSecurityFilePlanDescriptorBaseImpl{}

type BaseSecurityFilePlanDescriptorBaseImpl struct {
	// Unique string that defines the name for the file plan descriptor associated with a particular retention label.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityFilePlanDescriptorBaseImpl) SecurityFilePlanDescriptorBase() BaseSecurityFilePlanDescriptorBaseImpl {
	return s
}

var _ SecurityFilePlanDescriptorBase = RawSecurityFilePlanDescriptorBaseImpl{}

// RawSecurityFilePlanDescriptorBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityFilePlanDescriptorBaseImpl struct {
	securityFilePlanDescriptorBase BaseSecurityFilePlanDescriptorBaseImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawSecurityFilePlanDescriptorBaseImpl) SecurityFilePlanDescriptorBase() BaseSecurityFilePlanDescriptorBaseImpl {
	return s.securityFilePlanDescriptorBase
}

func UnmarshalSecurityFilePlanDescriptorBaseImplementation(input []byte) (SecurityFilePlanDescriptorBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityFilePlanDescriptorBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanAppliedCategory") {
		var out SecurityFilePlanAppliedCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanAppliedCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanAuthority") {
		var out SecurityFilePlanAuthority
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanAuthority: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanCitation") {
		var out SecurityFilePlanCitation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanCitation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanDepartment") {
		var out SecurityFilePlanDepartment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanDepartment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanReference") {
		var out SecurityFilePlanReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.filePlanSubcategory") {
		var out SecurityFilePlanSubcategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityFilePlanSubcategory: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityFilePlanDescriptorBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityFilePlanDescriptorBaseImpl: %+v", err)
	}

	return RawSecurityFilePlanDescriptorBaseImpl{
		securityFilePlanDescriptorBase: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
