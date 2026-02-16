package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegateMeetingMessageDeliveryOptions string

const (
	DelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal DelegateMeetingMessageDeliveryOptions = "sendToDelegateAndInformationToPrincipal"
	DelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal              DelegateMeetingMessageDeliveryOptions = "sendToDelegateAndPrincipal"
	DelegateMeetingMessageDeliveryOptions_SendToDelegateOnly                      DelegateMeetingMessageDeliveryOptions = "sendToDelegateOnly"
)

func PossibleValuesForDelegateMeetingMessageDeliveryOptions() []string {
	return []string{
		string(DelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal),
		string(DelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal),
		string(DelegateMeetingMessageDeliveryOptions_SendToDelegateOnly),
	}
}

func (s *DelegateMeetingMessageDeliveryOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegateMeetingMessageDeliveryOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegateMeetingMessageDeliveryOptions(input string) (*DelegateMeetingMessageDeliveryOptions, error) {
	vals := map[string]DelegateMeetingMessageDeliveryOptions{
		"sendtodelegateandinformationtoprincipal": DelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal,
		"sendtodelegateandprincipal":              DelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal,
		"sendtodelegateonly":                      DelegateMeetingMessageDeliveryOptions_SendToDelegateOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegateMeetingMessageDeliveryOptions(input)
	return &out, nil
}
