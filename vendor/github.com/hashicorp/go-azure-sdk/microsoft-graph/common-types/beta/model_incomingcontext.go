package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncomingContext struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The id of the participant that is under observation. Read-only.
	ObservedParticipantId nullable.Type[string] `json:"observedParticipantId,omitempty"`

	// The identity that the call is happening on behalf of.
	OnBehalfOf IdentitySet `json:"onBehalfOf"`

	// The id of the participant that triggered the incoming call. Read-only.
	SourceParticipantId nullable.Type[string] `json:"sourceParticipantId,omitempty"`

	// The identity that transferred the call.
	Transferor IdentitySet `json:"transferor"`
}

var _ json.Marshaler = IncomingContext{}

func (s IncomingContext) MarshalJSON() ([]byte, error) {
	type wrapper IncomingContext
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IncomingContext: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IncomingContext: %+v", err)
	}

	delete(decoded, "observedParticipantId")
	delete(decoded, "sourceParticipantId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IncomingContext: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IncomingContext{}

func (s *IncomingContext) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId               *string               `json:"@odata.id,omitempty"`
		ODataType             *string               `json:"@odata.type,omitempty"`
		ObservedParticipantId nullable.Type[string] `json:"observedParticipantId,omitempty"`
		SourceParticipantId   nullable.Type[string] `json:"sourceParticipantId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ObservedParticipantId = decoded.ObservedParticipantId
	s.SourceParticipantId = decoded.SourceParticipantId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IncomingContext into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["onBehalfOf"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OnBehalfOf' for 'IncomingContext': %+v", err)
		}
		s.OnBehalfOf = impl
	}

	if v, ok := temp["transferor"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Transferor' for 'IncomingContext': %+v", err)
		}
		s.Transferor = impl
	}

	return nil
}
