package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatDetectionDetail struct {
	// Indicates the confidence level in the threat detection.
	ConfidenceLevel nullable.Type[string] `json:"confidenceLevel,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates if the account has priority protection enabled.
	PriorityAccountProtection nullable.Type[string] `json:"priorityAccountProtection,omitempty"`

	// Lists the detected threats.
	Threats nullable.Type[string] `json:"threats,omitempty"`
}
