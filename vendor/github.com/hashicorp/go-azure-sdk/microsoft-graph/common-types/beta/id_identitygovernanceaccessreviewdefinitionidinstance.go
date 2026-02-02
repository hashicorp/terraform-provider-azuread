package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance
type IdentityGovernanceAccessReviewDefinitionIdInstanceId struct {
	AccessReviewScheduleDefinitionId string
	AccessReviewInstanceId           string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string) IdentityGovernanceAccessReviewDefinitionIdInstanceId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceId{
		AccessReviewScheduleDefinitionId: accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:           accessReviewInstanceId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance (%s)", strings.Join(components, "\n"))
}
