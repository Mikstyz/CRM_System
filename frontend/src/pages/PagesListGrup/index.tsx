import { useState } from 'react'
import { FilterGrup } from '../../features/FilterGrup'
import { GroupCard } from '../../widgets/GroupCard'
import { Group } from '../../entities/group/lib/types'
import { LinkDocument } from '@/shared/ui/LinkDocument'

const initialGroups: Group[] = [
  {
    id: 'g1',
    name: '1ИСП11 - 45',
    isExpanded: false,
    disciplines: [
      { id: 'd1', title: 'Проектирование баз данных' },
      { id: 'd2', title: 'Проектирование баз данных' },
      { id: 'd3', title: 'Проектирование баз данных' },
      { id: 'd4', title: 'Проектирование баз данных' },
      { id: 'd5', title: 'Проектирование баз данных' },
      { id: 'd6', title: 'Проектирование баз данных' },
      { id: 'd7', title: 'Проектирование баз данных' },
      { id: 'd8', title: 'Проектирование баз данных' },
      { id: 'd9', title: 'Проектирование баз данных' },
      { id: 'd10', title: 'Проектирование баз данных' },
      { id: 'd11', title: 'Проектирование баз данных' },
      { id: 'd12', title: 'Проектирование баз данных' },
      { id: 'd13', title: 'Проектирование баз данных' },
      { id: 'd14', title: 'Проектирование баз данных' },
      { id: 'd15', title: 'Проектирование баз данных' },
      { id: 'd16', title: 'Проектирование баз данных' },
      { id: 'd17', title: 'Проектирование баз данных' },
      { id: 'd18', title: 'Проектирование баз данных' },
      { id: 'd19', title: 'Проектирование баз данных' },
      { id: 'd20', title: 'Проектирование баз данных' },
      { id: 'd21', title: 'Проектирование баз данных' },
      { id: 'd22', title: 'Проектирование баз данных' },
      { id: 'd23', title: 'Проектирование баз данных' },
      { id: 'd24', title: 'Проектирование баз данных' },
      { id: 'd25', title: 'Проектирование баз данных' },
      { id: 'd26', title: 'Проектирование баз данных' },
      { id: 'd27', title: 'Проектирование баз данных' },
      { id: 'd28', title: 'Проектирование баз данных' },
      { id: 'd29', title: 'Проектирование баз данных' },
      { id: 'd30', title: 'Проектирование баз данных' },
      { id: 'd31', title: 'Проектирование баз данных' },
      { id: 'd32', title: 'Проектирование баз данных' },
      { id: 'd33', title: 'Проектирование баз данных' },
      { id: 'd34', title: 'Проектирование баз данных' },
      { id: 'd35', title: 'Проектирование баз данных' },
      { id: 'd36', title: 'Проектирование баз данных' },
      { id: 'd37', title: 'Проектирование баз данных' },
    ],
  },
  {
    id: 'g2',
    name: '2ИСП9 - 45',
    isExpanded: false,
    disciplines: [{ id: 'd5', title: 'Проектирование баз данных' }],
  },
]

export function PagesListGrup() {
  const [groups, setGroups] = useState<Group[]>(initialGroups)

  // Пример функции для обновления группы по id (сменить isExpanded, изменить дисциплины и т.д.)
  const updateGroup = (groupId: string, newData: Partial<Group>) => {
    setGroups((prev) => prev.map((g) => (g.id === groupId ? { ...g, ...newData } : g)))
  }

  // Пример добавления новой дисциплины
  const addDiscipline = (groupId: string) => {
    const newDisc = {
      id: Math.random().toString(36).slice(2),
      title: 'Новая дисциплина',
    }
    setGroups((prev) => prev.map((g) => (g.id === groupId ? { ...g, disciplines: [...g.disciplines, newDisc] } : g)))
  }

  // Пример удаления дисциплины
  const deleteDiscipline = (groupId: string, disciplineId: string) => {
    setGroups((prev) =>
      prev.map((g) =>
        g.id === groupId
          ? {
              ...g,
              disciplines: g.disciplines.filter((d) => d.id !== disciplineId),
            }
          : g,
      ),
    )
  }

  // Пример удаления всей группы
  const deleteGroup = (groupId: string) => {
    setGroups((prev) => prev.filter((g) => g.id !== groupId))
  }

  return (
    <div className="p-4">
      <div className="flex gap-2">
        <h1 className="text-xl font-bold mb-4">Управление группами</h1>
        <LinkDocument href="#">Открыть xlsx групп</LinkDocument>
      </div>

      <FilterGrup />

      <div className="mt-6 space-y-4">
        {groups.map((group) => (
          <GroupCard
            key={group.id}
            group={group}
            onToggleExpand={() => updateGroup(group.id, { isExpanded: !group.isExpanded })}
            onAddDiscipline={() => addDiscipline(group.id)}
            onDeleteGroup={() => deleteGroup(group.id)}
            onDeleteDiscipline={(discId) => deleteDiscipline(group.id, discId)}
          />
        ))}
      </div>
    </div>
  )
}
