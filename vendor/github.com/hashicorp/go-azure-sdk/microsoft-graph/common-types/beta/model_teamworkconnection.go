package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConnection struct {
	// Indicates whether a component/peripheral is connected/disconnected or its state is unknown. The possible values are:
	// unknown, connected, disconnected, unknownFutureValue.
	ConnectionStatus *TeamworkConnectionStatus `json:"connectionStatus,omitempty"`

	// Time at which the state was last changed. For example, indicates connected since when the state is connected and
	// disconnected since when the state is disconnected.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
