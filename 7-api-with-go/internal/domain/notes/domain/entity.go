package domain

import "time"

type Note struct {
	title       string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}
