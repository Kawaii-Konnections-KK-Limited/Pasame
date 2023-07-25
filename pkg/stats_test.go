package pkg

import (
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestGetConnectedIPs(t *testing.T) {
	t.Run("Connected Users", func(t *testing.T) {
		ls := LogStorage{}
		s := Stats{}

		ls.SetLogsFromFile("..\\example.log")
		s.SetStorage(&ls)

		ti, _ := time.Parse(time.DateTime, "2023-07-22 13:54:53")

		r := s.GetConnectedIPs(ti, time.Minute*1)

		log.Debugf("%+v", r)
	})
}
