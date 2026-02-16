package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppTroubleshootingHistoryItem = MobileAppTroubleshootingAppUpdateHistory{}

type MobileAppTroubleshootingAppUpdateHistory struct {

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

func (s MobileAppTroubleshootingAppUpdateHistory) MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl {
	return BaseMobileAppTroubleshootingHistoryItemImpl{
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
		OccurrenceDateTime:          s.OccurrenceDateTime,
		TroubleshootingErrorDetails: s.TroubleshootingErrorDetails,
	}
}

var _ json.Marshaler = MobileAppTroubleshootingAppUpdateHistory{}

func (s MobileAppTroubleshootingAppUpdateHistory) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppTroubleshootingAppUpdateHistory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppTroubleshootingAppUpdateHistory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppTroubleshootingAppUpdateHistory: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppTroubleshootingAppUpdateHistory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppTroubleshootingAppUpdateHistory: %+v", err)
	}

	return encoded, nil
}
