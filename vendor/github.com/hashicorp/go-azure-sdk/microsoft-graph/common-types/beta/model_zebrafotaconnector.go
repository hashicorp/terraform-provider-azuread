package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ZebraFotaConnector{}

type ZebraFotaConnector struct {
	// Complete account enrollment authorization URL. This corresponds to verificationuricomplete in the Zebra API
	// documentations.
	EnrollmentAuthorizationUrl *string `json:"enrollmentAuthorizationUrl,omitempty"`

	// Tenant enrollment token from Zebra. The token is used to enroll Zebra devices in the FOTA Service via app config.
	EnrollmentToken *string `json:"enrollmentToken,omitempty"`

	// Flag indicating if required Firmware Over-the-Air (FOTA) Apps have been approved.
	FotaAppsApproved *bool `json:"fotaAppsApproved,omitempty"`

	// Date and time when the account was last synched with Zebra
	LastSyncDateTime *string `json:"lastSyncDateTime,omitempty"`

	// Represents various states for Zebra FOTA connector.
	State *ZebraFotaConnectorState `json:"state,omitempty"`

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

func (s ZebraFotaConnector) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ZebraFotaConnector{}

func (s ZebraFotaConnector) MarshalJSON() ([]byte, error) {
	type wrapper ZebraFotaConnector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ZebraFotaConnector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ZebraFotaConnector: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.zebraFotaConnector"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ZebraFotaConnector: %+v", err)
	}

	return encoded, nil
}
