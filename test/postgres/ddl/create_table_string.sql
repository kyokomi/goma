DROP TABLE IF EXISTS goma_string_types;

CREATE TABLE goma_string_types (
  id                 BIGINT        PRIMARY KEY,
  text_columns       TEXT          NOT NULL,
  char_columns       CHAR(8)       NOT NULL,
  varchar_columns    VARCHAR(255)  NOT NULL
);
