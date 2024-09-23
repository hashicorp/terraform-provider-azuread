package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesCurrentExportData struct {
	// The name of the onPremises client machine that ran the last export.
	ClientMachineName nullable.Type[string] `json:"clientMachineName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The count of pending adds from on-premises directory.
	PendingObjectsAddition nullable.Type[int64] `json:"pendingObjectsAddition,omitempty"`

	// The count of pending deletes from on-premises directory.
	PendingObjectsDeletion nullable.Type[int64] `json:"pendingObjectsDeletion,omitempty"`

	// The count of pending updates from on-premises directory.
	PendingObjectsUpdate nullable.Type[int64] `json:"pendingObjectsUpdate,omitempty"`

	// The name of the dirsync service account that is configured to connect to the directory.
	ServiceAccount nullable.Type[string] `json:"serviceAccount,omitempty"`

	// The count of updated links during the current directory sync export run.
	SuccessfulLinksProvisioningCount nullable.Type[int64] `json:"successfulLinksProvisioningCount,omitempty"`

	// The count of objects that were successfully provisioned during the current directory sync export run.
	SuccessfulObjectsProvisioningCount nullable.Type[int64] `json:"successfulObjectsProvisioningCount,omitempty"`

	// The total number of objects in the AAD Connector Space.
	TotalConnectorSpaceObjects nullable.Type[int64] `json:"totalConnectorSpaceObjects,omitempty"`
}
