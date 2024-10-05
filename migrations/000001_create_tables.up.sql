CREATE TABLE IF NOT EXISTS files
(
    id          SERIAL PRIMARY KEY,
    uuid        uuid         NOT NULL,
    content     VARCHAR(255) NOT NULL,
    size        INTEGER      NOT NULL,
    mime_type   VARCHAR(255) NOT NULL,
    origin_name VARCHAR(255) NOT NULL,
    created_by  INTEGER      NOT NULL,
    deleted_at  TIMESTAMP    NULL,
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP    NOT NULL
);

CREATE UNIQUE INDEX files_uuid_unique_index ON files (uuid);
CREATE INDEX files_deleted_at_not_null_index ON files (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    uuid       uuid                          not null,
    email      VARCHAR(255)                  NULL,
    first_name VARCHAR(255)                  NULL,
    last_name  VARCHAR(255)                  NULL,
    role       VARCHAR(50)                   NOT NULL,
    password   VARCHAR(255)                  NULL,
    avatar_id  INTEGER references files (id) NULL,
    deleted_at TIMESTAMP                     NULL,
    created_at TIMESTAMP                     NOT NULL,
    updated_at TIMESTAMP                     NOT NULL
);

ALTER TABLE files ADD CONSTRAINT files_created_by_foreign FOREIGN KEY (created_by) REFERENCES users (id);

CREATE UNIQUE INDEX users_uuid_unique_index ON users (uuid);
CREATE INDEX users_deleted_at_not_null_index ON users (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE IF NOT EXISTS courses
(
    id          SERIAL PRIMARY KEY,
    uuid        uuid                          NOT NULL,
    name        VARCHAR(255)                  NOT NULL,
    description text                          NOT NULL,
    color       VARCHAR(50)                   NOT NULL,
    icon        VARCHAR(50)                   NOT NULL,
    status      VARCHAR(50)                   NOT NULL,
    created_by  INTEGER references users (id) not null,
    deleted_at  TIMESTAMP                     NULL,
    created_at  TIMESTAMP                     NOT NULL,
    updated_at  TIMESTAMP                     NOT NULL
);

CREATE UNIQUE INDEX courses_uuid_unique_index ON courses (uuid);
CREATE INDEX courses_deleted_at_not_null_index ON courses (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE IF NOT EXISTS modules
(
    id         SERIAL PRIMARY KEY,
    uuid       uuid                            NOT NULL,
    name       varchar(255)                    not null,
    status     VARCHAR(50)                     NOT NULL,
    course_id  INTEGER references courses (id) not null,
    created_by INTEGER references users (id)   not null,
    deleted_at TIMESTAMP                       NULL,
    created_at TIMESTAMP                       NOT NULL,
    updated_at TIMESTAMP                       NOT NULL
);

CREATE UNIQUE INDEX modules_uuid_unique_index ON modules (uuid);
CREATE INDEX modules_deleted_at_not_null_index ON modules (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE IF NOT EXISTS questions
(
    id                SERIAL PRIMARY KEY,
    uuid              uuid                              NOT NULL,
    title             text                              not null,
    content           text                              null,
    explanation       text                              null,
    moderation_reason TEXT                              null,
    type              VARCHAR(50)                       NOT NULL,
    status            VARCHAR(50)                       NOT NULL,
    course_id         INTEGER references courses (id)   not null,
    module_id         INTEGER references modules (id)   null,
    created_by        INTEGER references users (id)     not null,
    moderated_by      INTEGER references users (id)     null,
    prev_question_id  INTEGER references questions (id) null,
    next_question_id  INTEGER references questions (id) null,
    options           jsonb                             not null,
    deleted_at        TIMESTAMP                         NULL,
    created_at        TIMESTAMP                         NOT NULL,
    updated_at        TIMESTAMP                         NOT NULL
);

CREATE UNIQUE INDEX questions_uuid_unique_index ON questions (uuid);
CREATE INDEX questions_deleted_at_not_null_index ON questions (deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX questions_course_id_index ON questions (course_id);

CREATE TABLE IF NOT EXISTS test_sessions
(
    id               SERIAL PRIMARY KEY,
    uuid             uuid                              NOT NULL,
    name             VARCHAR(255)                      NOT NULL,
    type             VARCHAR(50)                       NOT NULL,
    user_id          INTEGER references users (id)     not null,
    course_id        INTEGER references courses (id)   null,
    question_ids     integer[]                         not null,
    last_question_id INTEGER references questions (id) NULL,
    deleted_at       TIMESTAMP                         NULL,
    created_at       TIMESTAMP                         NOT NULL,
    updated_at       TIMESTAMP                         NOT NULL
);

CREATE UNIQUE INDEX test_sessions_uuid_unique_index ON test_sessions (uuid);
CREATE INDEX test_sessions_deleted_at_not_null_index ON test_sessions (deleted_at) WHERE deleted_at IS NOT NULL;

CREATE TABLE IF NOT EXISTS user_answers
(
    id              SERIAL PRIMARY KEY,
    test_session_id INTEGER references test_sessions (id) NOT NULL,
    question_id     INTEGER references questions (id)     NOT NULL,
    answer_data     JSONB                                 NOT NULL,
    is_correct      BOOLEAN                               not null,
    answered_at     TIMESTAMP                             not null
);

CREATE INDEX user_answers_test_session_id_index ON user_answers (test_session_id);