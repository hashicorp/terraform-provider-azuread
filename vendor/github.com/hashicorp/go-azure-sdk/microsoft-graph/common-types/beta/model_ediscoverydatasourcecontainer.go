package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceContainer interface {
	Entity
	EdiscoveryDataSourceContainer() BaseEdiscoveryDataSourceContainerImpl
}

var _ EdiscoveryDataSourceContainer = BaseEdiscoveryDataSourceContainerImpl{}

type BaseEdiscoveryDataSourceContainerImpl struct {
	// Created date and time of the dataSourceContainer entity.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Display name of the dataSourceContainer entity.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	HoldStatus         *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
	LastIndexOperation *EdiscoveryCaseIndexOperation   `json:"lastIndexOperation,omitempty"`

	// Last modified date and time of the dataSourceContainer.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time that the dataSourceContainer was released from the case.
	ReleasedDateTime nullable.Type[string] `json:"releasedDateTime,omitempty"`

	// Latest status of the dataSourceContainer. Possible values are: Active, Released.
	Status *EdiscoveryDataSourceContainerStatus `json:"status,omitempty"`

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

func (s BaseEdiscoveryDataSourceContainerImpl) EdiscoveryDataSourceContainer() BaseEdiscoveryDataSourceContainerImpl {
	return s
}

func (s BaseEdiscoveryDataSourceContainerImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EdiscoveryDataSourceContainer = RawEdiscoveryDataSourceContainerImpl{}

// RawEdiscoveryDataSourceContainerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEdiscoveryDataSourceContainerImpl struct {
	ediscoveryDataSourceContainer BaseEdiscoveryDataSourceContainerImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawEdiscoveryDataSourceContainerImpl) EdiscoveryDataSourceContainer() BaseEdiscoveryDataSourceContainerImpl {
	return s.ediscoveryDataSourceContainer
}

func (s RawEdiscoveryDataSourceContainerImpl) Entity() BaseEntityImpl {
	return s.ediscoveryDataSourceContainer.Entity()
}

var _ json.Marshaler = BaseEdiscoveryDataSourceContainerImpl{}

func (s BaseEdiscoveryDataSourceContainerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEdiscoveryDataSourceContainerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEdiscoveryDataSourceContainerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEdiscoveryDataSourceContainerImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.dataSourceContainer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEdiscoveryDataSourceContainerImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalEdiscoveryDataSourceContainerImplementation(input []byte) (EdiscoveryDataSourceContainer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryDataSourceContainer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.custodian") {
		var out EdiscoveryCustodian
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryCustodian: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.noncustodialDataSource") {
		var out EdiscoveryNoncustodialDataSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryNoncustodialDataSource: %+v", err)
		}
		return out, nil
	}

	var parent BaseEdiscoveryDataSourceContainerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEdiscoveryDataSourceContainerImpl: %+v", err)
	}

	return RawEdiscoveryDataSourceContainerImpl{
		ediscoveryDataSourceContainer: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
