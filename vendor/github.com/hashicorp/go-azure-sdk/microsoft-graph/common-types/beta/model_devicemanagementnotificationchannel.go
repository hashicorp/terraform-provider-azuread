package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationChannel struct {
	// The type of the notification channel. The possible values are: portal, email, phoneCall, sms, unknownFutureValue.
	NotificationChannelType *DeviceManagementNotificationChannelType `json:"notificationChannelType,omitempty"`

	// Information about the notification receivers, such as locale and contact information. For example, en-us for locale
	// and serena.davis@contoso.com for contact information.
	NotificationReceivers *[]DeviceManagementNotificationReceiver `json:"notificationReceivers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
