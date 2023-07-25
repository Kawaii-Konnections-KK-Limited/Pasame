package pkg

import (
	"time"

	"golang.org/x/exp/slices"
)

type Stats struct {
	storage *LogStorage
}

type ConnectedIPs struct {
	Owner string
	IPs   []string
}

func (s *Stats) SetStorage(ls *LogStorage) {
	s.storage = ls
}

func (s *Stats) GetOnlineUsers(at time.Time, offset time.Duration) (userList []string) {

	logsAt := s.storage.LogsAt(at, offset)

	for i := 0; i < len(logsAt); i++ {
		if slices.Contains(userList, logsAt[i].Email) {
			continue
		}
		userList = append(userList, logsAt[i].Email)
	}

	return userList
}

func (s *Stats) GetConnectedIPs(at time.Time, offset time.Duration) (c []ConnectedIPs) {
	return c
}
