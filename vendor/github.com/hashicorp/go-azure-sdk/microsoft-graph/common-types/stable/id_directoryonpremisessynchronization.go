package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryOnPremisesSynchronizationId{}

// DirectoryOnPremisesSynchronizationId is a struct representing the Resource ID for a Directory On Premises Synchronization
type DirectoryOnPremisesSynchronizationId struct {
	OnPremisesDirectorySynchronizationId string
}

// NewDirectoryOnPremisesSynchronizationID returns a new DirectoryOnPremisesSynchronizationId struct
func NewDirectoryOnPremisesSynchronizationID(onPremisesDirectorySynchronizationId string) DirectoryOnPremisesSynchronizationId {
	return DirectoryOnPremisesSynchronizationId{
		OnPremisesDirectorySynchronizationId: onPremisesDirectorySynchronizationId,
	}
}

// ParseDirectoryOnPremisesSynchronizationID parses 'input' into a DirectoryOnPremisesSynchronizationId
func ParseDirectoryOnPremisesSynchronizationID(input string) (*DirectoryOnPremisesSynchronizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOnPremisesSynchronizationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOnPremisesSynchronizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryOnPremisesSynchronizationIDInsensitively parses 'input' case-insensitively into a DirectoryOnPremisesSynchronizationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryOnPremisesSynchronizationIDInsensitively(input string) (*DirectoryOnPremisesSynchronizationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryOnPremisesSynchronizationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryOnPremisesSynchronizationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryOnPremisesSynchronizationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnPremisesDirectorySynchronizationId, ok = input.Parsed["onPremisesDirectorySynchronizationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onPremisesDirectorySynchronizationId", input)
	}

	return nil
}

// ValidateDirectoryOnPremisesSynchronizationID checks that 'input' can be parsed as a Directory On Premises Synchronization ID
func ValidateDirectoryOnPremisesSynchronizationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryOnPremisesSynchronizationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory On Premises Synchronization ID
func (id DirectoryOnPremisesSynchronizationId) ID() string {
	fmtString := "/directory/onPremisesSynchronization/%s"
	return fmt.Sprintf(fmtString, id.OnPremisesDirectorySynchronizationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory On Premises Synchronization ID
func (id DirectoryOnPremisesSynchronizationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("onPremisesSynchronization", "onPremisesSynchronization", "onPremisesSynchronization"),
		resourceids.UserSpecifiedSegment("onPremisesDirectorySynchronizationId", "onPremisesDirectorySynchronizationId"),
	}
}

// String returns a human-readable description of this Directory On Premises Synchronization ID
func (id DirectoryOnPremisesSynchronizationId) String() string {
	components := []string{
		fmt.Sprintf("On Premises Directory Synchronization: %q", id.OnPremisesDirectorySynchronizationId),
	}
	return fmt.Sprintf("Directory On Premises Synchronization (%s)", strings.Join(components, "\n"))
}
