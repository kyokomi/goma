update goma_string_types set
    id = :id
,   text_columns = :text_columns
,   tinytext_columns = :tinytext_columns
,   mediumtext_columns = :mediumtext_columns
,   longtext_columns = :longtext_columns
,   char_columns = :char_columns
,   varchar_columns = :varchar_columns
,   enum_columns = :enum_columns
 where
    id = :id

