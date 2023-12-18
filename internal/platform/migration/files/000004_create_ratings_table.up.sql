CREATE TABLE IF NOT EXISTS ratings (
  user_id INT NOT NULL,
  dish_id INT NOT NULL,
  rate SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL,

  CONSTRAINT "ratings_pkey" PRIMARY KEY ("user_id","dish_id"),
  CONSTRAINT "ratings_fk_user_id" FOREIGN KEY ("user_id") REFERENCES users("id"),
  CONSTRAINT "ratings_fk_dish_id" FOREIGN KEY ("dish_id") REFERENCES dishes("id")
)
