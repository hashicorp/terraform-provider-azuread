package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CallRecordsCallRecord{}

type CallRecordsCallRecord struct {
	// UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Meeting URL associated to the call. May not be available for a peerToPeer call record type.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

	// UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing,
	// data, screenSharing, unknownFutureValue.
	Modalities *[]CallRecordsModality `json:"modalities,omitempty"`

	// The organizing party's identity. The organizer property is deprecated and will stop returning data on June 30, 2026.
	// Going forward, use the organizer_v2 relationship.
	Organizer IdentitySet `json:"organizer"`

	// Identity of the organizer of the call. This relationship is expanded by default in callRecord methods.
	Organizerv2 *CallRecordsOrganizer `json:"organizer_v2,omitempty"`

	// List of distinct identities involved in the call. Limited to 130 entries. The participants property is deprecated and
	// will stop returning data on June 30, 2026. Going forward, use the participants_v2 relationship.
	Participants *[]IdentitySet `json:"participants,omitempty"`

	// List of distinct participants in the call.
	Participantsv2 *[]CallRecordsParticipant `json:"participants_v2,omitempty"`

	// List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls
	// typically have at least one session per participant. Read-only. Nullable.
	Sessions *[]CallRecordsSession `json:"sessions,omitempty"`

	// UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime *string `json:"startDateTime,omitempty"`

	Type *CallRecordsCallType `json:"type,omitempty"`

	// Monotonically increasing version of the call record. Higher version call records with the same ID include additional
	// data compared to the lower version.
	Version *int64 `json:"version,omitempty"`

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

func (s CallRecordsCallRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsCallRecord{}

func (s CallRecordsCallRecord) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsCallRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsCallRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsCallRecord: %+v", err)
	}

	delete(decoded, "sessions")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.callRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsCallRecord: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CallRecordsCallRecord{}

func (s *CallRecordsCallRecord) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EndDateTime          *string                   `json:"endDateTime,omitempty"`
		JoinWebUrl           nullable.Type[string]     `json:"joinWebUrl,omitempty"`
		LastModifiedDateTime *string                   `json:"lastModifiedDateTime,omitempty"`
		Modalities           *[]CallRecordsModality    `json:"modalities,omitempty"`
		Organizerv2          *CallRecordsOrganizer     `json:"organizer_v2,omitempty"`
		Participantsv2       *[]CallRecordsParticipant `json:"participants_v2,omitempty"`
		Sessions             *[]CallRecordsSession     `json:"sessions,omitempty"`
		StartDateTime        *string                   `json:"startDateTime,omitempty"`
		Type                 *CallRecordsCallType      `json:"type,omitempty"`
		Version              *int64                    `json:"version,omitempty"`
		Id                   *string                   `json:"id,omitempty"`
		ODataId              *string                   `json:"@odata.id,omitempty"`
		ODataType            *string                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EndDateTime = decoded.EndDateTime
	s.JoinWebUrl = decoded.JoinWebUrl
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Modalities = decoded.Modalities
	s.Organizerv2 = decoded.Organizerv2
	s.Participantsv2 = decoded.Participantsv2
	s.Sessions = decoded.Sessions
	s.StartDateTime = decoded.StartDateTime
	s.Type = decoded.Type
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CallRecordsCallRecord into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["organizer"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Organizer' for 'CallRecordsCallRecord': %+v", err)
		}
		s.Organizer = impl
	}

	if v, ok := temp["participants"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Participants into list []json.RawMessage: %+v", err)
		}

		output := make([]IdentitySet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentitySetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Participants' for 'CallRecordsCallRecord': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Participants = &output
	}

	return nil
}
