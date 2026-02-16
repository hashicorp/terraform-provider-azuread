package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UploadSession interface {
	UploadSession() BaseUploadSessionImpl
}

var _ UploadSession = BaseUploadSessionImpl{}

type BaseUploadSessionImpl struct {
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

func (s BaseUploadSessionImpl) UploadSession() BaseUploadSessionImpl {
	return s
}

var _ UploadSession = RawUploadSessionImpl{}

// RawUploadSessionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUploadSessionImpl struct {
	uploadSession BaseUploadSessionImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawUploadSessionImpl) UploadSession() BaseUploadSessionImpl {
	return s.uploadSession
}

func UnmarshalUploadSessionImplementation(input []byte) (UploadSession, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UploadSession into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.engagementUploadSession") {
		var out EngagementUploadSession
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EngagementUploadSession: %+v", err)
		}
		return out, nil
	}

	var parent BaseUploadSessionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUploadSessionImpl: %+v", err)
	}

	return RawUploadSessionImpl{
		uploadSession: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
