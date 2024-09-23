package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DirectoryAudit{}

type DirectoryAudit struct {
	// Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ge, le) and $orderby.
	ActivityDateTime *string `json:"activityDateTime,omitempty"`

	// Indicates the activity name or the operation name (for example 'Create User', 'Add member to group'). For a list of
	// activities logged, refer to Microsoft Entra audit log categories and activities. Supports $filter (eq, startswith).
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`

	// Indicates more details on the activity.
	AdditionalDetails *[]KeyValue `json:"additionalDetails,omitempty"`

	// Indicates which resource category that's targeted by the activity. For example: UserManagement, GroupManagement,
	// ApplicationManagement, RoleManagement. For a list of categories for activities logged, refer to Microsoft Entra audit
	// log categories and activities.
	Category *string `json:"category,omitempty"`

	// Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs
	// across services. Supports $filter (eq).
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	InitiatedBy *AuditActivityInitiator `json:"initiatedBy,omitempty"`

	// Indicates information on which service initiated the activity (For example: Self-service Password Management, Core
	// Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management. Supports $filter (eq).
	LoggedByService nullable.Type[string] `json:"loggedByService,omitempty"`

	// Indicates the type of operation that was performed. The possible values include but aren't limited to the following:
	// Add, Assign, Update, Unassign, and Delete.
	OperationType nullable.Type[string] `json:"operationType,omitempty"`

	// Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.
	Result *OperationResult `json:"result,omitempty"`

	// Indicates the reason for failure if the result is failure or timeout.
	ResultReason nullable.Type[string] `json:"resultReason,omitempty"`

	// Information about the resource that changed due to the activity. Supports $filter (eq) for id and displayName; and
	// $filter (startswith) for displayName.
	TargetResources *[]TargetResource `json:"targetResources,omitempty"`

	// Type of user agent used by a user in the activity.
	UserAgent nullable.Type[string] `json:"userAgent,omitempty"`

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

func (s DirectoryAudit) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DirectoryAudit{}

func (s DirectoryAudit) MarshalJSON() ([]byte, error) {
	type wrapper DirectoryAudit
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DirectoryAudit: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DirectoryAudit: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.directoryAudit"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DirectoryAudit: %+v", err)
	}

	return encoded, nil
}
