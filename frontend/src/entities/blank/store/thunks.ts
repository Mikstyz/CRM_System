import { Student } from "@/entities/student/types";
import { Id } from "@/shared/types";
import {
  CreateStudent,
  DeleteStudent,
  GenerateFilledPDF,
  InfStudentByGroup,
  UpdateStudentByID,
} from "@wails/go/main/App";
import { Group } from "@/entities/group/types";
import { Semester } from "@/entities/discipline/types";
import { createAsyncThunk } from "@reduxjs/toolkit";
import { toast } from "react-toastify";
import { BrowserOpenURL } from "@wails/runtime";

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
    if (!groupId) return rejectWithValue("Не указан ID группы");
    const res = await InfStudentByGroup({
      GroupId: groupId,
    });
    const studentsAll = res.Student;
    if (res.code === 200 && Array.isArray(studentsAll)) {
      return studentsAll.map((s) => {
        return {
          id: s.id,
          fullName: s.full_name,
          company: s.Enterprise,
          startDateWork: s.WorkStartDate,
          position: s.JobTitle,
        };
      });
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
  student: Omit<Student, "id">;
  groupId: Id;
};
type CreateStudentThunksResponse = Student;
export const createStudentThunks = createAsyncThunk<
  CreateStudentThunksResponse,
  CreateStudentThunksParams,
  ThunkConfig
>(
  "userFiles/createStudent",
  async (
    { student, groupId }: CreateStudentThunksParams,
    { rejectWithValue },
  ) => {
    try {
      if (!groupId) return rejectWithValue("Не указан ID группы");
      const res = await CreateStudent({
        FullName: student.fullName,
        GroupId: groupId,
        Enterprise: student.company ?? "",
        WorkStartDate: student.startDateWork ?? "",
        JobTitle: student.position ?? "",
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
      if (!groupId || !student.id)
        return rejectWithValue("Не указан ID группы");
      const res = await UpdateStudentByID({
        StudId: student.id,
        NewFullName: student.fullName,
        NewGroupId: groupId,
        NewEnterprise: student.company ?? "",
        NewWorkStartDate: student.startDateWork ?? "",
        NewJobTitle: student.position ?? "",
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
    if (!studentId) return rejectWithValue("Не указан ID student");
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

/* -------- unified save OR update -------- */
type SaveStudentArgs = {
  groupId: Id;
  student: Omit<Student, "id"> & { id?: Id };
};
export const saveOrUpdateStudentThunks = createAsyncThunk<
  Student,
  SaveStudentArgs,
  ThunkConfig
>(
  "blank/saveOrUpdate",
  async ({ student, groupId }: SaveStudentArgs, { dispatch }) => {
    if (student.id === undefined) {
      return await dispatch(createStudentThunks({ student, groupId })).unwrap();
    }
    return await dispatch(
      updateStudentThunks({ student, groupId } as {
        student: Student;
        groupId: Id;
      }),
    ).unwrap();
  },
);

/* -------- PDF generation (save first) -------- */
interface GeneratePdfArgs {
  group: Group;
  student: Student;
  semester: Semester;
}
export const generatePdfThunks = createAsyncThunk<
  void,
  GeneratePdfArgs,
  ThunkConfig
>(
  "blank/generatePdf",
  async ({ group, student, semester }, { rejectWithValue }) => {
    try {
      console.log("generatePdf", {
        group,
        student,
        semester,
      });
      const res = await GenerateFilledPDF({
        StudentName: student.fullName,
        GroupId: Number(group.id),
        Semester: Number(semester),
        Enterprise: student.company || "",
        WorkStartDate: student.startDateWork || "",
        JobTitle: student.position || "",
      });

      if (res.code !== 200 || !res.File) {
        throw new Error(res.error || "Ошибка генерации PDF");
      }
      if (res.Path) {
        toast.success("Успешно сгенерирован: нажать, чтобы открыть", {
          onClick: () => BrowserOpenURL(`file://${res.Path}`),
        });
      }
    } catch (e) {
      return rejectWithValue((e as Error).message);
    }
  },
);
