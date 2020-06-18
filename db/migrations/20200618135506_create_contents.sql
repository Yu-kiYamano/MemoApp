
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE contents (
    id int AUTO_INCREMENT,
    content varchar(100),
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE contents;