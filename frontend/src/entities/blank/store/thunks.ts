import { createAsyncThunk } from "@reduxjs/toolkit";
import { Student } from "@/entities/student/types";
import { Id } from "@/shared/types";
import {
  CreateStudent,
  DeleteStudent,
  InfStudentByGroup,
  UpdateStudentByID,
} from "@wails/go/main/App";

interface ThunkConfig {
  rejectValue: string;
}

// GET AllStudentGroup
type GetAllStudentGroupResponse = Student[];
type GetAllStudentGroupParams = Id;
export const getAllStudentGroupThunks = createAsyncThunk<
  GetAllStudentGroupResponse,
  GetAllStudentGroupParams,
  ThunkConfig
>("userFiles/getAllStudentGroup", async (groupId, { rejectWithValue }) => {
  try {
    const res = await InfStudentByGroup({
      GroupId: groupId,
    });
    const studentsAll = res.Student;
    if (res.code === 200 && Array.isArray(studentsAll)) {
      // return studentsAll.map((s) => {
      //   return {
      //     id: s.id,
      //     fullName: s.full_name,
      //     company: s.speciality,
      //     startDateWork: s.startDateWork,
      //     position: s.speciality
      //   }
      // })
      return [];
    } else {
      console.error("Ошибка при получении студентов", res?.error);
      return rejectWithValue("Ошибка при получении студентов");
    }
  } catch (error) {
    console.error("Error:", error);
    return rejectWithValue(
      `Ошибка при получении студентов: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// CreateStudent
type CreateStudentThunksParams = {
  student: Student;
  groupId: Id;
};
type CreateStudentThunksResponse = Student;
export const createStudentThunks = createAsyncThunk<
  CreateStudentThunksResponse,
  CreateStudentThunksParams,
  ThunkConfig
>(
  "userFiles/createStudent",
  async ({ student, groupId }, { rejectWithValue }) => {
    try {
      const res = await CreateStudent({
        FullName: student.fullName,
        GroupId: groupId,
      });
      if (res.code === 200 && res.Id) {
        return {
          id: res.Id,
          fullName: student.fullName,
          company: student.company,
          startDateWork: student.startDateWork,
          position: student.position,
        };
      } else {
        console.error("Ошибка при сохранении студента", res?.error);
        return rejectWithValue("Ошибка при сохранении студента");
      }
    } catch (error) {
      console.error("Error:", error);
      return rejectWithValue(
        `Ошибка при сохранении студента: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);

// UpdateStudentByID
type UpdateStudentThunksParams = {
  student: Student;
  groupId: Id;
};
type UpdateStudentThunksResponse = Student;
export const updateStudentThunks = createAsyncThunk<
  UpdateStudentThunksResponse,
  UpdateStudentThunksParams,
  ThunkConfig
>(
  "userFiles/updateStudent",
  async ({ student, groupId }, { rejectWithValue }) => {
    try {
      const res = await UpdateStudentByID({
        StudId: student.id,
        NewFullName: student.fullName,
        NewGroupId: groupId,
      });
      if (res.code === 200 && res.Id) {
        return {
          id: res.Id,
          fullName: student.fullName,
          company: student.company,
          startDateWork: student.startDateWork,
          position: student.position,
        };
      } else {
        console.error("Ошибка при обновлении студента", res?.error);
        return rejectWithValue("Ошибка при обновлении студента");
      }
    } catch (error) {
      console.error("Error:", error);
      return rejectWithValue(
        `Ошибка при обновлении студента: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);

// DeleteStudent
type DeleteStudentThunksParams = Id;
type DeleteStudentThunksResponse = Id;
export const deleteStudentThunks = createAsyncThunk<
  DeleteStudentThunksResponse,
  DeleteStudentThunksParams,
  ThunkConfig
>("userFiles/deleteStudent", async (studentId, { rejectWithValue }) => {
  try {
    const res = await DeleteStudent({
      StudId: studentId,
    });
    if (res.code === 200) {
      return studentId;
    } else {
      console.error("Ошибка при удалении студента", res?.error);
      return rejectWithValue("Ошибка при удалении студента");
    }
  } catch (error) {
    console.error("Error:", error);
    return rejectWithValue(
      `Ошибка при удалении студента: ${(error as Error).message || "Unknown error"}`,
    );
  }
});
