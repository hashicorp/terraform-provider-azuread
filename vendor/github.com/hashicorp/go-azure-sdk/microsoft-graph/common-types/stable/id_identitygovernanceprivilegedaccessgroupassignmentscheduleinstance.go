package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{}

// IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId is a struct representing the Resource ID for a Identity Governance Privileged Access Group Assignment Schedule Instance
type IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId struct {
	PrivilegedAccessGroupAssignmentScheduleInstanceId string
}

// NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID returns a new IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId struct
func NewIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(privilegedAccessGroupAssignmentScheduleInstanceId string) IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId {
	return IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{
		PrivilegedAccessGroupAssignmentScheduleInstanceId: privilegedAccessGroupAssignmentScheduleInstanceId,
	}
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID parses 'input' into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceIDInsensitively(input string) (*IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PrivilegedAccessGroupAssignmentScheduleInstanceId, ok = input.Parsed["privilegedAccessGroupAssignmentScheduleInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privilegedAccessGroupAssignmentScheduleInstanceId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID checks that 'input' can be parsed as a Identity Governance Privileged Access Group Assignment Schedule Instance ID
func ValidateIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Access Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.PrivilegedAccessGroupAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Access Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentScheduleInstances", "assignmentScheduleInstances", "assignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("privilegedAccessGroupAssignmentScheduleInstanceId", "privilegedAccessGroupAssignmentScheduleInstanceId"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Access Group Assignment Schedule Instance ID
func (id IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Privileged Access Group Assignment Schedule Instance: %q", id.PrivilegedAccessGroupAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Identity Governance Privileged Access Group Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
