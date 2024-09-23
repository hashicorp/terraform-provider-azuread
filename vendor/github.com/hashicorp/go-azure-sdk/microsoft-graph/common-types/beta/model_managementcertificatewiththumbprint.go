package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementCertificateWithThumbprint struct {
	// The Base 64 encoded management certificate
	Certificate nullable.Type[string] `json:"certificate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The thumbprint of the management certificate
	Thumbprint nullable.Type[string] `json:"thumbprint,omitempty"`
}
