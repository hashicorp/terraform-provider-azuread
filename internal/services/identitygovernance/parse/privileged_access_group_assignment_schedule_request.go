package parse

import (
	"fmt"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type PrivilegedAccessGroupAssignmentScheduleRequestId struct {
	RequestId string
}

func NewPrivilegedAccessGroupAssignmentScheduleRequestID(requestId string) *PrivilegedAccessGroupAssignmentScheduleRequestId {
	return &PrivilegedAccessGroupAssignmentScheduleRequestId{
		RequestId: requestId,
	}
}

func ParsePrivilegedAccessGroupAssignmentScheduleRequestID(idString string) (*PrivilegedAccessGroupAssignmentScheduleRequestId, error) {
	if _, err := validation.IsUUID(idString, "RequestId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RequestId: %+v", err)
	}

	return &PrivilegedAccessGroupAssignmentScheduleRequestId{
		RequestId: idString,
	}, nil
}

func (id *PrivilegedAccessGroupAssignmentScheduleRequestId) ID() string {
	return id.RequestId
}

func (id *PrivilegedAccessGroupAssignmentScheduleRequestId) String() string {
	return fmt.Sprintf("Privileged Access Group Assigment Schedule Request ID: %q", id.RequestId)
}
