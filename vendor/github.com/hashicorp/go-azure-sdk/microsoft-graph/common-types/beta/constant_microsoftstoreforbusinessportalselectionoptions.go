package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftStoreForBusinessPortalSelectionOptions string

const (
	MicrosoftStoreForBusinessPortalSelectionOptions_CompanyPortal MicrosoftStoreForBusinessPortalSelectionOptions = "companyPortal"
	MicrosoftStoreForBusinessPortalSelectionOptions_None          MicrosoftStoreForBusinessPortalSelectionOptions = "none"
	MicrosoftStoreForBusinessPortalSelectionOptions_PrivateStore  MicrosoftStoreForBusinessPortalSelectionOptions = "privateStore"
)

func PossibleValuesForMicrosoftStoreForBusinessPortalSelectionOptions() []string {
	return []string{
		string(MicrosoftStoreForBusinessPortalSelectionOptions_CompanyPortal),
		string(MicrosoftStoreForBusinessPortalSelectionOptions_None),
		string(MicrosoftStoreForBusinessPortalSelectionOptions_PrivateStore),
	}
}

func (s *MicrosoftStoreForBusinessPortalSelectionOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftStoreForBusinessPortalSelectionOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftStoreForBusinessPortalSelectionOptions(input string) (*MicrosoftStoreForBusinessPortalSelectionOptions, error) {
	vals := map[string]MicrosoftStoreForBusinessPortalSelectionOptions{
		"companyportal": MicrosoftStoreForBusinessPortalSelectionOptions_CompanyPortal,
		"none":          MicrosoftStoreForBusinessPortalSelectionOptions_None,
		"privatestore":  MicrosoftStoreForBusinessPortalSelectionOptions_PrivateStore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftStoreForBusinessPortalSelectionOptions(input)
	return &out, nil
}
