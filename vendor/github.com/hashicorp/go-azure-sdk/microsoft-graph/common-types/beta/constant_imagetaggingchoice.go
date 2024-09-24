package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageTaggingChoice string

const (
	ImageTaggingChoice_Basic    ImageTaggingChoice = "basic"
	ImageTaggingChoice_Disabled ImageTaggingChoice = "disabled"
	ImageTaggingChoice_Enhanced ImageTaggingChoice = "enhanced"
)

func PossibleValuesForImageTaggingChoice() []string {
	return []string{
		string(ImageTaggingChoice_Basic),
		string(ImageTaggingChoice_Disabled),
		string(ImageTaggingChoice_Enhanced),
	}
}

func (s *ImageTaggingChoice) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImageTaggingChoice(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImageTaggingChoice(input string) (*ImageTaggingChoice, error) {
	vals := map[string]ImageTaggingChoice{
		"basic":    ImageTaggingChoice_Basic,
		"disabled": ImageTaggingChoice_Disabled,
		"enhanced": ImageTaggingChoice_Enhanced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImageTaggingChoice(input)
	return &out, nil
}
