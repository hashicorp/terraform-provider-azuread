package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreApprovalDetail struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Permissions       PreApprovedPermissions `json:"permissions"`
	ScopeType         *ResourceScopeType     `json:"scopeType,omitempty"`
	SensitivityLabels ScopeSensitivityLabels `json:"sensitivityLabels"`
}

var _ json.Unmarshaler = &PreApprovalDetail{}

func (s *PreApprovalDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string            `json:"@odata.id,omitempty"`
		ODataType *string            `json:"@odata.type,omitempty"`
		ScopeType *ResourceScopeType `json:"scopeType,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ScopeType = decoded.ScopeType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PreApprovalDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["permissions"]; ok {
		impl, err := UnmarshalPreApprovedPermissionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Permissions' for 'PreApprovalDetail': %+v", err)
		}
		s.Permissions = impl
	}

	if v, ok := temp["sensitivityLabels"]; ok {
		impl, err := UnmarshalScopeSensitivityLabelsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SensitivityLabels' for 'PreApprovalDetail': %+v", err)
		}
		s.SensitivityLabels = impl
	}

	return nil
}
