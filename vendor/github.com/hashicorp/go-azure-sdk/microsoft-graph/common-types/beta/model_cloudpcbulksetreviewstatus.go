package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudPCBulkAction = CloudPCBulkSetReviewStatus{}

type CloudPCBulkSetReviewStatus struct {
	ReviewStatus *CloudPCReviewStatus `json:"reviewStatus,omitempty"`

	// Fields inherited from CloudPCBulkAction

	// Run summary of this bulk action.
	ActionSummary *CloudPCBulkActionSummary `json:"actionSummary,omitempty"`

	CloudPCIds *[]string `json:"cloudPcIds,omitempty"`

	// The date and time when the bulk action was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Name of the bulk action.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates the user principal name (UPN) of the user who initiated this bulk action. Read-only.
	InitiatedByUserPrincipalName nullable.Type[string] `json:"initiatedByUserPrincipalName,omitempty"`

	// Indicates whether the bulk action is scheduled according to the maintenance window. When true, the bulk action uses
	// the maintenance window to schedule the action; false means that the bulk action doesn't use the maintenance window.
	// The default value is false.
	ScheduledDuringMaintenanceWindow nullable.Type[bool] `json:"scheduledDuringMaintenanceWindow,omitempty"`

	// Indicates the status of bulk actions. Possible values are pending, succeeded, failed, unknownFutureValue. The default
	// value is pending. Read-only.
	Status *CloudPCBulkActionStatus `json:"status,omitempty"`

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

func (s CloudPCBulkSetReviewStatus) CloudPCBulkAction() BaseCloudPCBulkActionImpl {
	return BaseCloudPCBulkActionImpl{
		ActionSummary:                    s.ActionSummary,
		CloudPCIds:                       s.CloudPCIds,
		CreatedDateTime:                  s.CreatedDateTime,
		DisplayName:                      s.DisplayName,
		InitiatedByUserPrincipalName:     s.InitiatedByUserPrincipalName,
		ScheduledDuringMaintenanceWindow: s.ScheduledDuringMaintenanceWindow,
		Status:                           s.Status,
		Id:                               s.Id,
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
	}
}

func (s CloudPCBulkSetReviewStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCBulkSetReviewStatus{}

func (s CloudPCBulkSetReviewStatus) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCBulkSetReviewStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCBulkSetReviewStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCBulkSetReviewStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcBulkSetReviewStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCBulkSetReviewStatus: %+v", err)
	}

	return encoded, nil
}
