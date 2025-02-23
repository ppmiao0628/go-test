package remote_pro_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestRemotePro(t *testing.T) {
	m := cm.CreateConcurrentMap(10)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
