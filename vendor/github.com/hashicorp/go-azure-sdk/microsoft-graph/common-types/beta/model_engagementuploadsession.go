package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UploadSession = EngagementUploadSession{}

type EngagementUploadSession struct {
	// The ID of the session.
	Id *string `json:"id,omitempty"`

	// Fields inherited from UploadSession

	// The date and time in UTC that the upload session expires. The complete file must be uploaded before this expiration
	// time is reached.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// When uploading files to document libraries, this property is a collection of byte ranges that the server is missing
	// for the file. These ranges are zero-indexed and of the format, '{start}-{end}' (for example '0-26' to indicate the
	// first 27 bytes of the file). When uploading files as Outlook attachments, instead of a collection of ranges, this
	// property always indicates a single value '{start}', the location in the file where the next upload should begin.
	NextExpectedRanges *[]string `json:"nextExpectedRanges,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The URL endpoint that accepts PUT requests for byte ranges of the file.
	UploadUrl nullable.Type[string] `json:"uploadUrl,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EngagementUploadSession) UploadSession() BaseUploadSessionImpl {
	return BaseUploadSessionImpl{
		ExpirationDateTime: s.ExpirationDateTime,
		NextExpectedRanges: s.NextExpectedRanges,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
		UploadUrl:          s.UploadUrl,
	}
}

var _ json.Marshaler = EngagementUploadSession{}

func (s EngagementUploadSession) MarshalJSON() ([]byte, error) {
	type wrapper EngagementUploadSession
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EngagementUploadSession: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EngagementUploadSession: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.engagementUploadSession"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EngagementUploadSession: %+v", err)
	}

	return encoded, nil
}
