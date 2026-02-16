package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{}

// DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId is a struct representing the Resource ID for a Directory Public Key Infrastructure Certificate Based Auth Configuration
type DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId struct {
	CertificateBasedAuthPkiId string
}

// NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID returns a new DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId struct
func NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(certificateBasedAuthPkiId string) DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId {
	return DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{
		CertificateBasedAuthPkiId: certificateBasedAuthPkiId,
	}
}

// ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID parses 'input' into a DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId
func ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(input string) (*DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIDInsensitively parses 'input' case-insensitively into a DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIDInsensitively(input string) (*DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CertificateBasedAuthPkiId, ok = input.Parsed["certificateBasedAuthPkiId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedAuthPkiId", input)
	}

	return nil
}

// ValidateDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID checks that 'input' can be parsed as a Directory Public Key Infrastructure Certificate Based Auth Configuration ID
func ValidateDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Public Key Infrastructure Certificate Based Auth Configuration ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId) ID() string {
	fmtString := "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/%s"
	return fmt.Sprintf(fmtString, id.CertificateBasedAuthPkiId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Public Key Infrastructure Certificate Based Auth Configuration ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("publicKeyInfrastructure", "publicKeyInfrastructure", "publicKeyInfrastructure"),
		resourceids.StaticSegment("certificateBasedAuthConfigurations", "certificateBasedAuthConfigurations", "certificateBasedAuthConfigurations"),
		resourceids.UserSpecifiedSegment("certificateBasedAuthPkiId", "certificateBasedAuthPkiId"),
	}
}

// String returns a human-readable description of this Directory Public Key Infrastructure Certificate Based Auth Configuration ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Based Auth Pki: %q", id.CertificateBasedAuthPkiId),
	}
	return fmt.Sprintf("Directory Public Key Infrastructure Certificate Based Auth Configuration (%s)", strings.Join(components, "\n"))
}
