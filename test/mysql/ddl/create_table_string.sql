DROP TABLE IF EXISTS goma_string_types;
/* string test */
CREATE TABLE goma_string_types (
  id                 BIGINT                          PRIMARY KEY,
  text_columns       TEXT                            NOT NULL,
  tinytext_columns   TINYTEXT                        NOT NULL,
  mediumtext_columns MEDIUMTEXT                      NOT NULL,
  longtext_columns   LONGTEXT                        NOT NULL,
  char_columns       CHAR(8)                         NOT NULL,
  varchar_columns    VARCHAR(255)                    NOT NULL,
  enum_columns       ENUM('OPEN','CLOSE','FOREVER')  NOT NULL
);
