/**
 * Этот файл содержит конфигурацию Axios для взаимодействия с API.
 */
import axios from 'axios'
import { API_ENDPOINTS } from './constEndpoints'

export const axiosInstance = axios.create({
  baseURL: API_ENDPOINTS.common.baseUrl, // Базовый URL для всех запросов
  timeout: 10000, // Тайм-аут запросов
  headers: {
    'Content-Type': 'application/json',
  },
})

/**
 * Перехватчик запросов
 *
 * @param {object} config - Конфигурация текущего запроса.
 * @returns {object} Модифицированная конфигурация запроса.
 */
axiosInstance.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

/**
 * Перехватчик ответов
 *
 * Обрабатывает успешные ответы и ошибки:
 * - Возвращает только данные из успешных ответов.
 *
 * @param {object} response - Ответ сервера.
 * @returns {object} Данные ответа.
 * @throws {string|object} Сообщение об ошибке или данные ошибки.
 */
axiosInstance.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    console.log('Error server:', error)
    if (error.response) {
      console.error('Ошибка сервера: ', error.response)
    } else if (error.request) {
      console.error('Нет ответа от сервера: ', error.request)
    } else {
      console.error('Ошибка конфигурации запроса: ', error.message)
    }
    return error.response
  },
)
