package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsInstallationOptionsForWindows struct {
	// Specifies whether users can install Microsoft 365 apps, including Skype for Business, on their Windows devices. The
	// default value is true.
	IsMicrosoft365AppsEnabled *bool `json:"isMicrosoft365AppsEnabled,omitempty"`

	// Specifies whether users can install Microsoft Project on their Windows devices. The default value is true.
	IsProjectEnabled *bool `json:"isProjectEnabled,omitempty"`

	// Specifies whether users can install Skype for Business (standalone) on their Windows devices. The default value is
	// true.
	IsSkypeForBusinessEnabled *bool `json:"isSkypeForBusinessEnabled,omitempty"`

	// Specifies whether users can install Visio on their Windows devices. The default value is true.
	IsVisioEnabled *bool `json:"isVisioEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
