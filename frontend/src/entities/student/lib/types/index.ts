export interface UserState {
  listStudents: []
}

export interface Student {
  id: string
  fullName: string
  groups: string
  semester: number
  company: string | null
  startDateWorkd: Date
  position: string
}
