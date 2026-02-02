package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessPackageAssignmentPolicy{}

type AccessPackageAssignmentPolicy struct {
	// The access package with this policy. Read-only. Nullable. Supports $expand.
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`

	AccessPackageCatalog *AccessPackageCatalog `json:"accessPackageCatalog,omitempty"`

	// Identifier of the access package.
	AccessPackageId nullable.Type[string] `json:"accessPackageId,omitempty"`

	// Represents the settings for email notifications for requests to an access package.
	AccessPackageNotificationSettings *AccessPackageNotificationSettings `json:"accessPackageNotificationSettings,omitempty"`

	// Who must review, and how often, the assignments to the access package from this policy. This property is null if
	// reviews aren't required.
	AccessReviewSettings *AssignmentReviewSettings `json:"accessReviewSettings,omitempty"`

	// Indicates whether a user can extend the access package assignment duration after approval.
	CanExtend nullable.Type[bool] `json:"canExtend,omitempty"`

	CreatedBy nullable.Type[string] `json:"createdBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
	CustomExtensionHandlers *[]CustomExtensionHandler `json:"customExtensionHandlers,omitempty"`

	// The collection of stages when to execute one or more custom access package workflow extensions. Supports $expand.
	CustomExtensionStageSettings *[]CustomExtensionStageSetting `json:"customExtensionStageSettings,omitempty"`

	// The description of the policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the policy. Supports $filter (eq).
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The number of days in which assignments from this policy last until they're expired.
	DurationInDays nullable.Type[int64] `json:"durationInDays,omitempty"`

	// The expiration date for assignments created in this policy. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	ModifiedBy nullable.Type[string] `json:"modifiedBy,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

	// Questions that are posed to the requestor.
	Questions *[]AccessPackageQuestion `json:"questions,omitempty"`

	// Who must approve requests for access package in this policy.
	RequestApprovalSettings *ApprovalSettings `json:"requestApprovalSettings,omitempty"`

	// Who can request this access package from this policy.
	RequestorSettings *RequestorSettings `json:"requestorSettings,omitempty"`

	// Settings for verifiable credentials set up through the Microsoft Entra Verified I D service. These settings represent
	// the verifiable credentials that a requestor of an access package in this policy can present to be assigned the access
	// package.
	VerifiableCredentialSettings *VerifiableCredentialSettings `json:"verifiableCredentialSettings,omitempty"`

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
		AccessPackage                     *AccessPackage                     `json:"accessPackage,omitempty"`
		AccessPackageCatalog              *AccessPackageCatalog              `json:"accessPackageCatalog,omitempty"`
		AccessPackageId                   nullable.Type[string]              `json:"accessPackageId,omitempty"`
		AccessPackageNotificationSettings *AccessPackageNotificationSettings `json:"accessPackageNotificationSettings,omitempty"`
		AccessReviewSettings              *AssignmentReviewSettings          `json:"accessReviewSettings,omitempty"`
		CanExtend                         nullable.Type[bool]                `json:"canExtend,omitempty"`
		CreatedBy                         nullable.Type[string]              `json:"createdBy,omitempty"`
		CreatedDateTime                   nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CustomExtensionHandlers           *[]CustomExtensionHandler          `json:"customExtensionHandlers,omitempty"`
		CustomExtensionStageSettings      *[]CustomExtensionStageSetting     `json:"customExtensionStageSettings,omitempty"`
		Description                       nullable.Type[string]              `json:"description,omitempty"`
		DisplayName                       nullable.Type[string]              `json:"displayName,omitempty"`
		DurationInDays                    nullable.Type[int64]               `json:"durationInDays,omitempty"`
		ExpirationDateTime                nullable.Type[string]              `json:"expirationDateTime,omitempty"`
		ModifiedBy                        nullable.Type[string]              `json:"modifiedBy,omitempty"`
		ModifiedDateTime                  nullable.Type[string]              `json:"modifiedDateTime,omitempty"`
		RequestApprovalSettings           *ApprovalSettings                  `json:"requestApprovalSettings,omitempty"`
		RequestorSettings                 *RequestorSettings                 `json:"requestorSettings,omitempty"`
		VerifiableCredentialSettings      *VerifiableCredentialSettings      `json:"verifiableCredentialSettings,omitempty"`
		Id                                *string                            `json:"id,omitempty"`
		ODataId                           *string                            `json:"@odata.id,omitempty"`
		ODataType                         *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessPackage = decoded.AccessPackage
	s.AccessPackageCatalog = decoded.AccessPackageCatalog
	s.AccessPackageId = decoded.AccessPackageId
	s.AccessPackageNotificationSettings = decoded.AccessPackageNotificationSettings
	s.AccessReviewSettings = decoded.AccessReviewSettings
	s.CanExtend = decoded.CanExtend
	s.CreatedBy = decoded.CreatedBy
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CustomExtensionHandlers = decoded.CustomExtensionHandlers
	s.CustomExtensionStageSettings = decoded.CustomExtensionStageSettings
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.DurationInDays = decoded.DurationInDays
	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.ModifiedBy = decoded.ModifiedBy
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.RequestApprovalSettings = decoded.RequestApprovalSettings
	s.RequestorSettings = decoded.RequestorSettings
	s.VerifiableCredentialSettings = decoded.VerifiableCredentialSettings
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

	return nil
}
