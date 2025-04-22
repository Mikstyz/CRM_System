import { Runtime } from "@wails/runtime";

// универсальный вызов с перехватом Code / Error
export async function call<T>(fn: string, ...args: unknown[]): Promise<T> {
    try {
        // window.backend.{Package}.{Method} генерирует Wails
        // например "backend.Api.Inf_AllGroup"
        // NB: Runtime.Call тоже можно, но window.backend короче
        //  @ts-ignore  – его нет в typings
        const res = await window.backend.Api[fn](...args);
        if (res.Code !== 0) throw new Error(res.Error || "Unknown backend error");
        return res as T;
    } catch (err) {
        console.error(`Wails call ${fn} failed:`, err);
        throw err;
    }
}
