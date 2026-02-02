package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdFederationConfigurationId{}

// DomainIdFederationConfigurationId is a struct representing the Resource ID for a Domain Id Federation Configuration
type DomainIdFederationConfigurationId struct {
	DomainId                   string
	InternalDomainFederationId string
}

// NewDomainIdFederationConfigurationID returns a new DomainIdFederationConfigurationId struct
func NewDomainIdFederationConfigurationID(domainId string, internalDomainFederationId string) DomainIdFederationConfigurationId {
	return DomainIdFederationConfigurationId{
		DomainId:                   domainId,
		InternalDomainFederationId: internalDomainFederationId,
	}
}

// ParseDomainIdFederationConfigurationID parses 'input' into a DomainIdFederationConfigurationId
func ParseDomainIdFederationConfigurationID(input string) (*DomainIdFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdFederationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdFederationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDomainIdFederationConfigurationIDInsensitively parses 'input' case-insensitively into a DomainIdFederationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDomainIdFederationConfigurationIDInsensitively(input string) (*DomainIdFederationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdFederationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdFederationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DomainIdFederationConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DomainId, ok = input.Parsed["domainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainId", input)
	}

	if id.InternalDomainFederationId, ok = input.Parsed["internalDomainFederationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "internalDomainFederationId", input)
	}

	return nil
}

// ValidateDomainIdFederationConfigurationID checks that 'input' can be parsed as a Domain Id Federation Configuration ID
func ValidateDomainIdFederationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainIdFederationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Id Federation Configuration ID
func (id DomainIdFederationConfigurationId) ID() string {
	fmtString := "/domains/%s/federationConfiguration/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.InternalDomainFederationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Id Federation Configuration ID
func (id DomainIdFederationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("domains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainId"),
		resourceids.StaticSegment("federationConfiguration", "federationConfiguration", "federationConfiguration"),
		resourceids.UserSpecifiedSegment("internalDomainFederationId", "internalDomainFederationId"),
	}
}

// String returns a human-readable description of this Domain Id Federation Configuration ID
func (id DomainIdFederationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Internal Domain Federation: %q", id.InternalDomainFederationId),
	}
	return fmt.Sprintf("Domain Id Federation Configuration (%s)", strings.Join(components, "\n"))
}
