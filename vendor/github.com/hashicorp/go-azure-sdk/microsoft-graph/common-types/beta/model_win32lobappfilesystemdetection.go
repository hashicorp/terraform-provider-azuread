package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppDetection = Win32LobAppFileSystemDetection{}

type Win32LobAppFileSystemDetection struct {
	// A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`

	// Contains all supported file system detection type.
	DetectionType *Win32LobAppFileSystemDetectionType `json:"detectionType,omitempty"`

	// The file or folder detection value
	DetectionValue nullable.Type[string] `json:"detectionValue,omitempty"`

	// The file or folder name to detect Win32 Line of Business (LoB) app
	FileOrFolderName nullable.Type[string] `json:"fileOrFolderName,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`

	// The file or folder path to detect Win32 Line of Business (LoB) app
	Path nullable.Type[string] `json:"path,omitempty"`

	// Fields inherited from Win32LobAppDetection

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Win32LobAppFileSystemDetection) Win32LobAppDetection() BaseWin32LobAppDetectionImpl {
	return BaseWin32LobAppDetectionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32LobAppFileSystemDetection{}

func (s Win32LobAppFileSystemDetection) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppFileSystemDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppFileSystemDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppFileSystemDetection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppFileSystemDetection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppFileSystemDetection: %+v", err)
	}

	return encoded, nil
}
