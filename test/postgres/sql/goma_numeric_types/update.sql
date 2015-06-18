update goma_numeric_types set
    id = :id
,   bool_columns = :bool_columns
,   smallint_columns = :smallint_columns
,   int_columns = :int_columns
,   integer_columns = :integer_columns
,   serial_columns = :serial_columns
,   decimal_columns = :decimal_columns
,   numeric_columns = :numeric_columns
,   float_columns = :float_columns
 where
    id = :id

