import axios from 'axios';
import { type ApiResponse } from './types';

const API_URL = import.meta.env.VITE_APP_API_URL;

export const anonGet = async (
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  return axios.get(reqUrl, { params });
};

export const anonPost = async (
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  return axios.post(reqUrl, body);
};

export const anonPut = async (
  url: string,
  body?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  return axios.put(reqUrl, body);
};

export const anonDelete = async (
  url: string,
  params?: Record<string, string | number>,
): Promise<ApiResponse> => {
  const reqUrl = `${API_URL}${url}`;

  return axios.delete(reqUrl, { params });
};
