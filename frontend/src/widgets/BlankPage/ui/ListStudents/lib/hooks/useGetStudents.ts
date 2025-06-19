import { useAppSelector } from "@/shared/lib/hooks/redux.ts";
import { selectBlank } from "@/entities/blank/store/selectors.ts";
import { useEffect, useState } from "react";
import { useFormContext } from "react-hook-form";
import { FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { Student } from "@/entities/student/types";

export const useGetStudents = () => {
  const { watch } = useFormContext<FormValuesBlank>();
  const [query, setQuery] = useState("");
  const [filteredData, setFilteredData] = useState<Student[]>([]);
  const { studentsData } = useAppSelector(selectBlank);

  useEffect(() => {
    const value = watch("studentName");
    if (value !== "" && typeof value === "string") {
      setQuery(value);
    } else {
      setFilteredData(studentsData);
    }
  }, [watch("studentName")]);

  useEffect(() => {
    if (!query) {
      setFilteredData(studentsData);
    } else {
      const lowerQuery = query.toLowerCase();
      setFilteredData(
        studentsData.filter((item) =>
          item.fullName.toLowerCase().includes(lowerQuery),
        ),
      );
    }
  }, [query, studentsData]);

  return [filteredData];
};
