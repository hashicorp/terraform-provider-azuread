package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentPermission string

const (
	ReferenceAttachmentPermission_AnonymousEdit    ReferenceAttachmentPermission = "anonymousEdit"
	ReferenceAttachmentPermission_AnonymousView    ReferenceAttachmentPermission = "anonymousView"
	ReferenceAttachmentPermission_Edit             ReferenceAttachmentPermission = "edit"
	ReferenceAttachmentPermission_OrganizationEdit ReferenceAttachmentPermission = "organizationEdit"
	ReferenceAttachmentPermission_OrganizationView ReferenceAttachmentPermission = "organizationView"
	ReferenceAttachmentPermission_Other            ReferenceAttachmentPermission = "other"
	ReferenceAttachmentPermission_View             ReferenceAttachmentPermission = "view"
)

func PossibleValuesForReferenceAttachmentPermission() []string {
	return []string{
		string(ReferenceAttachmentPermission_AnonymousEdit),
		string(ReferenceAttachmentPermission_AnonymousView),
		string(ReferenceAttachmentPermission_Edit),
		string(ReferenceAttachmentPermission_OrganizationEdit),
		string(ReferenceAttachmentPermission_OrganizationView),
		string(ReferenceAttachmentPermission_Other),
		string(ReferenceAttachmentPermission_View),
	}
}

func (s *ReferenceAttachmentPermission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReferenceAttachmentPermission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReferenceAttachmentPermission(input string) (*ReferenceAttachmentPermission, error) {
	vals := map[string]ReferenceAttachmentPermission{
		"anonymousedit":    ReferenceAttachmentPermission_AnonymousEdit,
		"anonymousview":    ReferenceAttachmentPermission_AnonymousView,
		"edit":             ReferenceAttachmentPermission_Edit,
		"organizationedit": ReferenceAttachmentPermission_OrganizationEdit,
		"organizationview": ReferenceAttachmentPermission_OrganizationView,
		"other":            ReferenceAttachmentPermission_Other,
		"view":             ReferenceAttachmentPermission_View,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceAttachmentPermission(input)
	return &out, nil
}
