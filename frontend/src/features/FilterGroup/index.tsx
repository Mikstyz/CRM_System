import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { InputFilterGroup } from "./ui/InputFilterGroup";
import { setFilters } from "@/features/FilterGroup/store/groupFiltersSlice.ts";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import { useDeferredValue, useEffect } from "react";
import { useDebouncedCallback } from "use-debounce";
import { shallowEqual } from "react-redux";
import { Filters, filterSchema, FiltersRaw } from "./schema";
import { ButtonPush } from "@/shared/ui/ButtonPush";
import { GraduatesToggle } from "@/features/FilterGroup/ui/GraduatesToggle";
import { headerCreateGroup } from "@/features/FilterGroup/lib/headers/headerCreateGroup.ts";
import { headerCleansingForm } from "@/features/FilterGroup/lib/headers/headerCleansingForm.ts";
import { RootState } from "@/app/store";
import { ErrorMassage } from "@/shared/ui/ErrorMassage";
import { clearErrors } from "@/entities/group/store";
import { Course, Graduates } from "@/entities/group/types";

// TODO:BAG Не работает фильтрация по Курс GraduatesToggle
export function FilterGroup({ groupsLength }: { groupsLength: number }) {
  const dispatch = useAppDispatch();
  const lastSent = useAppSelector((s: RootState) => s.groupFilters);
  const { error } = useAppSelector((s: RootState) => s.groups);

  const {
    register,
    watch,
    setError,
    setValue,
    getValues,
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
      const cleaned = Object.fromEntries(
        Object.entries(parsed.data).map(([k, v]) =>
          v === "" || v === undefined ? [k, null] : [k, v],
        ),
      ) as Filters;
      pushFilters(cleaned);
    }
  }, [parsed, pushFilters]);

  useEffect(() => {
    if (error) {
      dispatch(clearErrors());
    }
  }, [dispatch, groupsLength]);

  return (
    <>
      <aside className="border p-4 rounded-lg mb-4 w-full max-w-xl">
        <div className="flex justify-between mb-1">
          <h2 className="font-semibold mb-2">Фильтрация</h2>
          {(getValues("course") ||
            getValues("graduates") ||
            getValues("groupNumber") ||
            getValues("specialty")) && (
            <button
              type="button"
              className="ml-auto text-sm text-gray-500 hover:text-gray-700 bg-gray-100 hover:bg-gray-200 px-2 py-1 rounded-md"
              onClick={() => headerCleansingForm({ setError, setValue })}
            >
              Очистить форму
            </button>
          )}
        </div>

        <form className="flex flex-wrap gap-2">
          {/* Курс */}
          <GraduatesToggle
            title="Курс"
            variant={["1", "2", "3", "4"]}
            onChange={(num) => {
              setValue("course", num);
              setError("course", {});
            }}
            value={getValues("course")}
            error={errors.course?.message}
          />

          {/* Специальность */}
          <InputFilterGroup
            title="Специальность"
            inputProps={{
              type: "text",
              placeholder: "ИСП",
              ...register("specialty"),
            }}
            error={errors.specialty?.message}
          />

          {/* Выпускники */}
          <GraduatesToggle
            title="Выпускники (9/11)"
            variant={["9", "11"]}
            onChange={(num) => {
              setValue("graduates", num);
              setError("graduates", {});
            }}
            value={getValues("graduates")}
            error={errors.graduates?.message}
          />

          {/* Номер группы */}
          <InputFilterGroup
            title="Номер группы"
            inputProps={{
              type: "number",
              placeholder: "45",
              ...register("groupNumber"),
            }}
            error={errors.groupNumber?.message}
          />
        </form>
      </aside>
      <ErrorMassage error={error} className="items-start" />
      {groupsLength <= 0 && (
        <ButtonPush
          onClick={() =>
            headerCreateGroup({
              course: getValues("course") as Course,
              specialty: getValues("specialty"),
              graduates: getValues("graduates") as Graduates,
              groupNumber: getValues("groupNumber") as number,
              dispatch,
              setError,
            })
          }
        >
          Добавить группу
        </ButtonPush>
      )}
    </>
  );
}
