-- +goose Up
-- +goose StatementBegin
CREATE TABLE meeting_state (
    id int PRIMARY KEY NOT NULL,
    state_description VARCHAR(20) NOT NULL
);

CREATE TABLE  meetings (
                  id UUID PRIMARY KEY NOT NULL,
                  gathering_place_id UUID REFERENCES gathering_places(id),
                  initiators_id  UUID REFERENCES users(id),
                  time_start     timestamp with time zone,
                  time_end       timestamp with time zone,
                  max_participants  int,
                  state        int REFERENCES meeting_state(id)
);
CREATE TABLE meetings_history (
                                  user_id UUID REFERENCES users(id),
                                  meeting_id UUID REFERENCES meetings(id)
);
ALTER TABLE users
ADD current_meeting_id UUID REFERENCES meetings(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS users
DROP COLUMN current_meeting_id;
DROP TABLE IF EXISTS meetings_history;
DROP TABLE IF EXISTS meetings;
DROP TABLE IF EXISTS meeting_state;
-- +goose StatementEnd
