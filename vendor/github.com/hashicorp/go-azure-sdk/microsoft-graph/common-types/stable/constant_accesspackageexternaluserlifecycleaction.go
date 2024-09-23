package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageExternalUserLifecycleAction string

const (
	AccessPackageExternalUserLifecycleAction_BlockSignIn          AccessPackageExternalUserLifecycleAction = "blockSignIn"
	AccessPackageExternalUserLifecycleAction_BlockSignInAndDelete AccessPackageExternalUserLifecycleAction = "blockSignInAndDelete"
	AccessPackageExternalUserLifecycleAction_None                 AccessPackageExternalUserLifecycleAction = "none"
)

func PossibleValuesForAccessPackageExternalUserLifecycleAction() []string {
	return []string{
		string(AccessPackageExternalUserLifecycleAction_BlockSignIn),
		string(AccessPackageExternalUserLifecycleAction_BlockSignInAndDelete),
		string(AccessPackageExternalUserLifecycleAction_None),
	}
}

func (s *AccessPackageExternalUserLifecycleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageExternalUserLifecycleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageExternalUserLifecycleAction(input string) (*AccessPackageExternalUserLifecycleAction, error) {
	vals := map[string]AccessPackageExternalUserLifecycleAction{
		"blocksignin":          AccessPackageExternalUserLifecycleAction_BlockSignIn,
		"blocksigninanddelete": AccessPackageExternalUserLifecycleAction_BlockSignInAndDelete,
		"none":                 AccessPackageExternalUserLifecycleAction_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageExternalUserLifecycleAction(input)
	return &out, nil
}
