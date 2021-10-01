CREATE VIEW view_teacher_courses AS
SELECT f.name AS faculty,
c.name AS course,
t.first_name,
t.last_name,
t.job_title
FROM teacher_courses tc
JOIN courses c ON c.id = tc.course_id
JOIN teachers t ON t.id = tc.teacher_id
INNER JOIN faculties f ON f.id = t.faculty_id;