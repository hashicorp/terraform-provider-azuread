package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementExchangeConnector{}

type DeviceManagementExchangeConnector struct {
	// The name of the server hosting the Exchange Connector.
	ConnectorServerName nullable.Type[string] `json:"connectorServerName,omitempty"`

	// An alias assigned to the Exchange server
	ExchangeAlias nullable.Type[string] `json:"exchangeAlias,omitempty"`

	// The type of Exchange Connector.
	ExchangeConnectorType *DeviceManagementExchangeConnectorType `json:"exchangeConnectorType,omitempty"`

	// Exchange Organization to the Exchange server
	ExchangeOrganization nullable.Type[string] `json:"exchangeOrganization,omitempty"`

	// Last sync time for the Exchange Connector
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// Email address used to configure the Service To Service Exchange Connector.
	PrimarySmtpAddress nullable.Type[string] `json:"primarySmtpAddress,omitempty"`

	// The name of the Exchange server.
	ServerName nullable.Type[string] `json:"serverName,omitempty"`

	// The current status of the Exchange Connector.
	Status *DeviceManagementExchangeConnectorStatus `json:"status,omitempty"`

	// The version of the ExchangeConnectorAgent
	Version nullable.Type[string] `json:"version,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementExchangeConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementExchangeConnector{}

func (s DeviceManagementExchangeConnector) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementExchangeConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementExchangeConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementExchangeConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementExchangeConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementExchangeConnector: %+v", err)
	}

	return encoded, nil
}
