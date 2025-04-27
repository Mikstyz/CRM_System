import { FiltersRaw } from "@/features/FilterGroup/schema";
import { UseFormSetError, UseFormSetValue } from "react-hook-form";

interface Arguments {
  setError: UseFormSetError<FiltersRaw>;
  setValue: UseFormSetValue<FiltersRaw>;
}
export const headerCleansingForm = ({ setError, setValue }: Arguments) => {
  setError("root", {});

  setError("course", {});
  setError("graduates", {});
  setError("groupNumber", {});
  setError("specialty", {});
  setValue("course", undefined);
  setValue("graduates", undefined);
  setValue("groupNumber", undefined);
  setValue("specialty", undefined);
};
