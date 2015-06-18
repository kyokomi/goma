select
  id
, bool_columns
, smallint_columns
, int_columns
, integer_columns
, serial_columns
, decimal_columns
, numeric_columns
, float_columns
FROM
  goma_numeric_types
WHERE
  id = :id

