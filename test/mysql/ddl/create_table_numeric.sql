DROP TABLE IF EXISTS goma_numeric_types;
/* numeric test */
CREATE TABLE goma_numeric_types (
  id                BIGINT     PRIMARY KEY,
  tinyint_columns   TINYINT    NOT NULL,
  bool_columns      BOOL       NOT NULL,
  smallint_columns  SMALLINT   NOT NULL,
  mediumint_columns MEDIUMINT  NOT NULL,
  int_columns       INT        NOT NULL,
  integer_columns   INTEGER    NOT NULL,
  serial_columns    SERIAL     NOT NULL,
  decimal_columns   DECIMAL    NOT NULL,
  numeric_columns   NUMERIC    NOT NULL,
  float_columns     FLOAT      NOT NULL,
  double_columns    DOUBLE     NOT NULL
);
