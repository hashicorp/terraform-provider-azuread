package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoverySourceCollection{}

type EdiscoverySourceCollection struct {
	// Adds the results of the sourceCollection to the specified reviewSet.
	AddToReviewSetOperation *EdiscoveryAddToReviewSetOperation `json:"addToReviewSetOperation,omitempty"`

	// Adds an additional source to the sourceCollection.
	AdditionalSources *[]EdiscoveryDataSource `json:"additionalSources,omitempty"`

	// The query string in KQL (Keyword Query Language) query. For details, see Keyword queries and search conditions for
	// Content Search and eDiscovery. You can refine searches by using fields paired with values; for example,
	// subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
	ContentQuery nullable.Type[string] `json:"contentQuery,omitempty"`

	// The user who created the sourceCollection.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time the sourceCollection was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Custodian sources that are included in the sourceCollection.
	CustodianSources *[]EdiscoveryDataSource `json:"custodianSources,omitempty"`

	// When specified, the collection spans across a service for an entire workload. Possible values are: none,
	// allTenantMailboxes, allTenantSites, allCaseCustodians, allCaseNoncustodialDataSources.
	DataSourceScopes *EdiscoveryDataSourceScopes `json:"dataSourceScopes,omitempty"`

	// The description of the sourceCollection.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the sourceCollection.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The last estimate operation associated with the sourceCollection.
	LastEstimateStatisticsOperation *EdiscoveryEstimateStatisticsOperation `json:"lastEstimateStatisticsOperation,omitempty"`

	// The last user who modified the sourceCollection.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The last date and time the sourceCollection was modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// noncustodialDataSource sources that are included in the sourceCollection
	NoncustodialSources *[]EdiscoveryNoncustodialDataSource `json:"noncustodialSources,omitempty"`

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

func (s EdiscoverySourceCollection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoverySourceCollection{}

func (s EdiscoverySourceCollection) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoverySourceCollection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoverySourceCollection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoverySourceCollection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.sourceCollection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoverySourceCollection: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoverySourceCollection{}

func (s *EdiscoverySourceCollection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddToReviewSetOperation         *EdiscoveryAddToReviewSetOperation     `json:"addToReviewSetOperation,omitempty"`
		ContentQuery                    nullable.Type[string]                  `json:"contentQuery,omitempty"`
		CreatedDateTime                 nullable.Type[string]                  `json:"createdDateTime,omitempty"`
		DataSourceScopes                *EdiscoveryDataSourceScopes            `json:"dataSourceScopes,omitempty"`
		Description                     nullable.Type[string]                  `json:"description,omitempty"`
		DisplayName                     nullable.Type[string]                  `json:"displayName,omitempty"`
		LastEstimateStatisticsOperation *EdiscoveryEstimateStatisticsOperation `json:"lastEstimateStatisticsOperation,omitempty"`
		LastModifiedDateTime            nullable.Type[string]                  `json:"lastModifiedDateTime,omitempty"`
		NoncustodialSources             *[]EdiscoveryNoncustodialDataSource    `json:"noncustodialSources,omitempty"`
		Id                              *string                                `json:"id,omitempty"`
		ODataId                         *string                                `json:"@odata.id,omitempty"`
		ODataType                       *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddToReviewSetOperation = decoded.AddToReviewSetOperation
	s.ContentQuery = decoded.ContentQuery
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DataSourceScopes = decoded.DataSourceScopes
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastEstimateStatisticsOperation = decoded.LastEstimateStatisticsOperation
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.NoncustodialSources = decoded.NoncustodialSources
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoverySourceCollection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["additionalSources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AdditionalSources into list []json.RawMessage: %+v", err)
		}

		output := make([]EdiscoveryDataSource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEdiscoveryDataSourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AdditionalSources' for 'EdiscoverySourceCollection': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AdditionalSources = &output
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EdiscoverySourceCollection': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["custodianSources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CustodianSources into list []json.RawMessage: %+v", err)
		}

		output := make([]EdiscoveryDataSource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEdiscoveryDataSourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CustodianSources' for 'EdiscoverySourceCollection': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CustodianSources = &output
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EdiscoverySourceCollection': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
