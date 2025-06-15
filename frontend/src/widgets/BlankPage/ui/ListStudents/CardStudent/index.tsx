import classNames from "classnames";
import { Student } from "@/entities/student/types";

interface CardStudentProps {
  student: Student;
  handleSelect: (student: Student) => void;
  onDelete: (id: number) => void;
}

export function CardStudent({
  student,
  handleSelect,
  onDelete,
}: CardStudentProps) {
  return (
    <div
      className={classNames(
        "group relative px-4 py-2 cursor-pointer hover:bg-gray-100 border-b overflow-hidden whitespace-nowrap text-ellipsis",
      )}
      onClick={() => handleSelect(student)}
    >
      {student.fullName}
      <button
        type="button"
        title="Удалить студента"
        onClick={(e) => {
          e.stopPropagation();
          onDelete(student.id);
        }}
        className="hidden cursor-pointer group-hover:block absolute right-2 top-1/2 -translate-y-1/2 text-red-500"
      >
        <div
          className="inline-flex h-6 w-6 items-center justify-center
         rounded-full bg-gray-200 text-gray-700
         hover:bg-red-500 hover:text-white
         focus:outline-none focus:ring-2 focus:ring-red-500/50
         transition"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            className="h-4 w-4"
          >
            <path d="M4.22 4.22a.75.75 0 0 1 1.06 0L10 8.94l4.72-4.72a.75.75 0 1 1 1.06 1.06L11.06 10l4.72 4.72a.75.75 0 1 1-1.06 1.06L10 11.06l-4.72 4.72a.75.75 0 0 1-1.06-1.06L8.94 10 4.22 5.28a.75.75 0 0 1 0-1.06z" />
          </svg>
        </div>
      </button>
    </div>
  );
}
