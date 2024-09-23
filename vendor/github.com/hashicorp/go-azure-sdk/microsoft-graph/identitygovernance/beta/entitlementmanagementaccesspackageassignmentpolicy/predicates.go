package entitlementmanagementaccesspackageassignmentpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type AccessPackageAssignmentPolicyOperationPredicate struct {
}

func (p AccessPackageAssignmentPolicyOperationPredicate) Matches(input beta.AccessPackageAssignmentPolicy) bool {

	return true
}
