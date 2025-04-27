```markdown

## Методы

### Студенты

#### InfAllStudent
- **Описание**: Возвращает список всех студентов.
- **Входные данные**:
  ```json
  {}
```

- **Ответ**:
    - **200**: Успех — возвращается список студентов.
        
        ```json
        {
          "Code": 200,
          "Students": [
            {
              "Id": "string",
              "FullName": "string",
              "Course": "string",
              "Graduates": "string",
              "Speciality": "string",
              "GroupNum": "string",
              "Semester": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — проблема на сервере (например, сбой базы данных).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### InfStudentByID

- **Описание**: Возвращает данные студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudentID": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращаются данные студента.
        
        ```json
        {
          "Code": 200,
          "Student": {
            "Id": "string",
            "FullName": "string",
            "Course": "string",
            "Graduates": "string",
            "Speciality": "string",
            "GroupNum": "string",
            "Semester": "string"
          }
        }
        ```
        
    - **500**: Ошибка — студент не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### InfStudentByGroup

- **Описание**: Возвращает список студентов по данным группы (курс, специальность, номер группы, семестр).
- **Входные данные**:
    
    ```json
    {
      "Course": "string",
      "Speciality": "string",
      "GroupNum": "string",
      "Semester": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается список студентов группы.
        
        ```json
        {
          "Code": 200,
          "Students": [
            {
              "Id": "string",
              "FullName": "string",
              "Course": "string",
              "Graduates": "string",
              "Speciality": "string",
              "GroupNum": "string",
              "Semester": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### CreateStudent

- **Описание**: Создаёт нового студента.
- **Входные данные**:
    
    ```json
    {
      "FullName": "string",
      "Course": "string",
      "Graduates": "string",
      "Speciality": "string",
      "GroupNum": "string",
      "Semester": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — студент создан, возвращаются его данные.
        
        ```json
        {
          "Code": 200,
          "Id": "string",
          "FullName": "string",
          "Course": "string",
          "Graduates": "string",
          "Speciality": "string",
          "GroupNum": "string",
          "Semester": "string"
        }
        ```
        
    - **500**: Ошибка — не удалось создать студента (например, дубликат или сбой базы).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### UpdateStudentByID

- **Описание**: Обновляет данные студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudId": "string",
      "NewFullName": "string",
      "NewCourse": "string",
      "NewGraduates": "string",
      "NewSpeciality": "string",
      "NewNumber": "string",
      "NewSemester": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — данные студента обновлены.
        
        ```json
        {
          "Code": 200,
          "NewId": "string",
          "NewFullName": "string",
          "NewCourse": "string",
          "NewGraduates": "string",
          "NewSpeciality": "string",
          "NewGroupNum": "string",
          "NewSemester": "string"
        }
        ```
        
    - **500**: Ошибка — студент не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### DeleteStudent

- **Описание**: Удаляет студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — студент удалён.
        
        ```json
        {
          "Code": 200
        }
        ```
        
    - **500**: Ошибка — студент не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

### Работодатели

#### InfEmpByStudentID

- **Описание**: Возвращает список работодателей для студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudentId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается список работодателей.
        
        ```json
        {
          "Code": 200,
          "Employers": [
            {
              "Id": "string",
              "Enterprise": "string",
              "WorkStartDate": "string",
              "JobTitle": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — данные не найдены или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### UpdateEmpByStudentID

- **Описание**: Обновляет данные работодателя для студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudId": "string",
      "NewEnterprise": "string",
      "NewWorkStartDate": "string",
      "NewJobTitle": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — данные работодателя обновлены.
        
        ```json
        {
          "Code": 200,
          "StudId": "string",
          "NewEnterprise": "string",
          "NewWorkStartDate": "string",
          "NewJobTitle": "string"
        }
        ```
        
    - **500**: Ошибка — студент не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### DeleteEmpByStudentID

- **Описание**: Удаляет данные работодателя для студента по его ID.
- **Входные данные**:
    
    ```json
    {
      "StudId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — данные работодателя удалены.
        
        ```json
        {
          "Code": 200
        }
        ```
        
    - **500**: Ошибка — данные не найдены или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

### Группы

#### InfAllGroup

- **Описание**: Возвращает список всех групп.
- **Входные данные**:
    
    ```json
    {}
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается список групп.
        
        ```json
        {
          "Code": 200,
          "Groups": [
            {
              "Id": "string",
              "Course": "string",
              "Groudates": "string",
              "Speciality": "string",
              "Number": "string",
              "Semester": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### CreateGroup

- **Описание**: Создаёт новую группу (две записи с семестрами 1 и 2).
- **Входные данные**:
    
    ```json
    {
      "Course": "string",
      "Graduates": "string",
      "Speciality": "string",
      "Number": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — группа создана, возвращаются данные двух групп.
        
        ```json
        {
          "Code": 200,
          "Groups": [
            {
              "Id": "string",
              "Course": "string",
              "Groudates": "string",
              "Speciality": "string",
              "Number": "string",
              "Semester": "1"
            },
            {
              "Id": "string",
              "Course": "string",
              "Groudates": "string",
              "Speciality": "string",
              "Number": "string",
              "Semester": "2"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — не удалось создать группу (например, дубликат или сбой базы).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### UpdateGroupByID

- **Описание**: Обновляет данные группы по её ID.
- **Входные данные**:
    
    ```json
    {
      "GroupId": "string",
      "NewCourse": "string",
      "NewGraduates": "string",
      "NewSpeciality": "string",
      "NewNumber": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — группа обновлена.
        
        ```json
        {
          "Code": 200,
          "Id": "string",
          "Course": "string",
          "Graduates": "string",
          "Speciality": "string",
          "Number": "string"
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "не удалось обновить группу"
        }
        ```
        

#### DeleteGroupByID

- **Описание**: Удаляет группу по данным (курс, выпускники, специальность, номер).
- **Входные данные**:
    
    ```json
    {
      "Course": "string",
      "Graduates": "string",
      "Speciality": "string",
      "Number": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — группа удалена.
        
        ```json
        {
          "Code": 200
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### GetGroupIDByInfo

- **Описание**: Возвращает ID группы по данным (курс, выпускники, специальность, номер, семестр).
- **Входные данные**:
    
    ```json
    {
      "Course": "string",
      "Graduates": "string",
      "Speciality": "string",
      "Number": "string",
      "Semester": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается ID группы.
        
        ```json
        {
          "Code": 200,
          "Id": "string"
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### DuplicateGroupAllData

- **Описание**: Дублирует данные группы (создаёт две новые записи с семестрами 1 и 2).
- **Входные данные**:
    
    ```json
    {
      "Course": "string",
      "Graduates": "string",
      "Speciality": "string",
      "Number": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — группы созданы, возвращаются их данные.
        
        ```json
        {
          "Code": 200,
          "Groups": [
            {
              "Id": "string",
              "Course": "string",
              "Groudates": "string",
              "Speciality": "string",
              "Number": "string",
              "Semester": "1"
            },
            {
              "Id": "string",
              "Course": "string",
              "Groudates": "string",
              "Speciality": "string",
              "Number": "string",
              "Semester": "2"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — не удалось создать группы (например, дубликат или сбой базы).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

### Предметы

#### InfSubjectByGroupID

- **Описание**: Возвращает список предметов для группы по её ID.
- **Входные данные**:
    
    ```json
    {
      "GroupId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается список предметов.
        
        ```json
        {
          "Code": 200,
          "Subject": [
            {
              "Id": "string",
              "Name": "string",
              "GroupId": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### InfDisciplinesByGroupData

- **Описание**: Возвращает список дисциплин по данным группы (специальность, номер группы, курс, выпускники).
- **Входные данные**:
    
    ```json
    {
      "Speciality": "string",
      "GroupNum": "string",
      "Course": "string",
      "Groudates": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — возвращается список дисциплин.
        
        ```json
        {
          "Code": 200,
          "Disciplines": [
            {
              "Id": "string",
              "Name": "string"
            }
          ]
        }
        ```
        
    - **500**: Ошибка — данные не найдены или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### AddSubjectByGroupID

- **Описание**: Добавляет новый предмет для группы.
- **Входные данные**:
    
    ```json
    {
      "GroupId": "string",
      "NewSubject": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — предмет добавлен.
        
        ```json
        {
          "Code": 200,
          "Id": "string",
          "Subject": "string",
          "GroupId": "string"
        }
        ```
        
    - **500**: Ошибка — не удалось добавить предмет (например, сбой базы).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### UpdateSubjectByID

- **Описание**: Обновляет предмет по его ID.
- **Входные данные**:
    
    ```json
    {
      "SubjectId": "string",
      "NewSubject": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — предмет обновлён.
        
        ```json
        {
          "Code": 200,
          "Id": "string",
          "Subject": "string"
        }
        ```
        
    - **500**: Ошибка — предмет не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "не удалось обновить предмет"
        }
        ```
        

#### DeleteSubjectByID

- **Описание**: Удаляет предмет по его ID.
- **Входные данные**:
    
    ```json
    {
      "SubjectId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — предмет удалён.
        
        ```json
        {
          "Code": 200
        }
        ```
        
    - **500**: Ошибка — предмет не найден или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

#### DeleteAllSubjectsByGroupID

- **Описание**: Удаляет все предметы для группы по её ID.
- **Входные данные**:
    
    ```json
    {
      "GroupId": "string"
    }
    ```
    
- **Ответ**:
    - **200**: Успех — все предметы группы удалены.
        
        ```json
        {
          "Code": 200
        }
        ```
        
    - **500**: Ошибка — группа не найдена или сбой сервера.
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
        

### PDF

#### GenerateFilledPDF

- **Описание**: Генерирует заполненный PDF-документ.
- **Входные данные**:
    
    ```json
    {
      "data": "object" // структура зависит от реализации
    }
    ```
    
- **Ответ**:
    - **200**: Успех — PDF-документ сгенерирован.
        
        ```json
        {
          "Code": 200,
          "File": "string" // например, base64-строка или путь к файлу
        }
        ```
        
    - **500**: Ошибка — не удалось сгенерировать PDF (например, неверные данные или сбой сервера).
        
        ```json
        {
          "Code": 500,
          "Error": "описание ошибки"
        }
        ```
```
