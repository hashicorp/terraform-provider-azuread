package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{}

// DirectoryCertificateAuthorityMutualTlsOauthConfigurationId is a struct representing the Resource ID for a Directory Certificate Authority Mutual Tls Oauth Configuration
type DirectoryCertificateAuthorityMutualTlsOauthConfigurationId struct {
	MutualTlsOauthConfigurationId string
}

// NewDirectoryCertificateAuthorityMutualTlsOauthConfigurationID returns a new DirectoryCertificateAuthorityMutualTlsOauthConfigurationId struct
func NewDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(mutualTlsOauthConfigurationId string) DirectoryCertificateAuthorityMutualTlsOauthConfigurationId {
	return DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{
		MutualTlsOauthConfigurationId: mutualTlsOauthConfigurationId,
	}
}

// ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationID parses 'input' into a DirectoryCertificateAuthorityMutualTlsOauthConfigurationId
func ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(input string) (*DirectoryCertificateAuthorityMutualTlsOauthConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationIDInsensitively parses 'input' case-insensitively into a DirectoryCertificateAuthorityMutualTlsOauthConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationIDInsensitively(input string) (*DirectoryCertificateAuthorityMutualTlsOauthConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityMutualTlsOauthConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryCertificateAuthorityMutualTlsOauthConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MutualTlsOauthConfigurationId, ok = input.Parsed["mutualTlsOauthConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mutualTlsOauthConfigurationId", input)
	}

	return nil
}

// ValidateDirectoryCertificateAuthorityMutualTlsOauthConfigurationID checks that 'input' can be parsed as a Directory Certificate Authority Mutual Tls Oauth Configuration ID
func ValidateDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCertificateAuthorityMutualTlsOauthConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Certificate Authority Mutual Tls Oauth Configuration ID
func (id DirectoryCertificateAuthorityMutualTlsOauthConfigurationId) ID() string {
	fmtString := "/directory/certificateAuthorities/mutualTlsOauthConfigurations/%s"
	return fmt.Sprintf(fmtString, id.MutualTlsOauthConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Certificate Authority Mutual Tls Oauth Configuration ID
func (id DirectoryCertificateAuthorityMutualTlsOauthConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("certificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.StaticSegment("mutualTlsOauthConfigurations", "mutualTlsOauthConfigurations", "mutualTlsOauthConfigurations"),
		resourceids.UserSpecifiedSegment("mutualTlsOauthConfigurationId", "mutualTlsOauthConfigurationId"),
	}
}

// String returns a human-readable description of this Directory Certificate Authority Mutual Tls Oauth Configuration ID
func (id DirectoryCertificateAuthorityMutualTlsOauthConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Mutual Tls Oauth Configuration: %q", id.MutualTlsOauthConfigurationId),
	}
	return fmt.Sprintf("Directory Certificate Authority Mutual Tls Oauth Configuration (%s)", strings.Join(components, "\n"))
}
