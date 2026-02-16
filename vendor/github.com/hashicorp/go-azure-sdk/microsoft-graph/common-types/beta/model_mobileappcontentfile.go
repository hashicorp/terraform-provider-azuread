package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MobileAppContentFile{}

type MobileAppContentFile struct {
	// Indicates the Azure Storage URI that the file is uploaded to. Created by the service upon receiving a valid
	// mobileAppContentFile. Read-only. This property is read-only.
	AzureStorageUri nullable.Type[string] `json:"azureStorageUri,omitempty"`

	// Indicates the date and time when the Azure storage URI expires, in ISO 8601 format. For example, midnight UTC on Jan
	// 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
	AzureStorageUriExpirationDateTime nullable.Type[string] `json:"azureStorageUriExpirationDateTime,omitempty"`

	// Indicates created date and time associated with app content file, in ISO 8601 format. For example, midnight UTC on
	// Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Read-only. This property is read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// A value indicating whether the file is committed. A committed app content file has been fully uploaded and validated
	// by the Intune service. TRUE means that app content file is committed, FALSE means that app content file is not
	// committed. Defaults to FALSE. Read-only. This property is read-only.
	IsCommitted *bool `json:"isCommitted,omitempty"`

	// Indicates whether this content file is a dependency for the main content file. TRUE means that the content file is a
	// dependency, FALSE means that the content file is not a dependency and is the main content file. Defaults to FALSE.
	IsDependency *bool `json:"isDependency,omitempty"`

	// A value indicating whether the file is a framework file. To be deprecated.
	IsFrameworkFile *bool `json:"isFrameworkFile,omitempty"`

	// Indicates the manifest information, containing file metadata.
	Manifest nullable.Type[string] `json:"manifest,omitempty"`

	// Indicates the name of the file.
	Name nullable.Type[string] `json:"name,omitempty"`

	// Indicates the original size of the file, in bytes.
	Size *int64 `json:"size,omitempty"`

	// Indicates the size of the file after encryption, in bytes.
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
