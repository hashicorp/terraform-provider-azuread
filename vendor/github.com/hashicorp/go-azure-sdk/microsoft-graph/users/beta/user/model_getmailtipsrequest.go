package user

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMailTipsRequest struct {
	EmailAddresses  *[]string          `json:"EmailAddresses,omitempty"`
	MailTipsOptions *beta.MailTipsType `json:"MailTipsOptions,omitempty"`
}
