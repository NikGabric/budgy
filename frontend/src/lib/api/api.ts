import axios from 'axios';
import { env } from '$env/dynamic/public';

const API_URL = env.PUBLIC_API_URL + '/api/v1';

export const get = async (path: string) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'GET'
	});
	return res.data;
};

export const post = async (path: string, data: any) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'POST',
    data
	});
	return res.data;
};

export const put = async (path: string) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'PUT'
	});
	return res.data;
};

export const deleteReq = async (path: string) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'DELETE'
	});
	return res.data;
};

export const getAuth = async (token: string, path: string, params?: any)=> {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'GET',
		headers: {
			Authorization: `Bearer ${token}`
		},
		params
	});
	return res.data;
};

export const postAuth = async (token: string, path: string, data: any) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'POST',
		headers: {
			Authorization: `Bearer ${token}`
		},
		data
	});
	return res.data;
};

export const putAuth = async (token:string, path: string) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'PUT',
		headers: {
			Authorization: `Bearer ${token}`
		},
	});
	return res.data;
};

export const deleteAuth = async (token: string, path: string) => {
	const res = await axios.request({
		url: `${API_URL}${path}`,
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${token}`
		},
	});
	return res.data;
};
