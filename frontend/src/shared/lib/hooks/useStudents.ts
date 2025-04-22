// /src/shared/lib/hooks/useStudents.ts
import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "@/shared/lib/hooks/redux";
import {fetchStudents} from "@/entities/student/studentSlice.ts";

export function useStudents() {
    const dispatch = useAppDispatch();
    const { list, loading } = useAppSelector((s) => s.students);

    useEffect(() => { dispatch(fetchStudents()); }, [dispatch]);

    return { students: list, loading };
}
