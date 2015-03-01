drop table goma_string_types;
/* string test */
CREATE TABLE goma_string_types (
  id                 BIGINT PRIMARY KEY,
  text_columns       TEXT,
  char_columns       CHAR(8),
  varchar_columns    VARCHAR(255)
);
