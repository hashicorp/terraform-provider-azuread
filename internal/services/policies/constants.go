// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package policies

const (
	RoleDefinitionIdMember = "member"
	RoleDefinitionIdOwner  = "owner"
)

const authenticationStrengthPolicyResourceName = "azuread_authentication_strength_policy"

var possibleValuesForRoleDefinitionId = []string{RoleDefinitionIdMember, RoleDefinitionIdOwner}
