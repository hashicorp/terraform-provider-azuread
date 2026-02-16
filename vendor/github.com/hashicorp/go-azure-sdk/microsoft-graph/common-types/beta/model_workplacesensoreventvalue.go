package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkplaceSensorEventValue struct {
	// The type of possible sensor event value. The possible values are: badgeIn, badgeOut, unknownFutureValue.
	EventType *WorkplaceSensorEventType `json:"eventType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier of a user. It could be an email or a Microsoft Entra ID.
	User *EmailIdentity `json:"user,omitempty"`
}
