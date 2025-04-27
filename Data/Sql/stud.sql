-- Создание таблицы студентов
CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FullName TEXT NOT NULL,
    Speciality TEXT NOT NULL,
    GroupNum INTEGER NOT NULL,
    Course INTEGER NOT NULL,
    Groudates INTEGER NOT NULL DEFAULT 1
);

-- Создание таблицы работодателей
CREATE TABLE IF NOT EXISTS employers (
    studid INTEGER NOT NULL,
    enterprise TEXT NULL,
    workstartdate TEXT NULL,
    jobtitle TEXT NULL,
    FOREIGN KEY (studid) REFERENCES students(id) ON DELETE CASCADE
);

-- Создание таблицы групп
CREATE TABLE IF NOT EXISTS einf_groups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Course INTEGER NOT NULL,
    Speciality TEXT NOT NULL,
    Groudates INTEGER NOT NULL,
    GroupNum INTEGER NOT NULL
);

-- Создание таблицы связи группа-учебный предмет
CREATE TABLE IF NOT EXISTS group_subject (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    subject_name TEXT NOT NULL,
    semester INTEGER NOT NULL CHECK (semester IN (1, 2)),
    FOREIGN KEY (group_id) REFERENCES einf_groups(id) ON DELETE CASCADE
);

-- Индекс на столбец group_id для ускорения поиска
CREATE INDEX IF NOT EXISTS idx_group_subject_group_id ON group_subject(group_id);

-- Индекс на столбцы Speciality, GroupNum, Course для ускорения поиска
CREATE INDEX IF NOT EXISTS idx_students_group_data ON students (Speciality, GroupNum, Course);

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
    OR OLD.Groudates != NEW.Groudates)
BEGIN
    UPDATE students 
    SET Speciality = NEW.Speciality, 
        GroupNum = NEW.GroupNum, 
        Course = NEW.Course, 
        Groudates = NEW.Groudates 
    WHERE Speciality = OLD.Speciality 
    AND GroupNum = OLD.GroupNum 
    AND Course = OLD.Course;
END;

-- Триггер создаёт пустую запись в таблице employers для нового студента
CREATE TRIGGER IF NOT EXISTS trigger_create_empty_employer
AFTER INSERT ON students
FOR EACH ROW
BEGIN
    INSERT INTO employers (studid, enterprise, workstartdate, jobtitle)
    VALUES (NEW.id, NULL, NULL, NULL);
END;

-- Триггер удаляет запись из employers при удалении студента
CREATE TRIGGER IF NOT EXISTS trigger_delete_employer_on_student_delete
AFTER DELETE ON students
FOR EACH ROW
BEGIN
    DELETE FROM employers
    WHERE studid = OLD.id;
END;

-- Триггер для удаления предметов при удалении группы
CREATE TRIGGER IF NOT EXISTS trigger_delete_subjects_for_group
AFTER DELETE ON einf_groups
FOR EACH ROW
BEGIN
    DELETE FROM group_subject
    WHERE group_id = OLD.Id;
END;