CREATE VIEW  teacher_faculties as
select 
f.name as "faculty",
t.first_name,
t.last_name,
t.job_title
from teachers t
join faculties f on f.id = t.faculty_id;