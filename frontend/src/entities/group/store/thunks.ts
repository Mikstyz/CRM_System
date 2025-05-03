import { createAsyncThunk } from "@reduxjs/toolkit";
import {
  Course,
  DateNameGroup,
  Graduates,
  Group,
} from "@/entities/group/types";
import { Id } from "@/shared/types";
import { Discipline, Semester } from "@/entities/discipline/types";
import {
  AddSubjectByGroupID,
  AppInFGroupAndSubject,
  CreateGroup,
  DeleteGroupByID,
  DeleteSubjectByID,
  UpdateGroupByID,
  UpdateSubjectByID,
} from "@wails/go/main/App";
import { convResDataInGroups } from "@/entities/group/lib/helper/isDataAndValidationGroupRes.ts";

interface ThunkConfig {
  rejectValue: string;
}

// ========= GROUPS =========

// GET GROUPS
type GetGroupsParams = void;
type GetGroupsResponse = Group[];
export const getGroupsThunks = createAsyncThunk<
  GetGroupsResponse,
  GetGroupsParams,
  ThunkConfig
>("userFiles/getGroups", async (_, { rejectWithValue }) => {
  try {
    const res = await AppInFGroupAndSubject({
      Switch: true,
    });
    console.log("res", res);
    const groupsAll = res.groupsAndSubject;
    if (res.code === 200 && Array.isArray(groupsAll)) {
      return convResDataInGroups(groupsAll);
    } else {
      console.error("Ошибка при получнии групп", res?.error);
      return rejectWithValue("Ошибка при получении групп");
    }
  } catch (error) {
    console.error("Error fetching groups:", error);
    return rejectWithValue(
      `Error fetching groups: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// CREATE GROUPS
type CreateGroupsParams = {
  dateNameGroup: DateNameGroup;
};
type CreateGroupsResponse = Group;
export const createGroupsThunks = createAsyncThunk<
  CreateGroupsResponse,
  CreateGroupsParams,
  ThunkConfig
>("userFiles/createGroups", async ({ dateNameGroup }, { rejectWithValue }) => {
  try {
    const res = await CreateGroup({
      course: Number(dateNameGroup.course),
      groudates: Number(dateNameGroup.graduates),
      speciality: dateNameGroup.specialty,
      group_num: dateNameGroup.groupNumber,
    });
    console.log("res createGroups", res);
    const Group = res?.Group;
    if (res.code === 200 && Group?.Id) {
      return {
        id: Group.Id,
        name: `${dateNameGroup.course}${dateNameGroup.specialty}${dateNameGroup.graduates}-${dateNameGroup.groupNumber}`,
        dateNameGroup: {
          course: Group.Course.toString() as Course,
          specialty: Group.Speciality,
          graduates: Group.Groudates.toString() as Graduates,
          groupNumber: Group.Number,
        },
        disciplines: {
          1: [],
          2: [],
        },
      };
    } else {
      return rejectWithValue("Ошибка при создании группы");
    }
  } catch (error) {
    console.error("Error creating group:", error);
    return rejectWithValue(
      `Error creating group: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// UPDATE GROUPS
type UpdateGroupsParams = Group;
type UpdateGroupsResponse = Omit<Group, "disciplines">;
export const updateGroupsThunks = createAsyncThunk<
  UpdateGroupsResponse,
  UpdateGroupsParams,
  ThunkConfig
>("userFiles/updateGroups", async (group, { rejectWithValue }) => {
  try {
    const res = await UpdateGroupByID({
      group_id: group.id,
      new_course: Number(group.dateNameGroup.course),
      new_groudates: Number(group.dateNameGroup.graduates),
      new_speciality: group.dateNameGroup.specialty,
      new_group_num: group.dateNameGroup.groupNumber,
    });
    if (res?.code === 200 && res?.id) {
      return {
        id: res.id,
        name: `${group.dateNameGroup.course}${group.dateNameGroup.specialty}${group.dateNameGroup.graduates}-${group.dateNameGroup.groupNumber}`,
        dateNameGroup: {
          course: group.dateNameGroup.course,
          specialty: group.dateNameGroup.specialty,
          graduates: group.dateNameGroup.graduates,
          groupNumber: group.dateNameGroup.groupNumber,
        },
      };
    } else {
      console.error("Ошибка при обновлении группы", res?.error);
      return rejectWithValue(
        `Failed to update group: ${res?.error || "Unknown error"}`,
      );
    }
  } catch (error) {
    console.error("Error updating group:", error);
    return rejectWithValue(
      `Error updating group: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

//  DELETE GROUPS
type DeleteGroupsParams = Id;
type DeleteGroupsResponse = Id;
export const deleteGroupsThunks = createAsyncThunk<
  DeleteGroupsResponse,
  DeleteGroupsParams,
  ThunkConfig
>("userFiles/deleteGroups", async (groupsId, { rejectWithValue }) => {
  try {
    const res = await DeleteGroupByID({
      group_id: groupsId,
    });
    if (res?.code === 200) {
      return groupsId;
    } else {
      console.error("Ошибка удаления группы", res?.error);
      return rejectWithValue(
        `Failed to delete group: ${res?.error || "Unknown error"}`,
      );
    }
  } catch (error) {
    console.error("Error deleting group:", error);
    return rejectWithValue(
      `Error deleting group: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// DUPLICATE GROUP
type DuplicateGroupsParams = Id;
type DuplicateGroupsResponse = Group;
export const duplicateGroupThunks = createAsyncThunk<
  DuplicateGroupsResponse,
  DuplicateGroupsParams,
  ThunkConfig
>("userFiles/duplicateGroup", async (groupId, { rejectWithValue }) => {
  try {
    // const res = await DuplicateGroupAllData({ groupId });
    // const resGroup = res.Group;
    // // if (res?.code === 200 && resGroup?.Id) {
    // //   return {
    // //     id: resGroup.Id,
    // //     name: `${resGroup.Course}${resGroup.Speciality}${resGroup.Groudates}-${resGroup.Number}`,
    // //     dateNameGroup: {
    // //       course: String(resGroup.Course) as Course,
    // //       specialty: resGroup.Speciality,
    // //       graduates: String(resGroup.Groudates) as Graduates,
    // //       groupNumber: resGroup.Number,
    // //     },
    // //     disciplines: {
    // //       "1": resGroup.Subject.firstSemester,
    // //       "2": resGroup.Subject.secondSemester,
    // //     },
    // //   };
    // } else {
    //   console.error("Ошибка дублирования группы", res?.error);
    //   return rejectWithValue(
    //     `Failed to duplicate group: ${res?.error || "Unknown error"}`,
    //   );
    // }
    return rejectWithValue(`Failed to duplicate group: Unknown error`);
  } catch (error) {
    console.error("Error duplicating group:", error);
    return rejectWithValue(
      `Error duplicating group: ${(error as Error).message}`,
    );
  }
});

// ========= disciplines =========

// GET DISCIPLINES
// type GetDisciplinesParams = Id;
// type GetDisciplinesResponse = {
//   groupId: Id;
//   disciplines: Disciplines;
// };
// export const getDisciplinesThunks = createAsyncThunk<
//   GetDisciplinesResponse,
//   GetDisciplinesParams,
//   ThunkConfig
// >("userFiles/getDisciplines", async ({groupId, newTitle, semester}, { rejectWithValue }) => {
//   // try {
//   //   const res = await AddSubjectByGroupID({
//   //     group_id: groupId,
//   //     new_subject: newTitle,
//   //     Semester: semester
//   //   })
//   //   console.log("AddSubjectByGroupID", res)
//   //   if (res?.code === 200 && res.id) {
//   //     return {
//   //       groupId,
//   //       disciplineId: res.id,
//   //       semester,
//   //       newTitle: res.subject_name
//   //     };
//   //   } else {
//   //     console.error("Ошибка при получении дисциплин", res?.error);
//   //     return rejectWithValue(
//   //       `Failed to get disciplines: ${res?.error || "Unknown error"}`,
//   //     );
//   //   }
//   // } catch (error) {
//   //   console.error("Error fetching DISCIPLINES:", error);
//   // return rejectWithValue(
//   //   `Error fetching DISCIPLINES: ${(error as Error).message || "Unknown error"}`,
//   // );
//   console.error("Error fetching DISCIPLINES:");
//
//     return rejectWithValue(
//       `Error fetching DISCIPLINES:  || "Unknown error"}`,
//     );
//   // }
// });

// ADD DISCIPLINES
type AddDisciplinesParams = {
  groupId: Id;
  semester: Semester;
  newTitle: string;
};
type AddDisciplinesResponse = {
  groupId: Id;
  semester: Semester;
  disc: {
    id: Id;
    title: string;
  };
};
export const addDisciplinesThunks = createAsyncThunk<
  AddDisciplinesResponse,
  AddDisciplinesParams,
  ThunkConfig
>(
  "userFiles/addDisciplines",
  async ({ groupId, newTitle, semester }, { rejectWithValue }) => {
    try {
      const res = await AddSubjectByGroupID({
        group_id: groupId,
        new_subject: newTitle,
        Semester: semester,
      });
      console.log("AddSubjectByGroupID", res);
      if (res?.code === 200 && res?.id) {
        return {
          groupId,
          semester,
          disc: {
            id: res.id,
            title: res.subject_name || "",
          },
        };
      } else {
        console.error("Ошибка добавления дисциплины", res?.error);
        return rejectWithValue(
          `Failed to add discipline: ${res?.error || "Unknown error"}`,
        );
      }
    } catch (error) {
      console.error("Error adding discipline:", error);
      return rejectWithValue(
        `Error adding discipline: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);

// UPDATE DISCIPLINES
type UpdateDisciplinesParams = {
  groupId: Id;
  semester: Semester;
  newTitle: string;
  discId: Id;
};
type UpdateDisciplinesResponse = {
  groupId: Id;
  semester: Semester;
  disc: Discipline;
};
export const updateDisciplinesThunks = createAsyncThunk<
  UpdateDisciplinesResponse,
  UpdateDisciplinesParams,
  ThunkConfig
>(
  "userFiles/updateDiscipline",
  async ({ groupId, semester, discId, newTitle }, { rejectWithValue }) => {
    try {
      const res = await UpdateSubjectByID({
        subject_id: discId,
        new_subject: newTitle,
      });
      if (res?.code === 200) {
        return {
          groupId,
          semester,
          disc: {
            id: discId,
            title: newTitle,
          },
        };
      } else {
        console.error("Ошибка обновления дисциплины", res?.error);
        return rejectWithValue(
          `Failed to update discipline: ${res?.error || "Unknown error"}`,
        );
      }
    } catch (error) {
      console.error("Error updating discipline:", error);
      return rejectWithValue(
        `Error updating discipline: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);

// DELETE DISCIPLINES
type DeleteDisciplinesParams = {
  groupId: Id;
  semester: Semester;
  discId: Id;
};
type DeleteDisciplinesResponse = {
  groupId: Id;
  semester: Semester;
  discId: Id;
};
export const deleteDisciplinesThunks = createAsyncThunk<
  DeleteDisciplinesResponse,
  DeleteDisciplinesParams,
  ThunkConfig
>(
  "userFiles/deleteDisciplines",
  async ({ groupId, semester, discId }, { rejectWithValue }) => {
    try {
      const res = await DeleteSubjectByID({
        subject_id: discId,
      });
      if (res?.code === 200) {
        return { groupId, semester, discId };
      } else {
        console.log("Ошибка удаления дисциплины", res?.error);
        return rejectWithValue(
          `Failed to delete discipline: ${res?.error || "Unknown error"}`,
        );
      }
    } catch (error) {
      console.error("Error deleting discipline:", error);
      return rejectWithValue(
        `Error deleting discipline: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);
