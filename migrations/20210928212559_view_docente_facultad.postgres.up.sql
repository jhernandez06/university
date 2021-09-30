CREATE VIEW  teachers_facultads as
select 
f.nombre as "facultad",
t.nombre,
t.apellido,
t.titulo
from teachers t
join facultads f on f.id = t.facultad_id;