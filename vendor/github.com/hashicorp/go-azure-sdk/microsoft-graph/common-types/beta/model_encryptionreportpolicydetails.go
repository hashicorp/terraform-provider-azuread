package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionReportPolicyDetails struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Policy Id for Encryption Report
	PolicyId nullable.Type[string] `json:"policyId,omitempty"`

	// Policy Name for Encryption Report
	PolicyName nullable.Type[string] `json:"policyName,omitempty"`
}
