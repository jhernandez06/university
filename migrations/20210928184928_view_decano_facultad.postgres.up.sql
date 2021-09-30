create view decano_facultads as
select d.nombre, d.apellido, f.nombre as facultad_nombre,  f.ubicacion, f.numero
from facultads f
join decanoes d on d.id = f.decano_id;