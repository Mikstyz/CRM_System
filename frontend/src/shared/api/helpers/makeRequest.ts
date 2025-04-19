import { BaseApiProps } from '@/shared/types/typeAPI.ts'
import { AxiosResponse } from 'axios'
import { axiosInstance } from '@/shared/api/axiosInstance.ts'

/**
 * Выполняет HTTP-запрос с использованием экземпляра axios.
 * *
 * @template TResponse - Тип данных ожидаемого ответа.
 * @template TRequest - Запрашивайте тип полезной нагрузки.
 * @param {BaseApiProps<TRequest>} props - Свойства запроса.
 * @returns {Promise<AxiosResponse<TResponse>>} - Обещание, разрешающееся с помощью ответа Axis.
 * @throws {Error} - Если указан неподдерживаемый HTTP-метод.
 */
export async function makeRequest<TResponse = unknown, TRequest = any>({
  method,
  url,
  body,
  config,
}: BaseApiProps<TRequest>): Promise<AxiosResponse<TResponse>> {
  switch (method) {
    case 'GET':
      return await axiosInstance.get<TResponse>(url, {
        ...config,
        params: body,
      })
    case 'DELETE':
      return await axiosInstance.delete<TResponse>(url, {
        ...config,
        data: body,
      })
    case 'POST':
      return await axiosInstance.post<TResponse>(url, body, config)
    case 'PUT':
      return await axiosInstance.put<TResponse>(url, body, config)
    default:
      throw new Error(`Неподдерживаемый метод: ${method}`)
  }
}
