// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import "fmt"

type ClaimsMappingPolicyAssignmentId struct {
	ObjectSubResourceId
	ServicePrincipalId    string
	ClaimsMappingPolicyId string
}

func NewClaimsMappingPolicyAssignmentID(ServicePolicyId, ClaimsMappingPolicyId string) ClaimsMappingPolicyAssignmentId {
	return ClaimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   NewObjectSubResourceID(ServicePolicyId, "claimsMappingPolicy", ClaimsMappingPolicyId),
		ServicePrincipalId:    ServicePolicyId,
		ClaimsMappingPolicyId: ClaimsMappingPolicyId,
	}
}

func ClaimsMappingPolicyAssignmentID(idString string) (*ClaimsMappingPolicyAssignmentId, error) {
	id, err := ObjectSubResourceID(idString, "claimsMappingPolicy")
	if err != nil {
		return nil, fmt.Errorf("unable to parse Claims Mapping Policy Assignment ID: %v", err)
	}

	return &ClaimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   *id,
		ServicePrincipalId:    id.objectId,
		ClaimsMappingPolicyId: id.subId,
	}, nil
}
