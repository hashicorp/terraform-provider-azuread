package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesCveInformation struct {
	// Identifies the number of the CVE. Read-only.
	Number *string `json:"number,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// URL to the full CVE information. Read-only.
	Url nullable.Type[string] `json:"url,omitempty"`
}

var _ json.Marshaler = WindowsUpdatesCveInformation{}

func (s WindowsUpdatesCveInformation) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesCveInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesCveInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesCveInformation: %+v", err)
	}

	delete(decoded, "number")
	delete(decoded, "url")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesCveInformation: %+v", err)
	}

	return encoded, nil
}
