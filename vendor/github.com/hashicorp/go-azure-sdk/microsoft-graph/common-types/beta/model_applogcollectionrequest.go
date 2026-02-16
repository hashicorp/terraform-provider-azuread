package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AppLogCollectionRequest{}

type AppLogCollectionRequest struct {
	// Time at which the upload log request reached a completed state if not completed yet NULL will be returned.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// List of log folders.
	CustomLogFolders *[]string `json:"customLogFolders,omitempty"`

	// Indicates error message if any during the upload process.
	ErrorMessage nullable.Type[string] `json:"errorMessage,omitempty"`

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

func (s AppLogCollectionRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AppLogCollectionRequest{}

func (s AppLogCollectionRequest) MarshalJSON() ([]byte, error) {
	type wrapper AppLogCollectionRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AppLogCollectionRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AppLogCollectionRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.appLogCollectionRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AppLogCollectionRequest: %+v", err)
	}

	return encoded, nil
}
