PRAGMA foreign_keys = ON;

-- Группы
CREATE TABLE IF NOT EXISTS einf_groups (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Well INTEGER NOT NULL,
    Semester INTEGER NOT NULL,
    Speciality TEXT NOT NULL,
    gClass INTEGER NOT NULL,
    GroupNum INTEGER NOT NULL
);

-- Студенты
CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    FullName TEXT NOT NULL,
    group_id INTEGER NOT NULL,
    FOREIGN KEY (group_id) REFERENCES einf_groups(Id) ON DELETE CASCADE
);

-- Таблица работодателей
CREATE TABLE IF NOT EXISTS employers (
    studid INTEGER NOT NULL,
    enterprise TEXT NULL,
    workstartdate DATE NULL,
    jobtitle TEXT NULL,
    FOREIGN KEY (studid) REFERENCES students(id) ON DELETE CASCADE
);

-- Таблица связь группа–предмет
CREATE TABLE IF NOT EXISTS group_subject (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    subject_name TEXT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES einf_groups(Id) ON DELETE CASCADE
);

-- Индексы
CREATE INDEX IF NOT EXISTS idx_group_subject_group_id ON group_subject(group_id);
CREATE INDEX IF NOT EXISTS idx_students_group_id ON students (group_id);
