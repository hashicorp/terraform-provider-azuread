package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CallRecordsParticipantBase = CallRecordsParticipant{}

type CallRecordsParticipant struct {

	// Fields inherited from CallRecordsParticipantBase

	// List of administrativeUnitInfo of the call participant.
	AdministrativeUnitInfos *[]CallRecordsAdministrativeUnitInfo `json:"administrativeUnitInfos,omitempty"`

	// The identity of the call participant.
	Identity *CommunicationsIdentitySet `json:"identity,omitempty"`

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

func (s CallRecordsParticipant) CallRecordsParticipantBase() BaseCallRecordsParticipantBaseImpl {
	return BaseCallRecordsParticipantBaseImpl{
		AdministrativeUnitInfos: s.AdministrativeUnitInfos,
		Identity:                s.Identity,
		Id:                      s.Id,
		ODataId:                 s.ODataId,
		ODataType:               s.ODataType,
	}
}

func (s CallRecordsParticipant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CallRecordsParticipant{}

func (s CallRecordsParticipant) MarshalJSON() ([]byte, error) {
	type wrapper CallRecordsParticipant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CallRecordsParticipant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CallRecordsParticipant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.callRecords.participant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CallRecordsParticipant: %+v", err)
	}

	return encoded, nil
}
