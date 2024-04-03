package test

import (
	"context"
	"iotvisual/processor/domain/frequency"
	"iotvisual/processor/domain/melody"
	"iotvisual/processor/domain/note"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMelody(t *testing.T) {
	ctx := context.Background()
	q := NewQueries(GetTestCollection())

	res, err := q.GetMelody(ctx, melody.ID("meme"))
	require.NoError(t, err)
	require.Equal(t, melody.ID("meme"), res.ID)
}

func TestInsertMelodies(t *testing.T) {
	ctx := context.Background()
	q := NewQueries(GetTestCollection())

	melodies := []melody.Melody{
		{
			ID: "meme",
			Sounds: []melody.Sound{
				{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 1, DurationMS: 2},
				{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 2, DurationMS: 1},
			},
		},
		{
			ID: "lmao",
			Sounds: []melody.Sound{
				{Note: note.NoteAd, Octave: frequency.Octave(0), Serial: 1, DurationMS: 1},
				{Note: note.NoteC, Octave: frequency.Octave(7), Serial: 2, DurationMS: 3},
			},
		},
	}
	err := q.InsertMelodies(ctx, melodies)
	require.NoError(t, err)
}
