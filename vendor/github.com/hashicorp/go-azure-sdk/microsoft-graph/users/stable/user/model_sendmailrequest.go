package user

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendMailRequest struct {
	Message         *stable.Message     `json:"Message,omitempty"`
	SaveToSentItems nullable.Type[bool] `json:"SaveToSentItems,omitempty"`
}
