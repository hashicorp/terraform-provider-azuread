package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendRemoteHelpSessionResponse struct {
	// Helper ACS User Token
	AcsHelperUserToken *string `json:"acsHelperUserToken,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Azure Pubsub Group Id
	PubSubHelperAccessUri *string `json:"pubSubHelperAccessUri,omitempty"`

	// Azure Pubsub Session Expiration Date Time.
	SessionExpirationDateTime *string `json:"sessionExpirationDateTime,omitempty"`

	// The unique identifier for a session
	SessionKey *string `json:"sessionKey,omitempty"`
}
