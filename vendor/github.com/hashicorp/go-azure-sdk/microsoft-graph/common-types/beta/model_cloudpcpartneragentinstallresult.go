package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResult struct {
	// Contains a detailed error message when the partner agent installation failed.
	ErrorMessage nullable.Type[string] `json:"errorMessage,omitempty"`

	// The status of a partner agent installation. Possible values are: installed, installFailed, installing, uninstalling,
	// uninstallFailed and licensed. Read-Only.
	InstallStatus *CloudPCPartnerAgentInstallStatus `json:"installStatus,omitempty"`

	// Indicates whether the partner agent is a third party. When true, the agent is a third-party (non-Microsoft) agent and
	// when false, the agent is a Microsoft agent or isn't known. The default value is false.
	IsThirdPartyPartner nullable.Type[bool] `json:"isThirdPartyPartner,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The name of the first-party or third-party partner agent. Possible values for third-party partners are Citrix, VMware
	// and HP. Read-Only.
	PartnerAgentName *CloudPCPartnerAgentName `json:"partnerAgentName,omitempty"`

	// Indicates whether the partner agent installation should be retried. The default value is false.
	Retriable nullable.Type[bool] `json:"retriable,omitempty"`
}
