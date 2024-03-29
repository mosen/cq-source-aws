# Table: aws_ec2_account_attributes

This table shows data for Amazon Elastic Compute Cloud (EC2) Account Attributes.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccountAttribute.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **attribute_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|partition|`utf8`|
|attribute_name|`utf8`|
|attribute_values|`json`|