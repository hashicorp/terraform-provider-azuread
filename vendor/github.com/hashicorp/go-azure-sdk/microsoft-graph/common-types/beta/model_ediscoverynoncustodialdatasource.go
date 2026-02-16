package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EdiscoveryDataSourceContainer = EdiscoveryNoncustodialDataSource{}

type EdiscoveryNoncustodialDataSource struct {
	// Indicates if hold is applied to noncustodial data source (such as mailbox or site).
	ApplyHoldToSource nullable.Type[bool] `json:"applyHoldToSource,omitempty"`

	// User source or SharePoint site data source as noncustodial data source.
	DataSource *EdiscoveryDataSource `json:"dataSource,omitempty"`

	// Fields inherited from EdiscoveryDataSourceContainer

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

func (s EdiscoveryNoncustodialDataSource) EdiscoveryDataSourceContainer() BaseEdiscoveryDataSourceContainerImpl {
	return BaseEdiscoveryDataSourceContainerImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		HoldStatus:           s.HoldStatus,
		LastIndexOperation:   s.LastIndexOperation,
		LastModifiedDateTime: s.LastModifiedDateTime,
		ReleasedDateTime:     s.ReleasedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EdiscoveryNoncustodialDataSource) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryNoncustodialDataSource{}

func (s EdiscoveryNoncustodialDataSource) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryNoncustodialDataSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryNoncustodialDataSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryNoncustodialDataSource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.noncustodialDataSource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryNoncustodialDataSource: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryNoncustodialDataSource{}

func (s *EdiscoveryNoncustodialDataSource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplyHoldToSource    nullable.Type[bool]                  `json:"applyHoldToSource,omitempty"`
		CreatedDateTime      nullable.Type[string]                `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]                `json:"displayName,omitempty"`
		HoldStatus           *EdiscoveryDataSourceHoldStatus      `json:"holdStatus,omitempty"`
		LastIndexOperation   *EdiscoveryCaseIndexOperation        `json:"lastIndexOperation,omitempty"`
		LastModifiedDateTime nullable.Type[string]                `json:"lastModifiedDateTime,omitempty"`
		ReleasedDateTime     nullable.Type[string]                `json:"releasedDateTime,omitempty"`
		Status               *EdiscoveryDataSourceContainerStatus `json:"status,omitempty"`
		Id                   *string                              `json:"id,omitempty"`
		ODataId              *string                              `json:"@odata.id,omitempty"`
		ODataType            *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplyHoldToSource = decoded.ApplyHoldToSource
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.HoldStatus = decoded.HoldStatus
	s.Id = decoded.Id
	s.LastIndexOperation = decoded.LastIndexOperation
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ReleasedDateTime = decoded.ReleasedDateTime
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryNoncustodialDataSource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataSource"]; ok {
		impl, err := UnmarshalEdiscoveryDataSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataSource' for 'EdiscoveryNoncustodialDataSource': %+v", err)
		}
		s.DataSource = &impl
	}

	return nil
}
