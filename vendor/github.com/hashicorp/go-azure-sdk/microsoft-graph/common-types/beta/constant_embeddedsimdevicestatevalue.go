package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMDeviceStateValue string

const (
	EmbeddedSIMDeviceStateValue_Deleted       EmbeddedSIMDeviceStateValue = "deleted"
	EmbeddedSIMDeviceStateValue_Deleting      EmbeddedSIMDeviceStateValue = "deleting"
	EmbeddedSIMDeviceStateValue_Error         EmbeddedSIMDeviceStateValue = "error"
	EmbeddedSIMDeviceStateValue_Failed        EmbeddedSIMDeviceStateValue = "failed"
	EmbeddedSIMDeviceStateValue_Installed     EmbeddedSIMDeviceStateValue = "installed"
	EmbeddedSIMDeviceStateValue_Installing    EmbeddedSIMDeviceStateValue = "installing"
	EmbeddedSIMDeviceStateValue_NotEvaluated  EmbeddedSIMDeviceStateValue = "notEvaluated"
	EmbeddedSIMDeviceStateValue_RemovedByUser EmbeddedSIMDeviceStateValue = "removedByUser"
)

func PossibleValuesForEmbeddedSIMDeviceStateValue() []string {
	return []string{
		string(EmbeddedSIMDeviceStateValue_Deleted),
		string(EmbeddedSIMDeviceStateValue_Deleting),
		string(EmbeddedSIMDeviceStateValue_Error),
		string(EmbeddedSIMDeviceStateValue_Failed),
		string(EmbeddedSIMDeviceStateValue_Installed),
		string(EmbeddedSIMDeviceStateValue_Installing),
		string(EmbeddedSIMDeviceStateValue_NotEvaluated),
		string(EmbeddedSIMDeviceStateValue_RemovedByUser),
	}
}

func (s *EmbeddedSIMDeviceStateValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmbeddedSIMDeviceStateValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmbeddedSIMDeviceStateValue(input string) (*EmbeddedSIMDeviceStateValue, error) {
	vals := map[string]EmbeddedSIMDeviceStateValue{
		"deleted":       EmbeddedSIMDeviceStateValue_Deleted,
		"deleting":      EmbeddedSIMDeviceStateValue_Deleting,
		"error":         EmbeddedSIMDeviceStateValue_Error,
		"failed":        EmbeddedSIMDeviceStateValue_Failed,
		"installed":     EmbeddedSIMDeviceStateValue_Installed,
		"installing":    EmbeddedSIMDeviceStateValue_Installing,
		"notevaluated":  EmbeddedSIMDeviceStateValue_NotEvaluated,
		"removedbyuser": EmbeddedSIMDeviceStateValue_RemovedByUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmbeddedSIMDeviceStateValue(input)
	return &out, nil
}
