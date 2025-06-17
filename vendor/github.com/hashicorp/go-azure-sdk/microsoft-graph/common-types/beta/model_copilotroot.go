package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopilotRoot struct {
	// The Microsoft 365 Copilot admin who can add or modify Copilot settings. Read-only. Nullable.
	Admin *CopilotAdmin `json:"admin,omitempty"`

	// The history of interactions between AI agents and users.
	InteractionHistory *AiInteractionHistory `json:"interactionHistory,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Settings *CopilotSetting `json:"settings,omitempty"`

	// The list of AI users or agents. Read-only. Nullable.
	Users *[]AiUser `json:"users,omitempty"`
}

var _ json.Marshaler = CopilotRoot{}

func (s CopilotRoot) MarshalJSON() ([]byte, error) {
	type wrapper CopilotRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CopilotRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CopilotRoot: %+v", err)
	}

	delete(decoded, "admin")
	delete(decoded, "users")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CopilotRoot: %+v", err)
	}

	return encoded, nil
}
