# Notes Module

This module handles note creation, updating, deletion, and retrieval.
Each note belongs to a user and supports tagging and archiving.

## Layers

- presentation/: HTTP handlers
- application/: Use-cases (CreateNote, UpdateNote)
- domain/: Entities and rules
- infrastructure/: Database, cache, file storage
