package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CertificateConnectorDetails{}

type CertificateConnectorDetails struct {
	// Connector name (set during enrollment).
	ConnectorName nullable.Type[string] `json:"connectorName,omitempty"`

	// Version of the connector installed.
	ConnectorVersion nullable.Type[string] `json:"connectorVersion,omitempty"`

	// Date/time when this connector was enrolled.
	EnrollmentDateTime *string `json:"enrollmentDateTime,omitempty"`

	// Date/time when this connector last connected to the service.
	LastCheckinDateTime *string `json:"lastCheckinDateTime,omitempty"`

	// Name of the machine hosting this connector service.
	MachineName nullable.Type[string] `json:"machineName,omitempty"`

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

func (s CertificateConnectorDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CertificateConnectorDetails{}

func (s CertificateConnectorDetails) MarshalJSON() ([]byte, error) {
	type wrapper CertificateConnectorDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CertificateConnectorDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CertificateConnectorDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.certificateConnectorDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CertificateConnectorDetails: %+v", err)
	}

	return encoded, nil
}
