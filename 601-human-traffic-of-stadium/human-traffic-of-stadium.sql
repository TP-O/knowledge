# Write your MySQL query statement below
select * from Stadium s1
where s1.people >= 100
    and (
        (exists (select 1 from Stadium s2 where s2.id = s1.id - 1 and s2.people >= 100)
            and exists (select 1 from Stadium s3 where s3.id = s1.id - 2 and s3.people >= 100))
        or
        (exists (select 1 from Stadium s2 where s2.id = s1.id + 1 and s2.people >= 100)
            and exists (select 1 from Stadium s3 where s3.id = s1.id + 2 and s3.people >= 100))
        or
        (exists (select 1 from Stadium s2 where s2.id = s1.id + 1 and s2.people >= 100)
            and exists (select 1 from Stadium s3 where s3.id = s1.id -1 and s3.people >= 100))
    )
order by s1.visit_date;