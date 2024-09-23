package entitlementmanagementaccesspackage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AccessPackageOperationPredicate struct {
}

func (p AccessPackageOperationPredicate) Matches(input beta.AccessPackage) bool {

	return true
}

type AccessPackageAssignmentRequestRequirementsOperationPredicate struct {
}

func (p AccessPackageAssignmentRequestRequirementsOperationPredicate) Matches(input beta.AccessPackageAssignmentRequestRequirements) bool {

	return true
}
