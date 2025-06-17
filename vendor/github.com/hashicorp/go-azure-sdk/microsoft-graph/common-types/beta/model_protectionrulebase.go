package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionRuleBase interface {
	Entity
	ProtectionRuleBase() BaseProtectionRuleBaseImpl
}

var _ ProtectionRuleBase = BaseProtectionRuleBaseImpl{}

type BaseProtectionRuleBaseImpl struct {
	// The identity of person who created the rule.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the rule.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains error details if an operation on a rule fails.
	Error *PublicError `json:"error,omitempty"`

	// true indicates that the protection rule is dynamic; false that it's static.
	IsAutoApplyEnabled nullable.Type[bool] `json:"isAutoApplyEnabled,omitempty"`

	// The identity of the person who last modified the rule.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification made to the rule.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The status of the protection rule. The possible values are: draft, active, completed, completedWithErrors,
	// unknownFutureValue, updateRequested, deleteRequested. Use the Prefer: include-unknown-enum-members request header to
	// get the following values in this evolvable enum: updateRequested , deleteRequested. The draft member is currently
	// unsupported.
	Status *ProtectionRuleStatus `json:"status,omitempty"`

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

func (s BaseProtectionRuleBaseImpl) ProtectionRuleBase() BaseProtectionRuleBaseImpl {
	return s
}

func (s BaseProtectionRuleBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ProtectionRuleBase = RawProtectionRuleBaseImpl{}

// RawProtectionRuleBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProtectionRuleBaseImpl struct {
	protectionRuleBase BaseProtectionRuleBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawProtectionRuleBaseImpl) ProtectionRuleBase() BaseProtectionRuleBaseImpl {
	return s.protectionRuleBase
}

func (s RawProtectionRuleBaseImpl) Entity() BaseEntityImpl {
	return s.protectionRuleBase.Entity()
}

var _ json.Marshaler = BaseProtectionRuleBaseImpl{}

func (s BaseProtectionRuleBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseProtectionRuleBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseProtectionRuleBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseProtectionRuleBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectionRuleBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseProtectionRuleBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseProtectionRuleBaseImpl{}

func (s *BaseProtectionRuleBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Error                *PublicError          `json:"error,omitempty"`
		IsAutoApplyEnabled   nullable.Type[bool]   `json:"isAutoApplyEnabled,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Status               *ProtectionRuleStatus `json:"status,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.IsAutoApplyEnabled = decoded.IsAutoApplyEnabled
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseProtectionRuleBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BaseProtectionRuleBaseImpl': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseProtectionRuleBaseImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalProtectionRuleBaseImplementation(input []byte) (ProtectionRuleBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectionRuleBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.driveProtectionRule") {
		var out DriveProtectionRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DriveProtectionRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mailboxProtectionRule") {
		var out MailboxProtectionRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MailboxProtectionRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.siteProtectionRule") {
		var out SiteProtectionRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SiteProtectionRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseProtectionRuleBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseProtectionRuleBaseImpl: %+v", err)
	}

	return RawProtectionRuleBaseImpl{
		protectionRuleBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
