/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from 'axios';

const API_URL = import.meta.env.VITE_APP_API_URL;

function toCamelCase<T>(obj: Record<string, any>): T {
  if (obj === null || typeof obj !== 'object') {
    return obj as T;
  }

  if (Array.isArray(obj)) {
    return obj.map((item) => toCamelCase(item)) as unknown as T;
  }

  const camelCaseObj: Record<string, any> = {};
  for (const key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      const camelCaseKey = key.replace(/_([a-z])/g, (_, letter) =>
        letter.toUpperCase(),
      );
      camelCaseObj[camelCaseKey] = toCamelCase(obj[key]);
    }
  }
  return camelCaseObj as T;
}

function toSnakeCase<T>(obj: Record<string, any>): T {
  if (obj === null || typeof obj !== 'object') {
    return obj as T;
  }

  if (Array.isArray(obj)) {
    return obj.map((item) => toSnakeCase(item)) as unknown as T;
  }

  const snakeCaseObj: Record<string, any> = {};
  for (const key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      const snakeCaseKey = key.replace(
        /([A-Z])/g,
        (_, letter) => `_${letter.toLowerCase()}`,
      );
      snakeCaseObj[snakeCaseKey] = toSnakeCase(obj[key]);
    }
  }
  return snakeCaseObj as T;
}

const apiClient = axios.create({
  baseURL: API_URL,
});

apiClient.interceptors.request.use(
  (config) => {
    if (config.data && typeof config.data === 'object') {
      config.data = toSnakeCase(config.data);
    }
    if (config.params && typeof config.params === 'object') {
      config.params = toSnakeCase(config.params);
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

apiClient.interceptors.response.use(
  (response) => {
    if (response.data && typeof response.data === 'object') {
      response.data = toCamelCase(response.data);
    }
    return response;
  },
  (error) => {
    return Promise.reject(error);
  },
);

export { apiClient };
