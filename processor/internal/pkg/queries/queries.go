package queries

import (
	"context"
	"iotvisual/processor/internal/domain/frequency"
	"iotvisual/processor/internal/domain/melody"
	"iotvisual/processor/internal/domain/note"

	"go.mongodb.org/mongo-driver/mongo"
)

type Queries struct {
	collection *mongo.Collection
}

func New(c *mongo.Collection) *Queries {
	return &Queries{
		collection: c,
	}
}

func (q *Queries) SeedDatabase() error {
	m := melody.Melody{
		ID: melody.ID("Master of Puppets"),
		Sounds: []melody.Sound{
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 1, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 2, DurationMS: 3000},
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 3, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 4, DurationMS: 3000},
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 5, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 6, DurationMS: 3000},
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 7, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 8, DurationMS: 3000},
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 9, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 10, DurationMS: 3000},
			{Note: note.NoteA, Octave: frequency.Octave(2), Serial: 11, DurationMS: 2000},
			{Note: note.NoteB, Octave: frequency.Octave(3), Serial: 12, DurationMS: 3000},
		},
	}
	return q.InsertMelodies(context.Background(), []melody.Melody{m})
}
