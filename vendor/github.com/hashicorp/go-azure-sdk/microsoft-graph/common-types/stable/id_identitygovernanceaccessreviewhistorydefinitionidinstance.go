package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{}

// IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId is a struct representing the Resource ID for a Identity Governance Access Review History Definition Id Instance
type IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId struct {
	AccessReviewHistoryDefinitionId string
	AccessReviewHistoryInstanceId   string
}

// NewIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID returns a new IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId struct
func NewIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(accessReviewHistoryDefinitionId string, accessReviewHistoryInstanceId string) IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId {
	return IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{
		AccessReviewHistoryDefinitionId: accessReviewHistoryDefinitionId,
		AccessReviewHistoryInstanceId:   accessReviewHistoryInstanceId,
	}
}

// ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID parses 'input' into a IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId
func ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(input string) (*IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceIDInsensitively(input string) (*IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewHistoryDefinitionId, ok = input.Parsed["accessReviewHistoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewHistoryDefinitionId", input)
	}

	if id.AccessReviewHistoryInstanceId, ok = input.Parsed["accessReviewHistoryInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewHistoryInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID checks that 'input' can be parsed as a Identity Governance Access Review History Definition Id Instance ID
func ValidateIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review History Definition Id Instance ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId) ID() string {
	fmtString := "/identityGovernance/accessReviews/historyDefinitions/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewHistoryDefinitionId, id.AccessReviewHistoryInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review History Definition Id Instance ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("historyDefinitions", "historyDefinitions", "historyDefinitions"),
		resourceids.UserSpecifiedSegment("accessReviewHistoryDefinitionId", "accessReviewHistoryDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewHistoryInstanceId", "accessReviewHistoryInstanceId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review History Definition Id Instance ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Access Review History Definition: %q", id.AccessReviewHistoryDefinitionId),
		fmt.Sprintf("Access Review History Instance: %q", id.AccessReviewHistoryInstanceId),
	}
	return fmt.Sprintf("Identity Governance Access Review History Definition Id Instance (%s)", strings.Join(components, "\n"))
}
