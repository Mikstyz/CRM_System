import { Stack, Grid } from "@mui/material";
import { FormProvider } from "react-hook-form";
import { TitleWidgets } from "./ui/ListStudents/ui/TitleWidgets";
import { FormCreatBlank } from "./ui/FormCreatBlank";
import { ListStudents } from "./ui/ListStudents";
import { useBlankForm } from "@/widgets/BlankPage/model/useBlankForm";
import styles from "./index.module.css";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { useEffect } from "react";
import { getAllStudentGroupThunks } from "@/entities/blank/store";

export function BlankPage() {
  const dispatch = useAppDispatch();
  const methods = useBlankForm();
  const { group } = useAppSelector(selectBlank);
  if (!group) return <p>Группа не добавлена</p>;

  useEffect(() => {
    dispatch(getAllStudentGroupThunks(group.id));
  }, []);

  return (
    <Stack spacing={3} className={styles.container}>
      <TitleWidgets group={group} />

      <Grid container spacing={2} columns={5}>
        <Grid size={3}>
          <FormProvider {...methods}>
            <FormCreatBlank />
          </FormProvider>
        </Grid>

        <Grid size={2}>
          <FormProvider {...methods}>
            <ListStudents />
          </FormProvider>
        </Grid>
      </Grid>
    </Stack>
  );
}
