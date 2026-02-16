package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageRequestState string

const (
	AccessPackageRequestState_Canceled           AccessPackageRequestState = "canceled"
	AccessPackageRequestState_Delivered          AccessPackageRequestState = "delivered"
	AccessPackageRequestState_Delivering         AccessPackageRequestState = "delivering"
	AccessPackageRequestState_DeliveryFailed     AccessPackageRequestState = "deliveryFailed"
	AccessPackageRequestState_Denied             AccessPackageRequestState = "denied"
	AccessPackageRequestState_PartiallyDelivered AccessPackageRequestState = "partiallyDelivered"
	AccessPackageRequestState_PendingApproval    AccessPackageRequestState = "pendingApproval"
	AccessPackageRequestState_Scheduled          AccessPackageRequestState = "scheduled"
	AccessPackageRequestState_Submitted          AccessPackageRequestState = "submitted"
)

func PossibleValuesForAccessPackageRequestState() []string {
	return []string{
		string(AccessPackageRequestState_Canceled),
		string(AccessPackageRequestState_Delivered),
		string(AccessPackageRequestState_Delivering),
		string(AccessPackageRequestState_DeliveryFailed),
		string(AccessPackageRequestState_Denied),
		string(AccessPackageRequestState_PartiallyDelivered),
		string(AccessPackageRequestState_PendingApproval),
		string(AccessPackageRequestState_Scheduled),
		string(AccessPackageRequestState_Submitted),
	}
}

func (s *AccessPackageRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageRequestState(input string) (*AccessPackageRequestState, error) {
	vals := map[string]AccessPackageRequestState{
		"canceled":           AccessPackageRequestState_Canceled,
		"delivered":          AccessPackageRequestState_Delivered,
		"delivering":         AccessPackageRequestState_Delivering,
		"deliveryfailed":     AccessPackageRequestState_DeliveryFailed,
		"denied":             AccessPackageRequestState_Denied,
		"partiallydelivered": AccessPackageRequestState_PartiallyDelivered,
		"pendingapproval":    AccessPackageRequestState_PendingApproval,
		"scheduled":          AccessPackageRequestState_Scheduled,
		"submitted":          AccessPackageRequestState_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageRequestState(input)
	return &out, nil
}
