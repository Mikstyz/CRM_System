import { useState } from 'react'
import { InputFilterGrup } from '../InputFilterGrup'

export function FilterGrup() {
  const [course, setCourse] = useState('1')
  const [specialty, setSpecialty] = useState('ИСП')
  const [graduates, setGraduates] = useState('11/11')
  const [groupNumber, setGroupNumber] = useState('45')

  return (
    <div className="border p-4 rounded-lg mb-4 w-[50%] min-w-min">
      <h2 className="font-semibold mb-2">Фильтрация</h2>
      <div className="flex flex-wrap content-start justify-between gap-1">
        <InputFilterGrup title="Курс" value={course} onChange={(e) => setCourse(e.target.value)} />
        <InputFilterGrup title="Специальность" value={specialty} onChange={(e) => setSpecialty(e.target.value)} />
        <InputFilterGrup title="Выпускники (9/11)" value={graduates} onChange={(e) => setGraduates(e.target.value)} />
        <InputFilterGrup title="Номер группы" value={groupNumber} onChange={(e) => setGroupNumber(e.target.value)} />
      </div>
    </div>
  )
}
