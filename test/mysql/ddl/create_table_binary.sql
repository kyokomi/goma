DROP TABLE IF EXISTS goma_binary_types;
/* binary test */
CREATE TABLE goma_binary_types (
  binary_id           BIGINT PRIMARY KEY,
  binary_columns      BINARY(3),
  tinyblob_columns    TINYBLOB,
  blob_columns        BLOB,
  mediumblob_columns  MEDIUMBLOB,
  longblob_columns    LONGBLOB,
  varbinary_columns   VARBINARY(10)
);
