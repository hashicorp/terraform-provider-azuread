package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppContentFile{}

type MobileAppContentFile struct {
	// The Azure Storage URI.
	AzureStorageUri nullable.Type[string] `json:"azureStorageUri,omitempty"`

	// The time the Azure storage Uri expires.
	AzureStorageUriExpirationDateTime nullable.Type[string] `json:"azureStorageUriExpirationDateTime,omitempty"`

	// The time the file was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A value indicating whether the file is committed.
	IsCommitted *bool `json:"isCommitted,omitempty"`

	// Indicates whether this content file is a dependency for the main content file. TRUE means that the content file is a
	// dependency, FALSE means that the content file is not a dependency and is the main content file. Defaults to FALSE.
	IsDependency *bool `json:"isDependency,omitempty"`

	// The manifest information.
	Manifest nullable.Type[string] `json:"manifest,omitempty"`

	// the file name.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The size of the file prior to encryption.
	Size *int64 `json:"size,omitempty"`

	// The size of the file after encryption.
	SizeEncrypted *int64 `json:"sizeEncrypted,omitempty"`

	// Contains properties for upload request states.
	UploadState *MobileAppContentFileUploadState `json:"uploadState,omitempty"`

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

func (s MobileAppContentFile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MobileAppContentFile{}

func (s MobileAppContentFile) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppContentFile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppContentFile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppContentFile: %+v", err)
	}

	delete(decoded, "azureStorageUri")
	delete(decoded, "azureStorageUriExpirationDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "isCommitted")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.mobileAppContentFile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppContentFile: %+v", err)
	}

	return encoded, nil
}
