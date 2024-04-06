package queries

import (
	"context"
	"fmt"
	"iotvisual/processor/internal/domain/melody"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetMelody достает полную мелодию по уникальному названию.
// Решение: хранить одну мелодию целиком в одном документе.
// Обоснование: массивы звуков в мелодии константны и невелики, вставка не планируется.
// Референс: https://www.mongodb.com/docs/atlas/schema-suggestions/avoid-unbounded-arrays/.
func (q *Queries) GetMelody(ctx context.Context, melodyName melody.ID) (melody.Melody, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var m melody.Melody
	if err := q.collection.FindOne(ctx, bson.D{
		{Key: "id", Value: melodyName},
	},
	).Decode(&m); err != nil {
		return melody.Melody{}, fmt.Errorf("decode: %w", err)
	}

	return m, nil
}

func (q *Queries) InsertMelodies(ctx context.Context, input []melody.Melody) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var mrsh []interface{}
	for _, v := range input {
		bytes, err := bson.Marshal(v)
		if err != nil {
			return err
		}
		mrsh = append(mrsh, bytes)
	}
	_, err := q.collection.InsertMany(ctx, mrsh)
	if err != nil {
		return err
	}
	return nil
}
