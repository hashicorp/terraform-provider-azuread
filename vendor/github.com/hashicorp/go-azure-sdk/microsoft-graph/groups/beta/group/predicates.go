package group

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DirectoryObjectOperationPredicate struct {
}

func (p DirectoryObjectOperationPredicate) Matches(input beta.DirectoryObject) bool {

	return true
}

type GroupOperationPredicate struct {
}

func (p GroupOperationPredicate) Matches(input beta.Group) bool {

	return true
}

type PasswordSingleSignOnCredentialSetOperationPredicate struct {
}

func (p PasswordSingleSignOnCredentialSetOperationPredicate) Matches(input beta.PasswordSingleSignOnCredentialSet) bool {

	return true
}

type ResourceSpecificPermissionGrantOperationPredicate struct {
}

func (p ResourceSpecificPermissionGrantOperationPredicate) Matches(input beta.ResourceSpecificPermissionGrant) bool {

	return true
}
