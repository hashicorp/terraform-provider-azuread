package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ConvertIdResultOperationPredicate struct {
}

func (p ConvertIdResultOperationPredicate) Matches(input stable.ConvertIdResult) bool {

	return true
}

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input stable.DirectoryObject) bool {

	return true
}

type ExtensionPropertyOperationPredicate struct {
}

func (p ExtensionPropertyOperationPredicate) Matches(input stable.ExtensionProperty) bool {

	return true
}

type MailTipsOperationPredicate struct {
}

func (p MailTipsOperationPredicate) Matches(input stable.MailTips) bool {

	return true
}

type UserOperationPredicate struct {
}

func (p UserOperationPredicate) Matches(input stable.User) bool {

	return true
}
