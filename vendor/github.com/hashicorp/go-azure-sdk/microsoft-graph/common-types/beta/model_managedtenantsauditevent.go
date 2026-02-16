package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsAuditEvent{}

type ManagedTenantsAuditEvent struct {
	// A string that uniquely represents the operation that occurred. Required. Read-only.
	Activity nullable.Type[string] `json:"activity,omitempty"`

	// The time when the activity occurred. Required. Read-only.
	ActivityDateTime nullable.Type[string] `json:"activityDateTime,omitempty"`

	// The identifier of the activity request that made the audit event. Required. Read-only.
	ActivityId nullable.Type[string] `json:"activityId,omitempty"`

	// A category that represents a logical grouping of activities. Required. Read-only.
	Category nullable.Type[string] `json:"category,omitempty"`

	// The HTTP verb that was used when making the API request. Required. Read-only.
	HttpVerb nullable.Type[string] `json:"httpVerb,omitempty"`

	// The IP address of where the activity was initiated. This may be an IPv4 or IPv6 address. Required. Read-only.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The identifier of the app that was used to make the request. Required. Read-only.
	InitiatedByAppId nullable.Type[string] `json:"initiatedByAppId,omitempty"`

	// The UPN of the user who initiated the activity. Required. Read-only.
	InitiatedByUpn nullable.Type[string] `json:"initiatedByUpn,omitempty"`

	// The identifier of the user who initiated the activity. Required. Read-only.
	InitiatedByUserId nullable.Type[string] `json:"initiatedByUserId,omitempty"`

	// The raw HTTP request body. Some sensitive information may be removed.
	RequestBody nullable.Type[string] `json:"requestBody,omitempty"`

	// The raw HTTP request URL. Required. Read-only.
	RequestUrl nullable.Type[string] `json:"requestUrl,omitempty"`

	// The collection of Microsoft Entra tenant identifiers for the managed tenants that were affected by a change, and is
	// formatted as a list of comma-separated values. Required. Read-only.
	TenantIds nullable.Type[string] `json:"tenantIds,omitempty"`

	// The collection of tenant names that were affected by a change, and is formatted as a list of comma-separated values.
	// Required. Read-only.
	TenantNames nullable.Type[string] `json:"tenantNames,omitempty"`

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

func (s ManagedTenantsAuditEvent) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsAuditEvent{}

func (s ManagedTenantsAuditEvent) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsAuditEvent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsAuditEvent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsAuditEvent: %+v", err)
	}

	delete(decoded, "activity")
	delete(decoded, "activityDateTime")
	delete(decoded, "activityId")
	delete(decoded, "category")
	delete(decoded, "httpVerb")
	delete(decoded, "initiatedByAppId")
	delete(decoded, "initiatedByUpn")
	delete(decoded, "initiatedByUserId")
	delete(decoded, "ipAddress")
	delete(decoded, "requestUrl")
	delete(decoded, "tenantIds")
	delete(decoded, "tenantNames")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.auditEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsAuditEvent: %+v", err)
	}

	return encoded, nil
}
