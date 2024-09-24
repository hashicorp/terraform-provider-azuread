package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRuleSchedule struct {
	// Timestamp of the custom detection rule's next scheduled run.
	NextRunDateTime nullable.Type[string] `json:"nextRunDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// How often the detection rule is set to run. The allowed values are: 0, 1H, 3H, 12H, or 24H. '0' signifies the rule is
	// run continuously.
	Period *string `json:"period,omitempty"`
}
