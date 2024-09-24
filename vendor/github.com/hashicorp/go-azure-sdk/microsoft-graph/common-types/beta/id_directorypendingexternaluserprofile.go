package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPendingExternalUserProfileId{}

// DirectoryPendingExternalUserProfileId is a struct representing the Resource ID for a Directory Pending External User Profile
type DirectoryPendingExternalUserProfileId struct {
	PendingExternalUserProfileId string
}

// NewDirectoryPendingExternalUserProfileID returns a new DirectoryPendingExternalUserProfileId struct
func NewDirectoryPendingExternalUserProfileID(pendingExternalUserProfileId string) DirectoryPendingExternalUserProfileId {
	return DirectoryPendingExternalUserProfileId{
		PendingExternalUserProfileId: pendingExternalUserProfileId,
	}
}

// ParseDirectoryPendingExternalUserProfileID parses 'input' into a DirectoryPendingExternalUserProfileId
func ParseDirectoryPendingExternalUserProfileID(input string) (*DirectoryPendingExternalUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPendingExternalUserProfileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPendingExternalUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryPendingExternalUserProfileIDInsensitively parses 'input' case-insensitively into a DirectoryPendingExternalUserProfileId
// note: this method should only be used for API response data and not user input
func ParseDirectoryPendingExternalUserProfileIDInsensitively(input string) (*DirectoryPendingExternalUserProfileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPendingExternalUserProfileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPendingExternalUserProfileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryPendingExternalUserProfileId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PendingExternalUserProfileId, ok = input.Parsed["pendingExternalUserProfileId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pendingExternalUserProfileId", input)
	}

	return nil
}

// ValidateDirectoryPendingExternalUserProfileID checks that 'input' can be parsed as a Directory Pending External User Profile ID
func ValidateDirectoryPendingExternalUserProfileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryPendingExternalUserProfileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Pending External User Profile ID
func (id DirectoryPendingExternalUserProfileId) ID() string {
	fmtString := "/directory/pendingExternalUserProfiles/%s"
	return fmt.Sprintf(fmtString, id.PendingExternalUserProfileId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Pending External User Profile ID
func (id DirectoryPendingExternalUserProfileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("pendingExternalUserProfiles", "pendingExternalUserProfiles", "pendingExternalUserProfiles"),
		resourceids.UserSpecifiedSegment("pendingExternalUserProfileId", "pendingExternalUserProfileId"),
	}
}

// String returns a human-readable description of this Directory Pending External User Profile ID
func (id DirectoryPendingExternalUserProfileId) String() string {
	components := []string{
		fmt.Sprintf("Pending External User Profile: %q", id.PendingExternalUserProfileId),
	}
	return fmt.Sprintf("Directory Pending External User Profile (%s)", strings.Join(components, "\n"))
}
