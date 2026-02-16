package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionUnitsBulkJobBase = SiteProtectionUnitsBulkAdditionJob{}

type SiteProtectionUnitsBulkAdditionJob struct {
	// The list of SharePoint site IDs to add to the SharePoint protection policy.
	SiteIds *[]string `json:"siteIds,omitempty"`

	// The list of SharePoint site URLs to add to the SharePoint protection policy.
	SiteWebUrls *[]string `json:"siteWebUrls,omitempty"`

	// Fields inherited from ProtectionUnitsBulkJobBase

	// The identity of person who created the job.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the job.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The name of the protection units bulk addition job.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Error details containing resource resolution failures, if any.
	Error *PublicError `json:"error,omitempty"`

	// The identity of the person who last modified the job.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification made to the job.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	Status *ProtectionUnitsBulkJobStatus `json:"status,omitempty"`

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

func (s SiteProtectionUnitsBulkAdditionJob) ProtectionUnitsBulkJobBase() BaseProtectionUnitsBulkJobBaseImpl {
	return BaseProtectionUnitsBulkJobBaseImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		DisplayName:          s.DisplayName,
		Error:                s.Error,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SiteProtectionUnitsBulkAdditionJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SiteProtectionUnitsBulkAdditionJob{}

func (s SiteProtectionUnitsBulkAdditionJob) MarshalJSON() ([]byte, error) {
	type wrapper SiteProtectionUnitsBulkAdditionJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SiteProtectionUnitsBulkAdditionJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SiteProtectionUnitsBulkAdditionJob: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.siteProtectionUnitsBulkAdditionJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SiteProtectionUnitsBulkAdditionJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SiteProtectionUnitsBulkAdditionJob{}

func (s *SiteProtectionUnitsBulkAdditionJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SiteIds              *[]string                     `json:"siteIds,omitempty"`
		SiteWebUrls          *[]string                     `json:"siteWebUrls,omitempty"`
		CreatedDateTime      nullable.Type[string]         `json:"createdDateTime,omitempty"`
		DisplayName          nullable.Type[string]         `json:"displayName,omitempty"`
		Error                *PublicError                  `json:"error,omitempty"`
		LastModifiedDateTime nullable.Type[string]         `json:"lastModifiedDateTime,omitempty"`
		Status               *ProtectionUnitsBulkJobStatus `json:"status,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.SiteIds = decoded.SiteIds
	s.SiteWebUrls = decoded.SiteWebUrls
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SiteProtectionUnitsBulkAdditionJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SiteProtectionUnitsBulkAdditionJob': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SiteProtectionUnitsBulkAdditionJob': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
