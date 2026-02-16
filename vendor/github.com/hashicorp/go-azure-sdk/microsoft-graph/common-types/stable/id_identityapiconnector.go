package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityApiConnectorId{}

// IdentityApiConnectorId is a struct representing the Resource ID for a Identity Api Connector
type IdentityApiConnectorId struct {
	IdentityApiConnectorId string
}

// NewIdentityApiConnectorID returns a new IdentityApiConnectorId struct
func NewIdentityApiConnectorID(identityApiConnectorId string) IdentityApiConnectorId {
	return IdentityApiConnectorId{
		IdentityApiConnectorId: identityApiConnectorId,
	}
}

// ParseIdentityApiConnectorID parses 'input' into a IdentityApiConnectorId
func ParseIdentityApiConnectorID(input string) (*IdentityApiConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityApiConnectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityApiConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityApiConnectorIDInsensitively parses 'input' case-insensitively into a IdentityApiConnectorId
// note: this method should only be used for API response data and not user input
func ParseIdentityApiConnectorIDInsensitively(input string) (*IdentityApiConnectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityApiConnectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityApiConnectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityApiConnectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IdentityApiConnectorId, ok = input.Parsed["identityApiConnectorId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityApiConnectorId", input)
	}

	return nil
}

// ValidateIdentityApiConnectorID checks that 'input' can be parsed as a Identity Api Connector ID
func ValidateIdentityApiConnectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityApiConnectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Api Connector ID
func (id IdentityApiConnectorId) ID() string {
	fmtString := "/identity/apiConnectors/%s"
	return fmt.Sprintf(fmtString, id.IdentityApiConnectorId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Api Connector ID
func (id IdentityApiConnectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("apiConnectors", "apiConnectors", "apiConnectors"),
		resourceids.UserSpecifiedSegment("identityApiConnectorId", "identityApiConnectorId"),
	}
}

// String returns a human-readable description of this Identity Api Connector ID
func (id IdentityApiConnectorId) String() string {
	components := []string{
		fmt.Sprintf("Identity Api Connector: %q", id.IdentityApiConnectorId),
	}
	return fmt.Sprintf("Identity Api Connector (%s)", strings.Join(components, "\n"))
}
