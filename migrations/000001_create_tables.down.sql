DROP INDEX user_answers_test_session_id_index;
DROP TABLE IF EXISTS user_answers;

DROP INDEX test_sessions_deleted_at_not_null_index;
DROP INDEX test_sessions_uuid_unique_index;
DROP TABLE IF EXISTS test_sessions;

DROP INDEX questions_deleted_at_not_null_index;
DROP INDEX questions_uuid_unique_index;
DROP TABLE IF EXISTS questions;

DROP INDEX modules_deleted_at_not_null_index;
DROP INDEX modules_uuid_unique_index;
DROP TABLE IF EXISTS modules;

DROP INDEX courses_deleted_at_not_null_index;
DROP INDEX courses_uuid_unique_index;
DROP TABLE IF EXISTS courses;

DROP INDEX users_deleted_at_not_null_index;
DROP INDEX users_uuid_unique_index;
ALTER TABLE files DROP CONSTRAINT files_created_by_foreign;
DROP TABLE IF EXISTS users;

DROP INDEX files_deleted_at_not_null_index;
DROP INDEX files_uuid_unique_index;
DROP TABLE IF EXISTS files;