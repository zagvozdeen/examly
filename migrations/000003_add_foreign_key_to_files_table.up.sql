ALTER TABLE files
    ADD CONSTRAINT files_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id);