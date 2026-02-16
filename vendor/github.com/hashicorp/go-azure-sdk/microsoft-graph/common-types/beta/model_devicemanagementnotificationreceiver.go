package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationReceiver struct {
	// The contact information about the notification receivers, such as an email address. Currently, only email and portal
	// notifications are supported. For portal notifications, contactInformation can be left blank. For email notifications,
	// contactInformation consists of an email address such as serena.davis@contoso.com.
	ContactInformation nullable.Type[string] `json:"contactInformation,omitempty"`

	// Defines the language and format in which the notification will be sent. Supported locale values are: en-us, cs-cz,
	// de-de, es-es, fr-fr, hu-hu, it-it, ja-jp, ko-kr, nl-nl, pl-pl, pt-br, pt-pt, ru-ru, sv-se, tr-tr, zh-cn, zh-tw.
	Locale nullable.Type[string] `json:"locale,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
