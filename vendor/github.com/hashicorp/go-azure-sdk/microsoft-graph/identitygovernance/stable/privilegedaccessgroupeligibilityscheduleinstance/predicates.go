package privilegedaccessgroupeligibilityscheduleinstance

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate struct {
}

func (p PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate) Matches(input stable.PrivilegedAccessGroupEligibilityScheduleInstance) bool {

	return true
}
