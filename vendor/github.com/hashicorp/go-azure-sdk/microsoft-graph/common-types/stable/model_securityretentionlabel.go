package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityRetentionLabel{}

type SecurityRetentionLabel struct {
	// Specifies the action to take on the labeled document after the period specified by the retentionDuration property
	// expires. The possible values are: none, delete, startDispositionReview, unknownFutureValue.
	ActionAfterRetentionPeriod *SecurityActionAfterRetentionPeriod `json:"actionAfterRetentionPeriod,omitempty"`

	// Specifies how the behavior of a document with this label should be during the retention period. The possible values
	// are: doNotRetain, retain, retainAsRecord, retainAsRegulatoryRecord, unknownFutureValue.
	BehaviorDuringRetentionPeriod *SecurityBehaviorDuringRetentionPeriod `json:"behaviorDuringRetentionPeriod,omitempty"`

	// Represents the user who created the retentionLabel.
	CreatedBy IdentitySet `json:"createdBy"`

	// Represents the date and time in which the retentionLabel is created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Specifies the locked or unlocked state of a record label when it is created.The possible values are: startLocked,
	// startUnlocked, unknownFutureValue.
	DefaultRecordBehavior *SecurityDefaultRecordBehavior `json:"defaultRecordBehavior,omitempty"`

	// Provides label information for the admin. Optional.
	DescriptionForAdmins nullable.Type[string] `json:"descriptionForAdmins,omitempty"`

	// Provides the label information for the user. Optional.
	DescriptionForUsers nullable.Type[string] `json:"descriptionForUsers,omitempty"`

	// Represents out-of-the-box values that provide more options to improve the manageability and organization of the
	// content you need to label.
	Descriptors *SecurityFilePlanDescriptor `json:"descriptors,omitempty"`

	// Unique string that defines a label name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// When action at the end of retention is chosen as 'dispositionReview', dispositionReviewStages specifies a sequential
	// set of stages with at least one reviewer in each stage.
	DispositionReviewStages *[]SecurityDispositionReviewStage `json:"dispositionReviewStages,omitempty"`

	// Specifies whether the label is currently being used.
	IsInUse nullable.Type[bool] `json:"isInUse,omitempty"`

	// Specifies the replacement label to be applied automatically after the retention period of the current label ends.
	LabelToBeApplied nullable.Type[string] `json:"labelToBeApplied,omitempty"`

	// The user who last modified the retentionLabel.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The latest date time when the retentionLabel was modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Specifies the number of days to retain the content.
	RetentionDuration SecurityRetentionDuration `json:"retentionDuration"`

	// Represents the type associated with a retention event.
	RetentionEventType *SecurityRetentionEventType `json:"retentionEventType,omitempty"`

	// Specifies whether the retention duration is calculated from the content creation date, labeled date, or last
	// modification date. The possible values are: dateLabeled, dateCreated, dateModified, dateOfEvent, unknownFutureValue.
	RetentionTrigger *SecurityRetentionTrigger `json:"retentionTrigger,omitempty"`

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

func (s SecurityRetentionLabel) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityRetentionLabel{}

func (s SecurityRetentionLabel) MarshalJSON() ([]byte, error) {
	type wrapper SecurityRetentionLabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityRetentionLabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityRetentionLabel: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.retentionLabel"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityRetentionLabel: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityRetentionLabel{}

func (s *SecurityRetentionLabel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActionAfterRetentionPeriod    *SecurityActionAfterRetentionPeriod    `json:"actionAfterRetentionPeriod,omitempty"`
		BehaviorDuringRetentionPeriod *SecurityBehaviorDuringRetentionPeriod `json:"behaviorDuringRetentionPeriod,omitempty"`
		CreatedDateTime               nullable.Type[string]                  `json:"createdDateTime,omitempty"`
		DefaultRecordBehavior         *SecurityDefaultRecordBehavior         `json:"defaultRecordBehavior,omitempty"`
		DescriptionForAdmins          nullable.Type[string]                  `json:"descriptionForAdmins,omitempty"`
		DescriptionForUsers           nullable.Type[string]                  `json:"descriptionForUsers,omitempty"`
		Descriptors                   *SecurityFilePlanDescriptor            `json:"descriptors,omitempty"`
		DisplayName                   nullable.Type[string]                  `json:"displayName,omitempty"`
		DispositionReviewStages       *[]SecurityDispositionReviewStage      `json:"dispositionReviewStages,omitempty"`
		IsInUse                       nullable.Type[bool]                    `json:"isInUse,omitempty"`
		LabelToBeApplied              nullable.Type[string]                  `json:"labelToBeApplied,omitempty"`
		LastModifiedDateTime          nullable.Type[string]                  `json:"lastModifiedDateTime,omitempty"`
		RetentionEventType            *SecurityRetentionEventType            `json:"retentionEventType,omitempty"`
		RetentionTrigger              *SecurityRetentionTrigger              `json:"retentionTrigger,omitempty"`
		Id                            *string                                `json:"id,omitempty"`
		ODataId                       *string                                `json:"@odata.id,omitempty"`
		ODataType                     *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActionAfterRetentionPeriod = decoded.ActionAfterRetentionPeriod
	s.BehaviorDuringRetentionPeriod = decoded.BehaviorDuringRetentionPeriod
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefaultRecordBehavior = decoded.DefaultRecordBehavior
	s.DescriptionForAdmins = decoded.DescriptionForAdmins
	s.DescriptionForUsers = decoded.DescriptionForUsers
	s.Descriptors = decoded.Descriptors
	s.DisplayName = decoded.DisplayName
	s.DispositionReviewStages = decoded.DispositionReviewStages
	s.IsInUse = decoded.IsInUse
	s.LabelToBeApplied = decoded.LabelToBeApplied
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.RetentionEventType = decoded.RetentionEventType
	s.RetentionTrigger = decoded.RetentionTrigger
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityRetentionLabel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SecurityRetentionLabel': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SecurityRetentionLabel': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["retentionDuration"]; ok {
		impl, err := UnmarshalSecurityRetentionDurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RetentionDuration' for 'SecurityRetentionLabel': %+v", err)
		}
		s.RetentionDuration = impl
	}

	return nil
}
