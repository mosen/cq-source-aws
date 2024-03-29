# Table: aws_docdb_events

This table shows data for Amazon DocumentDB Events.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Event.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **categories_concat**, **date**, **source_arn**, **source_identifier**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|categories_concat|`utf8`|
|date|`timestamp[us, tz=UTC]`|
|event_categories|`list<item: utf8, nullable>`|
|message|`utf8`|
|source_arn|`utf8`|
|source_identifier|`utf8`|
|source_type|`utf8`|