DROP TABLE IF EXISTS goma_date_types;
/* date test */
CREATE TABLE goma_date_types (
  id                  BIGINT PRIMARY KEY,
  datetime_columns    DATETIME,
  timestamp_columns   TIMESTAMP
);
