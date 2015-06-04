DROP TABLE IF EXISTS boards;

CREATE TABLE boards (
  id           serial       NOT NULL,
  name         varchar(255) NOT NULL,
  description  text         NOT NULL,
  created_at   timestamp    DEFAULT (now() at time zone 'utc'),
  updated_at   timestamp    DEFAULT (now() at time zone 'utc'),
  CONSTRAINT boards_pkey PRIMARY KEY(id)
);
