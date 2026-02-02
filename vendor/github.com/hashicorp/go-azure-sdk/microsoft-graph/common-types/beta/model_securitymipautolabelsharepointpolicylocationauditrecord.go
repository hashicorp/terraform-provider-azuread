package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAuditData = SecurityMipAutoLabelSharePointPolicyLocationAuditRecord{}

type SecurityMipAutoLabelSharePointPolicyLocationAuditRecord struct {

	// Fields inherited from SecurityAuditData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityMipAutoLabelSharePointPolicyLocationAuditRecord) SecurityAuditData() BaseSecurityAuditDataImpl {
	return BaseSecurityAuditDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityMipAutoLabelSharePointPolicyLocationAuditRecord{}

func (s SecurityMipAutoLabelSharePointPolicyLocationAuditRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityMipAutoLabelSharePointPolicyLocationAuditRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityMipAutoLabelSharePointPolicyLocationAuditRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMipAutoLabelSharePointPolicyLocationAuditRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.mipAutoLabelSharePointPolicyLocationAuditRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityMipAutoLabelSharePointPolicyLocationAuditRecord: %+v", err)
	}

	return encoded, nil
}
