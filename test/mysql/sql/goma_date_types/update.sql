update goma_date_types set
    id = :id
,   datetime_columns = :datetime_columns
,   timestamp_columns = :timestamp_columns
,   nil_datetime_columns = :nil_datetime_columns
 where
    id = :id

