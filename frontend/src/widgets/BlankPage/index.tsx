import { VirtualizedSearch } from "@/features/VirtualizedSearch";
import { EditableTitle } from "@/shared/ui/EditableTitle";

import { z } from "zod";
import { useForm, SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { useEffect, useState } from "react";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { setStudent } from "@/entities/blank/store";
import { Group } from "@/entities/group/types";

import { handleTitleGroupSave } from "@/shared/lib/headers/titleGroupSave.ts";
import { GraduatesToggle } from "@/shared/ui/GraduatesToggle";
import { Student } from "@/entities/student/types";
import {
  deleteStudentThunks,
  generatePdfThunks,
  saveOrUpdateStudentThunks,
} from "@/entities/blank/store/thunks.ts";

export function BlankPage({ group }: { group: Group | undefined }) {
  const [err, setErr] = useState<string | null>(null);
  const dispatch = useAppDispatch();
  const {
    selectStudent,
    // studentsData,
    semester,
  } = useAppSelector(selectBlank);
  const studentsData: Student[] = [
    {
      id: 1,
      fullName: "Кванов Иван Иванович",
      company: "ООО «Рога и копыта»",
      startDateWork: new Date().toISOString().slice(0, 10),
      position: "Стажер",
    },
    {
      id: 2,
      fullName: "Иванов Иван",
      company: "ООО «Рога и копыта»",
      startDateWork: new Date().toISOString().slice(0, 10),
      position: "Стажер",
    },
    {
      id: 3,
      fullName: "Ивов Иван Иванович",
      company: "ООО «Рога и копыта»",
      startDateWork: new Date().toISOString().slice(0, 10),
      position: "Стажер",
    },
  ];

  /* ---------- 1. Updated Zod schema to fix resolver type mismatch ---------- */
  const semesterEnum = ["1", "2"] as const;
  type Semester = (typeof semesterEnum)[number];

  // Updated preprocessing function with proper typing
  const toSemesterOrUndef = (v: unknown): Semester | undefined => {
    if (typeof v !== "string") return undefined;
    return semesterEnum.includes(v as Semester) ? (v as Semester) : undefined;
  };

  // Schema using refined preprocessing
  const schema = z.object({
    semester: z
      .custom<Semester | undefined>(toSemesterOrUndef, {
        message: "Invalid semester",
      })
      .optional(),
    studentName: z.string().min(1, "Обязательное поле"),
    company: z.string().optional(),
    startDate: z.string().optional(),
    position: z.string().min(2),
  });
  type FormValues = z.infer<typeof schema>;

  /* ---------- 2. useForm ---------- */
  const {
    register,
    handleSubmit,
    setValue,
    setError,
    getValues,
    formState: { errors },
  } = useForm<FormValues>({
    resolver: zodResolver(schema), // generic выводится автоматически
    defaultValues: {
      semester: "1",
      studentName: "",
      company: "",
      startDate: "",
      position: "",
    },
  });

  /* ---------- 3. обработчик ---------- */
  const onSubmit: SubmitHandler<FormValues> = (data) => {
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

  useEffect(() => {
    if (!selectStudent) return;

    setValue("studentName", selectStudent?.fullName);
    if (selectStudent?.startDateWork) {
      setValue("startDate", selectStudent?.startDateWork);
    }
    setValue("semester", semester);
    setValue("company", selectStudent?.company ?? "");
    setValue("position", selectStudent?.position ?? "");
  }, [setValue, semester]);

  const onSave = async (values: FormValues) => {
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

  const onDownload = async (values: FormValues) => {
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
      <div className="flex items-center justify-between">
        {group ? (
          <h1 className="text-2xl font-bold">
            Группа:
            <EditableTitle
              key={group?.name}
              initialValue={group?.name}
              onSave={async (value) => {
                const ok = await handleTitleGroupSave({
                  dispatch,
                  group,
                  value,
                });
                setErr(ok ? null : "Неверный формат имени");
              }}
              className="ml-1 inline-block text-2xl font-bold"
              error={err || undefined}
            />
          </h1>
        ) : (
          <h1 className="text-2xl font-bold">Группа не выбрана</h1>
        )}
      </div>

      <div className="flex gap-4">
        <div className="col-2">
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
        <div className="col-4">
          <VirtualizedSearch
            data={studentsData}
            placeholder="Введите ФИО..."
            maxDropdownHeight={200}
            onSelect={(s) => {
              setValue("studentName", s.fullName, { shouldValidate: true });
              dispatch(setStudent({ id: s.id, fullName: s.fullName }));
              setValue("company", s.company ?? "");
              setValue("startDate", s.startDateWork ?? "");
              setValue("position", s.position ?? "");
            }}
            error={errors.studentName?.message}
          />
        </div>
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
