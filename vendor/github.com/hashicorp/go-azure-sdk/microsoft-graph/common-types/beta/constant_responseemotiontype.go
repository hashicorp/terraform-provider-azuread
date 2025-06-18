package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResponseEmotionType string

const (
	ResponseEmotionType_Ambitious    ResponseEmotionType = "ambitious"
	ResponseEmotionType_Angry        ResponseEmotionType = "angry"
	ResponseEmotionType_Annoyed      ResponseEmotionType = "annoyed"
	ResponseEmotionType_Anxious      ResponseEmotionType = "anxious"
	ResponseEmotionType_Apathetic    ResponseEmotionType = "apathetic"
	ResponseEmotionType_Ashamed      ResponseEmotionType = "ashamed"
	ResponseEmotionType_Awed         ResponseEmotionType = "awed"
	ResponseEmotionType_Bored        ResponseEmotionType = "bored"
	ResponseEmotionType_Calm         ResponseEmotionType = "calm"
	ResponseEmotionType_Cheerful     ResponseEmotionType = "cheerful"
	ResponseEmotionType_Comfortable  ResponseEmotionType = "comfortable"
	ResponseEmotionType_Concerned    ResponseEmotionType = "concerned"
	ResponseEmotionType_Confident    ResponseEmotionType = "confident"
	ResponseEmotionType_Confused     ResponseEmotionType = "confused"
	ResponseEmotionType_Content      ResponseEmotionType = "content"
	ResponseEmotionType_Creative     ResponseEmotionType = "creative"
	ResponseEmotionType_Curious      ResponseEmotionType = "curious"
	ResponseEmotionType_Depressed    ResponseEmotionType = "depressed"
	ResponseEmotionType_Determined   ResponseEmotionType = "determined"
	ResponseEmotionType_Disappointed ResponseEmotionType = "disappointed"
	ResponseEmotionType_Energized    ResponseEmotionType = "energized"
	ResponseEmotionType_Excited      ResponseEmotionType = "excited"
	ResponseEmotionType_Exhausted    ResponseEmotionType = "exhausted"
	ResponseEmotionType_Focused      ResponseEmotionType = "focused"
	ResponseEmotionType_Frightened   ResponseEmotionType = "frightened"
	ResponseEmotionType_Frustrated   ResponseEmotionType = "frustrated"
	ResponseEmotionType_Fulfilled    ResponseEmotionType = "fulfilled"
	ResponseEmotionType_Glad         ResponseEmotionType = "glad"
	ResponseEmotionType_Grateful     ResponseEmotionType = "grateful"
	ResponseEmotionType_Happy        ResponseEmotionType = "happy"
	ResponseEmotionType_Hopeless     ResponseEmotionType = "hopeless"
	ResponseEmotionType_Hurt         ResponseEmotionType = "hurt"
	ResponseEmotionType_Included     ResponseEmotionType = "included"
	ResponseEmotionType_Inspired     ResponseEmotionType = "inspired"
	ResponseEmotionType_Jealous      ResponseEmotionType = "jealous"
	ResponseEmotionType_Lonely       ResponseEmotionType = "lonely"
	ResponseEmotionType_Miserable    ResponseEmotionType = "miserable"
	ResponseEmotionType_Motivated    ResponseEmotionType = "motivated"
	ResponseEmotionType_Nervous      ResponseEmotionType = "nervous"
	ResponseEmotionType_None         ResponseEmotionType = "none"
	ResponseEmotionType_Optimistic   ResponseEmotionType = "optimistic"
	ResponseEmotionType_Overwhelmed  ResponseEmotionType = "overwhelmed"
	ResponseEmotionType_Peaceful     ResponseEmotionType = "peaceful"
	ResponseEmotionType_Pensive      ResponseEmotionType = "pensive"
	ResponseEmotionType_Proud        ResponseEmotionType = "proud"
	ResponseEmotionType_Reserved     ResponseEmotionType = "reserved"
	ResponseEmotionType_Restless     ResponseEmotionType = "restless"
	ResponseEmotionType_Sad          ResponseEmotionType = "sad"
	ResponseEmotionType_Sensitive    ResponseEmotionType = "sensitive"
	ResponseEmotionType_Shocked      ResponseEmotionType = "shocked"
	ResponseEmotionType_Skeptical    ResponseEmotionType = "skeptical"
	ResponseEmotionType_Stressed     ResponseEmotionType = "stressed"
	ResponseEmotionType_Stuck        ResponseEmotionType = "stuck"
	ResponseEmotionType_Successful   ResponseEmotionType = "successful"
	ResponseEmotionType_Tired        ResponseEmotionType = "tired"
	ResponseEmotionType_Valuable     ResponseEmotionType = "valuable"
	ResponseEmotionType_Worthless    ResponseEmotionType = "worthless"
)

func PossibleValuesForResponseEmotionType() []string {
	return []string{
		string(ResponseEmotionType_Ambitious),
		string(ResponseEmotionType_Angry),
		string(ResponseEmotionType_Annoyed),
		string(ResponseEmotionType_Anxious),
		string(ResponseEmotionType_Apathetic),
		string(ResponseEmotionType_Ashamed),
		string(ResponseEmotionType_Awed),
		string(ResponseEmotionType_Bored),
		string(ResponseEmotionType_Calm),
		string(ResponseEmotionType_Cheerful),
		string(ResponseEmotionType_Comfortable),
		string(ResponseEmotionType_Concerned),
		string(ResponseEmotionType_Confident),
		string(ResponseEmotionType_Confused),
		string(ResponseEmotionType_Content),
		string(ResponseEmotionType_Creative),
		string(ResponseEmotionType_Curious),
		string(ResponseEmotionType_Depressed),
		string(ResponseEmotionType_Determined),
		string(ResponseEmotionType_Disappointed),
		string(ResponseEmotionType_Energized),
		string(ResponseEmotionType_Excited),
		string(ResponseEmotionType_Exhausted),
		string(ResponseEmotionType_Focused),
		string(ResponseEmotionType_Frightened),
		string(ResponseEmotionType_Frustrated),
		string(ResponseEmotionType_Fulfilled),
		string(ResponseEmotionType_Glad),
		string(ResponseEmotionType_Grateful),
		string(ResponseEmotionType_Happy),
		string(ResponseEmotionType_Hopeless),
		string(ResponseEmotionType_Hurt),
		string(ResponseEmotionType_Included),
		string(ResponseEmotionType_Inspired),
		string(ResponseEmotionType_Jealous),
		string(ResponseEmotionType_Lonely),
		string(ResponseEmotionType_Miserable),
		string(ResponseEmotionType_Motivated),
		string(ResponseEmotionType_Nervous),
		string(ResponseEmotionType_None),
		string(ResponseEmotionType_Optimistic),
		string(ResponseEmotionType_Overwhelmed),
		string(ResponseEmotionType_Peaceful),
		string(ResponseEmotionType_Pensive),
		string(ResponseEmotionType_Proud),
		string(ResponseEmotionType_Reserved),
		string(ResponseEmotionType_Restless),
		string(ResponseEmotionType_Sad),
		string(ResponseEmotionType_Sensitive),
		string(ResponseEmotionType_Shocked),
		string(ResponseEmotionType_Skeptical),
		string(ResponseEmotionType_Stressed),
		string(ResponseEmotionType_Stuck),
		string(ResponseEmotionType_Successful),
		string(ResponseEmotionType_Tired),
		string(ResponseEmotionType_Valuable),
		string(ResponseEmotionType_Worthless),
	}
}

func (s *ResponseEmotionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResponseEmotionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResponseEmotionType(input string) (*ResponseEmotionType, error) {
	vals := map[string]ResponseEmotionType{
		"ambitious":    ResponseEmotionType_Ambitious,
		"angry":        ResponseEmotionType_Angry,
		"annoyed":      ResponseEmotionType_Annoyed,
		"anxious":      ResponseEmotionType_Anxious,
		"apathetic":    ResponseEmotionType_Apathetic,
		"ashamed":      ResponseEmotionType_Ashamed,
		"awed":         ResponseEmotionType_Awed,
		"bored":        ResponseEmotionType_Bored,
		"calm":         ResponseEmotionType_Calm,
		"cheerful":     ResponseEmotionType_Cheerful,
		"comfortable":  ResponseEmotionType_Comfortable,
		"concerned":    ResponseEmotionType_Concerned,
		"confident":    ResponseEmotionType_Confident,
		"confused":     ResponseEmotionType_Confused,
		"content":      ResponseEmotionType_Content,
		"creative":     ResponseEmotionType_Creative,
		"curious":      ResponseEmotionType_Curious,
		"depressed":    ResponseEmotionType_Depressed,
		"determined":   ResponseEmotionType_Determined,
		"disappointed": ResponseEmotionType_Disappointed,
		"energized":    ResponseEmotionType_Energized,
		"excited":      ResponseEmotionType_Excited,
		"exhausted":    ResponseEmotionType_Exhausted,
		"focused":      ResponseEmotionType_Focused,
		"frightened":   ResponseEmotionType_Frightened,
		"frustrated":   ResponseEmotionType_Frustrated,
		"fulfilled":    ResponseEmotionType_Fulfilled,
		"glad":         ResponseEmotionType_Glad,
		"grateful":     ResponseEmotionType_Grateful,
		"happy":        ResponseEmotionType_Happy,
		"hopeless":     ResponseEmotionType_Hopeless,
		"hurt":         ResponseEmotionType_Hurt,
		"included":     ResponseEmotionType_Included,
		"inspired":     ResponseEmotionType_Inspired,
		"jealous":      ResponseEmotionType_Jealous,
		"lonely":       ResponseEmotionType_Lonely,
		"miserable":    ResponseEmotionType_Miserable,
		"motivated":    ResponseEmotionType_Motivated,
		"nervous":      ResponseEmotionType_Nervous,
		"none":         ResponseEmotionType_None,
		"optimistic":   ResponseEmotionType_Optimistic,
		"overwhelmed":  ResponseEmotionType_Overwhelmed,
		"peaceful":     ResponseEmotionType_Peaceful,
		"pensive":      ResponseEmotionType_Pensive,
		"proud":        ResponseEmotionType_Proud,
		"reserved":     ResponseEmotionType_Reserved,
		"restless":     ResponseEmotionType_Restless,
		"sad":          ResponseEmotionType_Sad,
		"sensitive":    ResponseEmotionType_Sensitive,
		"shocked":      ResponseEmotionType_Shocked,
		"skeptical":    ResponseEmotionType_Skeptical,
		"stressed":     ResponseEmotionType_Stressed,
		"stuck":        ResponseEmotionType_Stuck,
		"successful":   ResponseEmotionType_Successful,
		"tired":        ResponseEmotionType_Tired,
		"valuable":     ResponseEmotionType_Valuable,
		"worthless":    ResponseEmotionType_Worthless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResponseEmotionType(input)
	return &out, nil
}
