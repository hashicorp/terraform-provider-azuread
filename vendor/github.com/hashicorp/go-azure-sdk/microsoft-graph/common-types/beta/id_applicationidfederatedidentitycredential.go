package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ApplicationIdFederatedIdentityCredentialId{}

// ApplicationIdFederatedIdentityCredentialId is a struct representing the Resource ID for a Application Id Federated Identity Credential
type ApplicationIdFederatedIdentityCredentialId struct {
	ApplicationId                 string
	FederatedIdentityCredentialId string
}

// NewApplicationIdFederatedIdentityCredentialID returns a new ApplicationIdFederatedIdentityCredentialId struct
func NewApplicationIdFederatedIdentityCredentialID(applicationId string, federatedIdentityCredentialId string) ApplicationIdFederatedIdentityCredentialId {
	return ApplicationIdFederatedIdentityCredentialId{
		ApplicationId:                 applicationId,
		FederatedIdentityCredentialId: federatedIdentityCredentialId,
	}
}

// ParseApplicationIdFederatedIdentityCredentialID parses 'input' into a ApplicationIdFederatedIdentityCredentialId
func ParseApplicationIdFederatedIdentityCredentialID(input string) (*ApplicationIdFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdFederatedIdentityCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseApplicationIdFederatedIdentityCredentialIDInsensitively parses 'input' case-insensitively into a ApplicationIdFederatedIdentityCredentialId
// note: this method should only be used for API response data and not user input
func ParseApplicationIdFederatedIdentityCredentialIDInsensitively(input string) (*ApplicationIdFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApplicationIdFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ApplicationIdFederatedIdentityCredentialId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ApplicationIdFederatedIdentityCredentialId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.FederatedIdentityCredentialId, ok = input.Parsed["federatedIdentityCredentialId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", input)
	}

	return nil
}

// ValidateApplicationIdFederatedIdentityCredentialID checks that 'input' can be parsed as a Application Id Federated Identity Credential ID
func ValidateApplicationIdFederatedIdentityCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationIdFederatedIdentityCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Id Federated Identity Credential ID
func (id ApplicationIdFederatedIdentityCredentialId) ID() string {
	fmtString := "/applications/%s/federatedIdentityCredentials/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.FederatedIdentityCredentialId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Id Federated Identity Credential ID
func (id ApplicationIdFederatedIdentityCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationId"),
		resourceids.StaticSegment("federatedIdentityCredentials", "federatedIdentityCredentials", "federatedIdentityCredentials"),
		resourceids.UserSpecifiedSegment("federatedIdentityCredentialId", "federatedIdentityCredentialId"),
	}
}

// String returns a human-readable description of this Application Id Federated Identity Credential ID
func (id ApplicationIdFederatedIdentityCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Federated Identity Credential: %q", id.FederatedIdentityCredentialId),
	}
	return fmt.Sprintf("Application Id Federated Identity Credential (%s)", strings.Join(components, "\n"))
}
