package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreArtifactsBulkRequestBase interface {
	Entity
	RestoreArtifactsBulkRequestBase() BaseRestoreArtifactsBulkRequestBaseImpl
}

var _ RestoreArtifactsBulkRequestBase = BaseRestoreArtifactsBulkRequestBaseImpl{}

type BaseRestoreArtifactsBulkRequestBaseImpl struct {
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

func (s BaseRestoreArtifactsBulkRequestBaseImpl) RestoreArtifactsBulkRequestBase() BaseRestoreArtifactsBulkRequestBaseImpl {
	return s
}

func (s BaseRestoreArtifactsBulkRequestBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ RestoreArtifactsBulkRequestBase = RawRestoreArtifactsBulkRequestBaseImpl{}

// RawRestoreArtifactsBulkRequestBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRestoreArtifactsBulkRequestBaseImpl struct {
	restoreArtifactsBulkRequestBase BaseRestoreArtifactsBulkRequestBaseImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawRestoreArtifactsBulkRequestBaseImpl) RestoreArtifactsBulkRequestBase() BaseRestoreArtifactsBulkRequestBaseImpl {
	return s.restoreArtifactsBulkRequestBase
}

func (s RawRestoreArtifactsBulkRequestBaseImpl) Entity() BaseEntityImpl {
	return s.restoreArtifactsBulkRequestBase.Entity()
}

var _ json.Marshaler = BaseRestoreArtifactsBulkRequestBaseImpl{}

func (s BaseRestoreArtifactsBulkRequestBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseRestoreArtifactsBulkRequestBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseRestoreArtifactsBulkRequestBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseRestoreArtifactsBulkRequestBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.restoreArtifactsBulkRequestBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseRestoreArtifactsBulkRequestBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseRestoreArtifactsBulkRequestBaseImpl{}

func (s *BaseRestoreArtifactsBulkRequestBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
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

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DestinationType = decoded.DestinationType
	s.DisplayName = decoded.DisplayName
	s.Error = decoded.Error
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ProtectionTimePeriod = decoded.ProtectionTimePeriod
	s.ProtectionUnitIds = decoded.ProtectionUnitIds
	s.RestorePointPreference = decoded.RestorePointPreference
	s.Status = decoded.Status
	s.Tags = decoded.Tags
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseRestoreArtifactsBulkRequestBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseRestoreArtifactsBulkRequestBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseRestoreArtifactsBulkRequestBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalRestoreArtifactsBulkRequestBaseImplementation(input []byte) (RestoreArtifactsBulkRequestBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RestoreArtifactsBulkRequestBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveRestoreArtifactsBulkAdditionRequest") {
		var out DriveRestoreArtifactsBulkAdditionRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveRestoreArtifactsBulkAdditionRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxRestoreArtifactsBulkAdditionRequest") {
		var out MailboxRestoreArtifactsBulkAdditionRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxRestoreArtifactsBulkAdditionRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteRestoreArtifactsBulkAdditionRequest") {
		var out SiteRestoreArtifactsBulkAdditionRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteRestoreArtifactsBulkAdditionRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseRestoreArtifactsBulkRequestBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRestoreArtifactsBulkRequestBaseImpl: %+v", err)
	}

	return RawRestoreArtifactsBulkRequestBaseImpl{
		restoreArtifactsBulkRequestBase: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
