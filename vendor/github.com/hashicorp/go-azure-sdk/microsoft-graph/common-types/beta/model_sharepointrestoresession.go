package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreSessionBase = SharePointRestoreSession{}

type SharePointRestoreSession struct {
	// A collection of restore points and destination details that can be used to restore SharePoint sites.
	SiteRestoreArtifacts *[]SiteRestoreArtifact `json:"siteRestoreArtifacts,omitempty"`

	// A collection of SharePoint site URLs and destination details that can be used to restore SharePoint sites.
	SiteRestoreArtifactsBulkAdditionRequests *[]SiteRestoreArtifactsBulkAdditionRequest `json:"siteRestoreArtifactsBulkAdditionRequests,omitempty"`

	// Fields inherited from RestoreSessionBase

	// The time of completion of the restore session.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The identity of person who created the restore session.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the restore session.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains error details if the restore session fails or completes with an error.
	Error *PublicError `json:"error,omitempty"`

	// Identity of the person who last modified the restore session.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification of the restore session.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Indicates whether the restore session was created normally or by a bulk job.
	RestoreJobType *RestoreJobType `json:"restoreJobType,omitempty"`

	// The number of metadata artifacts that belong to this restore session.
	RestoreSessionArtifactCount *RestoreSessionArtifactCount `json:"restoreSessionArtifactCount,omitempty"`

	// Status of the restore session. The value is an aggregated status of the restored artifacts. The possible values are:
	// draft, activating, active, completedWithError, completed, unknownFutureValue, failed. Use the Prefer:
	// include-unknown-enum-members request header to get the following value in this evolvable enum: failed.
	Status *RestoreSessionStatus `json:"status,omitempty"`

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

func (s SharePointRestoreSession) RestoreSessionBase() BaseRestoreSessionBaseImpl {
	return BaseRestoreSessionBaseImpl{
		CompletedDateTime:           s.CompletedDateTime,
		CreatedBy:                   s.CreatedBy,
		CreatedDateTime:             s.CreatedDateTime,
		Error:                       s.Error,
		LastModifiedBy:              s.LastModifiedBy,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		RestoreJobType:              s.RestoreJobType,
		RestoreSessionArtifactCount: s.RestoreSessionArtifactCount,
		Status:                      s.Status,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s SharePointRestoreSession) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SharePointRestoreSession{}

func (s SharePointRestoreSession) MarshalJSON() ([]byte, error) {
	type wrapper SharePointRestoreSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharePointRestoreSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharePointRestoreSession: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharePointRestoreSession"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharePointRestoreSession: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SharePointRestoreSession{}

func (s *SharePointRestoreSession) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SiteRestoreArtifacts                     *[]SiteRestoreArtifact                     `json:"siteRestoreArtifacts,omitempty"`
		SiteRestoreArtifactsBulkAdditionRequests *[]SiteRestoreArtifactsBulkAdditionRequest `json:"siteRestoreArtifactsBulkAdditionRequests,omitempty"`
		CompletedDateTime                        nullable.Type[string]                      `json:"completedDateTime,omitempty"`
		CreatedDateTime                          nullable.Type[string]                      `json:"createdDateTime,omitempty"`
		Error                                    *PublicError                               `json:"error,omitempty"`
		LastModifiedDateTime                     nullable.Type[string]                      `json:"lastModifiedDateTime,omitempty"`
		RestoreJobType                           *RestoreJobType                            `json:"restoreJobType,omitempty"`
		RestoreSessionArtifactCount              *RestoreSessionArtifactCount               `json:"restoreSessionArtifactCount,omitempty"`
		Status                                   *RestoreSessionStatus                      `json:"status,omitempty"`
		Id                                       *string                                    `json:"id,omitempty"`
		ODataId                                  *string                                    `json:"@odata.id,omitempty"`
		ODataType                                *string                                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.SiteRestoreArtifacts = decoded.SiteRestoreArtifacts
	s.SiteRestoreArtifactsBulkAdditionRequests = decoded.SiteRestoreArtifactsBulkAdditionRequests
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RestoreJobType = decoded.RestoreJobType
	s.RestoreSessionArtifactCount = decoded.RestoreSessionArtifactCount
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharePointRestoreSession into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SharePointRestoreSession': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SharePointRestoreSession': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
