/* date test */
CREATE TABLE goma_date_types (
  id                  BIGINT PRIMARY KEY,
  date_columns        DATE,
  time_columns        TIME,
  datetime_columns    DATETIME,
  timestamp_columns   TIMESTAMP
);
