package pkg

import (
	"analyzer/internal"
	"sort"
	"time"
)

type LogStorage struct {
	logs []internal.Log
}

func (a *LogStorage) sort() {
	sort.Slice(a.logs[:], func(i, j int) bool {
		return a.logs[i].Time.Before(a.logs[j].Time)
	})
}

func (a *LogStorage) SetLogsFromFile(path string) {
	logs := internal.ReadFromFile(path)
	p := internal.Parser{}
	l := p.ProcessList(logs)

	a.SetLogs(l)
}

func (a *LogStorage) SetLogs(logs []internal.Log) {
	a.logs = logs
	a.sort()
}

func (a *LogStorage) AddLog(log internal.Log) {
	a.logs = append(a.logs, log)
	a.sort()
}

func (a *LogStorage) LogsAt(at time.Time, offset time.Duration) []internal.Log {
	from := at.Add(-offset)
	to := at
	l := make([]internal.Log, 0)
	for i := 0; i < len(a.logs); i++ {
		if a.logs[i].Time.Before(from) {
			continue
		}
		if a.logs[i].Time.After(to) {
			break
		}

		l = append(l, a.logs[i])
	}

	return l
}
