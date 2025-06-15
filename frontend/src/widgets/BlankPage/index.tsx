import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { useEffect } from "react";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { Group } from "@/entities/group/types";

import { ListStudents } from "src/widgets/BlankPage/ui/ListStudents";
import { TitleWidgets } from "@/widgets/BlankPage/ui/ListStudents/ui/TitleWidgets";
import { blankSchema, FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { FormCreatBlank } from "@/widgets/BlankPage/ui/FormCreatBlank";
import { getAllStudentGroupThunks } from "@/entities/blank/store/thunks.ts";

export function BlankPage({ group }: { group: Group }) {
  const dispatch = useAppDispatch();
  const { selectStudent, studentsData, semester, groupId } =
    useAppSelector(selectBlank); /*
 * {
    id: 18,
    fullName: "Иватов Иван",
    company: "ООО «Рога и копыта»",
    startDateWork: "2025-05-13",
    position: "Стажер",
  },
 * */
  console.log(studentsData);

  const {
    register,
    handleSubmit,
    setValue,
    setError,
    getValues,
    control,
    formState: { errors },
  } = useForm<FormValuesBlank>({
    resolver: zodResolver(blankSchema), // generic выводится автоматически
    defaultValues: {
      semester: "1",
      studentName: "",
      company: "",
      startDate: "",
      position: "",
    },
  });

  useEffect(() => {
    if (groupId) {
      dispatch(getAllStudentGroupThunks(groupId));
    }
  }, [dispatch, groupId]);

  useEffect(() => {
    if (!selectStudent) return;

    setValue("studentName", selectStudent?.fullName);
    if (selectStudent?.startDateWork) {
      setValue("startDate", selectStudent?.startDateWork);
    }
    setValue("semester", semester);
    setValue("company", selectStudent?.company ?? "");
    setValue("position", selectStudent?.position ?? "");
  }, [setValue, semester, selectStudent]);

  return (
    <section>
      <TitleWidgets group={group} />
      <div className="grid grid-cols-2 gap-3 divide-x-2">
        <div className="col">
          <FormCreatBlank
            setValue={setValue}
            register={register}
            selectStudent={selectStudent}
            getValues={getValues}
            setError={setError}
            errors={errors}
            handleSubmit={handleSubmit}
            group={group}
            control={control}
          />
        </div>
        <article className="col">
          <ListStudents
            setValue={setValue}
            studentsData={studentsData}
            errors={errors}
          />
        </article>
      </div>
    </section>
  );
}
