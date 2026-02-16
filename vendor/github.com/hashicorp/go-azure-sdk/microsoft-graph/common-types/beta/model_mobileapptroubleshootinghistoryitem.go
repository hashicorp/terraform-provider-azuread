package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingHistoryItem interface {
	MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl
}

var _ MobileAppTroubleshootingHistoryItem = BaseMobileAppTroubleshootingHistoryItemImpl{}

type BaseMobileAppTroubleshootingHistoryItemImpl struct {
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

func (s BaseMobileAppTroubleshootingHistoryItemImpl) MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl {
	return s
}

var _ MobileAppTroubleshootingHistoryItem = RawMobileAppTroubleshootingHistoryItemImpl{}

// RawMobileAppTroubleshootingHistoryItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMobileAppTroubleshootingHistoryItemImpl struct {
	mobileAppTroubleshootingHistoryItem BaseMobileAppTroubleshootingHistoryItemImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawMobileAppTroubleshootingHistoryItemImpl) MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl {
	return s.mobileAppTroubleshootingHistoryItem
}

func UnmarshalMobileAppTroubleshootingHistoryItemImplementation(input []byte) (MobileAppTroubleshootingHistoryItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppTroubleshootingHistoryItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingAppPolicyCreationHistory") {
		var out MobileAppTroubleshootingAppPolicyCreationHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingAppPolicyCreationHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingAppStateHistory") {
		var out MobileAppTroubleshootingAppStateHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingAppStateHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingAppTargetHistory") {
		var out MobileAppTroubleshootingAppTargetHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingAppTargetHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingAppUpdateHistory") {
		var out MobileAppTroubleshootingAppUpdateHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingAppUpdateHistory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.mobileAppTroubleshootingDeviceCheckinHistory") {
		var out MobileAppTroubleshootingDeviceCheckinHistory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MobileAppTroubleshootingDeviceCheckinHistory: %+v", err)
		}
		return out, nil
	}

	var parent BaseMobileAppTroubleshootingHistoryItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMobileAppTroubleshootingHistoryItemImpl: %+v", err)
	}

	return RawMobileAppTroubleshootingHistoryItemImpl{
		mobileAppTroubleshootingHistoryItem: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
