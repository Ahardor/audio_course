package test

import (
	"context"
	"iotvisual/processor/domain/frequency"
	"iotvisual/processor/domain/melody"
	"iotvisual/processor/domain/note"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetMelody(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	q := NewQueries(GetTestCollection())

	id := melody.ID(uuid.NewString())
	q.InsertMelodies(ctx, []melody.Melody{
		{
			ID: id,
			Sounds: []melody.Sound{
				{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 1, DurationMS: 2},
				{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 2, DurationMS: 1},
			},
		},
	})
	res, err := q.GetMelody(ctx, id)
	require.NoError(t, err)
	require.Equal(t, id, res.ID)
}

func TestInsertMelodies(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	q := NewQueries(GetTestCollection())

	id1 := melody.ID(uuid.NewString())
	id2 := melody.ID(uuid.NewString())
	melodies := []melody.Melody{
		{
			ID: id1,
			Sounds: []melody.Sound{
				{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 1, DurationMS: 2},
				{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 2, DurationMS: 1},
			},
		},
		{
			ID: id2,
			Sounds: []melody.Sound{
				{Note: note.NoteAd, Octave: frequency.Octave(0), Serial: 1, DurationMS: 1},
				{Note: note.NoteC, Octave: frequency.Octave(7), Serial: 2, DurationMS: 3},
			},
		},
	}
	err := q.InsertMelodies(ctx, melodies)
	require.NoError(t, err)
}
