package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailAttachment struct {
	// The detonation details of the attachment.
	DetonationDetails *SecurityDetonationDetails `json:"detonationDetails,omitempty"`

	// The name of the attachment in the email.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// The type of the attachment in the email.
	FileType nullable.Type[string] `json:"fileType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The SHA256 file hash of the attachment.
	Sha256 nullable.Type[string] `json:"sha256,omitempty"`

	// The threat name associated with the threat type.
	ThreatName nullable.Type[string] `json:"threatName,omitempty"`

	// The threat type associated with the attachment. The possible values are: unknown, spam, malware, phishing, none,
	// unknownFutureValue.
	ThreatType *SecurityThreatType `json:"threatType,omitempty"`
}
