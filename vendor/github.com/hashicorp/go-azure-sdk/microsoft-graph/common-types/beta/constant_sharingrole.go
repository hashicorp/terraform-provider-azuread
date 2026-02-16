package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingRole string

const (
	SharingRole_Edit           SharingRole = "edit"
	SharingRole_ManageList     SharingRole = "manageList"
	SharingRole_None           SharingRole = "none"
	SharingRole_RestrictedView SharingRole = "restrictedView"
	SharingRole_Review         SharingRole = "review"
	SharingRole_SubmitOnly     SharingRole = "submitOnly"
	SharingRole_View           SharingRole = "view"
)

func PossibleValuesForSharingRole() []string {
	return []string{
		string(SharingRole_Edit),
		string(SharingRole_ManageList),
		string(SharingRole_None),
		string(SharingRole_RestrictedView),
		string(SharingRole_Review),
		string(SharingRole_SubmitOnly),
		string(SharingRole_View),
	}
}

func (s *SharingRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingRole(input string) (*SharingRole, error) {
	vals := map[string]SharingRole{
		"edit":           SharingRole_Edit,
		"managelist":     SharingRole_ManageList,
		"none":           SharingRole_None,
		"restrictedview": SharingRole_RestrictedView,
		"review":         SharingRole_Review,
		"submitonly":     SharingRole_SubmitOnly,
		"view":           SharingRole_View,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingRole(input)
	return &out, nil
}
