package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutInstance struct {
	// Identification of the custom extension that was triggered at this instance.
	CustomExtensionId nullable.Type[string] `json:"customExtensionId,omitempty"`

	// Details provided by the logic app during the callback of the request instance.
	Detail nullable.Type[string] `json:"detail,omitempty"`

	// The unique run identifier for the logic app.
	ExternalCorrelationId nullable.Type[string] `json:"externalCorrelationId,omitempty"`

	// Unique identifier for the callout instance. Read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The status of the request to the custom extension. The possible values are: calloutSent, callbackReceived,
	// calloutFailed, callbackTimedOut, waitingForCallback, unknownFutureValue.
	Status *CustomExtensionCalloutInstanceStatus `json:"status,omitempty"`
}

var _ json.Marshaler = CustomExtensionCalloutInstance{}

func (s CustomExtensionCalloutInstance) MarshalJSON() ([]byte, error) {
	type wrapper CustomExtensionCalloutInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomExtensionCalloutInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionCalloutInstance: %+v", err)
	}

	delete(decoded, "id")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomExtensionCalloutInstance: %+v", err)
	}

	return encoded, nil
}
