package parse

import "fmt"

type ClaimsMappingPolicyAssignmentId struct {
	ObjectSubResourceId
	ServicePolicyId       string
	ClaimsMappingPolicyId string
}

func NewClaimsMappingPolicyAssignmentID(ServicePolicyId, ClaimsMappingPolicyId string) ClaimsMappingPolicyAssignmentId {
	return ClaimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   NewObjectSubResourceID(ServicePolicyId, "azuread_claims_mapping_policy", ClaimsMappingPolicyId),
		ServicePolicyId:       ServicePolicyId,
		ClaimsMappingPolicyId: ClaimsMappingPolicyId,
	}
}

func ClaimsMappingPolicyAssignmentID(idString string) (*ClaimsMappingPolicyAssignmentId, error) {
	id, err := ObjectSubResourceID(idString, "azuread_claims_mapping_policy")
	if err != nil {
		return nil, fmt.Errorf("unable to parse azuread_claims_mapping_policy ID: %v", err)
	}

	return &ClaimsMappingPolicyAssignmentId{
		ObjectSubResourceId:   *id,
		ServicePolicyId:       id.objectId,
		ClaimsMappingPolicyId: id.subId,
	}, nil
}
