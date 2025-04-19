import { Id } from '../types'

export const API_ENDPOINTS = {
  common: {
    baseUrl: import.meta.env.VITE_API_URL || 'https://api.example.com',
  },
  groups: {
    getGroup: (groupsId: Id) => `/api/Group/${groupsId}`,
    getList: '/api/Group/AllGroup',
    create: '/api/Group',
    update: (groupsId: Id) => `/api/Group/${groupsId}`,
    delete: (groupsId: Id) => `/api/Group/${groupsId}`,
  },
  disciplines: {
    getList: (groupsId: Id) => `/api/Disciplines/DisciplinesInGroup/${groupsId}`,
    create: '/api/Group',
    update: (id: Id) => `/api/Group/${id}`,
    delete: (id: Id) => `/api/Group/${id}`,
  },
  employerS: {
    downloadPdfDocs: '/PdfDocs',
  },
}
