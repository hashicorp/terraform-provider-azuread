package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationFlow struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the transfer methods in scope for the policy. The possible values are: none, deviceCodeFlow,
	// authenticationTransfer, unknownFutureValue. Default value is none.
	TransferMethod *ConditionalAccessTransferMethods `json:"transferMethod,omitempty"`
}
