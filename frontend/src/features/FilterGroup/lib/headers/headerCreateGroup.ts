import { AppDispatch } from "@/app/store";
import { UseFormSetError } from "react-hook-form";
import { FiltersRaw } from "@/features/FilterGroup/schema";
import { Course, Graduates } from "@/entities/group/types";
import { createGroupsThunks } from "@/entities/group/store/thunks.ts";

interface Arguments {
  course: Course | null | undefined;
  specialty: string | null | undefined;
  graduates: Graduates | null | undefined;
  groupNumber: number | null | undefined;
  dispatch: AppDispatch;
  setError: UseFormSetError<FiltersRaw>;
}

export const headerCreateGroup = async ({
  course,
  specialty,
  graduates,
  groupNumber,
  dispatch,
  setError,
}: Arguments) => {
  if (course && specialty && graduates && groupNumber) {
    const validCourse = ["1", "2", "3", "4"].includes(course.toString())
      ? course
      : undefined;
    if (!validCourse) {
      setError("course", {
        type: "custom",
        message: "Выберите допустимый курс (1, 2, 3, 4)",
      });
      return;
    }
    dispatch(
      createGroupsThunks({
        dateNameGroup: {
          course: validCourse,
          specialty,
          graduates,
          groupNumber: Number(groupNumber),
        },
      }),
    );
  } else {
    if (!course) {
      setError("course", {
        type: "custom",
        message: "Выберите курс",
      });
    }
    if (!specialty) {
      setError("specialty", {
        type: "custom",
        message: "Выберите специальность",
      });
    }
    if (!graduates) {
      setError("graduates", {
        type: "custom",
        message: "Выберите выпускников",
      });
    }
    if (!groupNumber) {
      setError("groupNumber", {
        type: "custom",
        message: "Выберите номер группы",
      });
    }
  }
};
