select
  id
, timestamp_columns
, nil_timestamp_columns
FROM
  goma_date_types
WHERE
  id = :id

