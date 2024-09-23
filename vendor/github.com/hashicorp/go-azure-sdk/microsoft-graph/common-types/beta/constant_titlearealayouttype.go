package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaLayoutType string

const (
	TitleAreaLayoutType_ColorBlock    TitleAreaLayoutType = "colorBlock"
	TitleAreaLayoutType_ImageAndTitle TitleAreaLayoutType = "imageAndTitle"
	TitleAreaLayoutType_Overlap       TitleAreaLayoutType = "overlap"
	TitleAreaLayoutType_Plain         TitleAreaLayoutType = "plain"
)

func PossibleValuesForTitleAreaLayoutType() []string {
	return []string{
		string(TitleAreaLayoutType_ColorBlock),
		string(TitleAreaLayoutType_ImageAndTitle),
		string(TitleAreaLayoutType_Overlap),
		string(TitleAreaLayoutType_Plain),
	}
}

func (s *TitleAreaLayoutType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTitleAreaLayoutType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTitleAreaLayoutType(input string) (*TitleAreaLayoutType, error) {
	vals := map[string]TitleAreaLayoutType{
		"colorblock":    TitleAreaLayoutType_ColorBlock,
		"imageandtitle": TitleAreaLayoutType_ImageAndTitle,
		"overlap":       TitleAreaLayoutType_Overlap,
		"plain":         TitleAreaLayoutType_Plain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaLayoutType(input)
	return &out, nil
}
