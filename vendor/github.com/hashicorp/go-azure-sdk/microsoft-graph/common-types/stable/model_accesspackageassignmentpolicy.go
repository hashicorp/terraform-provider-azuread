package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignmentPolicy{}

type AccessPackageAssignmentPolicy struct {
	// Access package containing this policy. Read-only. Supports $expand.
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`

	// Principals that can be assigned the access package through this policy. The possible values are: notSpecified,
	// specificDirectoryUsers, specificConnectedOrganizationUsers, specificDirectoryServicePrincipals, allMemberUsers,
	// allDirectoryUsers, allDirectoryServicePrincipals, allConfiguredConnectedOrganizationUsers, allExternalUsers,
	// unknownFutureValue.
	AllowedTargetScope *AllowedTargetScope `json:"allowedTargetScope,omitempty"`

	// This property is only present for an auto assignment policy; if absent, this is a request-based policy.
	AutomaticRequestSettings *AccessPackageAutomaticRequestSettings `json:"automaticRequestSettings,omitempty"`

	// Catalog of the access package containing this policy. Read-only.
	Catalog *AccessPackageCatalog `json:"catalog,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
	CustomExtensionStageSettings *[]CustomExtensionStageSetting `json:"customExtensionStageSettings,omitempty"`

	// The description of the policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the policy.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The expiration date for assignments created in this policy.
	Expiration *ExpirationPattern `json:"expiration,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// Questions that are posed to the requestor.
	Questions *[]AccessPackageQuestion `json:"questions,omitempty"`

	// Specifies the settings for approval of requests for an access package assignment through this policy. For example, if
	// approval is required for new requests.
	RequestApprovalSettings *AccessPackageAssignmentApprovalSettings `json:"requestApprovalSettings,omitempty"`

	// Provides additional settings to select who can create a request for an access package assignment through this policy,
	// and what they can include in their request.
	RequestorSettings *AccessPackageAssignmentRequestorSettings `json:"requestorSettings,omitempty"`

	// Settings for access reviews of assignments through this policy.
	ReviewSettings *AccessPackageAssignmentReviewSettings `json:"reviewSettings,omitempty"`

	// The principals that can be assigned access from an access package through this policy.
	SpecificAllowedTargets *[]SubjectSet `json:"specificAllowedTargets,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageAssignmentPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAssignmentPolicy{}

func (s AccessPackageAssignmentPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAssignmentPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAssignmentPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAssignmentPolicy: %+v", err)
	}

	delete(decoded, "accessPackage")
	delete(decoded, "catalog")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAssignmentPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAssignmentPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageAssignmentPolicy{}

func (s *AccessPackageAssignmentPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessPackage                *AccessPackage                            `json:"accessPackage,omitempty"`
		AllowedTargetScope           *AllowedTargetScope                       `json:"allowedTargetScope,omitempty"`
		AutomaticRequestSettings     *AccessPackageAutomaticRequestSettings    `json:"automaticRequestSettings,omitempty"`
		Catalog                      *AccessPackageCatalog                     `json:"catalog,omitempty"`
		CreatedDateTime              nullable.Type[string]                     `json:"createdDateTime,omitempty"`
		CustomExtensionStageSettings *[]CustomExtensionStageSetting            `json:"customExtensionStageSettings,omitempty"`
		Description                  nullable.Type[string]                     `json:"description,omitempty"`
		DisplayName                  nullable.Type[string]                     `json:"displayName,omitempty"`
		Expiration                   *ExpirationPattern                        `json:"expiration,omitempty"`
		ModifiedDateTime             nullable.Type[string]                     `json:"modifiedDateTime,omitempty"`
		RequestApprovalSettings      *AccessPackageAssignmentApprovalSettings  `json:"requestApprovalSettings,omitempty"`
		RequestorSettings            *AccessPackageAssignmentRequestorSettings `json:"requestorSettings,omitempty"`
		ReviewSettings               *AccessPackageAssignmentReviewSettings    `json:"reviewSettings,omitempty"`
		Id                           *string                                   `json:"id,omitempty"`
		ODataId                      *string                                   `json:"@odata.id,omitempty"`
		ODataType                    *string                                   `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackage = decoded.AccessPackage
	s.AllowedTargetScope = decoded.AllowedTargetScope
	s.AutomaticRequestSettings = decoded.AutomaticRequestSettings
	s.Catalog = decoded.Catalog
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomExtensionStageSettings = decoded.CustomExtensionStageSettings
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Expiration = decoded.Expiration
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.RequestApprovalSettings = decoded.RequestApprovalSettings
	s.RequestorSettings = decoded.RequestorSettings
	s.ReviewSettings = decoded.ReviewSettings
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAssignmentPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["questions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Questions into list []json.RawMessage: %+v", err)
		}

		output := make([]AccessPackageQuestion, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAccessPackageQuestionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Questions' for 'AccessPackageAssignmentPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Questions = &output
	}

	if v, ok := temp["specificAllowedTargets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SpecificAllowedTargets into list []json.RawMessage: %+v", err)
		}

		output := make([]SubjectSet, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSubjectSetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SpecificAllowedTargets' for 'AccessPackageAssignmentPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SpecificAllowedTargets = &output
	}

	return nil
}
