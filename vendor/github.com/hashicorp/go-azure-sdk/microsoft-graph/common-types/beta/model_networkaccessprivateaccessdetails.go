package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPrivateAccessDetails struct {
	AccessType       *NetworkaccessAccessType       `json:"accessType,omitempty"`
	ConnectionStatus *NetworkaccessConnectionStatus `json:"connectionStatus,omitempty"`
	ConnectorId      nullable.Type[string]          `json:"connectorId,omitempty"`
	ConnectorIp      nullable.Type[string]          `json:"connectorIp,omitempty"`
	ConnectorName    nullable.Type[string]          `json:"connectorName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ProcessingRegion       nullable.Type[string]                `json:"processingRegion,omitempty"`
	ThirdPartyTokenDetails *NetworkaccessThirdPartyTokenDetails `json:"thirdPartyTokenDetails,omitempty"`
}
