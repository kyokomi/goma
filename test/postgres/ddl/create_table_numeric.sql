DROP TABLE IF EXISTS goma_numeric_types;

CREATE TABLE goma_numeric_types (
  id                BIGINT   PRIMARY KEY,
  bool_columns      BOOL,
  smallint_columns  SMALLINT,   
  int_columns       INT,
  integer_columns   INTEGER,  
  serial_columns    SERIAL,   
  decimal_columns   DECIMAL,   
  numeric_columns   NUMERIC,   
  float_columns     FLOAT
);
