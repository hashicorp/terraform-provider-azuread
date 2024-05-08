package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type FallbackPublicClientId struct {
	ApplicationId string
}

func NewFallbackPublicClientID(applicationId string) *FallbackPublicClientId {
	return &FallbackPublicClientId{
		ApplicationId: applicationId,
	}
}

// ParseFallbackPublicClientID parses 'input' into an FallbackPublicClientId
func ParseFallbackPublicClientID(input string) (*FallbackPublicClientId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FallbackPublicClientId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &FallbackPublicClientId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	return id, nil
}

// ValidateFallbackPublicClientID checks that 'input' can be parsed as an Application ID
func ValidateFallbackPublicClientID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseFallbackPublicClientID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.ApplicationId, "ID")
}

func (id *FallbackPublicClientId) ID() string {
	fmtString := "/applications/%s/fallbackPublicClient"
	return fmt.Sprintf(fmtString, id.ApplicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *FallbackPublicClientId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("fallbackPublicClient", "fallbackPublicClient", "fallbackPublicClient"),
	}
}

func (id *FallbackPublicClientId) String() string {
	return fmt.Sprintf("Fallback Public Client (Application ID: %q)", id.ApplicationId)
}

func (id *FallbackPublicClientId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	return nil
}
