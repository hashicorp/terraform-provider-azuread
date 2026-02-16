package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationDelivery struct {
	// The number of users to whom mails couldn't be delivered.
	FailedMessageDeliveryCount nullable.Type[int64] `json:"failedMessageDeliveryCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of users whose email address was successfully resolved from target users.
	ResolvedTargetsCount nullable.Type[int64] `json:"resolvedTargetsCount,omitempty"`

	// The number of users who received a mail while the training campaign was in the 'in progress' state.
	SuccessfulMessageDeliveryCount nullable.Type[int64] `json:"successfulMessageDeliveryCount,omitempty"`
}
