package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewScheduleDefinition{}

type AccessReviewScheduleDefinition struct {
	// Defines the list of additional users or group members to be notified of the access review progress.
	AdditionalNotificationRecipients *[]AccessReviewNotificationRecipientItem `json:"additionalNotificationRecipients,omitempty"`

	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers are
	// notified to take action if no users are found from the list of reviewers specified. This could occur when either the
	// group owner is specified as the reviewer but the group owner doesn't exist, or manager is specified as reviewer but a
	// user's manager doesn't exist. Supports $select. Note: This property has been replaced by fallbackReviewers. However,
	// specifying either backupReviewers or fallbackReviewers automatically populates the same values to the other property.
	BackupReviewers *[]AccessReviewReviewerScope `json:"backupReviewers,omitempty"`

	// User who created this review. Read-only.
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`

	// Timestamp when the access review series was created. Supports $select. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description provided by review creators to provide more context of the review to admins. Supports $select.
	DescriptionForAdmins nullable.Type[string] `json:"descriptionForAdmins,omitempty"`

	// Description provided by review creators to provide more context of the review to reviewers. Reviewers see this
	// description in the email sent to them requesting their review. Email notifications support up to 256 characters.
	// Supports $select.
	DescriptionForReviewers nullable.Type[string] `json:"descriptionForReviewers,omitempty"`

	// Name of the access review series. Supports $select and $orderby. Required on create.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be
	// notified to take action if no users are found from the list of reviewers specified. This could occur when either the
	// group owner is specified as the reviewer but the group owner doesn't exist, or manager is specified as reviewer but a
	// user's manager doesn't exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. NOTE: The
	// value of this property will be ignored if fallback reviewers are assigned through the stageSettings property.
	FallbackReviewers *[]AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`

	// This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines
	// which Microsoft 365 groups are reviewed. Each group becomes a unique accessReviewInstance of the access review
	// series. For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring
	// instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
	InstanceEnumerationScope AccessReviewScope `json:"instanceEnumerationScope"`

	// Set of access reviews instances for this access review series. Access reviews that don't recur will only have one
	// instance; otherwise, there's an instance for each recurrence.
	Instances *[]AccessReviewInstance `json:"instances,omitempty"`

	// Timestamp when the access review series was last modified. Supports $select. Read-only.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// This collection of access review scopes is used to define who are the reviewers. The reviewers property is only
	// updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of
	// options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
	// NOTE: The value of this property will be ignored if reviewers are assigned through the stageSettings property.
	Reviewers *[]AccessReviewReviewerScope `json:"reviewers,omitempty"`

	// Defines the entities whose access is reviewed. For supported scopes, see accessReviewScope. Required on create.
	// Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope
	// of your access review definition using the Microsoft Graph API.
	Scope AccessReviewScope `json:"scope"`

	// The settings for an access review series, see type definition below. Supports $select. Required on create.
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`

	// Required only for a multi-stage access review to define the stages and their settings. You can break down each review
	// instance into up to three sequential stages, where each stage can have a different set of reviewers, fallback
	// reviewers, and settings. Stages are created sequentially based on the dependsOn property. Optional. When this
	// property is defined, its settings are used instead of the corresponding settings in the
	// accessReviewScheduleDefinition object and its settings, reviewers, and fallbackReviewers properties.
	StageSettings *[]AccessReviewStageSettings `json:"stageSettings,omitempty"`

	// This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted,
	// Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter
	// (eq only). Read-only.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s AccessReviewScheduleDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewScheduleDefinition{}

func (s AccessReviewScheduleDefinition) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewScheduleDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewScheduleDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewScheduleDefinition: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewScheduleDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewScheduleDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewScheduleDefinition{}

func (s *AccessReviewScheduleDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdditionalNotificationRecipients *[]AccessReviewNotificationRecipientItem `json:"additionalNotificationRecipients,omitempty"`
		BackupReviewers                  *[]AccessReviewReviewerScope             `json:"backupReviewers,omitempty"`
		CreatedDateTime                  nullable.Type[string]                    `json:"createdDateTime,omitempty"`
		DescriptionForAdmins             nullable.Type[string]                    `json:"descriptionForAdmins,omitempty"`
		DescriptionForReviewers          nullable.Type[string]                    `json:"descriptionForReviewers,omitempty"`
		DisplayName                      nullable.Type[string]                    `json:"displayName,omitempty"`
		FallbackReviewers                *[]AccessReviewReviewerScope             `json:"fallbackReviewers,omitempty"`
		Instances                        *[]AccessReviewInstance                  `json:"instances,omitempty"`
		LastModifiedDateTime             nullable.Type[string]                    `json:"lastModifiedDateTime,omitempty"`
		Reviewers                        *[]AccessReviewReviewerScope             `json:"reviewers,omitempty"`
		Settings                         *AccessReviewScheduleSettings            `json:"settings,omitempty"`
		StageSettings                    *[]AccessReviewStageSettings             `json:"stageSettings,omitempty"`
		Status                           nullable.Type[string]                    `json:"status,omitempty"`
		Id                               *string                                  `json:"id,omitempty"`
		ODataId                          *string                                  `json:"@odata.id,omitempty"`
		ODataType                        *string                                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdditionalNotificationRecipients = decoded.AdditionalNotificationRecipients
	s.BackupReviewers = decoded.BackupReviewers
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DescriptionForAdmins = decoded.DescriptionForAdmins
	s.DescriptionForReviewers = decoded.DescriptionForReviewers
	s.DisplayName = decoded.DisplayName
	s.FallbackReviewers = decoded.FallbackReviewers
	s.Instances = decoded.Instances
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Reviewers = decoded.Reviewers
	s.Settings = decoded.Settings
	s.StageSettings = decoded.StageSettings
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewScheduleDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'AccessReviewScheduleDefinition': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["instanceEnumerationScope"]; ok {
		impl, err := UnmarshalAccessReviewScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InstanceEnumerationScope' for 'AccessReviewScheduleDefinition': %+v", err)
		}
		s.InstanceEnumerationScope = impl
	}

	if v, ok := temp["scope"]; ok {
		impl, err := UnmarshalAccessReviewScopeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Scope' for 'AccessReviewScheduleDefinition': %+v", err)
		}
		s.Scope = impl
	}

	return nil
}
