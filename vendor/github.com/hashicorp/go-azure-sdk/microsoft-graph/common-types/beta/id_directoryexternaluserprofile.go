package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryExternalUserProfileId{}

// DirectoryExternalUserProfileId is a struct representing the Resource ID for a Directory External User Profile
type DirectoryExternalUserProfileId struct {
	ExternalUserProfileId string
}

// NewDirectoryExternalUserProfileID returns a new DirectoryExternalUserProfileId struct
func NewDirectoryExternalUserProfileID(externalUserProfileId string) DirectoryExternalUserProfileId {
	return DirectoryExternalUserProfileId{
		ExternalUserProfileId: externalUserProfileId,
	}
}

// ParseDirectoryExternalUserProfileID parses 'input' into a DirectoryExternalUserProfileId
func ParseDirectoryExternalUserProfileID(input string) (*DirectoryExternalUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryExternalUserProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryExternalUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryExternalUserProfileIDInsensitively parses 'input' case-insensitively into a DirectoryExternalUserProfileId
// note: this method should only be used for API response data and not user input
func ParseDirectoryExternalUserProfileIDInsensitively(input string) (*DirectoryExternalUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryExternalUserProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryExternalUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryExternalUserProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ExternalUserProfileId, ok = input.Parsed["externalUserProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "externalUserProfileId", input)
	}

	return nil
}

// ValidateDirectoryExternalUserProfileID checks that 'input' can be parsed as a Directory External User Profile ID
func ValidateDirectoryExternalUserProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryExternalUserProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory External User Profile ID
func (id DirectoryExternalUserProfileId) ID() string {
	fmtString := "/directory/externalUserProfiles/%s"
	return fmt.Sprintf(fmtString, id.ExternalUserProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory External User Profile ID
func (id DirectoryExternalUserProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("externalUserProfiles", "externalUserProfiles", "externalUserProfiles"),
		resourceids.UserSpecifiedSegment("externalUserProfileId", "externalUserProfileId"),
	}
}

// String returns a human-readable description of this Directory External User Profile ID
func (id DirectoryExternalUserProfileId) String() string {
	components := []string{
		fmt.Sprintf("External User Profile: %q", id.ExternalUserProfileId),
	}
	return fmt.Sprintf("Directory External User Profile (%s)", strings.Join(components, "\n"))
}
