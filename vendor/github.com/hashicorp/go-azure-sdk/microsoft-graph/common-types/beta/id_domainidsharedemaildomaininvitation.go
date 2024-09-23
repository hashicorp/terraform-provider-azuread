package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdSharedEmailDomainInvitationId{}

// DomainIdSharedEmailDomainInvitationId is a struct representing the Resource ID for a Domain Id Shared Email Domain Invitation
type DomainIdSharedEmailDomainInvitationId struct {
	DomainId                      string
	SharedEmailDomainInvitationId string
}

// NewDomainIdSharedEmailDomainInvitationID returns a new DomainIdSharedEmailDomainInvitationId struct
func NewDomainIdSharedEmailDomainInvitationID(domainId string, sharedEmailDomainInvitationId string) DomainIdSharedEmailDomainInvitationId {
	return DomainIdSharedEmailDomainInvitationId{
		DomainId:                      domainId,
		SharedEmailDomainInvitationId: sharedEmailDomainInvitationId,
	}
}

// ParseDomainIdSharedEmailDomainInvitationID parses 'input' into a DomainIdSharedEmailDomainInvitationId
func ParseDomainIdSharedEmailDomainInvitationID(input string) (*DomainIdSharedEmailDomainInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdSharedEmailDomainInvitationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdSharedEmailDomainInvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDomainIdSharedEmailDomainInvitationIDInsensitively parses 'input' case-insensitively into a DomainIdSharedEmailDomainInvitationId
// note: this method should only be used for API response data and not user input
func ParseDomainIdSharedEmailDomainInvitationIDInsensitively(input string) (*DomainIdSharedEmailDomainInvitationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdSharedEmailDomainInvitationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdSharedEmailDomainInvitationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DomainIdSharedEmailDomainInvitationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DomainId, ok = input.Parsed["domainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainId", input)
	}

	if id.SharedEmailDomainInvitationId, ok = input.Parsed["sharedEmailDomainInvitationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedEmailDomainInvitationId", input)
	}

	return nil
}

// ValidateDomainIdSharedEmailDomainInvitationID checks that 'input' can be parsed as a Domain Id Shared Email Domain Invitation ID
func ValidateDomainIdSharedEmailDomainInvitationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainIdSharedEmailDomainInvitationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Id Shared Email Domain Invitation ID
func (id DomainIdSharedEmailDomainInvitationId) ID() string {
	fmtString := "/domains/%s/sharedEmailDomainInvitations/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.SharedEmailDomainInvitationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Id Shared Email Domain Invitation ID
func (id DomainIdSharedEmailDomainInvitationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("domains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainId"),
		resourceids.StaticSegment("sharedEmailDomainInvitations", "sharedEmailDomainInvitations", "sharedEmailDomainInvitations"),
		resourceids.UserSpecifiedSegment("sharedEmailDomainInvitationId", "sharedEmailDomainInvitationId"),
	}
}

// String returns a human-readable description of this Domain Id Shared Email Domain Invitation ID
func (id DomainIdSharedEmailDomainInvitationId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Shared Email Domain Invitation: %q", id.SharedEmailDomainInvitationId),
	}
	return fmt.Sprintf("Domain Id Shared Email Domain Invitation (%s)", strings.Join(components, "\n"))
}
