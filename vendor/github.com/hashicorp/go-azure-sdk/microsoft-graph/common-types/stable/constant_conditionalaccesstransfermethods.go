package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessTransferMethods string

const (
	ConditionalAccessTransferMethods_AuthenticationTransfer               ConditionalAccessTransferMethods = "authenticationTransfer"
	ConditionalAccessTransferMethods_DeviceCodeFlow                       ConditionalAccessTransferMethods = "deviceCodeFlow"
	ConditionalAccessTransferMethods_AuthenticationTransferDeviceCodeFlow ConditionalAccessTransferMethods = "authenticationTransfer,deviceCodeFlow"
	ConditionalAccessTransferMethods_DeviceCodeFlowAuthenticationTransfer ConditionalAccessTransferMethods = "deviceCodeFlow,authenticationTransfer"
	ConditionalAccessTransferMethods_None                                 ConditionalAccessTransferMethods = "none"
)

func PossibleValuesForConditionalAccessTransferMethods() []string {
	return []string{
		string(ConditionalAccessTransferMethods_AuthenticationTransfer),
		string(ConditionalAccessTransferMethods_DeviceCodeFlow),
		string(ConditionalAccessTransferMethods_AuthenticationTransferDeviceCodeFlow),
		string(ConditionalAccessTransferMethods_DeviceCodeFlowAuthenticationTransfer),
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
		"authenticationTransfer":                ConditionalAccessTransferMethods_AuthenticationTransfer,
		"deviceCodeFlow":                        ConditionalAccessTransferMethods_DeviceCodeFlow,
		"uthenticationTransfer,deviceCodeFlow":  ConditionalAccessTransferMethods_AuthenticationTransferDeviceCodeFlow,
		"deviceCodeFlow,authenticationTransfer": ConditionalAccessTransferMethods_DeviceCodeFlowAuthenticationTransfer,
		"none":                                  ConditionalAccessTransferMethods_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessTransferMethods(input)
	return &out, nil
}
