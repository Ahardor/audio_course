package queries

import (
	"context"
	"fmt"
	"iotvisual/processor/domain/melody"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMelodyInput struct {
	MelodyName melody.ID
	Limit      int
	Offset     int
}

// GetMelody достает полную мелодию по уникальному названию.
// Решение: хранить одну мелодию целиком в одном документе.
// Обоснование: массивы звуков в мелодии константны и невелики, вставка не планируется.
// Референс: https://www.mongodb.com/docs/atlas/schema-suggestions/avoid-unbounded-arrays/.
func (q *Queries) GetMelody(ctx context.Context, input GetMelodyInput) (melody.Melody, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	var m melody.Melody
	if err := q.collection.FindOne(ctx,
		bson.D{
			{Key: "id", Value: input.MelodyName},
		},
		options.FindOne().SetProjection(bson.M{
			"$slice": []interface{}{input.Offset, input.Limit},
		}),
	).Decode(&m); err != nil {
		return melody.Melody{}, fmt.Errorf("decode: %w", err)
	}

	return m, nil
}

// {Key: "sounds", Value: bson.M{
// "$project": bson.M{
// 	"$slice": []interface{}{input.Offset, input.Limit},
// },
// }},
