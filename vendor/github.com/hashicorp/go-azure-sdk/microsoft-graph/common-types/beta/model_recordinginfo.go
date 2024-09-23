package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingInfo struct {
	// The participant who initiated the recording.
	InitiatedBy *ParticipantInfo `json:"initiatedBy,omitempty"`

	// The identities of recording initiator.
	Initiator IdentitySet `json:"initiator"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RecordingStatus *RecordingStatus `json:"recordingStatus,omitempty"`
}

var _ json.Unmarshaler = &RecordingInfo{}

func (s *RecordingInfo) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		InitiatedBy     *ParticipantInfo `json:"initiatedBy,omitempty"`
		ODataId         *string          `json:"@odata.id,omitempty"`
		ODataType       *string          `json:"@odata.type,omitempty"`
		RecordingStatus *RecordingStatus `json:"recordingStatus,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InitiatedBy = decoded.InitiatedBy
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecordingStatus = decoded.RecordingStatus

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RecordingInfo into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'RecordingInfo': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
