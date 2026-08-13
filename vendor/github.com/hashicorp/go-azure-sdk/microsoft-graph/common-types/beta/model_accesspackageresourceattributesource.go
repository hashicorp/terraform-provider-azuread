package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceAttributeSource interface {
	AccessPackageResourceAttributeSource() BaseAccessPackageResourceAttributeSourceImpl
}

var _ AccessPackageResourceAttributeSource = BaseAccessPackageResourceAttributeSourceImpl{}

type BaseAccessPackageResourceAttributeSourceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessPackageResourceAttributeSourceImpl) AccessPackageResourceAttributeSource() BaseAccessPackageResourceAttributeSourceImpl {
	return s
}

var _ AccessPackageResourceAttributeSource = RawAccessPackageResourceAttributeSourceImpl{}

// RawAccessPackageResourceAttributeSourceImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAccessPackageResourceAttributeSourceImpl struct {
	accessPackageResourceAttributeSource BaseAccessPackageResourceAttributeSourceImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawAccessPackageResourceAttributeSourceImpl) AccessPackageResourceAttributeSource() BaseAccessPackageResourceAttributeSourceImpl {
	return s.accessPackageResourceAttributeSource
}

func (s RawAccessPackageResourceAttributeSourceImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAccessPackageResourceAttributeSourceImplementation(input []byte) (AccessPackageResourceAttributeSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceAttributeSource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageResourceAttributeQuestion") {
		var out AccessPackageResourceAttributeQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageResourceAttributeQuestion: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessPackageResourceAttributeSourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessPackageResourceAttributeSourceImpl: %+v", err)
	}

	return RawAccessPackageResourceAttributeSourceImpl{
		accessPackageResourceAttributeSource: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
