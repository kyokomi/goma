/* string test */
CREATE TABLE goma_string_types (
  id                 BIGINT PRIMARY KEY,
  text               TEXT,
  char_columns       CHAR(8),
  varchar_columns    VARCHAR(255)
);
