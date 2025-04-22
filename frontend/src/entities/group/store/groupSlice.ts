// ── entities/group/store/groupSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Discipline, Semester } from "@/entities/discipline/types";
import { Group } from "../types";

interface GroupsState {
  list: Group[];
}

const initialState: GroupsState = { list: [] };

export const groupSlice = createSlice({
  name: "groups",
  initialState,
  reducers: {
    setGroups(state, { payload }: PayloadAction<Group[]>) {
      state.list = payload;
    },
    toggleExpand(state, { payload }: PayloadAction<string>) {
      const g = state.list.find((gr) => gr.id === payload);
      if (g) g.isExpanded = !g.isExpanded;
    },

    /** add / delete now carry `semester` */
    addDiscipline(
      state,
      {
        payload,
      }: PayloadAction<{
        groupId: string;
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
        groupId: string;
        semester: Semester;
        discId: string;
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
      { payload }: PayloadAction<{ groupId: string; patch: Partial<Group> }>,
    ) {
      const g = state.list.find((gr) => gr.id === payload.groupId);
      if (g) Object.assign(g, payload.patch);
    },

    addGroup(state, { payload }: PayloadAction<Group>) {
      state.list = [payload, ...state.list];
    },
    deleteGroup(state, { payload }: PayloadAction<string>) {
      state.list = state.list.filter((g) => g.id !== payload);
    },
    duplicateGroup(state, { payload }: PayloadAction<string>) {
      const src = state.list.find((g) => g.id === payload);
      if (!src) return;
      state.list = [
        ...state.list,
        {
          ...JSON.parse(JSON.stringify(src)), // deep clone
          id: crypto.randomUUID(),
          name: src.name + "0",
          isExpanded: false,
        },
      ];
    },
  },
});

export const {
  setGroups,
  toggleExpand,
  addDiscipline,
  deleteDiscipline,
  updateGroup,
  addGroup,
  deleteGroup,
  duplicateGroup,
} = groupSlice.actions;

export const groupReducer = groupSlice.reducer;
