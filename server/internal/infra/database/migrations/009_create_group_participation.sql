-- Write your migrate up statements here

/*
type GroupParticipation struct {
	GroupID   uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt *time.Time
}
*/

CREATE TABLE IF NOT EXISTS "group_participation" (
    group_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (group_id, user_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS "group_participation";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
