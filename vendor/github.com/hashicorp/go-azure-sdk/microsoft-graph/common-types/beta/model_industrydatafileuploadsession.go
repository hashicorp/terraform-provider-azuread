package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFileUploadSession struct {
	// The expiration date and time for the container. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ContainerExpirationDateTime *string `json:"containerExpirationDateTime,omitempty"`

	// The container ID where the files are uploaded.
	ContainerId *string `json:"containerId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The expiration date and time for the file upload session. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	SessionExpirationDateTime *string `json:"sessionExpirationDateTime,omitempty"`

	// The Azure Storage SAS URI to upload source files to.
	SessionUrl *string `json:"sessionUrl,omitempty"`
}

var _ json.Marshaler = IndustryDataFileUploadSession{}

func (s IndustryDataFileUploadSession) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataFileUploadSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataFileUploadSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataFileUploadSession: %+v", err)
	}

	delete(decoded, "containerExpirationDateTime")
	delete(decoded, "containerId")
	delete(decoded, "sessionExpirationDateTime")
	delete(decoded, "sessionUrl")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataFileUploadSession: %+v", err)
	}

	return encoded, nil
}
