package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSource interface {
	Entity
	EdiscoveryDataSource() BaseEdiscoveryDataSourceImpl
}

var _ EdiscoveryDataSource = BaseEdiscoveryDataSourceImpl{}

type BaseEdiscoveryDataSourceImpl struct {
	// The user who created the dataSource.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the dataSource was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name of the dataSource, and is the name of the SharePoint site.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	HoldStatus *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`

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

func (s BaseEdiscoveryDataSourceImpl) EdiscoveryDataSource() BaseEdiscoveryDataSourceImpl {
	return s
}

func (s BaseEdiscoveryDataSourceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EdiscoveryDataSource = RawEdiscoveryDataSourceImpl{}

// RawEdiscoveryDataSourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEdiscoveryDataSourceImpl struct {
	ediscoveryDataSource BaseEdiscoveryDataSourceImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawEdiscoveryDataSourceImpl) EdiscoveryDataSource() BaseEdiscoveryDataSourceImpl {
	return s.ediscoveryDataSource
}

func (s RawEdiscoveryDataSourceImpl) Entity() BaseEntityImpl {
	return s.ediscoveryDataSource.Entity()
}

var _ json.Marshaler = BaseEdiscoveryDataSourceImpl{}

func (s BaseEdiscoveryDataSourceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEdiscoveryDataSourceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEdiscoveryDataSourceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEdiscoveryDataSourceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.dataSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEdiscoveryDataSourceImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseEdiscoveryDataSourceImpl{}

func (s *BaseEdiscoveryDataSourceImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string]           `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]           `json:"displayName,omitempty"`
		HoldStatus      *EdiscoveryDataSourceHoldStatus `json:"holdStatus,omitempty"`
		Id              *string                         `json:"id,omitempty"`
		ODataId         *string                         `json:"@odata.id,omitempty"`
		ODataType       *string                         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.HoldStatus = decoded.HoldStatus
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseEdiscoveryDataSourceImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseEdiscoveryDataSourceImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalEdiscoveryDataSourceImplementation(input []byte) (EdiscoveryDataSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryDataSource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.siteSource") {
		var out EdiscoverySiteSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoverySiteSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.unifiedGroupSource") {
		var out EdiscoveryUnifiedGroupSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryUnifiedGroupSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.ediscovery.userSource") {
		var out EdiscoveryUserSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EdiscoveryUserSource: %+v", err)
		}
		return out, nil
	}

	var parent BaseEdiscoveryDataSourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEdiscoveryDataSourceImpl: %+v", err)
	}

	return RawEdiscoveryDataSourceImpl{
		ediscoveryDataSource: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
