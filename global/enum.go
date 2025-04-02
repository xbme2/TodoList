package global

type ImportanceType int
type CalendarService int

const (
	Light ImportanceType = iota
	Normal
	Important
	Critical
)

const (
	GoogleCalendar CalendarService = iota // 0: Google Calendar
	OutlookAPI                            // 1: Outlook API
)
