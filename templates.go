package ics

const ics = `BEGIN:VCALENDAR
PRODID:{{.ProdId}}
VERSION:2.0
{{range $ve := .Events}}{{$ve}}
{{end}}END:VCALENDAR`

const vevent = `BEGIN:VEVENT
LOCATION:{{.Location}}
DTSTAMP:{{.DtStamp}}
DTSTART:{{.DtStart}}
DTEND:{{.DtEnd}}{{range $ru := .RRule}}
RRULE:{{$ru}}{{end}}
SUMMARY:{{.Summary}}
UID:{{.UID}}
END:VEVENT`

type vEvent struct {
	*Event
	DtStamp     string
	DtEnd       string
	DtStart     string
	ExDate      []string
	Description string
}
