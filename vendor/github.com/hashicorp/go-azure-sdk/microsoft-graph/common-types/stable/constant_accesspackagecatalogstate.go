package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalogState string

const (
	AccessPackageCatalogState_Published   AccessPackageCatalogState = "published"
	AccessPackageCatalogState_Unpublished AccessPackageCatalogState = "unpublished"
)

func PossibleValuesForAccessPackageCatalogState() []string {
	return []string{
		string(AccessPackageCatalogState_Published),
		string(AccessPackageCatalogState_Unpublished),
	}
}

func (s *AccessPackageCatalogState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageCatalogState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageCatalogState(input string) (*AccessPackageCatalogState, error) {
	vals := map[string]AccessPackageCatalogState{
		"published":   AccessPackageCatalogState_Published,
		"unpublished": AccessPackageCatalogState_Unpublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCatalogState(input)
	return &out, nil
}
