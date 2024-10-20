CREATE TABLE IF NOT EXISTS  chats (
  chat_id   text PRIMARY KEY,
  user_id_1 text NOT NULL,
  user_id_2 text NOT NULL,
  content text NOT NULL
);

