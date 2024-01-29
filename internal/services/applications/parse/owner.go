package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type OwnerId struct {
	ApplicationId string
	OwnerId       string
}

func NewOwnerID(applicationId, ownerId string) *OwnerId {
	return &OwnerId{
		ApplicationId: applicationId,
		OwnerId:       ownerId,
	}
}

// ParseOwnerID parses 'input' into an OwnerId
func ParseOwnerID(input string) (*OwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &OwnerId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.OwnerId, ok = parsed.Parsed["ownerId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "ownerId", *parsed)
	}

	return id, nil
}

// ValidateOwnerID checks that 'input' can be parsed as an Application ID
func ValidateOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseOwnerID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.OwnerId, "ID")
}

func (id *OwnerId) ID() string {
	fmtString := "/applications/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.OwnerId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *OwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("owners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("ownerId", "11111111-1111-1111-1111-111111111111"),
	}
}

func (id *OwnerId) String() string {
	return fmt.Sprintf("Application Owner (Application ID: %q, Owner ID: %q)", id.ApplicationId, id.OwnerId)
}

func (id *OwnerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	if id.OwnerId, ok = input.Parsed["ownerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ownerId", input)
	}

	return nil
}
