package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Schedule
type IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId struct {
	PrivilegedAccessGroupAssignmentScheduleId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID(privilegedAccessGroupAssignmentScheduleId string) IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{
		PrivilegedAccessGroupAssignmentScheduleId: privilegedAccessGroupAssignmentScheduleId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Schedule ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentSchedules", "assignmentSchedules", "assignmentSchedules"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleId", "privilegedAccessGroupAssignmentScheduleId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Schedule ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule: %q", id.PrivilegedAccessGroupAssignmentScheduleId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Schedule (%s)", strings.Join(components, "\n"))
}
