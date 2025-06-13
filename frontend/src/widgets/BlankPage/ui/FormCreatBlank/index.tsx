import { GraduatesToggle } from "@/shared/ui/GraduatesToggle";
import { Semester } from "@/entities/discipline/types";
import { FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { Student } from "@/entities/student/types";
import {
  generatePdfThunks,
  saveOrUpdateStudentThunks,
} from "@/entities/blank/store/thunks.ts";
import {
  FieldErrors,
  SubmitHandler,
  UseFormGetValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormSetError,
  UseFormSetValue,
} from "react-hook-form";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { Group } from "@/entities/group/types";

interface FormCreatBlankProps {
  group: Group;
  selectStudent?: Student;
  setValue: UseFormSetValue<FormValuesBlank>;
  getValues: UseFormGetValues<FormValuesBlank>;
  errors: FieldErrors<FormValuesBlank>;
  setError: UseFormSetError<FormValuesBlank>;
  register: UseFormRegister<FormValuesBlank>;
  handleSubmit: UseFormHandleSubmit<FormValuesBlank>;
}

export function FormCreatBlank({
  group,
  setValue,
  selectStudent,
  getValues,
  setError,
  errors,
  register,
  handleSubmit,
}: FormCreatBlankProps) {
  const dispatch = useAppDispatch();

  /* ---------- 3. обработчик ---------- */
  const onSubmit: SubmitHandler<FormValuesBlank> = (data) => {
    if (!selectStudent?.id) {
      alert("Выберите студента!");
      return;
    }

    console.log({
      groupId: group?.id,
      selectStudentId: selectStudent?.id,
      studentName: data.studentName,
      semester: data.semester,
      company: data.company ?? null,
      startDate: data.startDate,
      position: data.position,
    });
    alert("Строка добавлена!");
  };

  const onSave = async (values: FormValuesBlank) => {
    if (!group) return;
    /* формируем объект */
    const student: Student = {
      id: selectStudent?.id ?? 0, // 0 → create
      fullName: values.studentName,
      company: values.company,
      startDateWork: values.startDate,
      position: values.position,
    };

    await dispatch(saveOrUpdateStudentThunks({ groupId: group.id, student }));
  };

  const onDownload = async (values: FormValuesBlank) => {
    if (!group) return;
    await onSave(values); // save OR update
    const latest = selectStudent
      ? { ...selectStudent, ...values }
      : { id: 0, ...values }; // fallback (не должно случиться)
    await dispatch(
      generatePdfThunks({
        group,
        student: latest as Student,
        semester: values.semester as Semester,
      }),
    );
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-4 space-y-4 w-full">
      <div>
        <GraduatesToggle
          title="Семестр:"
          variant={["1", "2"]}
          onChange={(num) => {
            setValue("semester", num as Semester);
            setError("semester", {});
          }}
          value={getValues("semester")}
          error={errors.semester?.message}
        />
      </div>

      <label className="block max-w-sm">
        <span className="font-semibold">Предприятие</span>
        <input
          type="text"
          {...register("company")}
          className="border rounded p-1 w-full"
          placeholder="ООО «Рога и копыта»"
        />
      </label>

      <label className="block max-w-sm">
        <span className="font-semibold">Дата начала</span>
        <input
          type="date"
          {...register("startDate")}
          className="border rounded p-1 w-full"
        />
        {errors.startDate && (
          <span className="text-red-500 text-xs">
            {errors.startDate.message}
          </span>
        )}
      </label>

      <label className="block max-w-sm">
        <span className="font-semibold">Должность</span>
        <input
          type="text"
          {...register("position")}
          className="border rounded p-1 w-full"
        />
        {errors.position && (
          <span className="text-red-500 text-xs">
            {errors.position.message}
          </span>
        )}
      </label>

      <div className="flex gap-4 pt-2">
        <button
          type="submit"
          className="bg-orange-500 hover:bg-orange-600 text-white font-semibold px-4 py-2 rounded"
          onClick={handleSubmit(onSave)}
        >
          Сохранить
        </button>
        <button
          type="button"
          className="bg-green-600 hover:bg-green-700 text-white font-semibold px-4 py-2 rounded"
          onClick={handleSubmit(onDownload)}
        >
          Скачать бланк
        </button>
      </div>
    </form>
  );
}
