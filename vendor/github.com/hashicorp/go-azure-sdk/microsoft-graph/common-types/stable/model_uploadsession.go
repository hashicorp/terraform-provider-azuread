package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UploadSession struct {
	// The date and time in UTC that the upload session will expire. The complete file must be uploaded before this
	// expiration time is reached.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// A collection of byte ranges that the server is missing for the file. These ranges are zero indexed and of the format
	// 'start-end' (for example '0-26' to indicate the first 27 bytes of the file). When uploading files as Outlook
	// attachments, instead of a collection of ranges, this property always indicates a single value '{start}', the location
	// in the file where the next upload should begin.
	NextExpectedRanges *[]string `json:"nextExpectedRanges,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The URL endpoint that accepts PUT requests for byte ranges of the file.
	UploadUrl nullable.Type[string] `json:"uploadUrl,omitempty"`
}
