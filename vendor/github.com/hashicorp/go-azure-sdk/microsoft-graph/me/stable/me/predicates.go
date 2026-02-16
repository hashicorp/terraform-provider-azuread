package me

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ConvertIdResultOperationPredicate struct {
}

func (p ConvertIdResultOperationPredicate) Matches(input stable.ConvertIdResult) bool {

	return true
}

type MailTipsOperationPredicate struct {
}

func (p MailTipsOperationPredicate) Matches(input stable.MailTips) bool {

	return true
}
