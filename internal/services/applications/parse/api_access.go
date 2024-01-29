package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type ApiAccessId struct {
	ApplicationId string
	ApiClientId   string
}

func NewApiAccessID(applicationId, apiClientId string) *ApiAccessId {
	return &ApiAccessId{
		ApplicationId: applicationId,
		ApiClientId:   apiClientId,
	}
}

// ParseApiAccessID parses 'input' into an ApiAccessId
func ParseApiAccessID(input string) (*ApiAccessId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ApiAccessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &ApiAccessId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.ApiClientId, ok = parsed.Parsed["apiClientId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiClientId", *parsed)
	}

	return id, nil
}

// ValidateApiAccessID checks that 'input' can be parsed as an Application ID
func ValidateApiAccessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseApiAccessID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.ApiClientId, "ID")
}

func (id *ApiAccessId) ID() string {
	fmtString := "/applications/%s/apiAccess/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.ApiClientId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *ApiAccessId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("apiAccess", "apiAccess", "apiAccess"),
		resourceids.UserSpecifiedSegment("apiClientId", "11111111-1111-1111-1111-111111111111"),
	}
}

func (id *ApiAccessId) String() string {
	return fmt.Sprintf("Application API Access (Application ID: %q, API Client ID: %q)", id.ApplicationId, id.ApiClientId)
}

func (id *ApiAccessId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.ApiClientId, ok = input.Parsed["roleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "apiClientId", input)
	}

	return nil
}
