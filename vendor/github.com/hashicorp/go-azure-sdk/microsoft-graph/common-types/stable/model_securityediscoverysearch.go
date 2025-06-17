package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecuritySearch = SecurityEdiscoverySearch{}

type SecurityEdiscoverySearch struct {
	// Adds the results of the eDiscovery search to the specified reviewSet.
	AddToReviewSetOperation *SecurityEdiscoveryAddToReviewSetOperation `json:"addToReviewSetOperation,omitempty"`

	// Adds an additional source to the eDiscovery search.
	AdditionalSources *[]SecurityDataSource `json:"additionalSources,omitempty"`

	// Custodian sources that are included in the eDiscovery search.
	CustodianSources *[]SecurityDataSource `json:"custodianSources,omitempty"`

	// When specified, the collection spans across a service for an entire workload. Possible values are: none,
	// allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
	DataSourceScopes *SecurityDataSourceScopes `json:"dataSourceScopes,omitempty"`

	// The last estimate operation associated with the eDiscovery search.
	LastEstimateStatisticsOperation *SecurityEdiscoveryEstimateOperation `json:"lastEstimateStatisticsOperation,omitempty"`

	// noncustodialDataSource sources that are included in the eDiscovery search
	NoncustodialSources *[]SecurityEdiscoveryNoncustodialDataSource `json:"noncustodialSources,omitempty"`

	// Fields inherited from SecuritySearch

	ContentQuery         nullable.Type[string] `json:"contentQuery,omitempty"`
	CreatedBy            IdentitySet           `json:"createdBy"`
	CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
	Description          nullable.Type[string] `json:"description,omitempty"`
	DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s SecurityEdiscoverySearch) SecuritySearch() BaseSecuritySearchImpl {
	return BaseSecuritySearchImpl{
		ContentQuery:         s.ContentQuery,
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DisplayName:          s.DisplayName,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SecurityEdiscoverySearch) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityEdiscoverySearch{}

func (s SecurityEdiscoverySearch) MarshalJSON() ([]byte, error) {
	type wrapper SecurityEdiscoverySearch
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityEdiscoverySearch: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityEdiscoverySearch: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.ediscoverySearch"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityEdiscoverySearch: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityEdiscoverySearch{}

func (s *SecurityEdiscoverySearch) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddToReviewSetOperation         *SecurityEdiscoveryAddToReviewSetOperation  `json:"addToReviewSetOperation,omitempty"`
		DataSourceScopes                *SecurityDataSourceScopes                   `json:"dataSourceScopes,omitempty"`
		LastEstimateStatisticsOperation *SecurityEdiscoveryEstimateOperation        `json:"lastEstimateStatisticsOperation,omitempty"`
		NoncustodialSources             *[]SecurityEdiscoveryNoncustodialDataSource `json:"noncustodialSources,omitempty"`
		ContentQuery                    nullable.Type[string]                       `json:"contentQuery,omitempty"`
		CreatedDateTime                 nullable.Type[string]                       `json:"createdDateTime,omitempty"`
		Description                     nullable.Type[string]                       `json:"description,omitempty"`
		DisplayName                     nullable.Type[string]                       `json:"displayName,omitempty"`
		LastModifiedDateTime            nullable.Type[string]                       `json:"lastModifiedDateTime,omitempty"`
		Id                              *string                                     `json:"id,omitempty"`
		ODataId                         *string                                     `json:"@odata.id,omitempty"`
		ODataType                       *string                                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddToReviewSetOperation = decoded.AddToReviewSetOperation
	s.DataSourceScopes = decoded.DataSourceScopes
	s.LastEstimateStatisticsOperation = decoded.LastEstimateStatisticsOperation
	s.NoncustodialSources = decoded.NoncustodialSources
	s.ContentQuery = decoded.ContentQuery
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityEdiscoverySearch into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["additionalSources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AdditionalSources into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityDataSource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityDataSourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AdditionalSources' for 'SecurityEdiscoverySearch': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AdditionalSources = &output
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityEdiscoverySearch': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["custodianSources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CustodianSources into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityDataSource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityDataSourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CustodianSources' for 'SecurityEdiscoverySearch': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CustodianSources = &output
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityEdiscoverySearch': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
