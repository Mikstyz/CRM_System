import { Button, Stack, TextField } from "@mui/material";
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { useFormContext, Controller } from "react-hook-form";
import { GraduatesToggle } from "@/shared/ui/GraduatesToggle";

import { BlankSchema, FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { Semester } from "@/entities/discipline/types";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { RootState } from "@/app/store";
import {
  generatePdfThunks,
  saveOrUpdateStudentThunks,
} from "@/entities/blank/store";
import dayjs from "dayjs";
import "dayjs/locale/ru";
import { ruRU } from "@mui/x-date-pickers/locales";

export function FormCreatBlank() {
  const dispatch = useAppDispatch();
  const { group } = useAppSelector((s: RootState) => s.blank);
  if (!group) return null;
  const { control, handleSubmit, setValue, watch, reset } =
    useFormContext<FormValuesBlank>();

  const onDownload = handleSubmit(async (dataSubmit: BlankSchema) => {
    try {
      const student = await dispatch(
        saveOrUpdateStudentThunks({
          groupId: group.id,
          student: {
            fullName: dataSubmit.studentName,
            company: dataSubmit.company,
            startDateWork: dataSubmit.startDate,
            position: dataSubmit.position,
          },
        }),
      ).unwrap();

      await dispatch(
        generatePdfThunks({
          group,
          student,
          semester: dataSubmit.semester,
        }),
      );

      reset();
    } catch (err) {
      console.error(err);
    }
  });

  return (
    <Stack spacing={2}>
      <GraduatesToggle<Semester>
        title="Семестр"
        variant={["1", "2"] as const}
        value={watch("semester")}
        onChange={(v) => setValue("semester", v ?? "1")}
      />

      <Controller
        name="studentName"
        control={control}
        render={({ field, fieldState }) => (
          <TextField
            {...field}
            label="ФИО студента"
            size="small"
            error={!!fieldState.error}
            helperText={fieldState.error?.message}
          />
        )}
      />

      <Controller
        name="company"
        control={control}
        render={({ field }) => (
          <TextField {...field} label="Предприятие" size="small" />
        )}
      />

      <LocalizationProvider
        dateAdapter={AdapterDayjs}
        adapterLocale="ru"
        localeText={
          ruRU.components.MuiLocalizationProvider.defaultProps.localeText
        }
      >
        <Controller
          name="startDate"
          control={control}
          render={({ field, fieldState }) => (
            <DatePicker
              label="Дата начала работы"
              views={["year", "month", "day"]}
              value={field.value ? dayjs(field.value) : null}
              onChange={(d) => field.onChange(d?.format("YYYY-MM-DD") ?? "")}
              format="DD.MM.YYYY"
              slotProps={{
                textField: {
                  size: "small",
                  error: !!fieldState.error,
                  helperText: fieldState.error?.message,
                },
              }}
            />
          )}
        />
      </LocalizationProvider>

      <Controller
        name="position"
        control={control}
        render={({ field, fieldState }) => (
          <TextField
            {...field}
            label="Должность"
            size="small"
            error={!!fieldState.error}
            helperText={fieldState.error?.message}
          />
        )}
      />

      <Button variant="contained" color="success" onClick={onDownload}>
        Скачать бланк
      </Button>
    </Stack>
  );
}
