package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRequirement = Win32LobAppFileSystemRequirement{}

type Win32LobAppFileSystemRequirement struct {
	// A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`

	// Contains all supported file system detection type.
	DetectionType *Win32LobAppFileSystemDetectionType `json:"detectionType,omitempty"`

	// The file or folder name to detect Win32 Line of Business (LoB) app
	FileOrFolderName nullable.Type[string] `json:"fileOrFolderName,omitempty"`

	// The file or folder path to detect Win32 Line of Business (LoB) app
	Path nullable.Type[string] `json:"path,omitempty"`

	// Fields inherited from Win32LobAppRequirement

	// The detection value
	DetectionValue nullable.Type[string] `json:"detectionValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Win32LobAppFileSystemRequirement) Win32LobAppRequirement() BaseWin32LobAppRequirementImpl {
	return BaseWin32LobAppRequirementImpl{
		DetectionValue: s.DetectionValue,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		Operator:       s.Operator,
	}
}

var _ json.Marshaler = Win32LobAppFileSystemRequirement{}

func (s Win32LobAppFileSystemRequirement) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppFileSystemRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppFileSystemRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppFileSystemRequirement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppFileSystemRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppFileSystemRequirement: %+v", err)
	}

	return encoded, nil
}
