package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
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

func (c *NotesCache) SetNote(ctx context.Context, id int64, note *domain.Note) error {
	key := prefix() + fmt.Sprintf("%d", id)

	data, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, time.Minute*2).Err()

}

func (c *NotesCache) GetNote(ctx context.Context, id int64) (*domain.Note, error) {
	key := prefix() + fmt.Sprintf("%d", id)

	data, err := c.client.Get(ctx, key).Result()

	var note domain.Note
	if err := json.Unmarshal([]byte(data), &note); err != nil {
		return nil, err
	}

	return &note, err
}

func (c *NotesCache) GetNotes(ctx context.Context, limit, offset int) ([]*domain.Note, error) {
	key := prefix() + fmt.Sprintf("limit:%d:offset:%d", limit, offset)

	data, err := c.client.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var storedNotes []*domain.Note
	if err := json.Unmarshal([]byte(data), &storedNotes); err != nil {
		return nil, err
	}

	return storedNotes, err
}

func (c *NotesCache) SetNotes(ctx context.Context, limit, offset int, notes []*domain.Note) error {
	key := prefix() + fmt.Sprintf("limit:%d:offset:%d", limit, offset)
	data, err := json.Marshal(notes)

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		return nil
	}

	return c.client.Set(ctx, key, data, time.Minute*2).Err()
}

func prefix() string {
	return "NOTES:"
}
