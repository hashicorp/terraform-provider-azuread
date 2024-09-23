package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Stage
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId struct {
	AccessReviewScheduleDefinitionId string
	AccessReviewInstanceId           string
	AccessReviewStageId              string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewStageId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{
		AccessReviewScheduleDefinitionId: accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:           accessReviewInstanceId,
		AccessReviewStageId:              accessReviewStageId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Stage ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Stage (%s)", strings.Join(components, "\n"))
}
