// export interface Group {
//   id: string
//   course: number
//   semester: number
//   specialty: string
//   graduation: string
//   groupsNum: string
// }

import { Discipline } from '@/entities/discipline/lib/types'

//////////////////////////////////////////////////

export interface Group {
  id: string
  name: string
  isExpanded: boolean
  disciplines: Discipline[]
}
