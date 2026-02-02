package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppLogUpload struct {
	// The Mobile Application Management (MAM) Logs Uploading Component. Such components can be the application itself, the
	// MAM SDK, and other on-device components that are capable of uploading diagnostic logs. Read-only.
	ManagedAppComponent nullable.Type[string] `json:"managedAppComponent,omitempty"`

	// The Mobile Application Management (MAM) Logs Uploading Component. Such components can be the application itself, the
	// MAM SDK, and other on-device components that are capable of uploading diagnostic logs. Read-only.
	ManagedAppComponentDescription nullable.Type[string] `json:"managedAppComponentDescription,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A provider-specific reference id for the uploaded logs. Read-only.
	ReferenceId nullable.Type[string] `json:"referenceId,omitempty"`

	// Represents the current status of the associated `managedAppLogCollectionRequest`.
	Status *ManagedAppLogUploadState `json:"status,omitempty"`
}

var _ json.Marshaler = ManagedAppLogUpload{}

func (s ManagedAppLogUpload) MarshalJSON() ([]byte, error) {
	type wrapper ManagedAppLogUpload
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedAppLogUpload: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedAppLogUpload: %+v", err)
	}

	delete(decoded, "managedAppComponent")
	delete(decoded, "managedAppComponentDescription")
	delete(decoded, "referenceId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedAppLogUpload: %+v", err)
	}

	return encoded, nil
}
