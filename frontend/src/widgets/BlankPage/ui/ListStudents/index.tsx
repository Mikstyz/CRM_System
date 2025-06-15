import { deleteStudentThunks, setStudent } from "@/entities/blank/store";
import { Student } from "@/entities/student/types";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { FieldErrors, UseFormSetValue } from "react-hook-form";
import { FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import {
  ChangeEvent,
  KeyboardEvent,
  useCallback,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import { CardStudent } from "@/widgets/BlankPage/ui/ListStudents/CardStudent";
import { Pagination } from "@/shared/ui/Pagination";

const PER_PAGE = 6;

interface ListStudentsProps {
  studentsData: Student[];
  setValue: UseFormSetValue<FormValuesBlank>;
  errors: FieldErrors<FormValuesBlank>;
}

export function ListStudents({
  studentsData,
  setValue,
  errors,
}: ListStudentsProps) {
  const dispatch = useAppDispatch();
  const onSelect = (s: Student) => {
    setValue("studentName", s.fullName, { shouldValidate: true });
    dispatch(setStudent({ id: s.id, fullName: s.fullName }));
    setValue("company", s.company ?? "");
    setValue("startDate", s.startDateWork ?? "");
    setValue("position", s.position ?? "");
  };

  const [query, setQuery] = useState("");
  const [allStudents, setAllStudents] = useState<Student[]>(studentsData);
  const [filteredData, setFilteredData] = useState<Student[]>(studentsData);

  const inputRef = useRef<HTMLInputElement | null>(null);
  const containerRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    setAllStudents(studentsData);
  }, [studentsData]);

  // Фильтруем результаты при изменении запроса
  useEffect(() => {
    if (!query) {
      setFilteredData(allStudents);
    } else {
      const lowerQuery = query.toLowerCase();
      setFilteredData(
        allStudents.filter((item) =>
          item.fullName.toLowerCase().includes(lowerQuery),
        ),
      );
    }
  }, [query, allStudents]);

  // Обработчик выбора элемента
  const handleSelect = useCallback(
    (student: Student) => {
      setQuery(student.fullName);
      onSelect?.(student);
    },
    [onSelect],
  );

  const handleDelete = useCallback(
    (id: number) => {
      setAllStudents((prev) => prev.filter((s) => s.id !== id));
      dispatch(deleteStudentThunks(id));
    },
    [dispatch],
  );

  // Обработчик ввода в поисковое поле
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
    onSelect?.({
      id: Date.now(),
      fullName: e.target.value,
      company: "",
      startDateWork: undefined,
      position: "",
    });
  };

  // При нажатии Enter выбираем первый элемент (при желании)
  const handleKeyDown = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter" && filteredData[0]) {
      handleSelect(filteredData[0]);
    }
  };

  const [page, setPage] = useState(1);

  useEffect(() => setPage(1), [filteredData]);

  const paginated = useMemo(() => {
    const start = (page - 1) * PER_PAGE;
    return filteredData.slice(start, start + PER_PAGE);
  }, [filteredData, page]);

  return (
    <div className="h-full relative">
      <div ref={containerRef} className="inline-block relative ">
        <span className="mr-1">Студенты:</span>
        {errors.studentName?.message && (
          <span className="text-red-500 text-xs">
            {errors.studentName?.message}
          </span>
        )}
        <div>
          <input
            ref={inputRef}
            type="text"
            className="
          border
          rounded
          px-3
          py-2
          focus:outline-none
          focus:ring-2
          focus:ring-blue-400
          w-64
        "
            placeholder="Введите ФИО..."
            value={query}
            onChange={handleChange}
            onKeyDown={handleKeyDown}
          />
        </div>
      </div>
      {paginated.length > 0 && (
        <>
          <div
            className="
            mt-1
            border
            bg-white
            rounded
            shadow-lg
            w-full
          "
          >
            {paginated.map((student: Student) => (
              <CardStudent
                key={student.id}
                student={student}
                handleSelect={handleSelect}
                onDelete={handleDelete}
              />
            ))}
          </div>
        </>
      )}
      <div className="absolute bottom-0 w-full">
        <Pagination
          page={page}
          pages={Math.ceil(filteredData.length / PER_PAGE)}
          onChange={setPage}
        />
      </div>
    </div>
  );
}
