package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAuditData = SecurityDisableConsentRecord{}

type SecurityDisableConsentRecord struct {

	// Fields inherited from SecurityAuditData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityDisableConsentRecord) SecurityAuditData() BaseSecurityAuditDataImpl {
	return BaseSecurityAuditDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityDisableConsentRecord{}

func (s SecurityDisableConsentRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityDisableConsentRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityDisableConsentRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityDisableConsentRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.disableConsentRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityDisableConsentRecord: %+v", err)
	}

	return encoded, nil
}
