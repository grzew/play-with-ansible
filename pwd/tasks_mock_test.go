package pwd

import "github.com/turkenh/play-with-ansible/pwd/types"

type mockTasks struct {
	schedule   func(s *types.Session)
	unschedule func(s *types.Session)
}

func (m *mockTasks) Schedule(s *types.Session) {
	if m.schedule != nil {
		m.schedule(s)
	}
}
func (m *mockTasks) Unschedule(s *types.Session) {
	if m.unschedule != nil {
		m.unschedule(s)
	}
}
