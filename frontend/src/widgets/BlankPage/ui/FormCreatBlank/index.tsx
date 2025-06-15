import { GraduatesToggle } from "@/shared/ui/GraduatesToggle";
import { Semester } from "@/entities/discipline/types";
import { FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { Student } from "@/entities/student/types";
import {
  generatePdfThunks,
  saveOrUpdateStudentThunks,
} from "@/entities/blank/store/thunks.ts";
import {
  Control,
  Controller,
  FieldErrors,
  UseFormGetValues,
  UseFormHandleSubmit,
  UseFormRegister,
  UseFormSetError,
  UseFormSetValue,
} from "react-hook-form";
import { useAppDispatch } from "@/shared/lib/hooks/redux.ts";
import { Group } from "@/entities/group/types";
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

interface FormCreatBlankProps {
  group: Group;
  selectStudent?: Student;
  setValue: UseFormSetValue<FormValuesBlank>;
  getValues: UseFormGetValues<FormValuesBlank>;
  errors: FieldErrors<FormValuesBlank>;
  setError: UseFormSetError<FormValuesBlank>;
  register: UseFormRegister<FormValuesBlank>;
  handleSubmit: UseFormHandleSubmit<FormValuesBlank>;
  control: Control<FormValuesBlank>;
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
  control,
}: FormCreatBlankProps) {
  const dispatch = useAppDispatch();

  const onSave = async (values: FormValuesBlank) => {
    if (!group) return;
    const student: Student = {
      id: selectStudent?.id ?? 0,
      fullName: values.studentName,
      company: values.company,
      startDateWork: values.startDate,
      position: values.position,
    };

    await dispatch(saveOrUpdateStudentThunks({ groupId: group.id, student }));
  };

  const onDownload = async (values: FormValuesBlank) => {
    if (!group) return;
    await onSave(values);
    const latest = selectStudent
      ? { ...selectStudent, ...values }
      : { id: 0, ...values };
    await dispatch(
      generatePdfThunks({
        group,
        student: latest as Student,
        semester: values.semester as Semester,
      }),
    );
  };

  return (
    <div className="pr-3 space-y-4 w-full">
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
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <Controller
            control={control}
            name="startDate"
            render={({ field }) => (
              <DatePicker
                onChange={(date) => {
                  const formatted = date ? date.format("YYYY-MM-DD") : "";
                  field.onChange(formatted);
                }}
                slotProps={{
                  textField: { className: "border rounded p-1 w-full" },
                }}
              />
            )}
          />
        </LocalizationProvider>
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
        <div>
          {errors.position && (
            <span className="text-red-500 text-xs ">
              {errors.position.message}
            </span>
          )}
        </div>
      </label>

      <div className="flex gap-4 pt-2">
        <button
          type="button"
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
    </div>
  );
}
