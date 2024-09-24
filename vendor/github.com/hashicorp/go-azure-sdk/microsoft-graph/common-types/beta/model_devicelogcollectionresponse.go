package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceLogCollectionResponse{}

type DeviceLogCollectionResponse struct {
	// The User Principal Name (UPN) of the user that enrolled the device.
	EnrolledByUser nullable.Type[string] `json:"enrolledByUser,omitempty"`

	// The error code, if any. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
	ErrorCode *int64 `json:"errorCode,omitempty"`

	// The DateTime of the expiration of the logs.
	ExpirationDateTimeUTC nullable.Type[string] `json:"expirationDateTimeUTC,omitempty"`

	// The UPN for who initiated the request.
	InitiatedByUserPrincipalName nullable.Type[string] `json:"initiatedByUserPrincipalName,omitempty"`

	// Indicates Intune device unique identifier.
	ManagedDeviceId *string `json:"managedDeviceId,omitempty"`

	// The DateTime the request was received.
	ReceivedDateTimeUTC nullable.Type[string] `json:"receivedDateTimeUTC,omitempty"`

	// The DateTime of the request.
	RequestedDateTimeUTC nullable.Type[string] `json:"requestedDateTimeUTC,omitempty"`

	// AppLogUploadStatus
	Status *AppLogUploadState `json:"status,omitempty"`

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

func (s DeviceLogCollectionResponse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceLogCollectionResponse{}

func (s DeviceLogCollectionResponse) MarshalJSON() ([]byte, error) {
	type wrapper DeviceLogCollectionResponse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceLogCollectionResponse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceLogCollectionResponse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceLogCollectionResponse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceLogCollectionResponse: %+v", err)
	}

	return encoded, nil
}
