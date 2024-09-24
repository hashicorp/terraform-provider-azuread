package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsEnrollmentStatusScreenSettings struct {
	// Allow or block user to use device before profile and app installation complete
	AllowDeviceUseBeforeProfileAndAppInstallComplete *bool `json:"allowDeviceUseBeforeProfileAndAppInstallComplete,omitempty"`

	// Allow the user to continue using the device on installation failure
	AllowDeviceUseOnInstallFailure *bool `json:"allowDeviceUseOnInstallFailure,omitempty"`

	// Allow or block log collection on installation failure
	AllowLogCollectionOnInstallFailure *bool `json:"allowLogCollectionOnInstallFailure,omitempty"`

	// Allow the user to retry the setup on installation failure
	BlockDeviceSetupRetryByUser *bool `json:"blockDeviceSetupRetryByUser,omitempty"`

	// Set custom error message to show upon installation failure
	CustomErrorMessage nullable.Type[string] `json:"customErrorMessage,omitempty"`

	// Show or hide installation progress to user
	HideInstallationProgress *bool `json:"hideInstallationProgress,omitempty"`

	// Set installation progress timeout in minutes
	InstallProgressTimeoutInMinutes nullable.Type[int64] `json:"installProgressTimeoutInMinutes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
