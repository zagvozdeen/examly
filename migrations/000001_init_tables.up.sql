CREATE TABLE IF NOT EXISTS files
(
    id         BIGSERIAL PRIMARY KEY,
    content    VARCHAR(255) NULL,
    deleted_at TIMESTAMP    NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL PRIMARY KEY,
    email       VARCHAR(255)                 NULL,
    first_name  VARCHAR(255)                 NULL,
    last_name   VARCHAR(255)                 NULL,
    role        VARCHAR(255)                 NOT NULL,
    password    VARCHAR(255)                 NULL,
    avatar_id   BIGINT references files (id) NULL,
    deleted_at  TIMESTAMP                    NULL,
    created_at  TIMESTAMP                    NOT NULL,
    updated_at  TIMESTAMP                    NOT NULL
);

CREATE TABLE IF NOT EXISTS courses
(
    id         BIGSERIAL PRIMARY KEY,
    uuid       uuid                         NOT NULL,
    name       VARCHAR(255)                 NOT NULL,
    user_id    BIGINT references users (id) not null,
    deleted_at TIMESTAMP                    NULL,
    created_at TIMESTAMP                    NOT NULL,
    updated_at TIMESTAMP                    NOT NULL
);

CREATE TABLE IF NOT EXISTS modules
(
    id         BIGSERIAL PRIMARY KEY,
    name       varchar(255)                   not null,
    course_id  BIGINT references courses (id) not null,
    deleted_at TIMESTAMP                      NULL,
    created_at TIMESTAMP                      NOT NULL,
    updated_at TIMESTAMP                      NOT NULL
);

CREATE TABLE IF NOT EXISTS questions
(
    id         BIGSERIAL PRIMARY KEY,
    content    text                           not null,
    type       VARCHAR(255)                   NOT NULL,
    course_id  BIGINT references courses (id) not null,
    module_id  BIGINT references modules (id) null,
    deleted_at TIMESTAMP                      NULL,
    created_at TIMESTAMP                      NOT NULL,
    updated_at TIMESTAMP                      NOT NULL
);

CREATE TABLE IF NOT EXISTS answers
(
    id          BIGSERIAL PRIMARY KEY,
    content     text                             not null,
    question_id BIGINT references questions (id) not null,
    is_true     bool                             not null,
    deleted_at  TIMESTAMP                        NULL,
    created_at  TIMESTAMP                        NOT NULL,
    updated_at  TIMESTAMP                        NOT NULL
);

CREATE TABLE IF NOT EXISTS user_courses
(
    id               BIGSERIAL PRIMARY KEY,
    uuid             uuid                         NOT NULL,
    name             VARCHAR(255)                 NOT NULL,
    user_id          BIGINT references users (id) not null,
    last_question_id BIGINT                       NULL,
    deleted_at       TIMESTAMP                    NULL,
    created_at       TIMESTAMP                    NOT NULL,
    updated_at       TIMESTAMP                    NOT NULL
);

CREATE TABLE IF NOT EXISTS user_modules
(
    id         BIGSERIAL PRIMARY KEY,
    name       varchar(255)                        not null,
    course_id  BIGINT references user_courses (id) not null,
    deleted_at TIMESTAMP                           NULL,
    created_at TIMESTAMP                           NOT NULL,
    updated_at TIMESTAMP                           NOT NULL
);

CREATE TABLE IF NOT EXISTS user_questions
(
    id         BIGSERIAL PRIMARY KEY,
    content    text                                not null,
    type       VARCHAR(255)                        NOT NULL,
    is_true    bool                                null,
    course_id  BIGINT references user_courses (id) not null,
    module_id  BIGINT references user_modules (id) null,
    deleted_at TIMESTAMP                           NULL,
    created_at TIMESTAMP                           NOT NULL,
    updated_at TIMESTAMP                           NOT NULL
);

CREATE TABLE IF NOT EXISTS user_answers
(
    id          BIGSERIAL PRIMARY KEY,
    content     text                                  not null,
    question_id BIGINT references user_questions (id) not null,
    is_true     bool                                  not null,
    is_chosen   bool                                  null,
    deleted_at  TIMESTAMP                             NULL,
    created_at  TIMESTAMP                             NOT NULL,
    updated_at  TIMESTAMP                             NOT NULL
);