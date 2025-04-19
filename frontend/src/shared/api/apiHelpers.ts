import { AxiosError, AxiosResponse } from 'axios'
import { BaseApiProps, BaseApiTrunksProps } from '@/shared/types/typeAPI.ts'
import { makeRequest } from './helpers/makeRequest'

/**
 * @template TResponse - Тип данных ожидаемого ответа.
 * @template TRequest - Запрашивайте тип полезной нагрузки.
 * @param props - Запросите конфигурацию.
 * @returns AxiosResponse<TResponse>
 */

/**
 * Выполняет HTTP-запрос с обработчиком отклонения.
 */
export async function executeApiRTK<TResponse = any, TRequest = any>(
  props: BaseApiTrunksProps<TRequest>,
): Promise<AxiosResponse<TResponse>> {
  try {
    return await makeRequest<TResponse, TRequest>(props)
  } catch (error: unknown) {
    const axiosError = error as AxiosError
    return props.rejectWithValue(axiosError?.response?.data || props.errorMessage)
  }
}

/**
 * Выполняет HTTP-запрос и выдает ошибку при сбое.
 */
export async function executeApiRequest<TResponse = any, TRequest = any>(
  props: BaseApiProps<TRequest>,
): Promise<AxiosResponse<TResponse>> {
  try {
    return await makeRequest<TResponse, TRequest>(props)
  } catch (error: unknown) {
    const axiosError = error as AxiosError
    throw new Error(String(axiosError.response?.data || props.errorMessage))
  }
}
