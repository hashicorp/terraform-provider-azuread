package parse

import (
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type ServicePrincipalId struct {
	ServicePrincipalId string
}

func NewServicePrincipalID(servicePrincipalId string) *ServicePrincipalId {
	return &ServicePrincipalId{
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseServicePrincipalID parses 'input' into an ServicePrincipalId
func ParseServicePrincipalID(input string) (*ServicePrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServicePrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := &ServicePrincipalId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return id, nil
}

// ValidateServicePrincipalID checks that 'input' can be parsed as an ServicePrincipal ID
func ValidateServicePrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseServicePrincipalID(v)
	if err != nil {
		errors = append(errors, err)
		return
	}

	return validation.IsUUID(id.ServicePrincipalId, "ID")
}

func (id *ServicePrincipalId) ID() string {
	fmtString := "/servicePrincipals/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this ID
func (id *ServicePrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("servicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "00000000-0000-0000-0000-000000000000"),
	}
}

func (id *ServicePrincipalId) String() string {
	return fmt.Sprintf("ServicePrincipal (Object ID: %q)", id.ServicePrincipalId)
}

func (id *ServicePrincipalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ServicePrincipalId, ok = input.Parsed["servicePrincipalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", input)
	}

	return nil
}
