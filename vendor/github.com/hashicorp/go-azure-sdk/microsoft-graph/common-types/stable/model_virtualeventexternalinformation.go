package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventExternalInformation struct {
	// Identifier of the application that hosts the externalEventId. Read-only.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// The identifier for a virtualEventExternalInformation object that associates the virtual event with an event ID in an
	// external application. This association bundles all the information (both supported and not supported in virtualEvent)
	// into one virtual event object. Optional. If set, the maximum supported length is 256 characters.
	ExternalEventId nullable.Type[string] `json:"externalEventId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = VirtualEventExternalInformation{}

func (s VirtualEventExternalInformation) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventExternalInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventExternalInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventExternalInformation: %+v", err)
	}

	delete(decoded, "applicationId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventExternalInformation: %+v", err)
	}

	return encoded, nil
}
