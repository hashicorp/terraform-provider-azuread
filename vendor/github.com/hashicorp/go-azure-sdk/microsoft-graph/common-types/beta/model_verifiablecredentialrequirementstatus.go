package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiableCredentialRequirementStatus interface {
	VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl
}

var _ VerifiableCredentialRequirementStatus = BaseVerifiableCredentialRequirementStatusImpl{}

type BaseVerifiableCredentialRequirementStatusImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseVerifiableCredentialRequirementStatusImpl) VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl {
	return s
}

var _ VerifiableCredentialRequirementStatus = RawVerifiableCredentialRequirementStatusImpl{}

// RawVerifiableCredentialRequirementStatusImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVerifiableCredentialRequirementStatusImpl struct {
	verifiableCredentialRequirementStatus BaseVerifiableCredentialRequirementStatusImpl
	Type                                  string
	Values                                map[string]interface{}
}

func (s RawVerifiableCredentialRequirementStatusImpl) VerifiableCredentialRequirementStatus() BaseVerifiableCredentialRequirementStatusImpl {
	return s.verifiableCredentialRequirementStatus
}

func UnmarshalVerifiableCredentialRequirementStatusImplementation(input []byte) (VerifiableCredentialRequirementStatus, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VerifiableCredentialRequirementStatus into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiableCredentialRequired") {
		var out VerifiableCredentialRequired
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiableCredentialRequired: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiableCredentialRetrieved") {
		var out VerifiableCredentialRetrieved
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiableCredentialRetrieved: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.verifiableCredentialVerified") {
		var out VerifiableCredentialVerified
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VerifiableCredentialVerified: %+v", err)
		}
		return out, nil
	}

	var parent BaseVerifiableCredentialRequirementStatusImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVerifiableCredentialRequirementStatusImpl: %+v", err)
	}

	return RawVerifiableCredentialRequirementStatusImpl{
		verifiableCredentialRequirementStatus: parent,
		Type:                                  value,
		Values:                                temp,
	}, nil

}
