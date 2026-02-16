package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdOwnerId{}

// ApplicationIdOwnerId is a struct representing the Resource ID for a Application Id Owner
type ApplicationIdOwnerId struct {
	ApplicationId     string
	DirectoryObjectId string
}

// NewApplicationIdOwnerID returns a new ApplicationIdOwnerId struct
func NewApplicationIdOwnerID(applicationId string, directoryObjectId string) ApplicationIdOwnerId {
	return ApplicationIdOwnerId{
		ApplicationId:     applicationId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseApplicationIdOwnerID parses 'input' into a ApplicationIdOwnerId
func ParseApplicationIdOwnerID(input string) (*ApplicationIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdOwnerIDInsensitively parses 'input' case-insensitively into a ApplicationIdOwnerId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdOwnerIDInsensitively(input string) (*ApplicationIdOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdOwnerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdOwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateApplicationIdOwnerID checks that 'input' can be parsed as a Application Id Owner ID
func ValidateApplicationIdOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Owner ID
func (id ApplicationIdOwnerId) ID() string {
	fmtString := "/applications/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Owner ID
func (id ApplicationIdOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Application Id Owner ID
func (id ApplicationIdOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Application Id Owner (%s)", strings.Join(components, "\n"))
}
