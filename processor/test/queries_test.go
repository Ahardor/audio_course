package test

import (
	"context"
	"iotvisual/processor/domain/melody"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMelody(t *testing.T) {
	ctx := context.Background()
	q := NewQueries(GetTestCollection())

	res, err := q.GetMelody(ctx, melody.ID("Smells Like Teen Spirit"))
	require.NoError(t, err)
	require.Equal(t, "Smells Like Teen Spirit", res.ID)
}
