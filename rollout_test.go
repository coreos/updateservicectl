package main

import (
	"bytes"
	"testing"

	"github.com/coreos/updateservicectl/client/update/v1"
)

func eq(a, b *update.Rollout) bool {
	if a.AppId != b.AppId || a.GroupId != b.GroupId {
		return false
	}
	if len(a.Rollout) != len(b.Rollout) {
		return false
	}
	for i, fa := range a.Rollout {
		fb := b.Rollout[i]
		if fa.Percent != fb.Percent {
			return false
		}
		if fa.Duration != fb.Duration {
			return false
		}
	}
	return true
}

func TestLinearRolloutGeneration(t *testing.T) {
	appId := "e96281a6-d1af-4bde-9a0a-97b76e56dc57"
	groupId := "stable"
	truth := &update.Rollout{
		AppId:   appId,
		GroupId: groupId,
		Rollout: []*update.Frame{
			&update.Frame{
				Percent:  10.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  20.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  30.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  40.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  50.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  60.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  70.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  80.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  90.0,
				Duration: 10,
			},
			&update.Frame{
				Percent:  100.0,
				Duration: 10,
			},
			// final frame to force 100%
			&update.Frame{
				Percent:  100.0,
				Duration: 0,
			},
		},
	}

	rollout := generateLinear(appId, groupId, 10, 100)

	if !eq(rollout, truth) {
		// generate nice frame output
		var e, g bytes.Buffer
		displayRollout(&e, truth)
		displayRollout(&g, rollout)
		t.Error("incorrect rollout generated.\n")
		t.Errorf("expected:\n%s\n", e.String())
		t.Errorf("got:\n%s\n", g.String())
	}
}
