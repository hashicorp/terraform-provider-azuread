package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapability struct {
	// Indicates the state of the supported action capability to perform a Cloud PC remote action. Possible values are:
	// enabled, disabled. Default value is enabled.
	ActionCapability *ActionCapability `json:"actionCapability,omitempty"`

	// The name of the supported Cloud PC remote action. Possible values are: unknown, restart, rename, restore, resize,
	// reprovision, troubleShoot, changeUserAccountType, placeUnderReview. Default value is unknown.
	ActionName *CloudPCRemoteActionName `json:"actionName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
