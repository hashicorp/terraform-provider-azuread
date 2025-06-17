package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheck struct {
	// More details about the health check or the recommended action. Read-only.
	AdditionalDetail nullable.Type[string] `json:"additionalDetail,omitempty"`

	// More details about the health check or the recommended action. Read-only. The additionalDetails property is
	// deprecated and stopped returning data on January 31, 2024. Goind forward, use the additionalDetail property.
	AdditionalDetails nullable.Type[string] `json:"additionalDetails,omitempty"`

	// The unique identifier of the health check item-related activities. This identifier can be useful in troubleshooting.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// The display name for this health check item.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The end time of the health check item. Read-only.
	EndDateTime *string `json:"endDateTime,omitempty"`

	// The type of error that occurred during this health check. For the list of possible values, see
	// cloudPcOnPremisesConnectionHealthCheckErrorType.
	ErrorType *CloudPCOnPremisesConnectionHealthCheckErrorType `json:"errorType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The recommended action to fix the corresponding error.
	RecommendedAction nullable.Type[string] `json:"recommendedAction,omitempty"`

	// The start time of the health check item. Read-only.
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
	delete(decoded, "additionalDetails")
	delete(decoded, "endDateTime")
	delete(decoded, "startDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCOnPremisesConnectionHealthCheck: %+v", err)
	}

	return encoded, nil
}
