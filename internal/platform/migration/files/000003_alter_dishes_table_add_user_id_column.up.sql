ALTER TABLE dishes
ADD COLUMN user_id int NOT NULL;

ALTER TABLE dishes
ADD CONSTRAINT dishes_fk_user_id FOREIGN KEY ("user_id") REFERENCES users("id");

