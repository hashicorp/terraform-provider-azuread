package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedEmailDeliveryDetail struct {
	// The delivery action of the email. The possible values are: unknown, deliveredToJunk, delivered, blocked, replaced,
	// unknownFutureValue.
	Action *SecurityDeliveryAction `json:"action,omitempty"`

	// The delivery location of the email. The possible values are: unknown, inboxfolder, junkFolder, deletedFolder,
	// quarantine, onpremexternal, failed, dropped, others, unknownFutureValue.
	Location *SecurityDeliveryLocation `json:"location,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
