package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SecurityAuditData = SecurityMeshWorldsAuditRecord{}

type SecurityMeshWorldsAuditRecord struct {

	// Fields inherited from SecurityAuditData

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityMeshWorldsAuditRecord) SecurityAuditData() BaseSecurityAuditDataImpl {
	return BaseSecurityAuditDataImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityMeshWorldsAuditRecord{}

func (s SecurityMeshWorldsAuditRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityMeshWorldsAuditRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityMeshWorldsAuditRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityMeshWorldsAuditRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.meshWorldsAuditRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityMeshWorldsAuditRecord: %+v", err)
	}

	return encoded, nil
}
