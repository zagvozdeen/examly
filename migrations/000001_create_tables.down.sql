DROP TABLE IF EXISTS user_experience;

DROP TABLE IF EXISTS question_tag;
DROP TABLE IF EXISTS tags;

DROP INDEX user_answers_test_session_id_index;
DROP TABLE IF EXISTS user_answers;

DROP INDEX test_sessions_deleted_at_not_null_index;
DROP INDEX test_sessions_uuid_unique_index;
DROP TABLE IF EXISTS test_sessions;

DROP INDEX questions_deleted_at_not_null_index;
DROP INDEX questions_uuid_unique_index;
DROP TABLE IF EXISTS questions;

DROP TABLE IF EXISTS modules;

DROP TABLE IF EXISTS courses;

ALTER TABLE files DROP CONSTRAINT files_created_by_foreign;
DROP TABLE IF EXISTS users;

DROP INDEX files_deleted_at_not_null_index;
DROP INDEX files_uuid_unique_index;
DROP TABLE IF EXISTS files;