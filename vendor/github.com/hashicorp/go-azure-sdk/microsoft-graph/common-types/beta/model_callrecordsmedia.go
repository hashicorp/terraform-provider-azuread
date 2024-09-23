package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMedia struct {
	// Device information associated with the callee endpoint of this media.
	CalleeDevice *CallRecordsDeviceInfo `json:"calleeDevice,omitempty"`

	// Network information associated with the callee endpoint of this media.
	CalleeNetwork *CallRecordsNetworkInfo `json:"calleeNetwork,omitempty"`

	// Device information associated with the caller endpoint of this media.
	CallerDevice *CallRecordsDeviceInfo `json:"callerDevice,omitempty"`

	// Network information associated with the caller endpoint of this media.
	CallerNetwork *CallRecordsNetworkInfo `json:"callerNetwork,omitempty"`

	// How the media was identified during media negotiation stage.
	Label nullable.Type[string] `json:"label,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Network streams associated with this media.
	Streams *[]CallRecordsMediaStream `json:"streams,omitempty"`
}
