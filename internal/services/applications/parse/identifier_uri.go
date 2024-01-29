package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type IdentifierUriId struct {
	ApplicationId string
	IdentifierUri string
}

func NewIdentifierUriID(applicationId, identifierUri string) *IdentifierUriId {
	return &IdentifierUriId{
		ApplicationId: applicationId,
		IdentifierUri: identifierUri,
	}
}

// ParseIdentifierUriID parses 'input' into an IdentifierUriId
func ParseIdentifierUriID(input string) (*IdentifierUriId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentifierUriId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &IdentifierUriId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.IdentifierUri, ok = parsed.Parsed["identifierUri"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identifierUri", *parsed)
	}

	return id, nil
}

// ValidateIdentifierUriID checks that 'input' can be parsed as an Application ID
func ValidateIdentifierUriID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentifierUriID(v); err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func (id *IdentifierUriId) ID() string {
	fmtString := "/applications/%s/identifierUris/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.IdentifierUri)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *IdentifierUriId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("identifierUris", "identifierUris", "identifierUris"),
		resourceids.UserSpecifiedSegment("identifierUri", "aHR0cHM6Ly9leGFtcGxlLm5ldC8="),
	}
}

func (id *IdentifierUriId) String() string {
	return fmt.Sprintf("Application IdentifierUri (Application ID: %q, IdentifierUri ID: %q)", id.ApplicationId, id.IdentifierUri)
}

func (id *IdentifierUriId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.IdentifierUri, ok = input.Parsed["identifierUri"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identifierUri", input)
	}

	return nil
}
