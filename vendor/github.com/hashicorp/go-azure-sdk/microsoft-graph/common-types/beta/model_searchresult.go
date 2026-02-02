package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchResult struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A callback URL that can be used to record telemetry information. The application should issue a GET on this URL if
	// the user interacts with this item to improve the quality of results.
	OnClickTelemetryUrl nullable.Type[string] `json:"onClickTelemetryUrl,omitempty"`
}
