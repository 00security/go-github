package scrape

import (
	"strings"
	"time"

	"github.com/00security/go-github/v41/github"
	"github.com/PuerkitoBio/goquery"
)

type SSOSession struct {
	ID             *string           `json:"id"`
	RevocationURL  *string           `json:"revocation_url,omitempty"`
	Location       *string           `json:"location,omitempty"`
	DeviceTags     []string          `json:"device_tags,omitempty"`
	ExpirationTime *github.Timestamp `json:"expiration_time,omitempty"`
}

func (c *Client) SSOSessions(org string, user string) ([]*SSOSession, error) {
	var sessions []*SSOSession
	var err error

	doc, err := c.get("/orgs/%s/people/%s/sso", org, user)
	if err != nil {
		return nil, err
	}

	doc.Find("main li.js-user-session").Each(func(i int, s *goquery.Selection) {
		session := new(SSOSession)
		formField := s.Find("form")
		if formField != nil {
			revocationURL, found := formField.Attr("action")
			if found {
				split := strings.Split(revocationURL, "sso_session/")
				if len(split) < 2 {
					session.ID = nil
				} else {
					session.ID = &split[1]
				}
			}
			session.RevocationURL = &revocationURL
		}
		loc := strings.TrimSpace(s.Find("strong").First().Text())
		session.Location = &loc

		browserField := s.Find(".d-block strong")
		if browserField != nil {
			deviceString, found := browserField.Attr("title")
			if found {
				split := strings.Split(deviceString, " ")
				session.DeviceTags = split
			}
		}

		expirationField := s.Find(".note relative-time")
		if expirationField != nil {
			expiration, found := expirationField.Attr("datetime")
			if found {
				var expirationTime time.Time
				expirationTime, err = time.Parse("2006-01-02T15:04:05Z", expiration)
				if err != nil {
					expirationTime = time.Time{}
				}
				ghTimestamp := github.Timestamp{Time: expirationTime}
				session.ExpirationTime = &ghTimestamp
			}
		}

		sessions = append(sessions, session)
	})

	return sessions, nil
}
