DROP TABLE IF EXISTS goma_date_types;
/* date test */
CREATE TABLE goma_date_types (
  id                  BIGINT PRIMARY KEY,
  date_columns        DATE,
  datetime_columns    DATETIME,
  timestamp_columns   TIMESTAMP
);
