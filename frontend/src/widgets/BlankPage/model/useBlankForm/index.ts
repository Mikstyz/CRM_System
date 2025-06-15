import { blankSchema, FormValuesBlank } from "@/widgets/BlankPage/model/schema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";

export function useBlankForm() {
  return useForm<FormValuesBlank>({
    resolver: zodResolver(blankSchema),
    defaultValues: {
      semester: "1",
      studentName: "",
      company: "",
      startDate: "",
      position: "",
    },
  });
}
