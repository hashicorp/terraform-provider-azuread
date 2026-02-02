package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesComplianceChangeRule = WindowsUpdatesContentApprovalRule{}

type WindowsUpdatesContentApprovalRule struct {
	// A filter to determine which content matches the rule on an ongoing basis.
	ContentFilter WindowsUpdatesContentFilter `json:"contentFilter"`

	// The time before the deployment starts represented in ISO 8601 format for durations.
	DurationBeforeDeploymentStart nullable.Type[string] `json:"durationBeforeDeploymentStart,omitempty"`

	// Fields inherited from WindowsUpdatesComplianceChangeRule

	// The date and time when the rule was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date and time when the rule was last evaluated.
	LastEvaluatedDateTime nullable.Type[string] `json:"lastEvaluatedDateTime,omitempty"`

	// The date and time when the rule was last modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesContentApprovalRule) WindowsUpdatesComplianceChangeRule() BaseWindowsUpdatesComplianceChangeRuleImpl {
	return BaseWindowsUpdatesComplianceChangeRuleImpl{
		CreatedDateTime:       s.CreatedDateTime,
		LastEvaluatedDateTime: s.LastEvaluatedDateTime,
		LastModifiedDateTime:  s.LastModifiedDateTime,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesContentApprovalRule{}

func (s WindowsUpdatesContentApprovalRule) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesContentApprovalRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesContentApprovalRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesContentApprovalRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.contentApprovalRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesContentApprovalRule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsUpdatesContentApprovalRule{}

func (s *WindowsUpdatesContentApprovalRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DurationBeforeDeploymentStart nullable.Type[string] `json:"durationBeforeDeploymentStart,omitempty"`
		CreatedDateTime               nullable.Type[string] `json:"createdDateTime,omitempty"`
		LastEvaluatedDateTime         nullable.Type[string] `json:"lastEvaluatedDateTime,omitempty"`
		LastModifiedDateTime          nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ODataId                       *string               `json:"@odata.id,omitempty"`
		ODataType                     *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DurationBeforeDeploymentStart = decoded.DurationBeforeDeploymentStart
	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastEvaluatedDateTime = decoded.LastEvaluatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsUpdatesContentApprovalRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["contentFilter"]; ok {
		impl, err := UnmarshalWindowsUpdatesContentFilterImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ContentFilter' for 'WindowsUpdatesContentApprovalRule': %+v", err)
		}
		s.ContentFilter = impl
	}

	return nil
}
