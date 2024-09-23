package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}

// MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a Me Online Meeting Id Meeting Attendance Report Attendance Record
type MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId struct {
	OnlineMeetingId    string
	AttendanceRecordId string
}

// NewMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID returns a new MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId struct
func NewMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(onlineMeetingId string, attendanceRecordId string) MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId {
	return MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{
		OnlineMeetingId:    onlineMeetingId,
		AttendanceRecordId: attendanceRecordId,
	}
}

// ParseMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID parses 'input' into a MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
func ParseMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(input string) (*MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.AttendanceRecordId, ok = input.Parsed["attendanceRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", input)
	}

	return nil
}

// ValidateMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a Me Online Meeting Id Meeting Attendance Report Attendance Record ID
func ValidateMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/me/onlineMeetings/%s/meetingAttendanceReport/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("meetingAttendanceReport", "meetingAttendanceReport", "meetingAttendanceReport"),
		resourceids.StaticSegment("attendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordId"),
	}
}

// String returns a human-readable description of this Me Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id MeOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("Me Online Meeting Id Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
