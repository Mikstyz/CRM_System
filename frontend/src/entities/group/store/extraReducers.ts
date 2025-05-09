import {
  addDisciplinesThunks,
  createGroupsThunks,
  deleteDisciplinesThunks,
  deleteGroupsThunks,
  duplicateGroupThunks,
  getGroupsThunks,
  updateDisciplinesThunks,
  updateGroupsThunks,
} from "@/entities/group/store/thunks.ts";
import {
  handlePending,
  handleRejected,
} from "@/shared/lib/helpers/StoreHandlers.ts";
import { GroupsState } from "@/entities/group/store/initialState.ts";
import { ActionReducerMapBuilder } from "@reduxjs/toolkit";
import { Id } from "@/shared/types";

export const groupsExtraReducers = (
  builder: ActionReducerMapBuilder<GroupsState>,
) => {
  // ========= GROUPS =========

  // GET GROUPS
  builder
    .addCase(getGroupsThunks.pending, handlePending<GroupsState>)
    .addCase(getGroupsThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.list = action.payload;
    })
    .addCase(getGroupsThunks.rejected, handleRejected<GroupsState>);

  // CREATE GROUP
  builder
    .addCase(createGroupsThunks.pending, handlePending<GroupsState>)
    .addCase(createGroupsThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.list = [action.payload, ...state.list];
    })
    .addCase(createGroupsThunks.rejected, handleRejected<GroupsState>);

  // UPDATE GROUP
  builder
    .addCase(updateGroupsThunks.pending, handlePending<GroupsState>)
    .addCase(updateGroupsThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.list = state.list.map((group) =>
        group.id === action.payload.id
          ? { ...group, ...action.payload }
          : group,
      );
    })
    .addCase(updateGroupsThunks.rejected, handleRejected<GroupsState>);

  //  DELETE GROUP
  builder
    .addCase(deleteGroupsThunks.pending, handlePending<GroupsState>)
    .addCase(deleteGroupsThunks.fulfilled, (state, action) => {
      state.loading = false;
      state.list = state.list.filter((group) => group.id !== action.payload);
    })
    .addCase(deleteGroupsThunks.rejected, handleRejected<GroupsState>);

  // DUPLICATE GROUP
  builder
    .addCase(duplicateGroupThunks.pending, handlePending<GroupsState>)
    .addCase(duplicateGroupThunks.fulfilled, (state, action) => {
      state.loading = false;
      const originalId = action.meta.arg as Id;
      const srcIndex = state.list.findIndex((g) => g.id === originalId);
      if (srcIndex === -1) {
        state.list.unshift(action.payload);
        return;
      }
      const exists = state.list.some((g) => g.id === action.payload.id);
      if (!exists) {
        state.list.splice(srcIndex + 1, 0, action.payload);
      }
      // const already = state.list.find((g) => g.id === action.payload.id);
      // if (!already) {
      //   state.list.unshift(action.payload);
      // }
      // const srcIndex = state.list.findIndex((g) => g.id === action.payload.id);
      // if (srcIndex === -1) return;
      // state.list.splice(srcIndex + 1, 0, action.payload);
    })
    .addCase(duplicateGroupThunks.rejected, handleRejected<GroupsState>);

  // ========= disciplines =========

  // ADD DISCIPLINE
  builder
    .addCase(addDisciplinesThunks.pending, handlePending<GroupsState>)
    .addCase(addDisciplinesThunks.fulfilled, (state, action) => {
      state.loading = false;
      const g = state.list.find((gr) => gr.id === action.payload.groupId);
      if (g) {
        g.disciplines[action.payload.semester].unshift(action.payload.disc);
      }
    })
    .addCase(addDisciplinesThunks.rejected, handleRejected<GroupsState>);

  // UPDATE DISCIPLINE
  builder
    .addCase(updateDisciplinesThunks.pending, handlePending<GroupsState>)
    .addCase(updateDisciplinesThunks.fulfilled, (state, action) => {
      state.loading = false;
      const g = state.list.find((gr) => gr.id === action.payload.groupId);
      if (g) {
        g.disciplines[action.payload.semester] = g.disciplines[
          action.payload.semester
        ].map((d) =>
          d.id === action.payload.disc.id ? action.payload.disc : d,
        );
      }
    })
    .addCase(updateDisciplinesThunks.rejected, handleRejected<GroupsState>);

  // DELETE DISCIPLINE
  builder
    .addCase(deleteDisciplinesThunks.pending, handlePending<GroupsState>)
    .addCase(deleteDisciplinesThunks.fulfilled, (state, action) => {
      state.loading = false;
      const g = state.list.find((gr) => gr.id === action.payload.groupId);
      if (g) {
        g.disciplines[action.payload.semester] = g.disciplines[
          action.payload.semester
        ].filter((d) => d.id !== action.payload.discId);
      }
    })
    .addCase(deleteDisciplinesThunks.rejected, handleRejected<GroupsState>);
};
