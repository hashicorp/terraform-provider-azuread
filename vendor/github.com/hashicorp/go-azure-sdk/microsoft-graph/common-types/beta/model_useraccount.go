package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccount struct {
	DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RiskScore  nullable.Type[string] `json:"riskScore,omitempty"`
	Service    nullable.Type[string] `json:"service,omitempty"`
	SigninName nullable.Type[string] `json:"signinName,omitempty"`
	Status     *AccountStatus        `json:"status,omitempty"`
}
