package ics

import (
	"bytes"
	"encoding/hex"
	"strings"
	"text/template"
	"time"
)

func Generate(prodId string, events ...*Event) (string, error) {
	obj := &generator{
		ProdId: prodId,
		Events: []string{},
	}

	eventTmpl, err := template.New("events").Parse(vevent)
	if err != nil {
		return "", err
	}
	for _, event := range events {

		for idx := range event.Attendees {
			if event.Attendees[idx].Rsvp == "" {
				event.Attendees[idx].Rsvp = Rsvp_False
			}
		}

		e := &vEvent{
			Event:       event,
			DtStamp:     FormatDateTime(time.Now()),
			DtEnd:       FormatDateTime(event.DtEnd),
			DtStart:     FormatDateTime(event.DtStart),
			ExDate:      make([]string, len(event.ExDate)),
			Description: strings.Join(strings.Split(event.Description, "\n"), `\n`),
		}
		for i, exd := range event.ExDate {
			e.ExDate[i] = FormatDateTime(exd)
		}
		event.dtStamp = e.DtStamp
		event.UID = hex.EncodeToString([]byte(e.UID))

		buf := &bytes.Buffer{}
		if err := eventTmpl.Execute(buf, e); err != nil {
			return "", err
		}

		obj.Events = append(obj.Events, buf.String())
	}

	buf := &bytes.Buffer{}
	icsTmpl, err := template.New("ics").Parse(ics)
	if err != nil {
		return "", err
	}
	if err := icsTmpl.Execute(buf, obj); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (event *Event) Generate(prodId string) (string, error) {
	obj := &generator{
		ProdId: prodId,
		Events: []string{},
	}

	for idx := range event.Attendees {
		if event.Attendees[idx].Rsvp == "" {
			event.Attendees[idx].Rsvp = Rsvp_False
		}
	}

	e := &vEvent{
		Event:       event,
		DtStamp:     FormatDateTime(time.Now()),
		DtEnd:       FormatDateTime(event.DtEnd),
		DtStart:     FormatDateTime(event.DtStart),
		ExDate:      make([]string, len(event.ExDate)),
		Description: strings.Join(strings.Split(event.Description, "\n"), `\n`),
	}
	for i, exd := range event.ExDate {
		e.ExDate[i] = FormatDateTime(exd)
	}
	event.dtStamp = e.DtStamp
	event.UID = hex.EncodeToString([]byte(e.UID))

	eventTmpl, err := template.New("events").Parse(vevent)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := eventTmpl.Execute(buf, e); err != nil {
		return "", err
	}

	obj.Events = append(obj.Events, buf.String())

	buf = &bytes.Buffer{}
	icsTmpl, err := template.New("ics").Parse(ics)
	if err != nil {
		return "", err
	}
	if err := icsTmpl.Execute(buf, obj); err != nil {
		return "", err
	}

	return buf.String(), nil
}

type generator struct {
	ProdId string
	Events []string
}

func FormatDateTime(t time.Time) string {
	return t.Format("20060102T150405")
}
