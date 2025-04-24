// ── entities/group/store/groupSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Discipline, Semester } from "@/entities/discipline/types";
import { Group } from "../types";
import { Id } from "@/shared/types";

interface GroupsState {
  list: Group[];
}

const initialState: GroupsState = { list: [] };

export const groupSlice = createSlice({
  name: "groups",
  initialState,
  reducers: {
    setGroups(state, { payload }: PayloadAction<Group[] | []>) {
      state.list = payload;
    },

    /** add / delete now carry `semester` */
    addDiscipline(
      state,
      {
        payload,
      }: PayloadAction<{
        groupId: Id;
        semester: Semester;
        disc: Discipline;
      }>,
    ) {
      const g = state.list.find((gr) => gr.id === payload.groupId);
      if (g) g.disciplines[payload.semester].push(payload.disc);
    },
    deleteDiscipline(
      state,
      {
        payload,
      }: PayloadAction<{
        groupId: Id;
        semester: Semester;
        discId: Id;
      }>,
    ) {
      const g = state.list.find((gr) => gr.id === payload.groupId);
      if (g) {
        g.disciplines[payload.semester] = g.disciplines[
          payload.semester
        ].filter((d) => d.id !== payload.discId);
      }
    },

    updateGroup(
      state,
      { payload }: PayloadAction<{ groupId: Id; patch: Partial<Group> }>,
    ) {
      const g = state.list.find((gr) => gr.id === payload.groupId);
      if (g) Object.assign(g, payload.patch);
    },

    addGroup(state, { payload }: PayloadAction<Group>) {
      state.list = [payload, ...state.list];
    },
    deleteGroup(state, { payload }: PayloadAction<Id>) {
      state.list = state.list.filter((g) => g.id !== payload);
    },
    duplicateGroup(state, { payload }: PayloadAction<Id>) {
      const srcIndex = state.list.findIndex((g) => g.id === payload);
      if (srcIndex === -1) return;
      const src = state.list[srcIndex];
      const duplicatedGroup = {
        ...JSON.parse(JSON.stringify(src)),
        id: crypto.randomUUID(),
        name: src.name + "0",
      };
      state.list.splice(srcIndex + 1, 0, duplicatedGroup);
    },
  },
});

export const {
  setGroups,
  addDiscipline,
  deleteDiscipline,
  updateGroup,
  addGroup,
  deleteGroup,
  duplicateGroup,
} = groupSlice.actions;

export const groupReducer = groupSlice.reducer;
