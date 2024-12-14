import axios from 'axios';
import { type ApiResponse } from './types';

const API_URL = import.meta.env.VITE_APP_API_URL;

export const anonGet = async (
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axios.get(reqUrl, { params });

  return resp.data;
};

export const anonPost = async (
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axios.post(reqUrl, body);

  return resp.data;
};

export const anonPut = async (
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axios.put(reqUrl, body);

  return resp.data;
};

export const anonDelete = async (
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  const resp = await axios.delete(reqUrl, { params });

  return resp.data;
};
