package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingInfo struct {
	// The identities of the recording initiator.
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
		ODataId         *string          `json:"@odata.id,omitempty"`
		ODataType       *string          `json:"@odata.type,omitempty"`
		RecordingStatus *RecordingStatus `json:"recordingStatus,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

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
