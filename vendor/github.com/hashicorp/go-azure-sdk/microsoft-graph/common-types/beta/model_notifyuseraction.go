package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DlpActionInfo = NotifyUserAction{}

type NotifyUserAction struct {
	// Timestamp when the notification action configuration was last modified.
	ActionLastModifiedDateTime nullable.Type[string] `json:"actionLastModifiedDateTime,omitempty"`

	// The body text of the email notification sent to users.
	EmailText nullable.Type[string] `json:"emailText,omitempty"`

	// The text of the policy tip displayed to the user within the application (For example, Outlook, Word).
	PolicyTip nullable.Type[string] `json:"policyTip,omitempty"`

	// List of email addresses or user identifiers designated to receive the notification email. Can include sender, owner,
	// manager, etc.
	Recipients *[]string `json:"recipients,omitempty"`

	// Fields inherited from DlpActionInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NotifyUserAction) DlpActionInfo() BaseDlpActionInfoImpl {
	return BaseDlpActionInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NotifyUserAction{}

func (s NotifyUserAction) MarshalJSON() ([]byte, error) {
	type wrapper NotifyUserAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NotifyUserAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NotifyUserAction: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.notifyUserAction"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NotifyUserAction: %+v", err)
	}

	return encoded, nil
}
