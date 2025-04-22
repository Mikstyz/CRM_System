# README

## О программе

Это официальный шаблон Wails Vanilla.

Вы можете настроить проект, отредактировав `wails.json`. Более подробную информацию о настройках проекта можно найти
здесь: https://wails.io/docs/reference/project-config

## Живая разработка

Чтобы запустить проект в режиме живой разработки, выполните команду `wails dev` в директории проекта. Это запустит сервер разработки Vite
который обеспечит очень быструю горячую перезагрузку изменений во фронтенде. Если вы хотите разрабатывать в браузере
и иметь доступ к своим методам Go, есть также сервер разработки, который работает на http://localhost:34115. Подключите
к нему в браузере, и вы сможете вызывать ваш Go-код из devtools.

## Сборка

Чтобы собрать перераспределяемый пакет в режиме производства, используйте `wails build`.


## Методы и документация к ним

## Методы студента

### Inf_AllStudent

**Метод**: `Inf_AllStudent() modeles.Inf_AllStudents`

**Возвращает**:

`Inf_AllStudents`:

- `Code` (int)
- `Students` (\[\]Student)
- `Error` (string)

### Inf_StudentByID

**Метод**: `Inf_StudentByID(studentID int) modeles.Inf_Student`

**Параметры**:

- `studentID` (int)

**Возвращает**:

`Inf_Student`:

- `Code` (int)
- `Student` (\[\]Student)
- `Error` (string)

### Create_Student

**Метод**: `Create_Student(fullName string, well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.Create`

**Параметры**:

- `fullName` (string)
- `well` (byte)
- `gClass` (byte)
- `speciality` (string)
- `groupNum` (int)
- `semester` (byte)

**Возвращает**:

`Create`:

- `Code` (int)
- `Id` (int)
- `Error` (string)

### Update_StudentById

**Метод**: `Update_StudentById(studId int, newFullName string, newWell byte, newClass byte, newSpeciality string, newGroupNum int, newSemester byte) modeles.Update`

**Параметры**:

- `studId` (int)
- `newFullName` (string)
- `newWell` (byte)
- `newClass` (byte)
- `newSpeciality` (string)
- `newGroupNum` (int)
- `newSemester` (byte)

**Возвращает**:

`Update`:

- `Code` (int)
- `Error` (string)

### Delete_Student

**Метод**: `Delete_Student(studId int) modeles.Delete`

**Параметры**:

- `studId` (int)

**Возвращает**:

`Delete`:

- `Code` (int)
- `Error` (string)

## Методы работодателя

### Inf_EmpByStudentId

**Метод**: `Inf_EmpByStudentId(StudentId int) modeles.Inf_Employer`

**Параметры**:

- `StudentId` (int)

**Возвращает**:

`Inf_Employer`:

- `Code` (int)
- `Employer` (EInfEmp)
- `Error` (string)

### Update_EmpbyStudentId

**Метод**: `Update_EmpbyStudentId(studId int, newEnterprise string, newWorkStartDate string, newJobTitle string) modeles.Update`

**Параметры**:

- `studId` (int)
- `newEnterprise` (string)
- `newWorkStartDate` (string)
- `newJobTitle` (string)

**Возвращает**:

`Update`:

- `Code` (int)
- `Error` (string)

### Delete_EmpbyStudentId

**Метод**: `Delete_EmpbyStudentId(studId int) modeles.Delete`

**Параметры**:

- `studId` (int)

**Возвращает**:

`Delete`:

- `Code` (int)
- `Error` (string)

## Методы группы

### Inf_AllGroup

**Метод**: `Inf_AllGroup() modeles.Inf_AllGroup`

**Возвращает**:

`Inf_AllGroup`:

- `Code` (int)
- `Groups` (\[\]EinfGroup)
- `Error` (string)

### Create_Group

**Метод**: `Create_Group(well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.Create`

**Параметры**:

- `well` (byte)
- `gClass` (byte)
- `speciality` (string)
- `groupNum` (int)
- `semester` (byte)

**Возвращает**:

`Create`:

- `Code` (int)
- `Id` (int)
- `Error` (string)

### Update_GroupById

**Метод**: `Update_GroupById(groupId int, newWell byte, newGClass byte, newSpeciality string, newGroupNum int, newSemester byte) modeles.Update`

**Параметры**:

- `groupId` (int)
- `newWell` (byte)
- `newGClass` (byte)
- `newSpeciality` (string)
- `newGroupNum` (int)
- `newSemester` (byte)

**Возвращает**:

`Update`:

- `Code` (int)
- `Error` (string)

### Delete_GroupById

**Метод**: `Delete_GroupById(groupId int) modeles.Delete`

**Параметры**:

- `groupId` (int)

**Возвращает**:

`Delete`:

- `Code` (int)
- `Error` (string)

### GetGroupId_GroupIdByInfo

**Метод**: `GetGroupId_GroupIdByInfo(well byte, gClass byte, speciality string, groupNum int, semester byte) modeles.GetId`

**Параметры**:

- `well` (byte)
- `gClass` (byte)
- `speciality` (string)
- `groupNum` (int)
- `semester` (byte)

**Возвращает**:

`GetId`:

- `Code` (int)
- `Id` (int)
- `Error` (string)

## Методы предметов

### Inf_SubjectByGroupId

**Метод**: `Inf_SubjectByGroupId(groupId int) modeles.Inf_Subject`

**Возвращает**:

`Inf_Subject`:

- `Code` (int)
- `Subject` (\[\]string)
- `Error` (string)

### Add_SubjectByGroupId

**Метод**: `Add_SubjectByGroupId(groupId int, newSubject string) modeles.Create`

**Параметры**:

- `groupId` (int)
- `newSubject` (string)

**Возвращает**:

`Create`:

- `Code` (int)
- `Id` (int)
- `Error` (string)

### Update_SubjectById

**Метод**: `Update_SubjectById(subjectId int, newSubject string) modeles.Update`

**Параметры**:

- `subjectId` (int)
- `newSubject` (string)

**Возвращает**:

`Update`:

- `Code` (int)
- `Error` (string)

### Delete_SubjectById

**Метод**: `Delete_SubjectById(subjectId int) modeles.Delete`

**Параметры**:

- `subjectId` (int)

**Возвращает**:

`Delete`:

- `Code` (int)
- `Error` (string)

### Delete_AllSubjectByGroupId

**Метод**: `Delete_AllSubjectByGroupId(groupID int) modeles.Delete`

**Параметры**:

- `groupID` (int)

**Возвращает**:

`Delete`:

- `Code` (int)
- `Error` (string)

## Генерация PDF

### GenerateFilledPDF

**Метод**: `GenerateFilledPDF(StudentId int, StudentName string, GroupId int) modeles.PdfDock`

**Параметры**:

- `StudentId` (int)
- `StudentName` (string)
- `GroupId` (int)

**Возвращает**:

`PdfDock`:

- `Code` (int)
- `File` (\[\]byte)
- `Error` (string)
