package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreArtifactsBulkRequestBase = SiteRestoreArtifactsBulkAdditionRequest{}

type SiteRestoreArtifactsBulkAdditionRequest struct {
	// The list of SharePoint site IDs that are added to the corresponding SharePoint restore session in a bulk operation.
	SiteIds *[]string `json:"siteIds,omitempty"`

	// The list of SharePoint site URLs that are added to the corresponding SharePoint restore session in a bulk operation.
	SiteWebUrls *[]string `json:"siteWebUrls,omitempty"`

	// Fields inherited from RestoreArtifactsBulkRequestBase

	// The identity of the person who created the bulk request.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time when the bulk request was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Indicates the restoration destination. The possible values are: new, inPlace, unknownFutureValue.
	DestinationType *DestinationType `json:"destinationType,omitempty"`

	// Name of the addition request.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Error details are populated for resource resolution failures.
	Error *PublicError `json:"error,omitempty"`

	// Identity of the person who last modified this entity.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp when this entity was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The start and end date and time of the protection period.
	ProtectionTimePeriod *TimePeriod `json:"protectionTimePeriod,omitempty"`

	// Indicates which protection units to restore. This property isn't implemented yet. Future value; don't use.
	ProtectionUnitIds *[]string `json:"protectionUnitIds,omitempty"`

	// Indicates which restore point to return. The possible values are: oldest, latest, unknownFutureValue.
	RestorePointPreference *RestorePointPreference `json:"restorePointPreference,omitempty"`

	Status *RestoreArtifactsBulkRequestStatus `json:"status,omitempty"`

	// The type of the restore point. The possible values are: none, fastRestore, unknownFutureValue.
	Tags *RestorePointTags `json:"tags,omitempty"`

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

func (s SiteRestoreArtifactsBulkAdditionRequest) RestoreArtifactsBulkRequestBase() BaseRestoreArtifactsBulkRequestBaseImpl {
	return BaseRestoreArtifactsBulkRequestBaseImpl{
		CreatedBy:              s.CreatedBy,
		CreatedDateTime:        s.CreatedDateTime,
		DestinationType:        s.DestinationType,
		DisplayName:            s.DisplayName,
		Error:                  s.Error,
		LastModifiedBy:         s.LastModifiedBy,
		LastModifiedDateTime:   s.LastModifiedDateTime,
		ProtectionTimePeriod:   s.ProtectionTimePeriod,
		ProtectionUnitIds:      s.ProtectionUnitIds,
		RestorePointPreference: s.RestorePointPreference,
		Status:                 s.Status,
		Tags:                   s.Tags,
		Id:                     s.Id,
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
	}
}

func (s SiteRestoreArtifactsBulkAdditionRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SiteRestoreArtifactsBulkAdditionRequest{}

func (s SiteRestoreArtifactsBulkAdditionRequest) MarshalJSON() ([]byte, error) {
	type wrapper SiteRestoreArtifactsBulkAdditionRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SiteRestoreArtifactsBulkAdditionRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SiteRestoreArtifactsBulkAdditionRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.siteRestoreArtifactsBulkAdditionRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SiteRestoreArtifactsBulkAdditionRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SiteRestoreArtifactsBulkAdditionRequest{}

func (s *SiteRestoreArtifactsBulkAdditionRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SiteIds                *[]string                          `json:"siteIds,omitempty"`
		SiteWebUrls            *[]string                          `json:"siteWebUrls,omitempty"`
		CreatedDateTime        nullable.Type[string]              `json:"createdDateTime,omitempty"`
		DestinationType        *DestinationType                   `json:"destinationType,omitempty"`
		DisplayName            nullable.Type[string]              `json:"displayName,omitempty"`
		Error                  *PublicError                       `json:"error,omitempty"`
		LastModifiedDateTime   nullable.Type[string]              `json:"lastModifiedDateTime,omitempty"`
		ProtectionTimePeriod   *TimePeriod                        `json:"protectionTimePeriod,omitempty"`
		ProtectionUnitIds      *[]string                          `json:"protectionUnitIds,omitempty"`
		RestorePointPreference *RestorePointPreference            `json:"restorePointPreference,omitempty"`
		Status                 *RestoreArtifactsBulkRequestStatus `json:"status,omitempty"`
		Tags                   *RestorePointTags                  `json:"tags,omitempty"`
		Id                     *string                            `json:"id,omitempty"`
		ODataId                *string                            `json:"@odata.id,omitempty"`
		ODataType              *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.SiteIds = decoded.SiteIds
	s.SiteWebUrls = decoded.SiteWebUrls
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DestinationType = decoded.DestinationType
	s.DisplayName = decoded.DisplayName
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ProtectionTimePeriod = decoded.ProtectionTimePeriod
	s.ProtectionUnitIds = decoded.ProtectionUnitIds
	s.RestorePointPreference = decoded.RestorePointPreference
	s.Status = decoded.Status
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SiteRestoreArtifactsBulkAdditionRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SiteRestoreArtifactsBulkAdditionRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SiteRestoreArtifactsBulkAdditionRequest': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
