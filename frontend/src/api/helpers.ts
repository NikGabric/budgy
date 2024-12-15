/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from 'axios';
import { type ApiResponse } from './types';

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

const axiosClient = axios.create({
  baseURL: API_URL,
});

axiosClient.interceptors.response.use(
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

export const anonGet = async <T = unknown>(
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axiosClient.get(reqUrl, { params });
  const isSuccess = resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};

export const anonPost = async <T = unknown>(
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axiosClient.post(reqUrl, body);
  const isSuccess = resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};

export const anonPut = async <T = unknown>(
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axiosClient.put(reqUrl, body);
  const isSuccess = resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};

export const anonDelete = async <T = unknown>(
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axiosClient.delete(reqUrl, { params });
  const isSuccess = resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};

export const authGet = async <T = unknown>(
  token: string,
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  if (!token) {
    throw new Error('Token missing from request!');
  }

  const reqUrl = `${API_URL}${url}`;

  const resp = await axiosClient.get(reqUrl, {
    params,
    headers: { Authorization: `Bearer ${token}` },
  });
  const isSuccess = resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};
