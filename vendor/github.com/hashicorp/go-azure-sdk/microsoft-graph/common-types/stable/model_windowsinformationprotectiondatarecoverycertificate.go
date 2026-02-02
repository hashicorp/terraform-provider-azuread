package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDataRecoveryCertificate struct {
	// Data recovery Certificate
	Certificate nullable.Type[string] `json:"certificate,omitempty"`

	// Data recovery Certificate description
	Description nullable.Type[string] `json:"description,omitempty"`

	// Data recovery Certificate expiration datetime
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Data recovery Certificate subject name
	SubjectName nullable.Type[string] `json:"subjectName,omitempty"`
}
