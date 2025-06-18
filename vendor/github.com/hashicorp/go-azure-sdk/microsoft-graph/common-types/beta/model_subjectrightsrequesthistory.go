package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistory struct {
	// Identity of the user who changed the subject rights request.
	ChangedBy IdentitySet `json:"changedBy"`

	// Data and time when the entity was changed.
	EventDateTime nullable.Type[string] `json:"eventDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport,
	// contentDeletion, caseResolved, unknownFutureValue, approval. Use the Prefer: include-unknown-enum-members request
	// header to get the following value(s) in this evolvable enum: approval.
	Stage *SubjectRightsRequestStage `json:"stage,omitempty"`

	// The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed,
	// unknownFutureValue.
	StageStatus *SubjectRightsRequestStageStatus `json:"stageStatus,omitempty"`

	// Type of history.
	Type nullable.Type[string] `json:"type,omitempty"`
}

var _ json.Unmarshaler = &SubjectRightsRequestHistory{}

func (s *SubjectRightsRequestHistory) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EventDateTime nullable.Type[string]            `json:"eventDateTime,omitempty"`
		ODataId       *string                          `json:"@odata.id,omitempty"`
		ODataType     *string                          `json:"@odata.type,omitempty"`
		Stage         *SubjectRightsRequestStage       `json:"stage,omitempty"`
		StageStatus   *SubjectRightsRequestStageStatus `json:"stageStatus,omitempty"`
		Type          nullable.Type[string]            `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EventDateTime = decoded.EventDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Stage = decoded.Stage
	s.StageStatus = decoded.StageStatus
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SubjectRightsRequestHistory into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["changedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ChangedBy' for 'SubjectRightsRequestHistory': %+v", err)
		}
		s.ChangedBy = impl
	}

	return nil
}
