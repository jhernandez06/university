CREATE VIEW view_teacher_courses AS
SELECT f.nombre AS facultad,
c.nombre AS course,
t.nombre,
t.apellido,
t.titulo
FROM teacher_courses tc
JOIN courses c ON c.id = tc.course_id
JOIN teachers t ON t.id = tc.teacher_id
INNER JOIN facultads f ON f.id = t.facultad_id;