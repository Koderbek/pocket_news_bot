CREATE TABLE IF NOT EXISTS chat_category
(
    chat_id     INTEGER NOT NULL,
    category_id SMALLINT NOT NULL REFERENCES category (id) ON DELETE CASCADE,
    name        VARCHAR(255)
);

ALTER TABLE chat_category ADD CONSTRAINT chat_cat_uniq UNIQUE (chat_id, category_id);
CREATE INDEX chat_idx ON chat_category (chat_id);
CREATE INDEX category_idx ON chat_category (category_id);