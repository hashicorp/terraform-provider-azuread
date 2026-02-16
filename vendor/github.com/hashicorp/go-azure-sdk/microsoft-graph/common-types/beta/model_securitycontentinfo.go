package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentInfo struct {
	// The format of the content to be labeled. Possible values are: file, email.
	ContentFormat nullable.Type[string] `json:"contentFormat,omitempty"`

	// Identifier used for Azure Information Protection Analytics.
	Identifier nullable.Type[string] `json:"identifier,omitempty"`

	// Existing Microsoft Purview Information Protection metadata is passed as key-value pairs, where the key is the
	// MSIPLabelGUID_PropName.
	Metadata *[]SecurityKeyValuePair `json:"metadata,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *SecurityContentState `json:"state,omitempty"`
}
