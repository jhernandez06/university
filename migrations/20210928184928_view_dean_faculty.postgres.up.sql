create view dean_faculties as
select d.first_name, d.last_name, f.name as faculty_name,  f.location, f.number
from faculties f
join deans d on d.id = f.dean_id;