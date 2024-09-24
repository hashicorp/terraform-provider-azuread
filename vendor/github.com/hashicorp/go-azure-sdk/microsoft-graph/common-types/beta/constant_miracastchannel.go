package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MiracastChannel string

const (
	MiracastChannel_Eight                MiracastChannel = "eight"
	MiracastChannel_Eleven               MiracastChannel = "eleven"
	MiracastChannel_Five                 MiracastChannel = "five"
	MiracastChannel_Forty                MiracastChannel = "forty"
	MiracastChannel_FortyEight           MiracastChannel = "fortyEight"
	MiracastChannel_FortyFour            MiracastChannel = "fortyFour"
	MiracastChannel_Four                 MiracastChannel = "four"
	MiracastChannel_Nine                 MiracastChannel = "nine"
	MiracastChannel_One                  MiracastChannel = "one"
	MiracastChannel_OneHundredFiftySeven MiracastChannel = "oneHundredFiftySeven"
	MiracastChannel_OneHundredFiftyThree MiracastChannel = "oneHundredFiftyThree"
	MiracastChannel_OneHundredFortyNine  MiracastChannel = "oneHundredFortyNine"
	MiracastChannel_OneHundredSixtyFive  MiracastChannel = "oneHundredSixtyFive"
	MiracastChannel_OneHundredSixtyOne   MiracastChannel = "oneHundredSixtyOne"
	MiracastChannel_Seven                MiracastChannel = "seven"
	MiracastChannel_Six                  MiracastChannel = "six"
	MiracastChannel_Ten                  MiracastChannel = "ten"
	MiracastChannel_ThirtySix            MiracastChannel = "thirtySix"
	MiracastChannel_Three                MiracastChannel = "three"
	MiracastChannel_Two                  MiracastChannel = "two"
	MiracastChannel_UserDefined          MiracastChannel = "userDefined"
)

func PossibleValuesForMiracastChannel() []string {
	return []string{
		string(MiracastChannel_Eight),
		string(MiracastChannel_Eleven),
		string(MiracastChannel_Five),
		string(MiracastChannel_Forty),
		string(MiracastChannel_FortyEight),
		string(MiracastChannel_FortyFour),
		string(MiracastChannel_Four),
		string(MiracastChannel_Nine),
		string(MiracastChannel_One),
		string(MiracastChannel_OneHundredFiftySeven),
		string(MiracastChannel_OneHundredFiftyThree),
		string(MiracastChannel_OneHundredFortyNine),
		string(MiracastChannel_OneHundredSixtyFive),
		string(MiracastChannel_OneHundredSixtyOne),
		string(MiracastChannel_Seven),
		string(MiracastChannel_Six),
		string(MiracastChannel_Ten),
		string(MiracastChannel_ThirtySix),
		string(MiracastChannel_Three),
		string(MiracastChannel_Two),
		string(MiracastChannel_UserDefined),
	}
}

func (s *MiracastChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMiracastChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMiracastChannel(input string) (*MiracastChannel, error) {
	vals := map[string]MiracastChannel{
		"eight":                MiracastChannel_Eight,
		"eleven":               MiracastChannel_Eleven,
		"five":                 MiracastChannel_Five,
		"forty":                MiracastChannel_Forty,
		"fortyeight":           MiracastChannel_FortyEight,
		"fortyfour":            MiracastChannel_FortyFour,
		"four":                 MiracastChannel_Four,
		"nine":                 MiracastChannel_Nine,
		"one":                  MiracastChannel_One,
		"onehundredfiftyseven": MiracastChannel_OneHundredFiftySeven,
		"onehundredfiftythree": MiracastChannel_OneHundredFiftyThree,
		"onehundredfortynine":  MiracastChannel_OneHundredFortyNine,
		"onehundredsixtyfive":  MiracastChannel_OneHundredSixtyFive,
		"onehundredsixtyone":   MiracastChannel_OneHundredSixtyOne,
		"seven":                MiracastChannel_Seven,
		"six":                  MiracastChannel_Six,
		"ten":                  MiracastChannel_Ten,
		"thirtysix":            MiracastChannel_ThirtySix,
		"three":                MiracastChannel_Three,
		"two":                  MiracastChannel_Two,
		"userdefined":          MiracastChannel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MiracastChannel(input)
	return &out, nil
}
