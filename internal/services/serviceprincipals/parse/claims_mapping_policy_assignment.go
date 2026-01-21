// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type claimsMappingPolicyAssignmentId struct {
	ObjectSubResourceId
	ServicePrincipalId    string
	ClaimsMappingPolicyId string
}

func NewClaimsMappingPolicyAssignmentID(servicePolicyId, claimsMappingPolicyId string) claimsMappingPolicyAssignmentId {
	return claimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   NewObjectSubResourceID(servicePolicyId, "claimsMappingPolicy", claimsMappingPolicyId),
		ServicePrincipalId:    servicePolicyId,
		ClaimsMappingPolicyId: claimsMappingPolicyId,
	}
}

func ClaimsMappingPolicyAssignmentID(idString string) (*claimsMappingPolicyAssignmentId, error) {
	id, err := ObjectSubResourceID(idString, "claimsMappingPolicy")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Claims Mapping Policy Assignment ID: %v", err)
	}

	return &claimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   *id,
		ServicePrincipalId:    id.objectId,
		ClaimsMappingPolicyId: id.subId,
	}, nil
}
