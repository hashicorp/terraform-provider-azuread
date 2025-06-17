package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionRuleBase = SiteProtectionRule{}

type SiteProtectionRule struct {
	// Contains a site expression. For examples, see siteExpression example.
	SiteExpression nullable.Type[string] `json:"siteExpression,omitempty"`

	// Fields inherited from ProtectionRuleBase

	// The identity of person who created the rule.
	CreatedBy IdentitySet `json:"createdBy"`

	// The time of creation of the rule.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Contains error details if an operation on a rule fails.
	Error *PublicError `json:"error,omitempty"`

	IsAutoApplyEnabled nullable.Type[bool] `json:"isAutoApplyEnabled,omitempty"`

	// The identity of the person who last modified the rule.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Timestamp of the last modification made to the rule.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The status of the protection rule. The possible values are: draft, active, completed, completedWithErrors,
	// unknownFutureValue. The draft member is currently unsupported.
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

func (s SiteProtectionRule) ProtectionRuleBase() BaseProtectionRuleBaseImpl {
	return BaseProtectionRuleBaseImpl{
		CreatedBy:            s.CreatedBy,
		CreatedDateTime:      s.CreatedDateTime,
		Error:                s.Error,
		IsAutoApplyEnabled:   s.IsAutoApplyEnabled,
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Status:               s.Status,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s SiteProtectionRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SiteProtectionRule{}

func (s SiteProtectionRule) MarshalJSON() ([]byte, error) {
	type wrapper SiteProtectionRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SiteProtectionRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SiteProtectionRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.siteProtectionRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SiteProtectionRule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SiteProtectionRule{}

func (s *SiteProtectionRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SiteExpression       nullable.Type[string] `json:"siteExpression,omitempty"`
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

	s.SiteExpression = decoded.SiteExpression
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Error = decoded.Error
	s.Id = decoded.Id
	s.IsAutoApplyEnabled = decoded.IsAutoApplyEnabled
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SiteProtectionRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SiteProtectionRule': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SiteProtectionRule': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
