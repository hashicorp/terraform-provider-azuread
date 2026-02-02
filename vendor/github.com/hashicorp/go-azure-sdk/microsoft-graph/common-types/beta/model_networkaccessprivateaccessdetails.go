package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPrivateAccessDetails struct {
	// Type of accessed application. Access type options: QuickAccess, PrivateAccess.
	AccessType *NetworkaccessAccessType `json:"accessType,omitempty"`

	// The unique identifier for Application segment ID from Azure AD.
	AppSegmentId nullable.Type[string] `json:"appSegmentId,omitempty"`

	// Status of a connection. Status options: Open, Active, Closed.
	ConnectionStatus *NetworkaccessConnectionStatus `json:"connectionStatus,omitempty"`

	// Private access connector ID.
	ConnectorId nullable.Type[string] `json:"connectorId,omitempty"`

	// Private access connector IP address.
	ConnectorIp nullable.Type[string] `json:"connectorIp,omitempty"`

	// Private access connector name.
	ConnectorName nullable.Type[string] `json:"connectorName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Region where the request was processed by the backend service.
	ProcessingRegion nullable.Type[string] `json:"processingRegion,omitempty"`

	// Details about third-party tokens used in the transaction.
	ThirdPartyTokenDetails *NetworkaccessThirdPartyTokenDetails `json:"thirdPartyTokenDetails,omitempty"`
}
