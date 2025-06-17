package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAuditData = SecurityMipAutoLabelProgressFeedbackAuditRecord{}

type SecurityMipAutoLabelProgressFeedbackAuditRecord struct {

	// Fields inherited from SecurityAuditData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityMipAutoLabelProgressFeedbackAuditRecord) SecurityAuditData() BaseSecurityAuditDataImpl {
	return BaseSecurityAuditDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityMipAutoLabelProgressFeedbackAuditRecord{}

func (s SecurityMipAutoLabelProgressFeedbackAuditRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityMipAutoLabelProgressFeedbackAuditRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityMipAutoLabelProgressFeedbackAuditRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMipAutoLabelProgressFeedbackAuditRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.mipAutoLabelProgressFeedbackAuditRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityMipAutoLabelProgressFeedbackAuditRecord: %+v", err)
	}

	return encoded, nil
}
