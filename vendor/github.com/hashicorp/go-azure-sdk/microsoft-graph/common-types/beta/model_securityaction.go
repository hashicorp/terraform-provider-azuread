package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityAction{}

type SecurityAction struct {
	// Reason for invoking this action.
	ActionReason nullable.Type[string] `json:"actionReason,omitempty"`

	// The Application ID of the calling application that submitted (POST) the action. The appId should be extracted from
	// the auth token and not entered manually by the calling application.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Azure tenant ID of the entity to determine which tenant the entity belongs to (multi-tenancy support). The
	// azureTenantId should be extracted from the auth token and not entered manually by the calling application.
	AzureTenantId nullable.Type[string] `json:"azureTenantId,omitempty"`

	// Unique client context string. Can have a maximum of 256 characters.
	ClientContext nullable.Type[string] `json:"clientContext,omitempty"`

	// Timestamp when the action was completed. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// Timestamp when the action is created. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Error info when the action fails.
	ErrorInfo *ResultInfo `json:"errorInfo,omitempty"`

	// Timestamp when this action was last updated. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastActionDateTime nullable.Type[string] `json:"lastActionDateTime,omitempty"`

	// Action name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Collection of parameters (key-value pairs) necessary to invoke the action, for example, URL or fileHash to block.).
	// Required.
	Parameters []KeyValuePair `json:"parameters"`

	// Collection of securityActionState to keep the history of an action.
	States *[]SecurityActionState `json:"states,omitempty"`

	// Status of the action. Possible values are: NotStarted, Running, Completed, Failed.
	Status *OperationStatus `json:"status,omitempty"`

	// The user principal name of the signed-in user that submitted (POST) the action. The user should be extracted from the
	// auth token and not entered manually by the calling application.
	User nullable.Type[string] `json:"user,omitempty"`

	// Complex Type containing details about the Security product/service vendor, provider, and sub-provider (for example,
	// vendor=Microsoft; provider=Windows Defender ATP; sub-provider=AppLocker).
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`

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

func (s SecurityAction) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAction{}

func (s SecurityAction) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.securityAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAction: %+v", err)
	}

	return encoded, nil
}
