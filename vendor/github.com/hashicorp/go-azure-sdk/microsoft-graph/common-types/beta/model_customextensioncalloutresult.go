package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationEventHandlerResult = CustomExtensionCalloutResult{}

type CustomExtensionCalloutResult struct {
	// When the API transaction was initiated, the date and time information uses ISO 8601 format and is always in UTC time.
	// Example: midnight on Jan 1, 2014, is reported as 2014-01-01T00:00:00Z.
	CalloutDateTime nullable.Type[string] `json:"calloutDateTime,omitempty"`

	// Identifier of the custom extension that was called.
	CustomExtensionId nullable.Type[string] `json:"customExtensionId,omitempty"`

	// Error code that was returned when the last API attempt failed.
	ErrorCode nullable.Type[int64] `json:"errorCode,omitempty"`

	// The HTTP status code that was returned by the target API endpoint after the last API attempt.
	HttpStatus nullable.Type[int64] `json:"httpStatus,omitempty"`

	// The number of API calls to the customer's API.
	NumberOfAttempts nullable.Type[int64] `json:"numberOfAttempts,omitempty"`

	// Fields inherited from AuthenticationEventHandlerResult

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CustomExtensionCalloutResult) AuthenticationEventHandlerResult() BaseAuthenticationEventHandlerResultImpl {
	return BaseAuthenticationEventHandlerResultImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomExtensionCalloutResult{}

func (s CustomExtensionCalloutResult) MarshalJSON() ([]byte, error) {
	type wrapper CustomExtensionCalloutResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomExtensionCalloutResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionCalloutResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customExtensionCalloutResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomExtensionCalloutResult: %+v", err)
	}

	return encoded, nil
}
