insert into aws_policy_results
with patch_compliance_status_groups as(
    select DISTINCT
        instance_arn,
        status
    from
        aws_ssm_instance_compliance_items
    where
        compliance_type = 'Patch'
)
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'Amazon EC2 instances managed by Systems Manager should have a patch compliance status of COMPLIANT after a patch installation' as title,
    aws_ssm_instances.account_id,
    aws_ssm_instances.arn,
    case when
        patch_compliance_status_groups.status is distinct from 'COMPLIANT'
     then 'fail' else 'pass' end as status
 from
     aws_ssm_instances
INNER join patch_compliance_status_groups 
    on aws_ssm_instances.arn = patch_compliance_status_groups.instance_arn