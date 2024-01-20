INSERT INTO t_courses VALUES (1, 'Java core');
INSERT INTO t_courses VALUES (2, 'Go lang');
INSERT INTO t_courses VALUES (3, 'Backend');
INSERT INTO t_courses VALUES (4, 'Algorithms and Data structures');

INSERT INTO t_users VALUES (7, 'FatherOfCEOs', 'Yerbol99', 'Yerbol', 'Baigarayev', 'Yerbol99@gmail.com');
INSERT INTO t_users VALUES (8, 'tauke_m', 'Taukekhan24', 'Taukekhan', 'Mustakov', 'Taukekhan24@gmail.com');
INSERT INTO t_users VALUES (9, 'azamatserek', 'AzamatS', 'Azamat', 'Serek', 'AzamatS@gmail.com');

INSERT INTO t_users_roles VALUES(7, 2);
INSERT INTO t_users_roles VALUES(8, 2);
INSERT INTO t_users_roles VALUES(9, 2);

INSERT INTO t_teachers VALUES(1, 7);
INSERT INTO t_teachers VALUES(2, 8);
INSERT INTO t_teachers VALUES(3, 9);


INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (1, 1, 1, '1:12:2');
INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (2, 1, 1, '2:10:2');
INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (3, 1, 1, '5:14:2');
INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (4, 1, 1, '1:09:2');

INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (5, 2, 2, '3:15:1');
INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (6, 2, 2, '4:16:2');

INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (7, 4, 3, '5:12:2');
INSERT INTO t_lessons (id, course_id, teacher_id, period) VALUES (8, 4, 3, '1:12:2');

INSERt INTO t_lesson_students(student_id, lesson_id) VALUES (1, 1);
INSERt INTO t_lesson_students(student_id, lesson_id) VALUES (1, 7);
INSERt INTO t_lesson_students(student_id, lesson_id) VALUES (1, 6);

INSERT INTO t_courses_grades_students(id, student_id, course_id) VALUES (1, 1, 1);
INSERT INTO t_courses_grades_students(id, student_id, course_id) VALUES (2, 1, 4);
INSERT INTO t_courses_grades_students(id, student_id, course_id) VALUES (3, 1, 2);


