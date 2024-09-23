package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrintConnector{}

type PrintConnector struct {
	// The connector's version.
	AppVersion *string `json:"appVersion,omitempty"`

	// The connector's device health.
	DeviceHealth *DeviceHealth `json:"deviceHealth,omitempty"`

	// The name of the connector.
	DisplayName *string `json:"displayName,omitempty"`

	// The connector machine's hostname.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// The physical and/or organizational location of the connector.
	Location *PrinterLocation `json:"location,omitempty"`

	Name nullable.Type[string] `json:"name,omitempty"`

	// The connector machine's operating system version.
	OperatingSystem *string `json:"operatingSystem,omitempty"`

	// The DateTimeOffset when the connector was registered.
	RegisteredDateTime *string `json:"registeredDateTime,omitempty"`

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

func (s PrintConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrintConnector{}

func (s PrintConnector) MarshalJSON() ([]byte, error) {
	type wrapper PrintConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintConnector: %+v", err)
	}

	return encoded, nil
}
