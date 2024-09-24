package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionUnitBase = MailboxProtectionUnit{}

type MailboxProtectionUnit struct {
	// The ID of the directory object.
	DirectoryObjectId nullable.Type[string] `json:"directoryObjectId,omitempty"`

	// Display name of the directory object.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Email address associated with the directory object.
	Email nullable.Type[string] `json:"email,omitempty"`

	// Fields inherited from ProtectionUnitBase

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

func (s MailboxProtectionUnit) ProtectionUnitBase() BaseProtectionUnitBaseImpl {
	return BaseProtectionUnitBaseImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Error:                s.Error,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		PolicyId:             s.PolicyId,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s MailboxProtectionUnit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MailboxProtectionUnit{}

func (s MailboxProtectionUnit) MarshalJSON() ([]byte, error) {
	type wrapper MailboxProtectionUnit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MailboxProtectionUnit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MailboxProtectionUnit: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "email")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mailboxProtectionUnit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MailboxProtectionUnit: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MailboxProtectionUnit{}

func (s *MailboxProtectionUnit) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DirectoryObjectId    nullable.Type[string] `json:"directoryObjectId,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		Email                nullable.Type[string] `json:"email,omitempty"`
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

	s.DirectoryObjectId = decoded.DirectoryObjectId
	s.DisplayName = decoded.DisplayName
	s.Email = decoded.Email
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PolicyId = decoded.PolicyId
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MailboxProtectionUnit into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'MailboxProtectionUnit': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'MailboxProtectionUnit': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
