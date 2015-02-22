/* numeric test */
CREATE TABLE goma_numeric_types (
  id                BIGINT   PRIMARY KEY,
  tinyint_columns   TINYINT,   
  bool_columns      BOOL,   
  smallint_columns  SMALLINT,   
  mediumint_columns MEDIUMINT,   
  int_columns       INT,   
  integer_columns   INTEGER,  
  serial_columns    SERIAL,   
  decimal_columns   DECIMAL,   
  numeric_columns   NUMERIC,   
  float_columns     FLOAT,   
  double_columns    DOUBLE 
);
