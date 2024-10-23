-- +goose Up
-- +goose StatementBegin
INSERT INTO meeting_state(id, state_description)
VALUES (0, 'Active'), (1, 'Cancelled'), (2, 'Archived');

INSERT INTO cuisine_type(id, type_description)
VALUES (0, 'FastFood'), (1, 'Russian'), (2, 'Eastern'), (3, 'Other');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd