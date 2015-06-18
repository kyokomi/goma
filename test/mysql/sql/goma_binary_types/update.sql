update goma_binary_types set
    binary_id = :binary_id
,   binary_columns = :binary_columns
,   tinyblob_columns = :tinyblob_columns
,   blob_columns = :blob_columns
,   mediumblob_columns = :mediumblob_columns
,   longblob_columns = :longblob_columns
,   varbinary_columns = :varbinary_columns
 where
    binary_id = :binary_id

