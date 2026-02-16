package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSessionBase interface {
	Entity
	RestoreSessionBase() BaseRestoreSessionBaseImpl
}

var _ RestoreSessionBase = BaseRestoreSessionBaseImpl{}

type BaseRestoreSessionBaseImpl struct {
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

func (s BaseRestoreSessionBaseImpl) RestoreSessionBase() BaseRestoreSessionBaseImpl {
	return s
}

func (s BaseRestoreSessionBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RestoreSessionBase = RawRestoreSessionBaseImpl{}

// RawRestoreSessionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRestoreSessionBaseImpl struct {
	restoreSessionBase BaseRestoreSessionBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawRestoreSessionBaseImpl) RestoreSessionBase() BaseRestoreSessionBaseImpl {
	return s.restoreSessionBase
}

func (s RawRestoreSessionBaseImpl) Entity() BaseEntityImpl {
	return s.restoreSessionBase.Entity()
}

var _ json.Marshaler = BaseRestoreSessionBaseImpl{}

func (s BaseRestoreSessionBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRestoreSessionBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRestoreSessionBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRestoreSessionBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.restoreSessionBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRestoreSessionBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseRestoreSessionBaseImpl{}

func (s *BaseRestoreSessionBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CompletedDateTime           nullable.Type[string]        `json:"completedDateTime,omitempty"`
		CreatedDateTime             nullable.Type[string]        `json:"createdDateTime,omitempty"`
		Error                       *PublicError                 `json:"error,omitempty"`
		LastModifiedDateTime        nullable.Type[string]        `json:"lastModifiedDateTime,omitempty"`
		RestoreJobType              *RestoreJobType              `json:"restoreJobType,omitempty"`
		RestoreSessionArtifactCount *RestoreSessionArtifactCount `json:"restoreSessionArtifactCount,omitempty"`
		Status                      *RestoreSessionStatus        `json:"status,omitempty"`
		Id                          *string                      `json:"id,omitempty"`
		ODataId                     *string                      `json:"@odata.id,omitempty"`
		ODataType                   *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RestoreJobType = decoded.RestoreJobType
	s.RestoreSessionArtifactCount = decoded.RestoreSessionArtifactCount
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRestoreSessionBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseRestoreSessionBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseRestoreSessionBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalRestoreSessionBaseImplementation(input []byte) (RestoreSessionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RestoreSessionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.exchangeRestoreSession") {
		var out ExchangeRestoreSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExchangeRestoreSession: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.oneDriveForBusinessRestoreSession") {
		var out OneDriveForBusinessRestoreSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OneDriveForBusinessRestoreSession: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointRestoreSession") {
		var out SharePointRestoreSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointRestoreSession: %+v", err)
		}
		return out, nil
	}

	var parent BaseRestoreSessionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRestoreSessionBaseImpl: %+v", err)
	}

	return RawRestoreSessionBaseImpl{
		restoreSessionBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
