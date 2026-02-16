package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingPresenters string

const (
	OnlineMeetingPresenters_Everyone        OnlineMeetingPresenters = "everyone"
	OnlineMeetingPresenters_Organization    OnlineMeetingPresenters = "organization"
	OnlineMeetingPresenters_Organizer       OnlineMeetingPresenters = "organizer"
	OnlineMeetingPresenters_RoleIsPresenter OnlineMeetingPresenters = "roleIsPresenter"
)

func PossibleValuesForOnlineMeetingPresenters() []string {
	return []string{
		string(OnlineMeetingPresenters_Everyone),
		string(OnlineMeetingPresenters_Organization),
		string(OnlineMeetingPresenters_Organizer),
		string(OnlineMeetingPresenters_RoleIsPresenter),
	}
}

func (s *OnlineMeetingPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingPresenters(input string) (*OnlineMeetingPresenters, error) {
	vals := map[string]OnlineMeetingPresenters{
		"everyone":        OnlineMeetingPresenters_Everyone,
		"organization":    OnlineMeetingPresenters_Organization,
		"organizer":       OnlineMeetingPresenters_Organizer,
		"roleispresenter": OnlineMeetingPresenters_RoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingPresenters(input)
	return &out, nil
}
