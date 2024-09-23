package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppTroubleshootingHistoryItem = MobileAppTroubleshootingAppStateHistory{}

type MobileAppTroubleshootingAppStateHistory struct {
	// Defines the Action Types for an Intune Application.
	ActionType *MobileAppActionType `json:"actionType,omitempty"`

	// Error code for the failure, empty if no failure.
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// Indicates the type of execution status of the device management script.
	RunState *RunState `json:"runState,omitempty"`

	// Fields inherited from MobileAppTroubleshootingHistoryItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time when the history item occurred.
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// Object containing detailed information about the error and its remediation.
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s MobileAppTroubleshootingAppStateHistory) MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl {
	return BaseMobileAppTroubleshootingHistoryItemImpl{
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
		OccurrenceDateTime:          s.OccurrenceDateTime,
		TroubleshootingErrorDetails: s.TroubleshootingErrorDetails,
	}
}

var _ json.Marshaler = MobileAppTroubleshootingAppStateHistory{}

func (s MobileAppTroubleshootingAppStateHistory) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppTroubleshootingAppStateHistory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppTroubleshootingAppStateHistory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppTroubleshootingAppStateHistory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppTroubleshootingAppStateHistory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppTroubleshootingAppStateHistory: %+v", err)
	}

	return encoded, nil
}
