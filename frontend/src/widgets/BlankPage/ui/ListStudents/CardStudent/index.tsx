import classNames from "classnames";
import { Student } from "@/entities/student/types";

interface CardStudentProps {
  student: Student;
  handleSelect: (student: Student) => void;
}

export function CardStudent({ student, handleSelect }: CardStudentProps) {
  return (
    <div
      className={classNames(
        "px-4 py-2 cursor-pointer hover:bg-gray-100 border-b overflow-hidden whitespace-nowrap text-ellipsis",
      )}
      onClick={() => handleSelect(student)}
    >
      {student.fullName}
      {/*
      {selectStudent?.id && (
        <button
          type="button"
          className="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
          onClick={() =>
            dispatch(deleteStudentThunks(selectStudent.id)).then(() =>
              setValue("studentName", ""),
            )
          }
        >
          Удалить
        </button>
      )}
      */}
    </div>
  );
}
