package parse

import (
	"fmt"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type PrivilegedAccessGroupEligibilityScheduleRequestId struct {
	RequestId string
}

func NewPrivilegedAccessGroupEligibilityScheduleRequestID(requestId string) *PrivilegedAccessGroupEligibilityScheduleRequestId {
	return &PrivilegedAccessGroupEligibilityScheduleRequestId{
		RequestId: requestId,
	}
}

func ParsePrivilegedAccessGroupEligibilityScheduleRequestID(idString string) (*PrivilegedAccessGroupEligibilityScheduleRequestId, error) {
	if _, err := validation.IsUUID(idString, "RequestId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RequestId: %+v", err)
	}

	return &PrivilegedAccessGroupEligibilityScheduleRequestId{
		RequestId: idString,
	}, nil
}

func (id *PrivilegedAccessGroupEligibilityScheduleRequestId) ID() string {
	return id.RequestId
}

func (id *PrivilegedAccessGroupEligibilityScheduleRequestId) String() string {
	return fmt.Sprintf("Privileged Access Group Assigment Schedule Request ID: %q", id.RequestId)
}
