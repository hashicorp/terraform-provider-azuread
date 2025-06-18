package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Certification struct {
	// URL that shows certification details for the application.
	CertificationDetailsUrl nullable.Type[string] `json:"certificationDetailsUrl,omitempty"`

	// The timestamp when the current certification for the application expires.
	CertificationExpirationDateTime nullable.Type[string] `json:"certificationExpirationDateTime,omitempty"`

	// Indicates whether the application is certified by Microsoft.
	IsCertifiedByMicrosoft nullable.Type[bool] `json:"isCertifiedByMicrosoft,omitempty"`

	// Indicates whether the application developer or publisher completed Publisher Attestation.
	IsPublisherAttested nullable.Type[bool] `json:"isPublisherAttested,omitempty"`

	// The timestamp when the certification for the application was most recently added or updated.
	LastCertificationDateTime nullable.Type[string] `json:"lastCertificationDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Marshaler = Certification{}

func (s Certification) MarshalJSON() ([]byte, error) {
	type wrapper Certification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Certification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Certification: %+v", err)
	}

	delete(decoded, "certificationDetailsUrl")
	delete(decoded, "isCertifiedByMicrosoft")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Certification: %+v", err)
	}

	return encoded, nil
}
