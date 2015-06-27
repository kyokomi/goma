DROP TABLE IF EXISTS goma_date_types;

CREATE TABLE goma_date_types (
  id                      BIGINT    PRIMARY KEY,
  timestamp_columns       TIMESTAMP NOT NULL,
  nil_timestamp_columns   TIMESTAMP
);
