package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsProductFamily string

const (
	CallRecordsProductFamily_AzureCommunicationServices CallRecordsProductFamily = "azureCommunicationServices"
	CallRecordsProductFamily_Lync                       CallRecordsProductFamily = "lync"
	CallRecordsProductFamily_SkypeForBusiness           CallRecordsProductFamily = "skypeForBusiness"
	CallRecordsProductFamily_Teams                      CallRecordsProductFamily = "teams"
	CallRecordsProductFamily_Unknown                    CallRecordsProductFamily = "unknown"
)

func PossibleValuesForCallRecordsProductFamily() []string {
	return []string{
		string(CallRecordsProductFamily_AzureCommunicationServices),
		string(CallRecordsProductFamily_Lync),
		string(CallRecordsProductFamily_SkypeForBusiness),
		string(CallRecordsProductFamily_Teams),
		string(CallRecordsProductFamily_Unknown),
	}
}

func (s *CallRecordsProductFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsProductFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsProductFamily(input string) (*CallRecordsProductFamily, error) {
	vals := map[string]CallRecordsProductFamily{
		"azurecommunicationservices": CallRecordsProductFamily_AzureCommunicationServices,
		"lync":                       CallRecordsProductFamily_Lync,
		"skypeforbusiness":           CallRecordsProductFamily_SkypeForBusiness,
		"teams":                      CallRecordsProductFamily_Teams,
		"unknown":                    CallRecordsProductFamily_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsProductFamily(input)
	return &out, nil
}
