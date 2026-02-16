package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Schedule Request
type IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId struct {
	PrivilegedAccessGroupAssignmentScheduleRequestId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID(privilegedAccessGroupAssignmentScheduleRequestId string) IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{
		PrivilegedAccessGroupAssignmentScheduleRequestId: privilegedAccessGroupAssignmentScheduleRequestId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleRequestId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleRequestId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Schedule Request ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentScheduleRequests", "assignmentScheduleRequests", "assignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleRequestId", "privilegedAccessGroupAssignmentScheduleRequestId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Schedule Request ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule Request: %q", id.PrivilegedAccessGroupAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
