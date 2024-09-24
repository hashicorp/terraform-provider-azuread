package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateDetail struct {
	// MobieApp identifier.
	ApplicationId nullable.Type[string] `json:"applicationId,omitempty"`

	// The admin provided or imported title of the app.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Human readable version of the application
	DisplayVersion nullable.Type[string] `json:"displayVersion,omitempty"`

	// A list of possible states for application status on an individual device. When devices contact the Intune service and
	// find targeted application enforcement intent, the status of the enforcement is recorded and becomes accessible in the
	// Graph API. Since the application status is identified during device interaction with the Intune service, status
	// records do not immediately appear upon application group assignment; it is created only after the assignment is
	// evaluated in the service and devices start receiving the policy during check-ins.
	InstallState *ResultantAppState `json:"installState,omitempty"`

	// Indicates the status of the mobile app on the device.
	MobileAppIntent *MobileAppIntent `json:"mobileAppIntent,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The supported platforms for the app.
	SupportedDeviceTypes *[]MobileAppSupportedDeviceType `json:"supportedDeviceTypes,omitempty"`
}
