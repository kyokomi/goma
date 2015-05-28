DROP TABLE IF EXISTS goma_string_types;
/* string test */
CREATE TABLE goma_string_types (
  id                 BIGINT PRIMARY KEY,
  text_columns       TEXT,
  tinytext_columns   TINYTEXT,
  mediumtext_columns MEDIUMTEXT,
  longtext_columns   LONGTEXT,
  char_columns       CHAR(8),
  varchar_columns    VARCHAR(255),
  enum_columns       ENUM('OPEN','CLOSE','FOREVER')
);
