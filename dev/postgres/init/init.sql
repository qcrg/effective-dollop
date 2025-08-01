CREATE TABLE coins (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL UNIQUE,
  observe BOOLEAN NOT NULL DEFAULT TRUE
);
INSERT INTO coins(name) VALUES ('bitcoin');
INSERT INTO coins(name) VALUES ('etherium');

CREATE TABLE currencies (
  id SERIAL PRIMARY KEY,
  name CHAR(3) NOT NULL UNIQUE
);
INSERT INTO currencies(name) VALUES ('usd');

CREATE TABLE records (
  id BIGSERIAL PRIMARY KEY,
  coin_id INT REFERENCES coins(id) NOT NULL,
  currency_id INT REFERENCES currencies(id) NOT NULL DEFAULT 1,
  timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  value BIGINT NOT NULL
);
INSERT INTO records(coin_id, currency_id, timestamp, value) VALUES (1, 1, '2025-08-01T10:54:00+00:00', 123);
INSERT INTO records(coin_id, currency_id, timestamp, value) VALUES (1, 1, '2025-08-01T10:55:00+00:00', 865);
