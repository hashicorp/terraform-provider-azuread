package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type KnownClientsId struct {
	ApplicationId string
}

func NewKnownClientsID(applicationId string) *KnownClientsId {
	return &KnownClientsId{
		ApplicationId: applicationId,
	}
}

// ParseKnownClientsID parses 'input' into an KnownClientsId
func ParseKnownClientsID(input string) (*KnownClientsId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KnownClientsId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &KnownClientsId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	return id, nil
}

// ValidateKnownClientsID checks that 'input' can be parsed as an Application ID
func ValidateKnownClientsID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseKnownClientsID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.ApplicationId, "ID")
}

func (id *KnownClientsId) ID() string {
	fmtString := "/applications/%s/knownClients"
	return fmt.Sprintf(fmtString, id.ApplicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *KnownClientsId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("applications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "00000000-0000-0000-0000-000000000000"),
		resourceids.StaticSegment("knownClients", "knownClients", "knownClients"),
	}
}

func (id *KnownClientsId) String() string {
	return fmt.Sprintf("Known Clients (Application ID: %q)", id.ApplicationId)
}

func (id *KnownClientsId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApplicationId, ok = input.Parsed["applicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applicationId", input)
	}

	return nil
}
