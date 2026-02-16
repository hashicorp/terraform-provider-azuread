package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentitySource interface {
	IdentitySource() BaseIdentitySourceImpl
}

var _ IdentitySource = BaseIdentitySourceImpl{}

type BaseIdentitySourceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentitySourceImpl) IdentitySource() BaseIdentitySourceImpl {
	return s
}

var _ IdentitySource = RawIdentitySourceImpl{}

// RawIdentitySourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentitySourceImpl struct {
	identitySource BaseIdentitySourceImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawIdentitySourceImpl) IdentitySource() BaseIdentitySourceImpl {
	return s.identitySource
}

func UnmarshalIdentitySourceImplementation(input []byte) (IdentitySource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentitySource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.azureActiveDirectoryTenant") {
		var out AzureActiveDirectoryTenant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureActiveDirectoryTenant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.crossCloudAzureActiveDirectoryTenant") {
		var out CrossCloudAzureActiveDirectoryTenant
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CrossCloudAzureActiveDirectoryTenant: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.domainIdentitySource") {
		var out DomainIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DomainIdentitySource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalDomainFederation") {
		var out ExternalDomainFederation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalDomainFederation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.socialIdentitySource") {
		var out SocialIdentitySource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SocialIdentitySource: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentitySourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentitySourceImpl: %+v", err)
	}

	return RawIdentitySourceImpl{
		identitySource: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
