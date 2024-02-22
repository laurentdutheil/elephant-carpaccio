package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "elephant_carpaccio/domain"
)

func TestBacklogScore_should_be_zero_if_no_story_done(t *testing.T) {
	backlog := Backlog{
		{Id: "id-1", ValuePoint: Score{Point: 1}, Done: false},
		{Id: "id-2", ValuePoint: Score{Point: 1}, Done: false},
	}

	score := backlog.Score()

	assert.Equal(t, Score{Point: 0}, score)
}

func TestBacklogScore_should_be_only_score_of_story_done(t *testing.T) {
	backlog := Backlog{
		{Id: "id-1", ValuePoint: Score{Point: 1}, Done: true},
		{Id: "id-2", ValuePoint: Score{Point: 1}, Done: true},
		{Id: "id-3", ValuePoint: Score{Point: 1}, Done: false},
	}

	score := backlog.Score()

	assert.Equal(t, Score{Point: 2}, score)
}
