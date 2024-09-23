package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetonationDetails struct {
	// The time of detonation.
	AnalysisDateTime nullable.Type[string] `json:"analysisDateTime,omitempty"`

	// The chain of detonation.
	DetonationChain *SecurityDetonationChain `json:"detonationChain,omitempty"`

	// All observables in the detonation tree.
	DetonationObservables *SecurityDetonationObservables `json:"detonationObservables,omitempty"`

	// The verdict of the detonation.
	DetonationVerdict nullable.Type[string] `json:"detonationVerdict,omitempty"`

	// The reason for the verdict of the detonation.
	DetonationVerdictReason nullable.Type[string] `json:"detonationVerdictReason,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
