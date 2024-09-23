package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFileSynchronizationVerificationMessage struct {
	// Detailed information about the message type.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Source file that contains the error.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of the message. Possible values are: error, warning, information.
	Type nullable.Type[string] `json:"type,omitempty"`
}
