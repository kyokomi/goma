DROP TABLE IF EXISTS goma_date_types;

CREATE TABLE goma_date_types (
  id                  BIGINT PRIMARY KEY,
  date_columns        DATE,
  timestamp_columns   TIMESTAMP
);
