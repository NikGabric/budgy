import axios from 'axios';
import { type ApiResponse } from './types';

const API_URL = import.meta.env.VITE_APP_API_URL;

export const anonGet = async <T = unknown>(
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse<T>> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axios.get(reqUrl, { params });
  const isSuccess =
    resp.status >= 200 && resp.status <= 299;

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

  const resp = await axios.post(reqUrl, body);
  const isSuccess =
    resp.status >= 200 && resp.status <= 299;

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

  const resp = await axios.put(reqUrl, body);
  const isSuccess =
    resp.status >= 200 && resp.status <= 299;

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

  const resp = await axios.delete(reqUrl, { params });
  const isSuccess =
    resp.status >= 200 && resp.status <= 299;

  return {
    data: resp.data,
    isSuccess,
    errorMsg: isSuccess ? undefined : resp.statusText,
  };
};
