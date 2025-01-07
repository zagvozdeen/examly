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
    id                 SERIAL PRIMARY KEY,
    uuid               uuid                          not null,
    email              VARCHAR(255)                  NULL,
    first_name         VARCHAR(255)                  NULL,
    last_name          VARCHAR(255)                  NULL,
    role               VARCHAR(50)                   NOT NULL,
    password           VARCHAR(255)                  NULL,
    avatar_id          INTEGER references files (id) NULL,
    description        TEXT                          NULL,
    company_name       VARCHAR(255)                  NULL,
    contact            VARCHAR(255)                  NULL,
    account            INTEGER                       NOT NULL DEFAULT 0,
    can_view_referrals BOOLEAN                       NOT NULL DEFAULT false,
    deleted_at         TIMESTAMP                     NULL,
    created_at         TIMESTAMP                     NOT NULL,
    updated_at         TIMESTAMP                     NOT NULL
);

ALTER TABLE files
    ADD CONSTRAINT files_created_by_foreign FOREIGN KEY (created_by) REFERENCES users (id);

CREATE TABLE IF NOT EXISTS courses
(
    id                SERIAL PRIMARY KEY,
    uuid              uuid                          NOT NULL,
    name              VARCHAR(255)                  NOT NULL,
    description       text                          NOT NULL,
    color             VARCHAR(50)                   NOT NULL,
    icon              VARCHAR(50)                   NOT NULL,
    status            VARCHAR(50)                   NOT NULL,
    moderation_reason TEXT                          null,
    created_by        INTEGER references users (id) not null,
    moderated_by      INTEGER references users (id) null,
    deleted_at        TIMESTAMP                     NULL,
    created_at        TIMESTAMP                     NOT NULL,
    updated_at        TIMESTAMP                     NOT NULL
);

CREATE TABLE IF NOT EXISTS modules
(
    id                SERIAL PRIMARY KEY,
    uuid              uuid                            NOT NULL,
    name              varchar(255)                    not null,
    status            VARCHAR(50)                     NOT NULL,
    moderation_reason TEXT                            null,
    course_id         INTEGER references courses (id) not null,
    created_by        INTEGER references users (id)   not null,
    moderated_by      INTEGER references users (id)   null,
    deleted_at        TIMESTAMP                       NULL,
    created_at        TIMESTAMP                       NOT NULL,
    updated_at        TIMESTAMP                       NOT NULL
);

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

CREATE TABLE IF NOT EXISTS tags
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO tags (id, name)
VALUES (1, 'Разработка ПО'),
       (2, 'Frontend'),
       (3, 'Backend'),
       (4, 'Проектирование баз данных'),
       (5, 'Разработка под iOS'),
       (6, 'Разработка под Android'),
       (7, 'System Design'),
       (8, 'Аналитика данных и ИИ'),
       (9, 'Разработка десктопных приложений'),
       (10, 'Создание игр'),
       (11, 'Системное администрирование и DevOps'),
       (12, 'Управление серверами и сетями'),
       (13, 'Контейнеризация и оркестрация (Docker, Kubernetes)'),
       (14, 'Непрерывная интеграция и развертывание (CI/CD)'),
       (15, 'Информационная безопасность'),
       (16, 'Кибербезопасность'),
       (17, 'Защита данных и нормативное соответствие'),
       (18, 'Управление угрозами и уязвимостями'),
       (19, 'Анализ данных и бизнес-аналитика'),
       (20, 'Машинное обучение и глубокое обучение'),
       (21, 'Обработка естественного языка'),
       (22, 'Сетевые технологии'),
       (23, 'Сетевое администрирование'),
       (24, 'Протоколы и сетевые архитектуры'),
       (25, 'Базы данных и хранение данных'),
       (26, 'Администрирование БД'),
       (27, 'Управление большими данными'),
       (28, 'Облачные технологии'),
       (29, 'Облачные платформы (AWS, Azure, Google Cloud)'),
       (30, 'Архитектура облачных решений'),
       (31, 'Управление облачными услугами'),
       (32, 'Проектный менеджмент в IT'),
       (33, 'Управление продуктом'),
       (34, 'Управление командами разработки'),
       (35, 'Разработка встроенных систем'),
       (36, 'Программирование микроконтроллеров'),
       (37, 'Создание IoT-устройств'),
       (38, 'Разработка VR и AR приложений'),
       (39, 'Технологии блокчейна и криптовалюта');

CREATE TABLE IF NOT EXISTS question_tag
(
    question_id INTEGER references questions (id) NOT NULL,
    tag_id      INTEGER references tags (id)      NOT NULL,
    PRIMARY KEY (question_id, tag_id)
);

CREATE TABLE IF NOT EXISTS user_experience
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id) NOT NULL,
    one        SMALLINT                      NOT NULL,
    two        SMALLINT                      NOT NULL,
    three      SMALLINT                      NOT NULL,
    four       VARCHAR(255)                  NOT NULL,
    five       SMALLINT                      NOT NULL,
    six        SMALLINT                      NOT NULL,
    seven      VARCHAR(255)                  NOT NULL,
    eight      VARCHAR(255)                  NOT NULL,
    nine       SMALLINT                      NOT NULL,
    ten        VARCHAR(255)                  NOT NULL,
    eleven     SMALLINT                      NOT NULL,
    twelve     VARCHAR(255)                  NOT NULL,
    thirteen   VARCHAR(255)                  NOT NULL,
    deleted_at TIMESTAMP                     NULL,
    created_at TIMESTAMP                     NOT NULL,
    updated_at TIMESTAMP                     NOT NULL
);