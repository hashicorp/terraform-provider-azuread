package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestorSettings struct {
	// False indicates that the requestor isn't permitted to include a schedule in their request.
	AllowCustomAssignmentSchedule nullable.Type[bool] `json:"allowCustomAssignmentSchedule,omitempty"`

	// True allows on-behalf-of requestors to create a request to add access for another principal.
	EnableOnBehalfRequestorsToAddAccess nullable.Type[bool] `json:"enableOnBehalfRequestorsToAddAccess,omitempty"`

	// True allows on-behalf-of requestors to create a request to remove access for another principal.
	EnableOnBehalfRequestorsToRemoveAccess nullable.Type[bool] `json:"enableOnBehalfRequestorsToRemoveAccess,omitempty"`

	// True allows on-behalf-of requestors to create a request to update access for another principal.
	EnableOnBehalfRequestorsToUpdateAccess nullable.Type[bool] `json:"enableOnBehalfRequestorsToUpdateAccess,omitempty"`

	// True allows requestors to create a request to add access for themselves.
	EnableTargetsToSelfAddAccess nullable.Type[bool] `json:"enableTargetsToSelfAddAccess,omitempty"`

	// True allows requestors to create a request to remove their access.
	EnableTargetsToSelfRemoveAccess nullable.Type[bool] `json:"enableTargetsToSelfRemoveAccess,omitempty"`

	// True allows requestors to create a request to update their access.
	EnableTargetsToSelfUpdateAccess nullable.Type[bool] `json:"enableTargetsToSelfUpdateAccess,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The principals who can request on-behalf-of others.
	OnBehalfRequestors *[]SubjectSet `json:"onBehalfRequestors,omitempty"`
}

var _ json.Unmarshaler = &AccessPackageAssignmentRequestorSettings{}

func (s *AccessPackageAssignmentRequestorSettings) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowCustomAssignmentSchedule          nullable.Type[bool] `json:"allowCustomAssignmentSchedule,omitempty"`
		EnableOnBehalfRequestorsToAddAccess    nullable.Type[bool] `json:"enableOnBehalfRequestorsToAddAccess,omitempty"`
		EnableOnBehalfRequestorsToRemoveAccess nullable.Type[bool] `json:"enableOnBehalfRequestorsToRemoveAccess,omitempty"`
		EnableOnBehalfRequestorsToUpdateAccess nullable.Type[bool] `json:"enableOnBehalfRequestorsToUpdateAccess,omitempty"`
		EnableTargetsToSelfAddAccess           nullable.Type[bool] `json:"enableTargetsToSelfAddAccess,omitempty"`
		EnableTargetsToSelfRemoveAccess        nullable.Type[bool] `json:"enableTargetsToSelfRemoveAccess,omitempty"`
		EnableTargetsToSelfUpdateAccess        nullable.Type[bool] `json:"enableTargetsToSelfUpdateAccess,omitempty"`
		ODataId                                *string             `json:"@odata.id,omitempty"`
		ODataType                              *string             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowCustomAssignmentSchedule = decoded.AllowCustomAssignmentSchedule
	s.EnableOnBehalfRequestorsToAddAccess = decoded.EnableOnBehalfRequestorsToAddAccess
	s.EnableOnBehalfRequestorsToRemoveAccess = decoded.EnableOnBehalfRequestorsToRemoveAccess
	s.EnableOnBehalfRequestorsToUpdateAccess = decoded.EnableOnBehalfRequestorsToUpdateAccess
	s.EnableTargetsToSelfAddAccess = decoded.EnableTargetsToSelfAddAccess
	s.EnableTargetsToSelfRemoveAccess = decoded.EnableTargetsToSelfRemoveAccess
	s.EnableTargetsToSelfUpdateAccess = decoded.EnableTargetsToSelfUpdateAccess
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentRequestorSettings into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["onBehalfRequestors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OnBehalfRequestors into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OnBehalfRequestors' for 'AccessPackageAssignmentRequestorSettings': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OnBehalfRequestors = &output
	}

	return nil
}
