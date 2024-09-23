package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetrieveRemoteHelpSessionResponse struct {
	// ACS Group Id
	AcsGroupId *string `json:"acsGroupId,omitempty"`

	// Helper ACS User Id
	AcsHelperUserId *string `json:"acsHelperUserId,omitempty"`

	// Helper ACS User Token
	AcsHelperUserToken *string `json:"acsHelperUserToken,omitempty"`

	// Sharer ACS User Id
	AcsSharerUserId *string `json:"acsSharerUserId,omitempty"`

	// Android Device Name
	DeviceName *string `json:"deviceName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Azure Pubsub Group Id
	PubSubGroupId *string `json:"pubSubGroupId,omitempty"`

	// Azure Pubsub Group Id
	PubSubHelperAccessUri *string `json:"pubSubHelperAccessUri,omitempty"`

	// Azure Pubsub Session Expiration Date Time.
	SessionExpirationDateTime *string `json:"sessionExpirationDateTime,omitempty"`

	// The unique identifier for a session
	SessionKey *string `json:"sessionKey,omitempty"`
}
