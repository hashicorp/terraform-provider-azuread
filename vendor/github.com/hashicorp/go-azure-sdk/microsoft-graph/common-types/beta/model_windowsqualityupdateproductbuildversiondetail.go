package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateProductBuildVersionDetail struct {
	// The build number of the product release, Allowed range is 0 - 2,147,483,647. For example: 19045. Read-only.
	BuildNumber *int64 `json:"buildNumber,omitempty"`

	// The major version of the product release, Allowed range is 0 - 2,147,483,647. For example: 10. Read-only.
	MajorVersionNumber *int64 `json:"majorVersionNumber,omitempty"`

	// The minor version of the product release, Allowed range is 0 - 2,147,483,647. For example: 0. Read-only.
	MinorVersionNumber *int64 `json:"minorVersionNumber,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The update build revision number of the product revision for the corresponding patch, Allowed range is 0 -
	// 2,147,483,647. For example: 4780. Read-only.
	UpdateBuildRevision *int64 `json:"updateBuildRevision,omitempty"`
}

var _ json.Marshaler = WindowsQualityUpdateProductBuildVersionDetail{}

func (s WindowsQualityUpdateProductBuildVersionDetail) MarshalJSON() ([]byte, error) {
	type wrapper WindowsQualityUpdateProductBuildVersionDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsQualityUpdateProductBuildVersionDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsQualityUpdateProductBuildVersionDetail: %+v", err)
	}

	delete(decoded, "buildNumber")
	delete(decoded, "majorVersionNumber")
	delete(decoded, "minorVersionNumber")
	delete(decoded, "updateBuildRevision")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsQualityUpdateProductBuildVersionDetail: %+v", err)
	}

	return encoded, nil
}
