package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionUnitBase interface {
	Entity
	ProtectionUnitBase() BaseProtectionUnitBaseImpl
}

var _ ProtectionUnitBase = BaseProtectionUnitBaseImpl{}

type BaseProtectionUnitBaseImpl struct {
	// The identity of person who created the protection unit.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the protection unit.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains error details if an error occurred while creating a protection unit.
	Error *PublicError `json:"error,omitempty"`

	// The identity of person who last modified the protection unit.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification of this protection unit.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The unique identifier of the protection policy based on which protection unit was created.
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// The status of the protection unit. The possible values are: protectRequested, protected, unprotectRequested,
	// unprotected, removeRequested, unknownFutureValue.
	Status *ProtectionUnitStatus `json:"status,omitempty"`

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

func (s BaseProtectionUnitBaseImpl) ProtectionUnitBase() BaseProtectionUnitBaseImpl {
	return s
}

func (s BaseProtectionUnitBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ProtectionUnitBase = RawProtectionUnitBaseImpl{}

// RawProtectionUnitBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionUnitBaseImpl struct {
	protectionUnitBase BaseProtectionUnitBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawProtectionUnitBaseImpl) ProtectionUnitBase() BaseProtectionUnitBaseImpl {
	return s.protectionUnitBase
}

func (s RawProtectionUnitBaseImpl) Entity() BaseEntityImpl {
	return s.protectionUnitBase.Entity()
}

var _ json.Marshaler = BaseProtectionUnitBaseImpl{}

func (s BaseProtectionUnitBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseProtectionUnitBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseProtectionUnitBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseProtectionUnitBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectionUnitBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseProtectionUnitBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseProtectionUnitBaseImpl{}

func (s *BaseProtectionUnitBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Error                *PublicError          `json:"error,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		PolicyId             nullable.Type[string] `json:"policyId,omitempty"`
		Status               *ProtectionUnitStatus `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.PolicyId = decoded.PolicyId
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseProtectionUnitBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseProtectionUnitBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseProtectionUnitBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalProtectionUnitBaseImplementation(input []byte) (ProtectionUnitBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionUnitBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionUnit") {
		var out DriveProtectionUnit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionUnit: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionUnit") {
		var out MailboxProtectionUnit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionUnit: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionUnit") {
		var out SiteProtectionUnit
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionUnit: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionUnitBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionUnitBaseImpl: %+v", err)
	}

	return RawProtectionUnitBaseImpl{
		protectionUnitBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
