package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegateAllowedActions struct {
	// Indicates whether the delegator or delegate allows participation in active calls.
	JoinActiveCalls nullable.Type[bool] `json:"joinActiveCalls,omitempty"`

	// Indicates whether the delegator or delegate allows calls to be made on their behalf.
	MakeCalls nullable.Type[bool] `json:"makeCalls,omitempty"`

	// Indicates whether the delegator or delegate allows the management of call and delegation settings.
	ManageCallAndDelegateSettings nullable.Type[bool] `json:"manageCallAndDelegateSettings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether the delegator or delegate allows held calls to be picked up.
	PickUpHeldCalls nullable.Type[bool] `json:"pickUpHeldCalls,omitempty"`

	// Indicates whether the delegator or delegate allows calls to be received on their behalf.
	ReceiveCalls nullable.Type[bool] `json:"receiveCalls,omitempty"`
}
