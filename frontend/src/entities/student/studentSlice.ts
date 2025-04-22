// import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
// // Import Wails backend methods for students
// import { Inf_AllStudent, Create_Student, Update_StudentById, Delete_Student } from '../../wailsjs/go/main/App';
//
// const initialState = {
//     list: [],            // list of student objects fetched from backend
//     loading: false,      // indicates if a request is in progress
//     error: null          // holds error message if a request fails
// };
//
// // Thunk to fetch all students from backend (replaces any mock data usage)
// export const fetchAllStudents = createAsyncThunk(
//     'students/fetchAll',
//     async () => {
//         const students = await Inf_AllStudent();  // call Go backend to get all students
//         return students;  // this will be the fulfilled action payload
//     }
// );
//
// // Thunk to create a new student via backend
// export const createStudent = createAsyncThunk(
//     'students/create',
//     async (newStudentData) => {
//         const createdStudent = await Create_Student(newStudentData);  // call backend to create student
//         return createdStudent;  // assuming backend returns the created student object (with new ID)
//     }
// );
//
// // Thunk to update an existing student by ID via backend
// export const updateStudent = createAsyncThunk(
//     'students/update',
//     async ({ id, updates }) => {
//         const updatedStudent = await Update_StudentById(id, updates);  // call backend to update student
//         return updatedStudent;  // backend returns updated student object
//     }
// );
//
// // Thunk to delete a student by ID via backend
// export const deleteStudent = createAsyncThunk(
//     'students/delete',
//     async (id) => {
//         await Delete_Student(id);   // call backend to delete student (no return expected on success)
//         return id;  // return the deleted student's ID to remove from state
//     }
// );
//
// const studentsSlice = createSlice({
//     name: 'students',
//     initialState,
//     reducers: {
//         // (No direct reducers for CRUD since we use thunks; you could add non-async reducers if needed)
//     },
//     extraReducers: (builder) => {
//         // Fetch all students lifecycle
//         builder.addCase(fetchAllStudents.pending, (state) => {
//             state.loading = true;
//             state.error = null;
//         });
//         builder.addCase(fetchAllStudents.fulfilled, (state, action) => {
//             state.loading = false;
//             state.list = action.payload;  // replace list with fetched students
//         });
//         builder.addCase(fetchAllStudents.rejected, (state, action) => {
//             state.loading = false;
//             state.error = action.error.message || 'Failed to load students';
//         });
//
//         // Create student lifecycle
//         builder.addCase(createStudent.pending, (state) => {
//             state.error = null;
//         });
//         builder.addCase(createStudent.fulfilled, (state, action) => {
//             state.list.push(action.payload);  // add the new student to list
//         });
//         builder.addCase(createStudent.rejected, (state, action) => {
//             state.error = action.error.message || 'Failed to create student';
//         });
//
//         // Update student lifecycle
//         builder.addCase(updateStudent.fulfilled, (state, action) => {
//             const updated = action.payload;
//             // find and update the student in list
//             const index = state.list.findIndex(s => s.id === updated.id);
//             if (index !== -1) {
//                 state.list[index] = updated;
//             }
//         });
//         builder.addCase(updateStudent.rejected, (state, action) => {
//             state.error = action.error.message || 'Failed to update student';
//         });
//
//         // Delete student lifecycle
//         builder.addCase(deleteStudent.fulfilled, (state, action) => {
//             const deletedId = action.payload;
//             state.list = state.list.filter(s => s.id !== deletedId);
//         });
//         builder.addCase(deleteStudent.rejected, (state, action) => {
//             state.error = action.error.message || 'Failed to delete student';
//         });
//     }
// });
//
// export const studentReducer = studentsSlice.reducer;
