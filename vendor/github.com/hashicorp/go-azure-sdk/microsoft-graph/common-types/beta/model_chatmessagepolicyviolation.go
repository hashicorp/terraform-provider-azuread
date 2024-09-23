package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolation struct {
	// The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender --
	// Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading
	// the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing
	// users within the organization to read the message.
	DlpAction *ChatMessagePolicyViolationDlpActionTypes `json:"dlpAction,omitempty"`

	// Justification text provided by the sender of the message when overriding a policy violation.
	JustificationText nullable.Type[string] `json:"justificationText,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information to display to the message sender about why the message was flagged as a violation.
	PolicyTip *ChatMessagePolicyViolationPolicyTip `json:"policyTip,omitempty"`

	// Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are:
	// NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content,
	// userAction isn't required.
	UserAction *ChatMessagePolicyViolationUserActionTypes `json:"userAction,omitempty"`

	// Indicates what actions the sender may take in response to the policy violation. Supported values are:
	// NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and
	// its rules, and allow readers to see the message again if it was hidden by dlpAction.AllowOverrideWithoutJustification
	// -- Allows the sender to override the DLP violation and allow readers to see the message again if the dlpAction hides
	// it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to
	// override the DLP violation and allow readers to see the message again if the dlpAction hides it, after providing an
	// explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
	VerdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes `json:"verdictDetails,omitempty"`
}
