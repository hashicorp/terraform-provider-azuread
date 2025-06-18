package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsUpdatesApprovalRule = WindowsUpdatesQualityUpdateApprovalRule{}

type WindowsUpdatesQualityUpdateApprovalRule struct {
	Cadence        *WindowsUpdatesQualityUpdateCadence        `json:"cadence,omitempty"`
	Classification *WindowsUpdatesQualityUpdateClassification `json:"classification,omitempty"`

	// Fields inherited from WindowsUpdatesApprovalRule

	DeferralInDays *int64 `json:"deferralInDays,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsUpdatesQualityUpdateApprovalRule) WindowsUpdatesApprovalRule() BaseWindowsUpdatesApprovalRuleImpl {
	return BaseWindowsUpdatesApprovalRuleImpl{
		DeferralInDays: s.DeferralInDays,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesQualityUpdateApprovalRule{}

func (s WindowsUpdatesQualityUpdateApprovalRule) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesQualityUpdateApprovalRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesQualityUpdateApprovalRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesQualityUpdateApprovalRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.qualityUpdateApprovalRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesQualityUpdateApprovalRule: %+v", err)
	}

	return encoded, nil
}
