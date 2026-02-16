package me

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TranslateExchangeIdsRequest struct {
	InputIds     *[]string                `json:"InputIds,omitempty"`
	SourceIdType *stable.ExchangeIdFormat `json:"SourceIdType,omitempty"`
	TargetIdType *stable.ExchangeIdFormat `json:"TargetIdType,omitempty"`
}
