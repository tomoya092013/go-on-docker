CREATE TABLE IF NOT EXISTS todos(
  id serial PRIMARY KEY,
  title TEXT,
  description TEXT,
  deadline DATE,
  status BOOLEAN,
  priority INTEGER
);