import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import clsx from "classnames";
import { InputFilterGroup } from "./ui/InputFilterGroup";
import { setFilters } from "@/entities/group/store/groupFiltersSlice";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { useDeferredValue, useEffect } from "react";
import { useDebouncedCallback } from "use-debounce";
import { shallowEqual } from "react-redux";
import { Filters, filterSchema, FiltersRaw } from "./schema";

export function FilterGroup() {
  const dispatch = useAppDispatch();
  const lastSent = useAppSelector((s) => s.groupFilters);

  const {
    register,
    watch,
    formState: { errors },
  } = useForm<FiltersRaw>({
    resolver: zodResolver(filterSchema),
    defaultValues: {
      course: undefined,
      specialty: undefined,
      graduates: undefined,
      groupNumber: undefined,
    },
    mode: "onChange",
  });

  const rawValues = watch();
  const deferredValues = useDeferredValue(rawValues);
  const parsed = z.object(filterSchema.shape).safeParse(deferredValues);

  const pushFilters = useDebouncedCallback((data: Filters) => {
    if (!shallowEqual(data, lastSent)) dispatch(setFilters(data));
  }, 150);
  useEffect(() => {
    if (parsed.success) {
      // пустые строки → null   (упрощаем логику selector’а)
      const cleaned = Object.fromEntries(
        Object.entries(parsed.data).map(([k, v]) =>
          v === "" || v === undefined ? [k, null] : [k, v],
        ),
      ) as Filters;
      pushFilters(cleaned);
    }
  }, [parsed, pushFilters]);

  return (
    <aside className="border p-4 rounded-lg mb-4 w-full max-w-xl">
      <h2 className="font-semibold mb-2">Фильтрация</h2>

      <form className="flex flex-wrap gap-2">
        {/* Курс */}
        <InputFilterGroup
          title="Курс"
          inputProps={{
            type: "number",
            placeholder: "1",
            ...register("course"),
            className: clsx(
              "border rounded px-3 py-2 w-full focus:outline-none focus:ring-1",
              errors.course ? "border-red-500" : "border-gray-300",
              !errors.course && watch("course") ? "focus:border-green-500" : "",
            ),
          }}
          error={errors.course?.message}
        />

        {/* Специальность */}
        <InputFilterGroup
          title="Специальность"
          inputProps={{
            type: "text",
            placeholder: "ИСП",
            ...register("specialty"),
            className: clsx(
              "border rounded px-3 py-2 w-full focus:outline-none focus:ring-1",
              errors.specialty ? "border-red-500" : "border-gray-300",
              !errors.specialty && watch("specialty")
                ? "focus:border-green-500"
                : "",
            ),
          }}
          error={errors.specialty?.message}
        />

        {/* Выпускники */}
        <InputFilterGroup
          title="Выпускники (9/11)"
          inputProps={{
            type: "text",
            placeholder: "9",
            ...register("graduates"),
            className: clsx(
              "border rounded px-3 py-2 w-full focus:outline-none focus:ring-1",
              errors.graduates ? "border-red-500" : "border-gray-300",
              !errors.graduates && watch("graduates")
                ? "focus:border-green-500"
                : "",
            ),
          }}
          error={errors.graduates?.message}
        />

        {/* Номер группы */}
        <InputFilterGroup
          title="Номер группы"
          inputProps={{
            type: "number",
            placeholder: "45",
            ...register("groupNumber"),
            className: clsx(
              "border rounded px-3 py-2 w-full focus:outline-none focus:ring-1",
              errors.groupNumber ? "border-red-500" : "border-gray-300",
              !errors.groupNumber && watch("groupNumber")
                ? "focus:border-green-500"
                : "",
            ),
          }}
          error={errors.groupNumber?.message}
        />
      </form>
    </aside>
  );
}
