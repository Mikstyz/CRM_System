import { createAsyncThunk } from "@reduxjs/toolkit";
import { InfAllStudent, InfStudentByID } from "@wails/go/main/App";
import { DateNameGroup } from "@/entities/group/types";
import { Student } from "@/entities/student/types";
import { Id } from "@/shared/types";

interface ThunkConfig {
  rejectValue: string;
}

//////////

// GET AllStudentGroup
type GetAllStudentGroupResponse = DateNameGroup;
type GetAllStudentGroupParams = Student[];
export const getAllStudentGroupThunks = createAsyncThunk<
  GetAllStudentGroupParams,
  GetAllStudentGroupResponse,
  ThunkConfig
>("userFiles/getAllStudent", async (_, { rejectWithValue }) => {
  try {
    const res = await InfAllStudent({});
    const studentsAll = res.students;

    if (res.code === 200 && Array.isArray(studentsAll)) {
      // return studentsAll.map((g): Group => {
      //   return {
      //     id: g.Id,
      //     name: `${g.Course}${g.Speciality}${g.Groudates}-${g.Number}`,
      //     dateNameGroup: {
      //       course: String(g.Course) as Course,
      //       specialty: g.Speciality,
      //       graduates: String(g.Groudates) as Graduates,
      //       groupNumber: g.Number,
      //     },
      //     disciplines: {
      //       "1": [],
      //       "2": [],
      //     },
      //   };
      // });
      return [];
    } else {
      console.error("Ошибка при получнии студентов", res?.error);
      return rejectWithValue("Ошибка при получении студентов");
    }
  } catch (error) {
    console.error("Error fetching groups:", error);
    return rejectWithValue(
      `шибка при получении студентов: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// GET InfStudentID
type GetInfStudentIDResponse = Id;
type GetInfStudentIDParams = Student[];
export const getInfStudentIDThunks = createAsyncThunk<
  GetInfStudentIDParams,
  GetInfStudentIDResponse,
  ThunkConfig
>("userFiles/getInfStudent", async (studentId, { rejectWithValue }) => {
  try {
    const res = await InfStudentByID({
      StudentID: studentId,
    });
    const student = res.Student;
    if (res.code === 200 && student?.id) {
      return [];
    } else {
      console.error("Ошибка при получнии студентов", res?.error);
      return rejectWithValue("Ошибка при получении студентов");
    }
  } catch (error) {
    console.error("Error fetching groups:", error);
    return rejectWithValue(
      `шибка при получении студентов: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// CreateStudent
// type CreateStudentThunksParams = Student[]
// type CreateStudentThunksResponse = Id
// export const createStudentThunks = createAsyncThunk<
//   CreateStudentThunksParams,
//   CreateStudentThunksResponse,
//   ThunkConfig
// >("userFiles/getAllStudent", async (_, { rejectWithValue }) => {
//   try {
// const res = await InfStudentByID({});
// const studentsAll = res.students;
//     if (res.code === 200 && Array.isArray(studentsAll)) {
//       return []
//     } else {
//       console.error("Ошибка при получнии студентов", res?.error);
//       return rejectWithValue("Ошибка при получении студентов");
//     }
//   } catch (error) {
//     console.error("Error fetching groups:", error);
//     return rejectWithValue(
//       `шибка при получении студентов: ${(error as Error).message || "Unknown error"}`,
//     );
//   }
// });
