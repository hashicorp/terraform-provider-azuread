package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesBuildVersionDetails struct {
	// The build number of the product release. Read-only.
	BuildNumber *int64 `json:"buildNumber,omitempty"`

	// The major version of the product release. Read-only.
	MajorVersion *int64 `json:"majorVersion,omitempty"`

	// The minor version of the product release. Read-only.
	MinorVersion *int64 `json:"minorVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The update build revision number of the product revision. Read-only.
	UpdateBuildRevision *int64 `json:"updateBuildRevision,omitempty"`
}

var _ json.Marshaler = WindowsUpdatesBuildVersionDetails{}

func (s WindowsUpdatesBuildVersionDetails) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesBuildVersionDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesBuildVersionDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesBuildVersionDetails: %+v", err)
	}

	delete(decoded, "buildNumber")
	delete(decoded, "majorVersion")
	delete(decoded, "minorVersion")
	delete(decoded, "updateBuildRevision")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesBuildVersionDetails: %+v", err)
	}

	return encoded, nil
}
