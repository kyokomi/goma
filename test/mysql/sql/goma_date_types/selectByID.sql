select
  id
, datetime_columns
, timestamp_columns
FROM
  goma_date_types
WHERE
  id = :id

