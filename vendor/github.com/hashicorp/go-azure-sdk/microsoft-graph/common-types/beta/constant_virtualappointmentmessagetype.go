package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualAppointmentMessageType string

const (
	VirtualAppointmentMessageType_Cancellation VirtualAppointmentMessageType = "cancellation"
	VirtualAppointmentMessageType_Confirmation VirtualAppointmentMessageType = "confirmation"
	VirtualAppointmentMessageType_Reschedule   VirtualAppointmentMessageType = "reschedule"
)

func PossibleValuesForVirtualAppointmentMessageType() []string {
	return []string{
		string(VirtualAppointmentMessageType_Cancellation),
		string(VirtualAppointmentMessageType_Confirmation),
		string(VirtualAppointmentMessageType_Reschedule),
	}
}

func (s *VirtualAppointmentMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualAppointmentMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualAppointmentMessageType(input string) (*VirtualAppointmentMessageType, error) {
	vals := map[string]VirtualAppointmentMessageType{
		"cancellation": VirtualAppointmentMessageType_Cancellation,
		"confirmation": VirtualAppointmentMessageType_Confirmation,
		"reschedule":   VirtualAppointmentMessageType_Reschedule,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualAppointmentMessageType(input)
	return &out, nil
}
