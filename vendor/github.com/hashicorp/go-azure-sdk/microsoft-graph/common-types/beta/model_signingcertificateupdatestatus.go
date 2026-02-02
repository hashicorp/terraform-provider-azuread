package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SigningCertificateUpdateStatus struct {
	// Status of the last certificate update. Read-only. For a list of statuses, see certificateUpdateResult status.
	CertificateUpdateResult nullable.Type[string] `json:"certificateUpdateResult,omitempty"`

	// Date and time in ISO 8601 format and in UTC time when the certificate was last updated. Read-only.
	LastRunDateTime nullable.Type[string] `json:"lastRunDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = SigningCertificateUpdateStatus{}

func (s SigningCertificateUpdateStatus) MarshalJSON() ([]byte, error) {
	type wrapper SigningCertificateUpdateStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SigningCertificateUpdateStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SigningCertificateUpdateStatus: %+v", err)
	}

	delete(decoded, "certificateUpdateResult")
	delete(decoded, "lastRunDateTime")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SigningCertificateUpdateStatus: %+v", err)
	}

	return encoded, nil
}
