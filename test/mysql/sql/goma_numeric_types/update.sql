update goma_numeric_types set
    id = :id
,   tinyint_columns = :tinyint_columns
,   bool_columns = :bool_columns
,   smallint_columns = :smallint_columns
,   mediumint_columns = :mediumint_columns
,   int_columns = :int_columns
,   integer_columns = :integer_columns
,   serial_columns = :serial_columns
,   decimal_columns = :decimal_columns
,   numeric_columns = :numeric_columns
,   float_columns = :float_columns
,   double_columns = :double_columns
 where
    id = :id

