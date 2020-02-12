CREATE TABLE sites (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  url VARCHAR (355),
  status INT NOT NULL,
  short_name VARCHAR (50),
  description VARCHAR (355),
  enabled BOOLEAN NOT NULL
);