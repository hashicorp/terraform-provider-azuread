package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceAttributeDestination interface {
	AccessPackageResourceAttributeDestination() BaseAccessPackageResourceAttributeDestinationImpl
}

var _ AccessPackageResourceAttributeDestination = BaseAccessPackageResourceAttributeDestinationImpl{}

type BaseAccessPackageResourceAttributeDestinationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessPackageResourceAttributeDestinationImpl) AccessPackageResourceAttributeDestination() BaseAccessPackageResourceAttributeDestinationImpl {
	return s
}

var _ AccessPackageResourceAttributeDestination = RawAccessPackageResourceAttributeDestinationImpl{}

// RawAccessPackageResourceAttributeDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessPackageResourceAttributeDestinationImpl struct {
	accessPackageResourceAttributeDestination BaseAccessPackageResourceAttributeDestinationImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawAccessPackageResourceAttributeDestinationImpl) AccessPackageResourceAttributeDestination() BaseAccessPackageResourceAttributeDestinationImpl {
	return s.accessPackageResourceAttributeDestination
}

func UnmarshalAccessPackageResourceAttributeDestinationImplementation(input []byte) (AccessPackageResourceAttributeDestination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageResourceAttributeDestination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageUserDirectoryAttributeStore") {
		var out AccessPackageUserDirectoryAttributeStore
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageUserDirectoryAttributeStore: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessPackageResourceAttributeDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessPackageResourceAttributeDestinationImpl: %+v", err)
	}

	return RawAccessPackageResourceAttributeDestinationImpl{
		accessPackageResourceAttributeDestination: parent,
		Type:   value,
		Values: temp,
	}, nil

}
