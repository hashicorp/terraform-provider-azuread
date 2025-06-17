package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCompromiseIndicator struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicator.
	Value nullable.Type[string] `json:"value,omitempty"`

	// .The possible values are: none, malware, phish, siteUnavailable, spam, decryptionFailed, unsupportedUriScheme,
	// unsupportedFileType, undefined, unknownFutureValue.
	Verdict *SecurityVerdictCategory `json:"verdict,omitempty"`
}
