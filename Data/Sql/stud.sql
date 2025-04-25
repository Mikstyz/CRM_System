CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FullName TEXT NOT NULL,
    Speciality TEXT NOT NULL,
    GroupNum INTEGER NOT NULL,
    Semester INTEGER NOT NULL,
    Course INTEGER NOT NULL,
    Groudates INTEGER NOT NULL DEFAULT 1
);

-- Создание таблицы работодателей
CREATE TABLE IF NOT EXISTS employers (
    studid INTEGER NOT NULL,
    enterprise TEXT NULL,
    workstartdate DATE NULL,
    jobtitle TEXT NULL,
    FOREIGN KEY (studid) REFERENCES students(id) ON DELETE CASCADE
);

-- Создание таблицы групп
CREATE TABLE IF NOT EXISTS einf_groups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Course INTEGER NOT NULL,
    Semester INTEGER NOT NULL,
    Speciality TEXT NOT NULL,
    Groudates INTEGER NOT NULL,
    GroupNum INTEGER NOT NULL
);

-- Создание таблицы связи группа-учебный предмет
CREATE TABLE IF NOT EXISTS group_subject (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    subject_name TEXT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES einf_groups(id) ON DELETE CASCADE
);

-- Индекс на столбцы group_id для ускорения поиска
CREATE INDEX IF NOT EXISTS idx_group_subject_group_id ON group_subject(group_id);

-- Индекс на столбцы Speciality, GroupNum, Course, Semester для ускорения поиска
CREATE INDEX IF NOT EXISTS idx_students_group_data ON students (Speciality, GroupNum, Course, Semester);

-- Триггер для удаления студентов при удалении группы
CREATE TRIGGER IF NOT EXISTS trigger_delete_students_for_group
AFTER DELETE ON einf_groups
FOR EACH ROW
BEGIN
    DELETE FROM students 
    WHERE Speciality = OLD.Speciality 
    AND GroupNum = OLD.GroupNum 
    AND Course = OLD.Course;
END;

-- Триггер для обновления данных студентов при изменении группы
CREATE TRIGGER IF NOT EXISTS trigger_update_student_group_data
AFTER UPDATE ON einf_groups
FOR EACH ROW
WHEN (OLD.Speciality != NEW.Speciality 
    OR OLD.GroupNum != NEW.GroupNum 
    OR OLD.Course != NEW.Course 
    OR OLD.Semester != NEW.Semester 
    OR OLD.Groudates != NEW.Groudates)
BEGIN
    UPDATE students 
    SET Speciality = NEW.Speciality, 
        GroupNum = NEW.GroupNum, 
        Course = NEW.Course, 
        Semester = NEW.Semester, 
        Groudates = NEW.Groudates 
    WHERE Speciality = OLD.Speciality 
    AND GroupNum = OLD.GroupNum 
    AND Course = OLD.Course 
    AND Semester = OLD.Semester;
END;