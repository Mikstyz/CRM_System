import { createAsyncThunk } from "@reduxjs/toolkit";
import {
  Course,
  DateNameGroup,
  Disciplines,
  Graduates,
  Group,
} from "@/entities/group/types";
import { Id } from "@/shared/types";
import { Discipline, Semester } from "@/entities/discipline/types";
import { InfAllGroup } from "@wails/go/main/App";

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
    const res = await InfAllGroup({});
    const groupsAll = res.groups;

    if (res.code === 200 && Array.isArray(groupsAll)) {
      return groupsAll.map((g): Group => {
        return {
          id: g.Id,
          name: `${g.Course}${g.Speciality}${g.Groudates}-${g.Number}`,
          dateNameGroup: {
            course: String(g.Course) as Course,
            specialty: g.Speciality,
            graduates: String(g.Groudates) as Graduates,
            groupNumber: g.Number,
          },
          disciplines: {
            "1": [],
            "2": [],
          },
        };
      });
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
    // const res = await CreateGroup({
    //   course: Number(dateNameGroup.course),
    //   groudates: Number(dateNameGroup.graduates),
    //   speciality: dateNameGroup.specialty,
    //   group_num: dateNameGroup.groupNumber,
    // });
    // const Groups = res.Groups;
    // if (res.code === 200 && Groups.Id) {
    //   return {
    //     id: Groups.Id,
    //     name: `${dateNameGroup.course}${dateNameGroup.specialty}${dateNameGroup.graduates}-${dateNameGroup.groupNumber}`,
    //     dateNameGroup: {
    //       course: Groups.Course.toString() as Course,
    //       specialty: Groups.Speciality,
    //       graduates: Groups.Groudates.toString() as Graduates,
    //       groupNumber: Groups.Number,
    //     },
    //     disciplines: {
    //       1: [],
    //       2: [],
    //     },
    //   };
    // } else {
    //   return rejectWithValue("Ошибка при создании группы");
    // }
    return rejectWithValue("Ошибка при создании группы");
  } catch (error) {
    console.error("Error creating group:", error);
    return rejectWithValue(
      `Error creating group: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// UPDATE GROUPS
type UpdateGroupsParams = Group;
type UpdateGroupsResponse = Group;
export const updateGroupsThunks = createAsyncThunk<
  UpdateGroupsResponse,
  UpdateGroupsParams,
  ThunkConfig
>("userFiles/updateGroups", async (group, { rejectWithValue }) => {
  try {
    console.log(rejectWithValue);
    return group;
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
    // const res = await DeleteGroupByID({ groupId });
    // if (res?.code === 200) {
    //   return groupsId;
    // } else {
    //   console.error("Ошибка удаления группы", res?.error);
    //   return rejectWithValue(
    //     `Failed to delete group: ${res?.error || "Unknown error"}`,
    //   );
    // }
    return rejectWithValue(`Failed to delete group: Unknown error`);
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
    // if (res?.code === 200) {
    //   return {
    //     id: res.Groups.Id,
    //     name: `${res.Groups.Course}${res.Groups.Speciality}${res.Groups.Groudates}-${res.Groups.Number}`,
    //     dateNameGroup: {
    //       course: res.Groups.Course.toString() as Course,
    //       specialty: res.Groups.Speciality,
    //       graduates: res.Groups.Groudates.toString() as Graduates,
    //       groupNumber: res.Groups.Number,
    //     },
    //     disciplines: {
    //       1: [],
    //       2: [],
    //     },
    //   };
    // } else {
    //   console.error("Ошибка дублирования группы", res?.error);
    // return rejectWithValue(
    //   `Failed to duplicate group: ${res?.error || "Unknown error"}`,
    // );
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
type GetDisciplinesParams = Id;
type GetDisciplinesResponse = {
  groupId: Id;
  disciplines: Disciplines;
};
export const getDisciplinesThunks = createAsyncThunk<
  GetDisciplinesResponse,
  GetDisciplinesParams,
  ThunkConfig
>("userFiles/getDisciplines", async (groupId, { rejectWithValue }) => {
  try {
    console.log({ rejectWithValue, groupId });
    return {
      groupId,
      disciplines: {
        1: [],
        2: [],
      },
    };
  } catch (error) {
    console.error("Error fetching DISCIPLINES:", error);
    return rejectWithValue(
      `Error fetching DISCIPLINES: ${(error as Error).message || "Unknown error"}`,
    );
  }
});

// ADD DISCIPLINES
type AddDisciplinesParams = {
  groupId: Id;
  semester: Semester;
  newTitle: string;
};
type AddDisciplinesResponse = {
  groupId: Id;
  semester: Semester;
  disc: Discipline;
};
export const addDisciplinesThunks = createAsyncThunk<
  AddDisciplinesResponse,
  AddDisciplinesParams,
  ThunkConfig
>(
  "userFiles/addDisciplines",
  async ({ groupId, semester, newTitle }, { rejectWithValue }) => {
    try {
      // const res = await AddSubjectByGroupID({
      //   GroupId: String(groupId),
      //   NewSubject: newTitle,
      // });
      // console.log("Add_SubjectByGroupId", res);
      // if (res?.group_id && res.subject_name) {
      //   return {
      //     groupId,
      //     semester,
      //     disc: {
      //       id: res.group_id,
      //       title: res.subject_name,
      //     },
      //   };
      // } else {
      //   console.error("Ошибка добавления дисциплины", res?.error);
      //   return rejectWithValue(
      //     `Failed to add discipline: ${res?.error || "Unknown error"}`,
      //   );
      // }
      return rejectWithValue(`Failed to add discipline: Unknown error`);
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
  disc: Discipline;
};
type UpdateDisciplinesResponse = {
  groupId: Id;
  disc: Discipline;
};
export const updateDisciplinesThunks = createAsyncThunk<
  UpdateDisciplinesResponse,
  UpdateDisciplinesParams,
  ThunkConfig
>(
  "userFiles/updateDisciplines",
  async ({ groupId, semester, disc }, { rejectWithValue }) => {
    try {
      console.log({ rejectWithValue, groupId, semester, disc });
      return {
        groupId,
        disc: {
          id: 1,
          title: "new",
        },
      };
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
      // const res = await DeleteSubjectByID({
      //   SubjectId: String(discId),
      // });
      // if (res?.code === 200) {
      //   return { groupId, semester, discId };
      // } else {
      //   console.log("Ошибка удаления дисциплины", res?.error);
      //   return rejectWithValue(
      //     `Failed to delete discipline: ${res?.error || "Unknown error"}`,
      //   );
      // }
      return rejectWithValue(`Failed to delete discipline: Unknown error`);
    } catch (error) {
      console.error("Error deleting discipline:", error);
      return rejectWithValue(
        `Error deleting discipline: ${(error as Error).message || "Unknown error"}`,
      );
    }
  },
);
