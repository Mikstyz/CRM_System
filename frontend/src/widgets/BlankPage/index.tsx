import {VirtualizedSearch} from "@/features/VirtualizedSearch";
import {EditableTitle} from "@/shared/ui/EditableTitle";

import {z} from "zod";
import {useForm, SubmitHandler} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {MIN_SEMESTER, MAX_SEMESTER} from "@/shared/const";
import {useAppDispatch, useAppSelector} from "@/shared/lib/hooks/redux.ts";
import {useEffect} from "react";
import {selectBlank} from "@/entities/blank/store/selectors.ts";
import {setStudent} from "@/entities/blank/store";
import {Group} from "@/entities/group/types";

import {handleTitleGroupSave} from "@/shared/lib/headers/titleGroupSave.ts";

export function BlankPage({group}: {
  group: Group | undefined
}) {
  const dispatch = useAppDispatch();
  const {
    selectStudent,
    studentsData,
    semester

  } = useAppSelector(selectBlank);

  /* ---------- 1. Updated Zod schema to fix resolver type mismatch ---------- */
  const semesterEnum = ["1", "2", "3", "4"] as const;
  type Semester = typeof semesterEnum[number];

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
    startDate: z.date(),
    position: z.string().min(2),
  });
  type FormValues = z.infer<typeof schema>;

  /* ---------- 2. useForm ---------- */
  const {
    register,
    handleSubmit,
    setValue,
    formState: {errors},
  } = useForm<FormValues>({
    resolver: zodResolver(schema), // generic выводится автоматически
    defaultValues: {
      semester: "1",
      studentName: "",
      company: "",
      startDate: new Date(),
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
    if (!selectStudent) return

    setValue("studentName", selectStudent?.fullName);
    if (selectStudent?.startDateWork) {
      setValue("startDate", selectStudent?.startDateWork);
    }
    setValue("semester", semester);
    setValue("company", selectStudent?.company ?? "");
    setValue("position", selectStudent?.position?? "");
  }, [setValue, semester]);

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="p-4 space-y-4 w-full">
      <div className="flex items-center justify-between">
        {
        group ?
        (<h1 className="text-2xl font-bold">
          Группа:
          <EditableTitle
            initialValue={group?.name}
            onSave={(value) => handleTitleGroupSave({ dispatch, group, value })}
            className="ml-1 inline-block text-2xl font-bold"
          />
        </h1>) : (
          <h1 className="text-2xl font-bold">
            Группа не выбрана
          </h1>
        )
        }
      </div>

      <div className="flex gap-4">
        <label className="flex flex-col w-24">
          <span className="font-semibold">Семестр</span>
          <input
            type="number"
            min={MIN_SEMESTER}
            max={MAX_SEMESTER}
            {...register("semester", {valueAsNumber: true})}
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
            setValue("studentName", s.fullName, {shouldValidate: true});
            // 2) сохраняем id + ФИО в blankSlice
            dispatch(setStudent({id: s.id, fullName: s.fullName}));
          }}
        />
      </div>
      {errors.studentName && (
        <span className="text-red-500 text-xs">{errors.studentName.message}</span>
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
