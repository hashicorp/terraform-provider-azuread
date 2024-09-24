package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}

// UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId is a struct representing the Resource ID for a User Id Online Meeting Id Meeting Attendance Report Attendance Record
type UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId struct {
	UserId             string
	OnlineMeetingId    string
	AttendanceRecordId string
}

// NewUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID returns a new UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId struct
func NewUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(userId string, onlineMeetingId string, attendanceRecordId string) UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId {
	return UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{
		UserId:             userId,
		OnlineMeetingId:    onlineMeetingId,
		AttendanceRecordId: attendanceRecordId,
	}
}

// ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID parses 'input' into a UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
func ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(input string) (*UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively parses 'input' case-insensitively into a UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordIDInsensitively(input string) (*UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnlineMeetingId, ok = input.Parsed["onlineMeetingId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", input)
	}

	if id.AttendanceRecordId, ok = input.Parsed["attendanceRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attendanceRecordId", input)
	}

	return nil
}

// ValidateUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID checks that 'input' can be parsed as a User Id Online Meeting Id Meeting Attendance Report Attendance Record ID
func ValidateUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/meetingAttendanceReport/attendanceRecords/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.AttendanceRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingId"),
		resourceids.StaticSegment("meetingAttendanceReport", "meetingAttendanceReport", "meetingAttendanceReport"),
		resourceids.StaticSegment("attendanceRecords", "attendanceRecords", "attendanceRecords"),
		resourceids.UserSpecifiedSegment("attendanceRecordId", "attendanceRecordId"),
	}
}

// String returns a human-readable description of this User Id Online Meeting Id Meeting Attendance Report Attendance Record ID
func (id UserIdOnlineMeetingIdMeetingAttendanceReportAttendanceRecordId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Attendance Record: %q", id.AttendanceRecordId),
	}
	return fmt.Sprintf("User Id Online Meeting Id Meeting Attendance Report Attendance Record (%s)", strings.Join(components, "\n"))
}
