CREATE TABLE users (
  id varchar(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email        VARCHAR(255) NOT NULL UNIQUE,
  password     VARCHAR(255) NOT NULL,
  premium bool DEFAULT TRUE NOT NULL,
  status bool DEFAULT TRUE NOT NULL,
  created_at timestamp DEFAULT now() NOT NULL
);