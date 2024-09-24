package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Dictionary interface {
	Dictionary() BaseDictionaryImpl
}

var _ Dictionary = BaseDictionaryImpl{}

type BaseDictionaryImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDictionaryImpl) Dictionary() BaseDictionaryImpl {
	return s
}

var _ Dictionary = RawDictionaryImpl{}

// RawDictionaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDictionaryImpl struct {
	dictionary BaseDictionaryImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawDictionaryImpl) Dictionary() BaseDictionaryImpl {
	return s.dictionary
}

func UnmarshalDictionaryImplementation(input []byte) (Dictionary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Dictionary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customAppScopeAttributesDictionary") {
		var out CustomAppScopeAttributesDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAppScopeAttributesDictionary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customMetadataDictionary") {
		var out CustomMetadataDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomMetadataDictionary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.fileStorageContainerCustomPropertyDictionary") {
		var out FileStorageContainerCustomPropertyDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FileStorageContainerCustomPropertyDictionary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.partner.security.additionalDataDictionary") {
		var out PartnerSecurityAdditionalDataDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartnerSecurityAdditionalDataDictionary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.plannerFormsDictionary") {
		var out PlannerFormsDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PlannerFormsDictionary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.resultTemplateDictionary") {
		var out ResultTemplateDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResultTemplateDictionary: %+v", err)
		}
		return out, nil
	}

	var parent BaseDictionaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDictionaryImpl: %+v", err)
	}

	return RawDictionaryImpl{
		dictionary: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
