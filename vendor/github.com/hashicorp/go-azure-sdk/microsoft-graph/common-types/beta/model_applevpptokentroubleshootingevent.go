package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementTroubleshootingEvent = AppleVppTokenTroubleshootingEvent{}

type AppleVppTokenTroubleshootingEvent struct {
	// Apple Volume Purchase Program Token Identifier.
	TokenId nullable.Type[string] `json:"tokenId,omitempty"`

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

func (s AppleVppTokenTroubleshootingEvent) DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl {
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

func (s AppleVppTokenTroubleshootingEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppleVppTokenTroubleshootingEvent{}

func (s AppleVppTokenTroubleshootingEvent) MarshalJSON() ([]byte, error) {
	type wrapper AppleVppTokenTroubleshootingEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppleVppTokenTroubleshootingEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppleVppTokenTroubleshootingEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appleVppTokenTroubleshootingEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppleVppTokenTroubleshootingEvent: %+v", err)
	}

	return encoded, nil
}
