package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityDataSourceContainer = SecurityEdiscoveryNoncustodialDataSource{}

type SecurityEdiscoveryNoncustodialDataSource struct {
	// User source or SharePoint site data source as non-custodial data source.
	DataSource *SecurityDataSource `json:"dataSource,omitempty"`

	// Operation entity that represents the latest indexing for the non-custodial data source.
	LastIndexOperation *SecurityEdiscoveryIndexOperation `json:"lastIndexOperation,omitempty"`

	// Fields inherited from SecurityDataSourceContainer

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

func (s SecurityEdiscoveryNoncustodialDataSource) SecurityDataSourceContainer() BaseSecurityDataSourceContainerImpl {
	return BaseSecurityDataSourceContainerImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		HoldStatus:           s.HoldStatus,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ReleasedDateTime:     s.ReleasedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoveryNoncustodialDataSource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoveryNoncustodialDataSource{}

func (s SecurityEdiscoveryNoncustodialDataSource) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoveryNoncustodialDataSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoveryNoncustodialDataSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoveryNoncustodialDataSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoveryNoncustodialDataSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoveryNoncustodialDataSource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoveryNoncustodialDataSource{}

func (s *SecurityEdiscoveryNoncustodialDataSource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastIndexOperation   *SecurityEdiscoveryIndexOperation  `json:"lastIndexOperation,omitempty"`
		CreatedDateTime      nullable.Type[string]              `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]              `json:"displayName,omitempty"`
		HoldStatus           *SecurityDataSourceHoldStatus      `json:"holdStatus,omitempty"`
		LastModifiedDateTime nullable.Type[string]              `json:"lastModifiedDateTime,omitempty"`
		ReleasedDateTime     nullable.Type[string]              `json:"releasedDateTime,omitempty"`
		Status               *SecurityDataSourceContainerStatus `json:"status,omitempty"`
		Id                   *string                            `json:"id,omitempty"`
		ODataId              *string                            `json:"@odata.id,omitempty"`
		ODataType            *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastIndexOperation = decoded.LastIndexOperation
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.HoldStatus = decoded.HoldStatus
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReleasedDateTime = decoded.ReleasedDateTime
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoveryNoncustodialDataSource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := UnmarshalSecurityDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'SecurityEdiscoveryNoncustodialDataSource': %+v", err)
		}
		s.DataSource = &impl
	}

	return nil
}
