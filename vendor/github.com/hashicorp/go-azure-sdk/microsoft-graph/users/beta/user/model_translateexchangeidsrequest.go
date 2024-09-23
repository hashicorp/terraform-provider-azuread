package user

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslateExchangeIdsRequest struct {
	InputIds     *[]string              `json:"InputIds,omitempty"`
	SourceIdType *beta.ExchangeIdFormat `json:"SourceIdType,omitempty"`
	TargetIdType *beta.ExchangeIdFormat `json:"TargetIdType,omitempty"`
}
