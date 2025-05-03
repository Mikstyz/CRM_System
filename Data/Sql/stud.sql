-- Создание таблицы групп
CREATE TABLE IF NOT EXISTS einf_groups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Course INTEGER NOT NULL,
    Speciality TEXT NOT NULL,
    Groudates INTEGER NOT NULL,
    GroupNum INTEGER NOT NULL
);

-- Создание таблицы студентов (обновлённая версия)
CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FullName TEXT NOT NULL,
    GroupId INTEGER NOT NULL,
    FOREIGN KEY (GroupId) REFERENCES einf_groups(Id) ON DELETE CASCADE
);

-- Создание таблицы работодателей
CREATE TABLE IF NOT EXISTS employers (
    studid INTEGER NOT NULL,
    enterprise TEXT NULL,
    workstartdate TEXT NULL,
    jobtitle TEXT NULL,
    FOREIGN KEY (studid) REFERENCES students(id) ON DELETE CASCADE
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

-- Индекс на столбец GroupId в students
CREATE INDEX IF NOT EXISTS idx_students_group_id ON students(GroupId);

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
