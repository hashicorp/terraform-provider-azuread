package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftAppDiagnosticDownloadRequest struct {
	// The list of files to download which is associated with the diagnostic.
	Files *[]string `json:"files,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique id for the request that serves as an identifer for the diagnostic to be downloaded.
	PowerliftId nullable.Type[string] `json:"powerliftId,omitempty"`
}
