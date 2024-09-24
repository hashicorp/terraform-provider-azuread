package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCvssSummary struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The CVSS severity rating for this vulnerability. The possible values are: none, low, medium, high, critical,
	// unknownFutureValue.
	Severity *SecurityVulnerabilitySeverity `json:"severity,omitempty"`

	// The CVSS vector string for this vulnerability.
	VectorString nullable.Type[string] `json:"vectorString,omitempty"`
}
