package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheck struct {
	// Additional details about the health check or the recommended action. For exmaple, the string value can be
	// download.microsoft.com:443;software-download.microsoft.com:443; Read-only.
	AdditionalDetail nullable.Type[string] `json:"additionalDetail,omitempty"`

	// The unique identifier of the health check item-related activities. This identifier can be useful in troubleshooting.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// The display name for this health check item.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The value cannot be modified and is automatically populated when the health check ends. The Timestamp type represents
	// date and time information using ISO 8601 format and is in Coordinated Universal Time (UTC). For example, midnight UTC
	// on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// The type of error that occurred during this health check. Possible values are:
	// endpointConnectivityCheckCloudPcUrlNotAllowListed, endpointConnectivityCheckWVDUrlNotAllowListed, etc. (The all
	// possible values can refer to cloudPcOnPremisesConnectionHealthCheckErrorType) Read-Only.
	ErrorType *CloudPCOnPremisesConnectionHealthCheckErrorType `json:"errorType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The recommended action to fix the corresponding error. For example, The Active Directory domain join check failed
	// because the password of the domain join user has expired. Read-Only.
	RecommendedAction nullable.Type[string] `json:"recommendedAction,omitempty"`

	// The value cannot be modified and is automatically populated when the health check starts. The Timestamp type
	// represents date and time information using ISO 8601 format and is in Coordinated Universal Time (UTC). For example,
	// midnight UTC on Jan 1, 2024 would look like this: '2024-01-01T00:00:00Z'. Returned by default. Read-only.
	StartDateTime *string `json:"startDateTime,omitempty"`

	Status *CloudPCOnPremisesConnectionStatus `json:"status,omitempty"`
}

var _ json.Marshaler = CloudPCOnPremisesConnectionHealthCheck{}

func (s CloudPCOnPremisesConnectionHealthCheck) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCOnPremisesConnectionHealthCheck
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCOnPremisesConnectionHealthCheck: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCOnPremisesConnectionHealthCheck: %+v", err)
	}

	delete(decoded, "additionalDetail")
	delete(decoded, "endDateTime")
	delete(decoded, "startDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCOnPremisesConnectionHealthCheck: %+v", err)
	}

	return encoded, nil
}
