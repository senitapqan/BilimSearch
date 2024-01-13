CREATE TABLE IF NOT EXISTS t_roles (
   id serial Primary key, 
   role_name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_students (
   id serial Primary key, 
   username varchar(50) unique not null, 
   password varchar(50) not null, 
   email varchar(255) not null, 
   name varchar(50) not null, 
   surname varchar(50) not null,
   role_id int not null,
   FOREIGN KEY (role_id) 
      REFERENCES t_roles(id)
);

CREATE TABLE IF NOT EXISTS t_teachers (
    id serial Primary key, 
    name varchar(50) not null, 
    surname varchar(50) not null, 
    email varchar(255) unique not null, 
    username varchar(50) unique not null, 
    password varchar(50) not null,
    role_id int not null, 
    FOREIGN KEY (role_id)
        REFERENCES t_roles(id)
);

CREATE TABLE IF NOT EXISTS t_courses (
    id serial Primary key, 
    name varchar(50) not null
);

CREATE TABLE IF NOT EXISTS t_lessons (
    id serial Primary key, 
    course_id int not null, 
    teacher_id int not null, 
    period varchar(50), 
    FOREIGN KEY (course_id) 
        REFERENCES t_courses(id),
    FOREIGN KEY (teacher_id)
        REFERENCES t_teachers(id)
);

CREATE TABLE IF NOT EXISTS t_lesson_students (
    student_id int not null, 
    lesson_id int not null, 
    Primary key (student_id, lesson_id), 
    FOREIGN KEY (student_id) REFERENCES t_students(id), 
    FOREIGN KEY (lesson_id) REFERENCES t_lessons(id)
);

CREATE TABLE IF NOt EXISTS t_lesson_items (
    id serial Primary key,
    lesson_id int not null, 
    date date,
    FOREIGN KEY (lesson_id) REFERENCES t_lessons(id)
);

CREATE TABLE IF NOT EXISTS t_attendance(
    id serial Primary key, 
    lesson_item_id int not null,
    FOREIGN KEY (lesson_item_id) REFERENCES t_lesson_items(id)
);

CREATE TABLE IF NOT EXISTS t_attendance_item(
    attendance_id int not null, 
    student_id int not NULL, 
    absence BOOLEAN default TRUE, 
    Primary key (attendance_id, student_id), 
    FOREIGN key (attendance_id) REFERENCES t_attendance(id), 
    FOREIGN key (student_id) REFERENCES t_students(id)
);

CREATE TABLE IF NOT EXISTS t_task(
    id serial Primary key, 
    lesson_item_id int not null,
    FOREIGN KEY (lesson_item_id) REFERENCES t_lesson_items(id)
);

CREATE TABLE IF NOT EXISTS t_task_item(
    task_id int not null, 
    student_id int not NULL, 
    grade int, 
    Primary key (task_id, student_id), 
    FOREIGN key (task_id) REFERENCES t_task(id), 
    FOREIGN key (student_id) REFERENCES t_students(id)
);

CREATE TABLE IF NOT EXISTS t_posts(
    id serial Primary key, 
    author int not null, 
    content TEXT not null,

    FOREIGN KEY (author) REFERENCES t_teachers (id)
);

CREATE TABLE IF NOT EXISTS t_courses_students(
    student_id int not NULL, 
    course_id int not NULL, 
    Primary KEY(student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES t_students(id),
    FOREIGN KEY (course_id) REFERENCES t_courses(id)
);

CREATE TABLE IF NOT EXISTS t_users (
    id serial Primary key,
    username varchar(50) unique not null, 
    password varchar(50) not null
);

CREATE TABLE IF NOt EXISTS t_users_roles (
    user_id int not null,
    role_id int not null, 
    Primary key(user_id, role_id), 
    FOREIGN key (user_id) REFERENCES t_users(id),
    FOREIGN key (role_id) REFERENCES t_roles(id)
);
