package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentState string

const (
	AccessPackageAssignmentState_Delivered          AccessPackageAssignmentState = "delivered"
	AccessPackageAssignmentState_Delivering         AccessPackageAssignmentState = "delivering"
	AccessPackageAssignmentState_DeliveryFailed     AccessPackageAssignmentState = "deliveryFailed"
	AccessPackageAssignmentState_Expired            AccessPackageAssignmentState = "expired"
	AccessPackageAssignmentState_PartiallyDelivered AccessPackageAssignmentState = "partiallyDelivered"
)

func PossibleValuesForAccessPackageAssignmentState() []string {
	return []string{
		string(AccessPackageAssignmentState_Delivered),
		string(AccessPackageAssignmentState_Delivering),
		string(AccessPackageAssignmentState_DeliveryFailed),
		string(AccessPackageAssignmentState_Expired),
		string(AccessPackageAssignmentState_PartiallyDelivered),
	}
}

func (s *AccessPackageAssignmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentState(input string) (*AccessPackageAssignmentState, error) {
	vals := map[string]AccessPackageAssignmentState{
		"delivered":          AccessPackageAssignmentState_Delivered,
		"delivering":         AccessPackageAssignmentState_Delivering,
		"deliveryfailed":     AccessPackageAssignmentState_DeliveryFailed,
		"expired":            AccessPackageAssignmentState_Expired,
		"partiallydelivered": AccessPackageAssignmentState_PartiallyDelivered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentState(input)
	return &out, nil
}
