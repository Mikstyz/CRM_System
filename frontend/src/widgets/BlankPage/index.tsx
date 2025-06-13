import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { useEffect } from "react";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { Group } from "@/entities/group/types";

import { ListStudents } from "src/widgets/BlankPage/ui/ListStudents";
import { TitleWidgets } from "@/widgets/BlankPage/ui/ListStudents/ui/TitleWidgets";
import { blankSchema, FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { FormCreatBlank } from "@/widgets/BlankPage/ui/FormCreatBlank";
import { studentsData } from "@/features/ListGroup/const";

export function BlankPage({ group }: { group: Group }) {
  const {
    selectStudent,
    // studentsData,
    semester,
  } = useAppSelector(selectBlank);

  const {
    register,
    handleSubmit,
    setValue,
    setError,
    getValues,
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
    if (!selectStudent) return;

    setValue("studentName", selectStudent?.fullName);
    if (selectStudent?.startDateWork) {
      setValue("startDate", selectStudent?.startDateWork);
    }
    setValue("semester", semester);
    setValue("company", selectStudent?.company ?? "");
    setValue("position", selectStudent?.position ?? "");
  }, [setValue, semester]);

  return (
    <section>
      <TitleWidgets group={group} />
      <div className="flex row">
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
