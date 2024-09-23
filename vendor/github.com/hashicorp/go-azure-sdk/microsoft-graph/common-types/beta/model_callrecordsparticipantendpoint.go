package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsEndpoint = CallRecordsParticipantEndpoint{}

type CallRecordsParticipantEndpoint struct {
	// Identity associated with the endpoint.
	AssociatedIdentity Identity `json:"associatedIdentity"`

	// CPU number of cores used by the media endpoint.
	CpuCoresCount nullable.Type[int64] `json:"cpuCoresCount,omitempty"`

	// CPU name used by the media endpoint.
	CpuName nullable.Type[string] `json:"cpuName,omitempty"`

	// CPU processor speed used by the media endpoint.
	CpuProcessorSpeedInMhz nullable.Type[int64] `json:"cpuProcessorSpeedInMhz,omitempty"`

	// The feedback provided by the user of this endpoint about the quality of the session.
	Feedback *CallRecordsUserFeedback `json:"feedback,omitempty"`

	// Identity associated with the endpoint. The identity property is deprecated and will stop returning data on June 30,
	// 2026. Going forward, use the associatedIdentity property.
	Identity IdentitySet `json:"identity"`

	// Name of the device used by the media endpoint.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Fields inherited from CallRecordsEndpoint

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// User-agent reported by this endpoint.
	UserAgent CallRecordsUserAgent `json:"userAgent"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CallRecordsParticipantEndpoint) CallRecordsEndpoint() BaseCallRecordsEndpointImpl {
	return BaseCallRecordsEndpointImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		UserAgent: s.UserAgent,
	}
}

var _ json.Marshaler = CallRecordsParticipantEndpoint{}

func (s CallRecordsParticipantEndpoint) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsParticipantEndpoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsParticipantEndpoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsParticipantEndpoint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.participantEndpoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsParticipantEndpoint: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordsParticipantEndpoint{}

func (s *CallRecordsParticipantEndpoint) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CpuCoresCount          nullable.Type[int64]     `json:"cpuCoresCount,omitempty"`
		CpuName                nullable.Type[string]    `json:"cpuName,omitempty"`
		CpuProcessorSpeedInMhz nullable.Type[int64]     `json:"cpuProcessorSpeedInMhz,omitempty"`
		Feedback               *CallRecordsUserFeedback `json:"feedback,omitempty"`
		Name                   nullable.Type[string]    `json:"name,omitempty"`
		ODataId                *string                  `json:"@odata.id,omitempty"`
		ODataType              *string                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CpuCoresCount = decoded.CpuCoresCount
	s.CpuName = decoded.CpuName
	s.CpuProcessorSpeedInMhz = decoded.CpuProcessorSpeedInMhz
	s.Feedback = decoded.Feedback
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRecordsParticipantEndpoint into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["associatedIdentity"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssociatedIdentity' for 'CallRecordsParticipantEndpoint': %+v", err)
		}
		s.AssociatedIdentity = impl
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'CallRecordsParticipantEndpoint': %+v", err)
		}
		s.Identity = impl
	}

	if v, ok := temp["userAgent"]; ok {
		impl, err := UnmarshalCallRecordsUserAgentImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'UserAgent' for 'CallRecordsParticipantEndpoint': %+v", err)
		}
		s.UserAgent = impl
	}

	return nil
}
