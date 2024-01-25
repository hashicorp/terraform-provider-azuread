package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type RedirectUrisId struct {
	ApplicationId string
	UriType       string
}

func NewRedirectUrisID(applicationId, uriType string) *RedirectUrisId {
	return &RedirectUrisId{
		ApplicationId: applicationId,
		UriType:       uriType,
	}
}

// ParseRedirectUrisID parses 'input' into an RedirectUrisId
func ParseRedirectUrisID(input string) (*RedirectUrisId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RedirectUrisId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &RedirectUrisId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.UriType, ok = parsed.Parsed["uriType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "uriType", *parsed)
	}

	return id, nil
}

// ValidateRedirectUrisID checks that 'input' can be parsed as an Application ID
func ValidateRedirectUrisID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	_, err := ParseRedirectUrisID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func (id *RedirectUrisId) ID() string {
	fmtString := "/applications/%s/redirectUris/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.UriType)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *RedirectUrisId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("redirectUris", "redirectUris", "redirectUris"),
		resourceids.UserSpecifiedSegment("uriType", "Web"),
	}
}

func (id *RedirectUrisId) String() string {
	return fmt.Sprintf("Application Redirect URIs (Application ID: %q, URI Type: %q)", id.ApplicationId, id.UriType)
}

func (id *RedirectUrisId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.UriType, ok = input.Parsed["uriType"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "uriType", input)
	}

	return nil
}
