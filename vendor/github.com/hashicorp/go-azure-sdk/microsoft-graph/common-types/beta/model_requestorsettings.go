package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestorSettings struct {
	// Indicates whether new requests are accepted on this policy.
	AcceptRequests nullable.Type[bool] `json:"acceptRequests,omitempty"`

	// The users who are allowed to request on this policy, which can be singleUser, groupMembers, and
	// connectedOrganizationMembers.
	AllowedRequestors *[]UserSet `json:"allowedRequestors,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Who can request. One of NoSubjects, SpecificDirectorySubjects, SpecificConnectedOrganizationSubjects,
	// AllConfiguredConnectedOrganizationSubjects, AllExistingConnectedOrganizationSubjects,
	// AllExistingDirectoryMemberUsers, AllExistingDirectorySubjects or AllExternalSubjects.
	ScopeType nullable.Type[string] `json:"scopeType,omitempty"`
}

var _ json.Unmarshaler = &RequestorSettings{}

func (s *RequestorSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AcceptRequests nullable.Type[bool]   `json:"acceptRequests,omitempty"`
		ODataId        *string               `json:"@odata.id,omitempty"`
		ODataType      *string               `json:"@odata.type,omitempty"`
		ScopeType      nullable.Type[string] `json:"scopeType,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AcceptRequests = decoded.AcceptRequests
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.ScopeType = decoded.ScopeType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RequestorSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["allowedRequestors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AllowedRequestors into list []json.RawMessage: %+v", err)
		}

		output := make([]UserSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalUserSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AllowedRequestors' for 'RequestorSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AllowedRequestors = &output
	}

	return nil
}
