package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAuditData = SecurityMicrosoftPurviewPrivacyAuditEvent{}

type SecurityMicrosoftPurviewPrivacyAuditEvent struct {

	// Fields inherited from SecurityAuditData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityMicrosoftPurviewPrivacyAuditEvent) SecurityAuditData() BaseSecurityAuditDataImpl {
	return BaseSecurityAuditDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityMicrosoftPurviewPrivacyAuditEvent{}

func (s SecurityMicrosoftPurviewPrivacyAuditEvent) MarshalJSON() ([]byte, error) {
	type wrapper SecurityMicrosoftPurviewPrivacyAuditEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityMicrosoftPurviewPrivacyAuditEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMicrosoftPurviewPrivacyAuditEvent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.microsoftPurviewPrivacyAuditEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityMicrosoftPurviewPrivacyAuditEvent: %+v", err)
	}

	return encoded, nil
}
