select
  id
, tinyint_columns
, bool_columns
, smallint_columns
, mediumint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
, double_columns
FROM
  goma_numeric_types
WHERE
  id = /* id */64

