package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportSuspiciousActivitySettings struct {
	IncludeTarget *IncludeTarget `json:"includeTarget,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *AdvancedConfigState `json:"state,omitempty"`

	// Specifies the number the user enters on their phone to report the MFA prompt as suspicious.
	VoiceReportingCode nullable.Type[int64] `json:"voiceReportingCode,omitempty"`
}
