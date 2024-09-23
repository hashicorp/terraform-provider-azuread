package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessTransferMethods string

const (
	ConditionalAccessTransferMethods_AuthenticationTransfer ConditionalAccessTransferMethods = "authenticationTransfer"
	ConditionalAccessTransferMethods_DeviceCodeFlow         ConditionalAccessTransferMethods = "deviceCodeFlow"
	ConditionalAccessTransferMethods_None                   ConditionalAccessTransferMethods = "none"
)

func PossibleValuesForConditionalAccessTransferMethods() []string {
	return []string{
		string(ConditionalAccessTransferMethods_AuthenticationTransfer),
		string(ConditionalAccessTransferMethods_DeviceCodeFlow),
		string(ConditionalAccessTransferMethods_None),
	}
}

func (s *ConditionalAccessTransferMethods) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessTransferMethods(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessTransferMethods(input string) (*ConditionalAccessTransferMethods, error) {
	vals := map[string]ConditionalAccessTransferMethods{
		"authenticationtransfer": ConditionalAccessTransferMethods_AuthenticationTransfer,
		"devicecodeflow":         ConditionalAccessTransferMethods_DeviceCodeFlow,
		"none":                   ConditionalAccessTransferMethods_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessTransferMethods(input)
	return &out, nil
}
