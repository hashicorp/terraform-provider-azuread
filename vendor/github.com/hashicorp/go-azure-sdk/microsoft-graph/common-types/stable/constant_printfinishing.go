package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintFinishing string

const (
	PrintFinishing_Bale               PrintFinishing = "bale"
	PrintFinishing_Bind               PrintFinishing = "bind"
	PrintFinishing_BindBottom         PrintFinishing = "bindBottom"
	PrintFinishing_BindLeft           PrintFinishing = "bindLeft"
	PrintFinishing_BindRight          PrintFinishing = "bindRight"
	PrintFinishing_BindTop            PrintFinishing = "bindTop"
	PrintFinishing_BookletMaker       PrintFinishing = "bookletMaker"
	PrintFinishing_Coat               PrintFinishing = "coat"
	PrintFinishing_Cover              PrintFinishing = "cover"
	PrintFinishing_Fold               PrintFinishing = "fold"
	PrintFinishing_FoldAccordion      PrintFinishing = "foldAccordion"
	PrintFinishing_FoldDoubleGate     PrintFinishing = "foldDoubleGate"
	PrintFinishing_FoldEngineeringZ   PrintFinishing = "foldEngineeringZ"
	PrintFinishing_FoldGate           PrintFinishing = "foldGate"
	PrintFinishing_FoldHalf           PrintFinishing = "foldHalf"
	PrintFinishing_FoldHalfZ          PrintFinishing = "foldHalfZ"
	PrintFinishing_FoldLeftGate       PrintFinishing = "foldLeftGate"
	PrintFinishing_FoldLetter         PrintFinishing = "foldLetter"
	PrintFinishing_FoldParallel       PrintFinishing = "foldParallel"
	PrintFinishing_FoldPoster         PrintFinishing = "foldPoster"
	PrintFinishing_FoldRightGate      PrintFinishing = "foldRightGate"
	PrintFinishing_FoldZ              PrintFinishing = "foldZ"
	PrintFinishing_Laminate           PrintFinishing = "laminate"
	PrintFinishing_None               PrintFinishing = "none"
	PrintFinishing_Punch              PrintFinishing = "punch"
	PrintFinishing_PunchBottomLeft    PrintFinishing = "punchBottomLeft"
	PrintFinishing_PunchBottomRight   PrintFinishing = "punchBottomRight"
	PrintFinishing_PunchDualBottom    PrintFinishing = "punchDualBottom"
	PrintFinishing_PunchDualLeft      PrintFinishing = "punchDualLeft"
	PrintFinishing_PunchDualRight     PrintFinishing = "punchDualRight"
	PrintFinishing_PunchDualTop       PrintFinishing = "punchDualTop"
	PrintFinishing_PunchQuadBottom    PrintFinishing = "punchQuadBottom"
	PrintFinishing_PunchQuadLeft      PrintFinishing = "punchQuadLeft"
	PrintFinishing_PunchQuadRight     PrintFinishing = "punchQuadRight"
	PrintFinishing_PunchQuadTop       PrintFinishing = "punchQuadTop"
	PrintFinishing_PunchTopLeft       PrintFinishing = "punchTopLeft"
	PrintFinishing_PunchTopRight      PrintFinishing = "punchTopRight"
	PrintFinishing_PunchTripleBottom  PrintFinishing = "punchTripleBottom"
	PrintFinishing_PunchTripleLeft    PrintFinishing = "punchTripleLeft"
	PrintFinishing_PunchTripleRight   PrintFinishing = "punchTripleRight"
	PrintFinishing_PunchTripleTop     PrintFinishing = "punchTripleTop"
	PrintFinishing_SaddleStitch       PrintFinishing = "saddleStitch"
	PrintFinishing_Staple             PrintFinishing = "staple"
	PrintFinishing_StapleBottomLeft   PrintFinishing = "stapleBottomLeft"
	PrintFinishing_StapleBottomRight  PrintFinishing = "stapleBottomRight"
	PrintFinishing_StapleDualBottom   PrintFinishing = "stapleDualBottom"
	PrintFinishing_StapleDualLeft     PrintFinishing = "stapleDualLeft"
	PrintFinishing_StapleDualRight    PrintFinishing = "stapleDualRight"
	PrintFinishing_StapleDualTop      PrintFinishing = "stapleDualTop"
	PrintFinishing_StapleTopLeft      PrintFinishing = "stapleTopLeft"
	PrintFinishing_StapleTopRight     PrintFinishing = "stapleTopRight"
	PrintFinishing_StapleTripleBottom PrintFinishing = "stapleTripleBottom"
	PrintFinishing_StapleTripleLeft   PrintFinishing = "stapleTripleLeft"
	PrintFinishing_StapleTripleRight  PrintFinishing = "stapleTripleRight"
	PrintFinishing_StapleTripleTop    PrintFinishing = "stapleTripleTop"
	PrintFinishing_StitchBottomEdge   PrintFinishing = "stitchBottomEdge"
	PrintFinishing_StitchEdge         PrintFinishing = "stitchEdge"
	PrintFinishing_StitchLeftEdge     PrintFinishing = "stitchLeftEdge"
	PrintFinishing_StitchRightEdge    PrintFinishing = "stitchRightEdge"
	PrintFinishing_StitchTopEdge      PrintFinishing = "stitchTopEdge"
	PrintFinishing_Trim               PrintFinishing = "trim"
	PrintFinishing_TrimAfterCopies    PrintFinishing = "trimAfterCopies"
	PrintFinishing_TrimAfterDocuments PrintFinishing = "trimAfterDocuments"
	PrintFinishing_TrimAfterJob       PrintFinishing = "trimAfterJob"
	PrintFinishing_TrimAfterPages     PrintFinishing = "trimAfterPages"
)

func PossibleValuesForPrintFinishing() []string {
	return []string{
		string(PrintFinishing_Bale),
		string(PrintFinishing_Bind),
		string(PrintFinishing_BindBottom),
		string(PrintFinishing_BindLeft),
		string(PrintFinishing_BindRight),
		string(PrintFinishing_BindTop),
		string(PrintFinishing_BookletMaker),
		string(PrintFinishing_Coat),
		string(PrintFinishing_Cover),
		string(PrintFinishing_Fold),
		string(PrintFinishing_FoldAccordion),
		string(PrintFinishing_FoldDoubleGate),
		string(PrintFinishing_FoldEngineeringZ),
		string(PrintFinishing_FoldGate),
		string(PrintFinishing_FoldHalf),
		string(PrintFinishing_FoldHalfZ),
		string(PrintFinishing_FoldLeftGate),
		string(PrintFinishing_FoldLetter),
		string(PrintFinishing_FoldParallel),
		string(PrintFinishing_FoldPoster),
		string(PrintFinishing_FoldRightGate),
		string(PrintFinishing_FoldZ),
		string(PrintFinishing_Laminate),
		string(PrintFinishing_None),
		string(PrintFinishing_Punch),
		string(PrintFinishing_PunchBottomLeft),
		string(PrintFinishing_PunchBottomRight),
		string(PrintFinishing_PunchDualBottom),
		string(PrintFinishing_PunchDualLeft),
		string(PrintFinishing_PunchDualRight),
		string(PrintFinishing_PunchDualTop),
		string(PrintFinishing_PunchQuadBottom),
		string(PrintFinishing_PunchQuadLeft),
		string(PrintFinishing_PunchQuadRight),
		string(PrintFinishing_PunchQuadTop),
		string(PrintFinishing_PunchTopLeft),
		string(PrintFinishing_PunchTopRight),
		string(PrintFinishing_PunchTripleBottom),
		string(PrintFinishing_PunchTripleLeft),
		string(PrintFinishing_PunchTripleRight),
		string(PrintFinishing_PunchTripleTop),
		string(PrintFinishing_SaddleStitch),
		string(PrintFinishing_Staple),
		string(PrintFinishing_StapleBottomLeft),
		string(PrintFinishing_StapleBottomRight),
		string(PrintFinishing_StapleDualBottom),
		string(PrintFinishing_StapleDualLeft),
		string(PrintFinishing_StapleDualRight),
		string(PrintFinishing_StapleDualTop),
		string(PrintFinishing_StapleTopLeft),
		string(PrintFinishing_StapleTopRight),
		string(PrintFinishing_StapleTripleBottom),
		string(PrintFinishing_StapleTripleLeft),
		string(PrintFinishing_StapleTripleRight),
		string(PrintFinishing_StapleTripleTop),
		string(PrintFinishing_StitchBottomEdge),
		string(PrintFinishing_StitchEdge),
		string(PrintFinishing_StitchLeftEdge),
		string(PrintFinishing_StitchRightEdge),
		string(PrintFinishing_StitchTopEdge),
		string(PrintFinishing_Trim),
		string(PrintFinishing_TrimAfterCopies),
		string(PrintFinishing_TrimAfterDocuments),
		string(PrintFinishing_TrimAfterJob),
		string(PrintFinishing_TrimAfterPages),
	}
}

func (s *PrintFinishing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintFinishing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintFinishing(input string) (*PrintFinishing, error) {
	vals := map[string]PrintFinishing{
		"bale":               PrintFinishing_Bale,
		"bind":               PrintFinishing_Bind,
		"bindbottom":         PrintFinishing_BindBottom,
		"bindleft":           PrintFinishing_BindLeft,
		"bindright":          PrintFinishing_BindRight,
		"bindtop":            PrintFinishing_BindTop,
		"bookletmaker":       PrintFinishing_BookletMaker,
		"coat":               PrintFinishing_Coat,
		"cover":              PrintFinishing_Cover,
		"fold":               PrintFinishing_Fold,
		"foldaccordion":      PrintFinishing_FoldAccordion,
		"folddoublegate":     PrintFinishing_FoldDoubleGate,
		"foldengineeringz":   PrintFinishing_FoldEngineeringZ,
		"foldgate":           PrintFinishing_FoldGate,
		"foldhalf":           PrintFinishing_FoldHalf,
		"foldhalfz":          PrintFinishing_FoldHalfZ,
		"foldleftgate":       PrintFinishing_FoldLeftGate,
		"foldletter":         PrintFinishing_FoldLetter,
		"foldparallel":       PrintFinishing_FoldParallel,
		"foldposter":         PrintFinishing_FoldPoster,
		"foldrightgate":      PrintFinishing_FoldRightGate,
		"foldz":              PrintFinishing_FoldZ,
		"laminate":           PrintFinishing_Laminate,
		"none":               PrintFinishing_None,
		"punch":              PrintFinishing_Punch,
		"punchbottomleft":    PrintFinishing_PunchBottomLeft,
		"punchbottomright":   PrintFinishing_PunchBottomRight,
		"punchdualbottom":    PrintFinishing_PunchDualBottom,
		"punchdualleft":      PrintFinishing_PunchDualLeft,
		"punchdualright":     PrintFinishing_PunchDualRight,
		"punchdualtop":       PrintFinishing_PunchDualTop,
		"punchquadbottom":    PrintFinishing_PunchQuadBottom,
		"punchquadleft":      PrintFinishing_PunchQuadLeft,
		"punchquadright":     PrintFinishing_PunchQuadRight,
		"punchquadtop":       PrintFinishing_PunchQuadTop,
		"punchtopleft":       PrintFinishing_PunchTopLeft,
		"punchtopright":      PrintFinishing_PunchTopRight,
		"punchtriplebottom":  PrintFinishing_PunchTripleBottom,
		"punchtripleleft":    PrintFinishing_PunchTripleLeft,
		"punchtripleright":   PrintFinishing_PunchTripleRight,
		"punchtripletop":     PrintFinishing_PunchTripleTop,
		"saddlestitch":       PrintFinishing_SaddleStitch,
		"staple":             PrintFinishing_Staple,
		"staplebottomleft":   PrintFinishing_StapleBottomLeft,
		"staplebottomright":  PrintFinishing_StapleBottomRight,
		"stapledualbottom":   PrintFinishing_StapleDualBottom,
		"stapledualleft":     PrintFinishing_StapleDualLeft,
		"stapledualright":    PrintFinishing_StapleDualRight,
		"stapledualtop":      PrintFinishing_StapleDualTop,
		"stapletopleft":      PrintFinishing_StapleTopLeft,
		"stapletopright":     PrintFinishing_StapleTopRight,
		"stapletriplebottom": PrintFinishing_StapleTripleBottom,
		"stapletripleleft":   PrintFinishing_StapleTripleLeft,
		"stapletripleright":  PrintFinishing_StapleTripleRight,
		"stapletripletop":    PrintFinishing_StapleTripleTop,
		"stitchbottomedge":   PrintFinishing_StitchBottomEdge,
		"stitchedge":         PrintFinishing_StitchEdge,
		"stitchleftedge":     PrintFinishing_StitchLeftEdge,
		"stitchrightedge":    PrintFinishing_StitchRightEdge,
		"stitchtopedge":      PrintFinishing_StitchTopEdge,
		"trim":               PrintFinishing_Trim,
		"trimaftercopies":    PrintFinishing_TrimAfterCopies,
		"trimafterdocuments": PrintFinishing_TrimAfterDocuments,
		"trimafterjob":       PrintFinishing_TrimAfterJob,
		"trimafterpages":     PrintFinishing_TrimAfterPages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintFinishing(input)
	return &out, nil
}
