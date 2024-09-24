package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceContainer interface {
	Entity
	SecurityDataSourceContainer() BaseSecurityDataSourceContainerImpl
}

var _ SecurityDataSourceContainer = BaseSecurityDataSourceContainerImpl{}

type BaseSecurityDataSourceContainerImpl struct {
	// Created date and time of the dataSourceContainer entity.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of the dataSourceContainer entity.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The hold status of the dataSourceContainer. The possible values are: notApplied, applied, applying, removing, partial
	HoldStatus *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`

	// Last modified date and time of the dataSourceContainer.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time that the dataSourceContainer was released from the case.
	ReleasedDateTime nullable.Type[string] `json:"releasedDateTime,omitempty"`

	// Latest status of the dataSourceContainer. Possible values are: Active, Released.
	Status *SecurityDataSourceContainerStatus `json:"status,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityDataSourceContainerImpl) SecurityDataSourceContainer() BaseSecurityDataSourceContainerImpl {
	return s
}

func (s BaseSecurityDataSourceContainerImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityDataSourceContainer = RawSecurityDataSourceContainerImpl{}

// RawSecurityDataSourceContainerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityDataSourceContainerImpl struct {
	securityDataSourceContainer BaseSecurityDataSourceContainerImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawSecurityDataSourceContainerImpl) SecurityDataSourceContainer() BaseSecurityDataSourceContainerImpl {
	return s.securityDataSourceContainer
}

func (s RawSecurityDataSourceContainerImpl) Entity() BaseEntityImpl {
	return s.securityDataSourceContainer.Entity()
}

var _ json.Marshaler = BaseSecurityDataSourceContainerImpl{}

func (s BaseSecurityDataSourceContainerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityDataSourceContainerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityDataSourceContainerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityDataSourceContainerImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.dataSourceContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityDataSourceContainerImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalSecurityDataSourceContainerImplementation(input []byte) (SecurityDataSourceContainer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDataSourceContainer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryCustodian") {
		var out SecurityEdiscoveryCustodian
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryCustodian: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.ediscoveryNoncustodialDataSource") {
		var out SecurityEdiscoveryNoncustodialDataSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityEdiscoveryNoncustodialDataSource: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityDataSourceContainerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityDataSourceContainerImpl: %+v", err)
	}

	return RawSecurityDataSourceContainerImpl{
		securityDataSourceContainer: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
