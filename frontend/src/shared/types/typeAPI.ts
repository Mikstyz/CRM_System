import { AxiosRequestConfig } from 'axios'

export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE'

export interface BaseApiProps<TRequest = any> {
  method: HttpMethod
  url: string
  body?: TRequest | object
  config?: AxiosRequestConfig
  errorMessage?: string
  setErrorServer?: () => void
}

type RejectWithValueFn = any

export interface BaseApiTrunksProps<TRequest = any> extends BaseApiProps<TRequest> {
  rejectWithValue: RejectWithValueFn
}
