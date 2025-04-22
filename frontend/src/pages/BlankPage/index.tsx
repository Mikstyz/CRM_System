import { VirtualizedSearch } from "@/features/VirtualizedSearch";
import { EditableTitle } from "@/shared/ui/EditableTitle";
import { LinkDocument } from "@/shared/ui/LinkDocument";

import { z } from "zod";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { MIN_SEMESTER, MAX_SEMESTER } from "@/shared/const";
import { selectBlank, setStudent } from "@/entities/blank/store/blankSlice";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { useEffect } from "react";

export function BlankPage({
  groupId,
  groupName,
}: {
  groupId: string | null;
  groupName: string;
}) {
  const dispatch = useAppDispatch();
  const {
    studentId,
    studentName,
    semester,
    company,
    startDate,
    position,
    studentsData,
  } = useAppSelector(selectBlank);

  /* ---------- 1. схема Zod ---------- */
  const schema = z.object({
    semester: z.coerce.number().min(1).max(2, "От 1 до 2"),
    student: z.string().min(1, "Обязательное поле"),
    company: z.string().optional(),
    startDate: z.string().regex(/^\d{4}-\d{2}-\d{2}$/, "YYYY‑MM‑DD"),
    position: z.string().min(2),
  });
  type FormValues = z.infer<typeof schema>;

  /* ---------- 2. useForm ---------- */
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<FormValues>({
    resolver: zodResolver(schema), // generic выводится автоматически
    defaultValues: {
      semester: 1,
      student: "",
      company: "",
      startDate: new Date().toISOString().slice(0, 10),
      position: "",
    },
  });

  /* ---------- 3. обработчик ---------- */
  const onSubmit: SubmitHandler<FormValues> = (data) => {
    if (!studentId) {
      alert("Выберите студента!");
      return;
    }

    console.log({
      groupId,
      studentId, // ← настоящий id
      student: data.student,
      semester: data.semester,
      company: data.company ?? null,
      startDate: data.startDate,
      position: data.position,
    });
    alert("Строка добавлена!");
  };

  useEffect(() => {
    setValue("student", studentName); // берём прямо из slice
    setValue("semester", semester);
    setValue("company", company ?? "");
    setValue("startDate", startDate);
    setValue("position", position);
  }, [setValue, studentName, semester, company, startDate, position]);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-4 space-y-4 w-full">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">
          Группа:
          <EditableTitle
            initialValue={groupName}
            onSave={() => {}}
            className="ml-1 inline-block text-2xl font-bold"
          />
        </h1>
        <LinkDocument href="#">Открыть XLSX студентов</LinkDocument>
      </div>

      <div className="flex gap-4">
        <label className="flex flex-col w-24">
          <span className="font-semibold">Семестр</span>
          <input
            type="number"
            min={MIN_SEMESTER}
            max={MAX_SEMESTER}
            {...register("semester", { valueAsNumber: true })}
            className="border rounded p-1"
          />
          {errors.semester && (
            <span className="text-red-500 text-xs">
              {errors.semester.message}
            </span>
          )}
        </label>

        <VirtualizedSearch
          data={studentsData}
          placeholder="Введите ФИО..."
          maxDropdownHeight={200}
          onSelect={(s) => {
            // 1) пишем ФИО в форму
            setValue("student", s.fullName, { shouldValidate: true });
            // 2) сохраняем id + ФИО в blankSlice
            dispatch(setStudent({ id: s.id, fullName: s.fullName }));
          }}
        />
      </div>
      {errors.student && (
        <span className="text-red-500 text-xs">{errors.student.message}</span>
      )}

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
        >
          Сохранить
        </button>
        <button
          type="button"
          className="bg-green-600 hover:bg-green-700 text-white font-semibold px-4 py-2 rounded"
          onClick={() => console.log("download")}
        >
          Скачать бланк
        </button>
      </div>
    </form>
  );
}
