package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpEvaluatePoliciesRequest struct {
	EvaluationInput  DlpEvaluationInput `json:"evaluationInput"`
	NotificationInfo DlpNotification    `json:"notificationInfo"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Target nullable.Type[string] `json:"target,omitempty"`
}

var _ json.Unmarshaler = &DlpEvaluatePoliciesRequest{}

func (s *DlpEvaluatePoliciesRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		Target    nullable.Type[string] `json:"target,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Target = decoded.Target

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DlpEvaluatePoliciesRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["evaluationInput"]; ok {
		impl, err := UnmarshalDlpEvaluationInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EvaluationInput' for 'DlpEvaluatePoliciesRequest': %+v", err)
		}
		s.EvaluationInput = impl
	}

	if v, ok := temp["notificationInfo"]; ok {
		impl, err := UnmarshalDlpNotificationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'NotificationInfo' for 'DlpEvaluatePoliciesRequest': %+v", err)
		}
		s.NotificationInfo = impl
	}

	return nil
}
