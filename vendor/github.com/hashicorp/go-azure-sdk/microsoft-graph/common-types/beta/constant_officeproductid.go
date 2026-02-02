package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeProductId string

const (
	OfficeProductId_O365BusinessRetail OfficeProductId = "o365BusinessRetail"
	OfficeProductId_O365ProPlusRetail  OfficeProductId = "o365ProPlusRetail"
	OfficeProductId_ProjectProRetail   OfficeProductId = "projectProRetail"
	OfficeProductId_VisioProRetail     OfficeProductId = "visioProRetail"
)

func PossibleValuesForOfficeProductId() []string {
	return []string{
		string(OfficeProductId_O365BusinessRetail),
		string(OfficeProductId_O365ProPlusRetail),
		string(OfficeProductId_ProjectProRetail),
		string(OfficeProductId_VisioProRetail),
	}
}

func (s *OfficeProductId) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOfficeProductId(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOfficeProductId(input string) (*OfficeProductId, error) {
	vals := map[string]OfficeProductId{
		"o365businessretail": OfficeProductId_O365BusinessRetail,
		"o365proplusretail":  OfficeProductId_O365ProPlusRetail,
		"projectproretail":   OfficeProductId_ProjectProRetail,
		"visioproretail":     OfficeProductId_VisioProRetail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeProductId(input)
	return &out, nil
}
