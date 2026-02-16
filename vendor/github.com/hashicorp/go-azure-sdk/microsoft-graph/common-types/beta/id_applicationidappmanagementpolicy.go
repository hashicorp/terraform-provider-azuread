package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdAppManagementPolicyId{}

// ApplicationIdAppManagementPolicyId is a struct representing the Resource ID for a Application Id App Management Policy
type ApplicationIdAppManagementPolicyId struct {
	ApplicationId         string
	AppManagementPolicyId string
}

// NewApplicationIdAppManagementPolicyID returns a new ApplicationIdAppManagementPolicyId struct
func NewApplicationIdAppManagementPolicyID(applicationId string, appManagementPolicyId string) ApplicationIdAppManagementPolicyId {
	return ApplicationIdAppManagementPolicyId{
		ApplicationId:         applicationId,
		AppManagementPolicyId: appManagementPolicyId,
	}
}

// ParseApplicationIdAppManagementPolicyID parses 'input' into a ApplicationIdAppManagementPolicyId
func ParseApplicationIdAppManagementPolicyID(input string) (*ApplicationIdAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a ApplicationIdAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdAppManagementPolicyIDInsensitively(input string) (*ApplicationIdAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdAppManagementPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdAppManagementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.AppManagementPolicyId, ok = input.Parsed["appManagementPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", input)
	}

	return nil
}

// ValidateApplicationIdAppManagementPolicyID checks that 'input' can be parsed as a Application Id App Management Policy ID
func ValidateApplicationIdAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id App Management Policy ID
func (id ApplicationIdAppManagementPolicyId) ID() string {
	fmtString := "/applications/%s/appManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.AppManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id App Management Policy ID
func (id ApplicationIdAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("appManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyId"),
	}
}

// String returns a human-readable description of this Application Id App Management Policy ID
func (id ApplicationIdAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
	}
	return fmt.Sprintf("Application Id App Management Policy (%s)", strings.Join(components, "\n"))
}
