package entitlementmanagementassignmentpolicy

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AccessPackageAssignmentPolicyOperationPredicate struct {
}

func (p AccessPackageAssignmentPolicyOperationPredicate) Matches(input stable.AccessPackageAssignmentPolicy) bool {

	return true
}
