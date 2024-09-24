package me

// Copyright (c) HashiCorp Inc. All rights reserved.
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
