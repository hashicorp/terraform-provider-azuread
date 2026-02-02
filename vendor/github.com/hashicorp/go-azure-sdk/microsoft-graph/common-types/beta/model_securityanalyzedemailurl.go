package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailUrl struct {
	// The method used to detect threats in the URL.
	DetectionMethod nullable.Type[string] `json:"detectionMethod,omitempty"`

	// Detonation data associated with the URL.
	DetonationDetails *SecurityDetonationDetails `json:"detonationDetails,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Details of entries in tenant allow/block list configured by tenant.
	TenantAllowBlockListDetailInfo nullable.Type[string] `json:"tenantAllowBlockListDetailInfo,omitempty"`

	// The type of threat associated with the URL. The possible values are: unknown, spam, malware, phishing, none,
	// unknownFutureValue.
	ThreatType *SecurityThreatType `json:"threatType,omitempty"`

	// The URL that is found in the email. This is full URL string, including query parameters.
	Url nullable.Type[string] `json:"url,omitempty"`
}
