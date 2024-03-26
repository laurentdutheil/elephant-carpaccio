package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestDefaultBacklogStoriesAreNotDoneAtBeginning(t *testing.T) {
	backlog := DefaultBacklog()

	for _, story := range backlog {
		assert.False(t, story.IsDone())
	}
}
