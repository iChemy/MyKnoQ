package domain

import "github.com/google/uuid"

type Event struct {
	ID          uuid.UUID
	Name        string
	Description string
	Open        bool
	// 参加対象グループ
	ParticipantGroup Group
	// 参加者の出欠
	// Group の Members と同じ長さであるべき (Groupの更新のタイミングによっては異なることがある)
	UserAttendances []UserAttendance
	Tags            []Tag
	CreatedBy       User
	Venue           Venue
	Model
}

type AttendanceState int

const (
	Unconfirmed AttendanceState = iota
	Attendance
	Absense
)

type UserAttendance struct {
	User            User
	AttendanceState AttendanceState
}
