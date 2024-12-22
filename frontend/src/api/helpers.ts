import { type ApiResponse } from './types';
import { apiClient } from './api-client';

export const anonGet = async <T = unknown>(
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const resp = await apiClient.get(url, { params });
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
  const resp = await apiClient.post(url, body);
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
  const resp = await apiClient.put(url, body);
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
  const resp = await apiClient.delete(url, { params });
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
  params?: Record<string, string | number | undefined>,
): Promise<ApiResponse<T>> => {
  if (!token) {
    throw new Error('Token missing from request!');
  }

  const resp = await apiClient.get(url, {
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
