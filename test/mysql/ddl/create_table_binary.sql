DROP TABLE IF EXISTS goma_binary_types;
/* binary test */
CREATE TABLE goma_binary_types (
  binary_id           BIGINT        PRIMARY KEY,
  binary_columns      BINARY(3)     NOT NULL,
  tinyblob_columns    TINYBLOB      NOT NULL,
  blob_columns        BLOB          NOT NULL,
  mediumblob_columns  MEDIUMBLOB    NOT NULL,
  longblob_columns    LONGBLOB      NOT NULL,
  varbinary_columns   VARBINARY(10) NOT NULL
);
