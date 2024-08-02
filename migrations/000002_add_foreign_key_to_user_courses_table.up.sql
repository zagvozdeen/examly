ALTER TABLE user_courses
    ADD CONSTRAINT user_courses_last_question_id_foreign FOREIGN KEY (last_question_id) REFERENCES user_questions (id);