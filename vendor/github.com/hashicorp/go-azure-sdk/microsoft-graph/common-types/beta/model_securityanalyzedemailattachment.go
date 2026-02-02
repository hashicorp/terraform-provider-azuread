package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailAttachment struct {
	// The detonation details of the attachment.
	DetonationDetails *SecurityDetonationDetails `json:"detonationDetails,omitempty"`

	// Extension of the file.
	FileExtension nullable.Type[string] `json:"fileExtension,omitempty"`

	// The name of the attachment in the email.
	FileName nullable.Type[string] `json:"fileName,omitempty"`

	// Size of the file.
	FileSize nullable.Type[int64] `json:"fileSize,omitempty"`

	// The type of the attachment in the email.
	FileType nullable.Type[string] `json:"fileType,omitempty"`

	// The threat name associated with the threat type.
	MalwareFamily nullable.Type[string] `json:"malwareFamily,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The SHA256 file hash of the attachment.
	Sha256 nullable.Type[string] `json:"sha256,omitempty"`

	// Details of entries in tenant allow/block list configured by tenant.
	TenantAllowBlockListDetailInfo nullable.Type[string] `json:"tenantAllowBlockListDetailInfo,omitempty"`

	// The threat type associated with the attachment. The possible values are: unknown, spam, malware, phishing, none,
	// unknownFutureValue.
	ThreatType *SecurityThreatType `json:"threatType,omitempty"`
}
