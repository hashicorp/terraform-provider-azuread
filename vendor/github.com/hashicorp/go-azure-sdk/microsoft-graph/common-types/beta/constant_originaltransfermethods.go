package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OriginalTransferMethods string

const (
	OriginalTransferMethods_AuthenticationTransfer OriginalTransferMethods = "authenticationTransfer"
	OriginalTransferMethods_DeviceCodeFlow         OriginalTransferMethods = "deviceCodeFlow"
	OriginalTransferMethods_None                   OriginalTransferMethods = "none"
)

func PossibleValuesForOriginalTransferMethods() []string {
	return []string{
		string(OriginalTransferMethods_AuthenticationTransfer),
		string(OriginalTransferMethods_DeviceCodeFlow),
		string(OriginalTransferMethods_None),
	}
}

func (s *OriginalTransferMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOriginalTransferMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOriginalTransferMethods(input string) (*OriginalTransferMethods, error) {
	vals := map[string]OriginalTransferMethods{
		"authenticationtransfer": OriginalTransferMethods_AuthenticationTransfer,
		"devicecodeflow":         OriginalTransferMethods_DeviceCodeFlow,
		"none":                   OriginalTransferMethods_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OriginalTransferMethods(input)
	return &out, nil
}
