# Table: aws_availability_zones

This table shows data for Availability Zones.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AvailabilityZone.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region_name**, **zone_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|enabled|`bool`|
|partition|`utf8`|
|region|`utf8`|
|group_name|`utf8`|
|messages|`json`|
|network_border_group|`utf8`|
|opt_in_status|`utf8`|
|parent_zone_id|`utf8`|
|parent_zone_name|`utf8`|
|region_name|`utf8`|
|state|`utf8`|
|zone_id|`utf8`|
|zone_name|`utf8`|
|zone_type|`utf8`|