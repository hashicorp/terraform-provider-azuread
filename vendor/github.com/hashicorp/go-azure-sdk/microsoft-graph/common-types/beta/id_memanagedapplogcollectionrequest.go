package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeManagedAppLogCollectionRequestId{}

// MeManagedAppLogCollectionRequestId is a struct representing the Resource ID for a Me Managed App Log Collection Request
type MeManagedAppLogCollectionRequestId struct {
	ManagedAppLogCollectionRequestId string
}

// NewMeManagedAppLogCollectionRequestID returns a new MeManagedAppLogCollectionRequestId struct
func NewMeManagedAppLogCollectionRequestID(managedAppLogCollectionRequestId string) MeManagedAppLogCollectionRequestId {
	return MeManagedAppLogCollectionRequestId{
		ManagedAppLogCollectionRequestId: managedAppLogCollectionRequestId,
	}
}

// ParseMeManagedAppLogCollectionRequestID parses 'input' into a MeManagedAppLogCollectionRequestId
func ParseMeManagedAppLogCollectionRequestID(input string) (*MeManagedAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeManagedAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a MeManagedAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseMeManagedAppLogCollectionRequestIDInsensitively(input string) (*MeManagedAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeManagedAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeManagedAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeManagedAppLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedAppLogCollectionRequestId, ok = input.Parsed["managedAppLogCollectionRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedAppLogCollectionRequestId", input)
	}

	return nil
}

// ValidateMeManagedAppLogCollectionRequestID checks that 'input' can be parsed as a Me Managed App Log Collection Request ID
func ValidateMeManagedAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed App Log Collection Request ID
func (id MeManagedAppLogCollectionRequestId) ID() string {
	fmtString := "/me/managedAppLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.ManagedAppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed App Log Collection Request ID
func (id MeManagedAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("managedAppLogCollectionRequests", "managedAppLogCollectionRequests", "managedAppLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("managedAppLogCollectionRequestId", "managedAppLogCollectionRequestId"),
	}
}

// String returns a human-readable description of this Me Managed App Log Collection Request ID
func (id MeManagedAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("Managed App Log Collection Request: %q", id.ManagedAppLogCollectionRequestId),
	}
	return fmt.Sprintf("Me Managed App Log Collection Request (%s)", strings.Join(components, "\n"))
}
