package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeInputType string

const (
	IdentityUserFlowAttributeInputType_CheckboxMultiSelect  IdentityUserFlowAttributeInputType = "checkboxMultiSelect"
	IdentityUserFlowAttributeInputType_DateTimeDropdown     IdentityUserFlowAttributeInputType = "dateTimeDropdown"
	IdentityUserFlowAttributeInputType_DropdownSingleSelect IdentityUserFlowAttributeInputType = "dropdownSingleSelect"
	IdentityUserFlowAttributeInputType_EmailBox             IdentityUserFlowAttributeInputType = "emailBox"
	IdentityUserFlowAttributeInputType_RadioSingleSelect    IdentityUserFlowAttributeInputType = "radioSingleSelect"
	IdentityUserFlowAttributeInputType_TextBox              IdentityUserFlowAttributeInputType = "textBox"
)

func PossibleValuesForIdentityUserFlowAttributeInputType() []string {
	return []string{
		string(IdentityUserFlowAttributeInputType_CheckboxMultiSelect),
		string(IdentityUserFlowAttributeInputType_DateTimeDropdown),
		string(IdentityUserFlowAttributeInputType_DropdownSingleSelect),
		string(IdentityUserFlowAttributeInputType_EmailBox),
		string(IdentityUserFlowAttributeInputType_RadioSingleSelect),
		string(IdentityUserFlowAttributeInputType_TextBox),
	}
}

func (s *IdentityUserFlowAttributeInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowAttributeInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowAttributeInputType(input string) (*IdentityUserFlowAttributeInputType, error) {
	vals := map[string]IdentityUserFlowAttributeInputType{
		"checkboxmultiselect":  IdentityUserFlowAttributeInputType_CheckboxMultiSelect,
		"datetimedropdown":     IdentityUserFlowAttributeInputType_DateTimeDropdown,
		"dropdownsingleselect": IdentityUserFlowAttributeInputType_DropdownSingleSelect,
		"emailbox":             IdentityUserFlowAttributeInputType_EmailBox,
		"radiosingleselect":    IdentityUserFlowAttributeInputType_RadioSingleSelect,
		"textbox":              IdentityUserFlowAttributeInputType_TextBox,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeInputType(input)
	return &out, nil
}
