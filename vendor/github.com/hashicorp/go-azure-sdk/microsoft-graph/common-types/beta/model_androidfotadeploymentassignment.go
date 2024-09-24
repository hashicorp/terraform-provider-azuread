package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidFotaDeploymentAssignment struct {
	// The Azure Active Directory (Azure AD) we are deploying firmware updates to (e.g.:
	// d93c8f48-bd42-4514-ba40-bc6b84780930). NOTE: Use this property moving forward because the existing property, target,
	// is deprecated.
	AssignmentTarget DeviceAndAppManagementAssignmentTarget `json:"assignmentTarget"`

	// The display name of the Azure AD security group used for the assignment.
	DisplayName *string `json:"displayName,omitempty"`

	// A unique identifier assigned to each Android FOTA Assignment entity
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The AAD Group we are deploying firmware updates to
	Target *AndroidFotaDeploymentAssignmentTarget `json:"target,omitempty"`
}

var _ json.Unmarshaler = &AndroidFotaDeploymentAssignment{}

func (s *AndroidFotaDeploymentAssignment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName *string                                `json:"displayName,omitempty"`
		Id          nullable.Type[string]                  `json:"id,omitempty"`
		ODataId     *string                                `json:"@odata.id,omitempty"`
		ODataType   *string                                `json:"@odata.type,omitempty"`
		Target      *AndroidFotaDeploymentAssignmentTarget `json:"target,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Target = decoded.Target

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AndroidFotaDeploymentAssignment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignmentTarget"]; ok {
		impl, err := UnmarshalDeviceAndAppManagementAssignmentTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignmentTarget' for 'AndroidFotaDeploymentAssignment': %+v", err)
		}
		s.AssignmentTarget = impl
	}

	return nil
}
