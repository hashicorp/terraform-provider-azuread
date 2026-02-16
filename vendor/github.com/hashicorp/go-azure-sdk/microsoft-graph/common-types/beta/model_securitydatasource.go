package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSource interface {
	Entity
	SecurityDataSource() BaseSecurityDataSourceImpl
}

var _ SecurityDataSource = BaseSecurityDataSourceImpl{}

type BaseSecurityDataSourceImpl struct {
	// The user who created the dataSource.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the dataSource was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The display name of the dataSource and is the name of the SharePoint site.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The hold status of the dataSource.The possible values are: notApplied, applied, applying, removing, partial
	HoldStatus *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`

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

func (s BaseSecurityDataSourceImpl) SecurityDataSource() BaseSecurityDataSourceImpl {
	return s
}

func (s BaseSecurityDataSourceImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ SecurityDataSource = RawSecurityDataSourceImpl{}

// RawSecurityDataSourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityDataSourceImpl struct {
	securityDataSource BaseSecurityDataSourceImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawSecurityDataSourceImpl) SecurityDataSource() BaseSecurityDataSourceImpl {
	return s.securityDataSource
}

func (s RawSecurityDataSourceImpl) Entity() BaseEntityImpl {
	return s.securityDataSource.Entity()
}

var _ json.Marshaler = BaseSecurityDataSourceImpl{}

func (s BaseSecurityDataSourceImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseSecurityDataSourceImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseSecurityDataSourceImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseSecurityDataSourceImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.dataSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseSecurityDataSourceImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseSecurityDataSourceImpl{}

func (s *BaseSecurityDataSourceImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime nullable.Type[string]         `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string]         `json:"displayName,omitempty"`
		HoldStatus      *SecurityDataSourceHoldStatus `json:"holdStatus,omitempty"`
		Id              *string                       `json:"id,omitempty"`
		ODataId         *string                       `json:"@odata.id,omitempty"`
		ODataType       *string                       `json:"@odata.type,omitempty"`
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
		return fmt.Errorf("unmarshaling BaseSecurityDataSourceImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseSecurityDataSourceImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}

func UnmarshalSecurityDataSourceImplementation(input []byte) (SecurityDataSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDataSource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.siteSource") {
		var out SecuritySiteSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySiteSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.unifiedGroupSource") {
		var out SecurityUnifiedGroupSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUnifiedGroupSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.userSource") {
		var out SecurityUserSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityUserSource: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityDataSourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityDataSourceImpl: %+v", err)
	}

	return RawSecurityDataSourceImpl{
		securityDataSource: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
