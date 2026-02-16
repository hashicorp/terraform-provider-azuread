package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedAppLogCollectionRequest{}

type ManagedAppLogCollectionRequest struct {
	// DateTime of when the log upload request was completed. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The unique identifier of the app instance for which diagnostic logs were collected. Read-only.
	ManagedAppRegistrationId nullable.Type[string] `json:"managedAppRegistrationId,omitempty"`

	// The user principal name associated with the request for the managed application log collection. Read-only.
	RequestedBy nullable.Type[string] `json:"requestedBy,omitempty"`

	// The user principal name associated with the request for the managed application log collection. Read-only.
	RequestedByUserPrincipalName nullable.Type[string] `json:"requestedByUserPrincipalName,omitempty"`

	// DateTime of when the log upload request was received. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this:
	// '2014-01-01T00:00:00Z'. Returned by default. Read-only.
	RequestedDateTime *string `json:"requestedDateTime,omitempty"`

	// Indicates the status for the app log collection request - pending, completed or failed. Default is pending.
	Status nullable.Type[string] `json:"status,omitempty"`

	// The collection of log upload results as reported by each component on the device. Such components can be the
	// application itself, the Mobile Application Management (MAM) SDK, and other on-device components that are requested to
	// upload diagnostic logs. Read-only.
	UploadedLogs *[]ManagedAppLogUpload `json:"uploadedLogs,omitempty"`

	// Represents the current consent status of the associated `managedAppLogCollectionRequest`.
	UserLogUploadConsent *ManagedAppLogUploadConsent `json:"userLogUploadConsent,omitempty"`

	// Version of the entity.
	Version nullable.Type[string] `json:"version,omitempty"`

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

func (s ManagedAppLogCollectionRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedAppLogCollectionRequest{}

func (s ManagedAppLogCollectionRequest) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAppLogCollectionRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAppLogCollectionRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppLogCollectionRequest: %+v", err)
	}

	delete(decoded, "completedDateTime")
	delete(decoded, "managedAppRegistrationId")
	delete(decoded, "requestedBy")
	delete(decoded, "requestedByUserPrincipalName")
	delete(decoded, "requestedDateTime")
	delete(decoded, "uploadedLogs")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedAppLogCollectionRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAppLogCollectionRequest: %+v", err)
	}

	return encoded, nil
}
