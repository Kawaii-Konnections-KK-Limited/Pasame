package internal

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type Log struct {
	Time     time.Time
	Protocol string
	SrcIP    string
	DestIP   string
	DestPort string
	Inbound  string
	Outbound string
	Email    string
}

type Parser struct {
	splitedLog []string
}

func (p *Parser) Process(log string) Log {
	log = strings.TrimSpace(log)
	p.splitedLog = strings.Split(log, " ")

	return Log{
		Time:     p.dateTime(),
		Protocol: p.protocol(),
		SrcIP:    p.srcIP(),
		DestIP:   p.destIP(),
		DestPort: p.destPort(),
		Inbound:  p.inbound(),
		Outbound: p.outbound(),
		Email:    p.email(),
	}
}

func (p *Parser) ProcessList(logs []string) []Log {
	l := make([]Log, len(logs))
	for i := 0; i < len(logs); i++ {
		l = append(l, p.Process(logs[i]))
	}
	return l
}

func (p *Parser) dateTime() time.Time {
	formatedDatetime := strings.Replace(p.splitedLog[0]+" "+p.splitedLog[1], "/", "-", -1)
	dateTime, err := time.Parse(time.DateTime, formatedDatetime)

	if err != nil {
		log.Errorf("Error parsing date-time: %+v", err)
	}

	return dateTime
}

func (p *Parser) srcIP() string {

	return strings.Split(p.splitedLog[2], ":")[0]
}

func (p *Parser) destIP() string {

	return strings.Split(p.splitedLog[4], ":")[1]
}

func (p *Parser) destPort() string {

	return strings.Split(p.splitedLog[4], ":")[2]
}

func (p *Parser) inbound() string {

	return p.splitedLog[5][1:]
}

func (p *Parser) outbound() string {
	s, _ := strings.CutSuffix(p.splitedLog[7], "]")

	return s
}

func (p *Parser) email() string {

	return p.splitedLog[len(p.splitedLog)-1]
}

func (p *Parser) protocol() string {

	return ""
}
