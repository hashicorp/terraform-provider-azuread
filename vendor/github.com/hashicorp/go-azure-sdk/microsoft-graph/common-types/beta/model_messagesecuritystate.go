package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageSecurityState struct {
	ConnectingIP            nullable.Type[string] `json:"connectingIP,omitempty"`
	DeliveryAction          nullable.Type[string] `json:"deliveryAction,omitempty"`
	DeliveryLocation        nullable.Type[string] `json:"deliveryLocation,omitempty"`
	Directionality          nullable.Type[string] `json:"directionality,omitempty"`
	InternetMessageId       nullable.Type[string] `json:"internetMessageId,omitempty"`
	MessageFingerprint      nullable.Type[string] `json:"messageFingerprint,omitempty"`
	MessageReceivedDateTime nullable.Type[string] `json:"messageReceivedDateTime,omitempty"`
	MessageSubject          nullable.Type[string] `json:"messageSubject,omitempty"`
	NetworkMessageId        nullable.Type[string] `json:"networkMessageId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
