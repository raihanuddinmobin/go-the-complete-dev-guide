package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"mobin.dev/internal/domain/notes/domain"
)

type NotesCache struct {
	client *redis.Client
}

func NewNotesCache(client *redis.Client) *NotesCache {
	return &NotesCache{client}
}

func (c NotesCache) SetNote(ctx context.Context, id int64, note *domain.Note) error {
	key := prefix() + fmt.Sprintf("%d", id)

	data, _ := json.Marshal(note)
	err := c.client.Set(ctx, key, data, time.Hour*2).Err()
	return err
}

func (c NotesCache) GetNote(ctx context.Context, id int64) (*domain.Note, error) {
	key := prefix() + fmt.Sprintf("%d", id)

	data, err := c.client.Get(ctx, key).Result()

	var note domain.Note
	if err := json.Unmarshal([]byte(data), &note); err != nil {
		return nil, err
	}

	return &note, err
}

func prefix() string {
	return "NOTES:"
}
