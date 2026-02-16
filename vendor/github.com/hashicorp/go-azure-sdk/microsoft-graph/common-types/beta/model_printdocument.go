package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PrintDocument{}

type PrintDocument struct {
	Configuration *PrinterDocumentConfiguration `json:"configuration,omitempty"`

	// The document's content (MIME) type. Read-only.
	ContentType nullable.Type[string] `json:"contentType,omitempty"`

	// The document's name. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	DownloadedDateTime nullable.Type[string] `json:"downloadedDateTime,omitempty"`

	// The document's size in bytes. Read-only.
	Size *int64 `json:"size,omitempty"`

	UploadedDateTime nullable.Type[string] `json:"uploadedDateTime,omitempty"`

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

func (s PrintDocument) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrintDocument{}

func (s PrintDocument) MarshalJSON() ([]byte, error) {
	type wrapper PrintDocument
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrintDocument: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrintDocument: %+v", err)
	}

	delete(decoded, "contentType")
	delete(decoded, "displayName")
	delete(decoded, "size")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printDocument"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrintDocument: %+v", err)
	}

	return encoded, nil
}
