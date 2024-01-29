package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

type OptionalClaimsId struct {
	ApplicationId string
}

func NewOptionalClaimsID(applicationId string) *OptionalClaimsId {
	return &OptionalClaimsId{
		ApplicationId: applicationId,
	}
}

// ParseOptionalClaimsID parses 'input' into an OptionalClaimsId
func ParseOptionalClaimsID(input string) (*OptionalClaimsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OptionalClaimsId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &OptionalClaimsId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	return id, nil
}

// ValidateOptionalClaimsID checks that 'input' can be parsed as an Application ID
func ValidateOptionalClaimsID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOptionalClaimsID(v); err != nil {
		errors = append(errors, err)
		return
	}

	return
}

func (id *OptionalClaimsId) ID() string {
	fmtString := "/applications/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *OptionalClaimsId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
	}
}

func (id *OptionalClaimsId) String() string {
	return fmt.Sprintf("Application Optional Claims (Application ID: %q)", id.ApplicationId)
}

func (id *OptionalClaimsId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	return nil
}
