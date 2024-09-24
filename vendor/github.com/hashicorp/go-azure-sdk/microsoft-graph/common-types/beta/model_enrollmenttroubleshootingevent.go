package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementTroubleshootingEvent = EnrollmentTroubleshootingEvent{}

type EnrollmentTroubleshootingEvent struct {
	// Azure AD device identifier.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// Possible ways of adding a mobile device to management.
	EnrollmentType *DeviceEnrollmentType `json:"enrollmentType,omitempty"`

	// Top level failure categories for enrollment.
	FailureCategory *DeviceEnrollmentFailureReason `json:"failureCategory,omitempty"`

	// Detailed failure reason.
	FailureReason nullable.Type[string] `json:"failureReason,omitempty"`

	// Device identifier created or collected by Intune.
	ManagedDeviceIdentifier nullable.Type[string] `json:"managedDeviceIdentifier,omitempty"`

	// Operating System.
	OperatingSystem nullable.Type[string] `json:"operatingSystem,omitempty"`

	// OS Version.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Identifier for the user that tried to enroll the device.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from DeviceManagementTroubleshootingEvent

	// A set of string key and string value pairs which provides additional information on the Troubleshooting event
	AdditionalInformation *[]KeyValuePair `json:"additionalInformation,omitempty"`

	// Id used for tracing the failure in the service.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Time when the event occurred .
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// Event Name corresponding to the Troubleshooting Event. It is an Optional field
	EventName nullable.Type[string] `json:"eventName,omitempty"`

	// Object containing detailed information about the error and its remediation.
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`

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

func (s EnrollmentTroubleshootingEvent) DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl {
	return BaseDeviceManagementTroubleshootingEventImpl{
		AdditionalInformation:       s.AdditionalInformation,
		CorrelationId:               s.CorrelationId,
		EventDateTime:               s.EventDateTime,
		EventName:                   s.EventName,
		TroubleshootingErrorDetails: s.TroubleshootingErrorDetails,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s EnrollmentTroubleshootingEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EnrollmentTroubleshootingEvent{}

func (s EnrollmentTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	type wrapper EnrollmentTroubleshootingEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnrollmentTroubleshootingEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnrollmentTroubleshootingEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enrollmentTroubleshootingEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnrollmentTroubleshootingEvent: %+v", err)
	}

	return encoded, nil
}
