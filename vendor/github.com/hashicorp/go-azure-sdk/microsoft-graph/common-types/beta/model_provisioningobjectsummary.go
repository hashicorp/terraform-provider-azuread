package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProvisioningObjectSummary{}

type ProvisioningObjectSummary struct {
	// Indicates the activity name or the operation name (for example, Create user, Add member to group). For a list of
	// activities logged, refer to Microsoft Entra activity list. This is deprecated. Please use provisioningAction instead.
	// Supports $filter (eq, contains).
	Action nullable.Type[string] `json:"action,omitempty"`

	// Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, gt, lt) and orderby.
	ActivityDateTime *string `json:"activityDateTime,omitempty"`

	// Unique ID of this change in this cycle. Supports $filter (eq, contains).
	ChangeId nullable.Type[string] `json:"changeId,omitempty"`

	// Unique ID per job iteration. Supports $filter (eq, contains).
	CycleId nullable.Type[string] `json:"cycleId,omitempty"`

	// Indicates how long this provisioning action took to finish. Measured in milliseconds. Supports $filter (eq, gt, lt).
	DurationInMilliseconds nullable.Type[int64] `json:"durationInMilliseconds,omitempty"`

	// Details of who initiated this provisioning. Supports $filter (eq, contains).
	InitiatedBy *Initiator `json:"initiatedBy,omitempty"`

	// The unique ID for the whole provisioning job. Supports $filter (eq, contains).
	JobId nullable.Type[string] `json:"jobId,omitempty"`

	// Details of each property that was modified in this provisioning action on this object.
	ModifiedProperties *[]ModifiedProperty `json:"modifiedProperties,omitempty"`

	// Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete,
	// disable, other and unknownFutureValue. For a list of activities logged, refer to Microsoft Entra activity list.
	// Supports $filter (eq, contains).
	ProvisioningAction *ProvisioningAction `json:"provisioningAction,omitempty"`

	// Details of provisioning status. Supports $filter (eq, contains) for status.
	ProvisioningStatusInfo *ProvisioningStatusInfo `json:"provisioningStatusInfo,omitempty"`

	// Details of each step in provisioning.
	ProvisioningSteps *[]ProvisioningStep `json:"provisioningSteps,omitempty"`

	// Represents the service principal used for provisioning. Supports $filter (eq) for id and name.
	ServicePrincipal *ProvisioningServicePrincipal `json:"servicePrincipal,omitempty"`

	// Details of source object being provisioned. Supports $filter (eq, contains) for identityType, id, and displayName.
	SourceIdentity *ProvisionedIdentity `json:"sourceIdentity,omitempty"`

	// Details of source system of the object being provisioned. Supports $filter (eq, contains) for displayName.
	SourceSystem *ProvisioningSystem `json:"sourceSystem,omitempty"`

	// Details of provisioning status. This is deprecated. Please use provisioningStatusInfo instead. Supports $filter (eq,
	// contains) for status.
	StatusInfo StatusBase `json:"statusInfo"`

	// Details of target object being provisioned. Supports $filter (eq, contains) for identityType, id, and displayName.
	TargetIdentity *ProvisionedIdentity `json:"targetIdentity,omitempty"`

	// Details of target system of the object being provisioned. Supports $filter (eq, contains) for displayName.
	TargetSystem *ProvisioningSystem `json:"targetSystem,omitempty"`

	// Unique Microsoft Entra tenant ID. Supports $filter (eq, contains).
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

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

func (s ProvisioningObjectSummary) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProvisioningObjectSummary{}

func (s ProvisioningObjectSummary) MarshalJSON() ([]byte, error) {
	type wrapper ProvisioningObjectSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProvisioningObjectSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProvisioningObjectSummary: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.provisioningObjectSummary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProvisioningObjectSummary: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ProvisioningObjectSummary{}

func (s *ProvisioningObjectSummary) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action                 nullable.Type[string]         `json:"action,omitempty"`
		ActivityDateTime       *string                       `json:"activityDateTime,omitempty"`
		ChangeId               nullable.Type[string]         `json:"changeId,omitempty"`
		CycleId                nullable.Type[string]         `json:"cycleId,omitempty"`
		DurationInMilliseconds nullable.Type[int64]          `json:"durationInMilliseconds,omitempty"`
		InitiatedBy            *Initiator                    `json:"initiatedBy,omitempty"`
		JobId                  nullable.Type[string]         `json:"jobId,omitempty"`
		ModifiedProperties     *[]ModifiedProperty           `json:"modifiedProperties,omitempty"`
		ProvisioningAction     *ProvisioningAction           `json:"provisioningAction,omitempty"`
		ProvisioningStatusInfo *ProvisioningStatusInfo       `json:"provisioningStatusInfo,omitempty"`
		ProvisioningSteps      *[]ProvisioningStep           `json:"provisioningSteps,omitempty"`
		ServicePrincipal       *ProvisioningServicePrincipal `json:"servicePrincipal,omitempty"`
		SourceIdentity         *ProvisionedIdentity          `json:"sourceIdentity,omitempty"`
		SourceSystem           *ProvisioningSystem           `json:"sourceSystem,omitempty"`
		TargetIdentity         *ProvisionedIdentity          `json:"targetIdentity,omitempty"`
		TargetSystem           *ProvisioningSystem           `json:"targetSystem,omitempty"`
		TenantId               nullable.Type[string]         `json:"tenantId,omitempty"`
		Id                     *string                       `json:"id,omitempty"`
		ODataId                *string                       `json:"@odata.id,omitempty"`
		ODataType              *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.ActivityDateTime = decoded.ActivityDateTime
	s.ChangeId = decoded.ChangeId
	s.CycleId = decoded.CycleId
	s.DurationInMilliseconds = decoded.DurationInMilliseconds
	s.InitiatedBy = decoded.InitiatedBy
	s.JobId = decoded.JobId
	s.ModifiedProperties = decoded.ModifiedProperties
	s.ProvisioningAction = decoded.ProvisioningAction
	s.ProvisioningStatusInfo = decoded.ProvisioningStatusInfo
	s.ProvisioningSteps = decoded.ProvisioningSteps
	s.ServicePrincipal = decoded.ServicePrincipal
	s.SourceIdentity = decoded.SourceIdentity
	s.SourceSystem = decoded.SourceSystem
	s.TargetIdentity = decoded.TargetIdentity
	s.TargetSystem = decoded.TargetSystem
	s.TenantId = decoded.TenantId
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProvisioningObjectSummary into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["statusInfo"]; ok {
		impl, err := UnmarshalStatusBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StatusInfo' for 'ProvisioningObjectSummary': %+v", err)
		}
		s.StatusInfo = impl
	}

	return nil
}
