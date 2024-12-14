import { anonPost } from './helpers';
import type { LoginUserDto } from './types';

export interface LoginUserResponse {
  token: string;
}

export const loginUser = async (
  loginUserDto: LoginUserDto,
): Promise<LoginUserResponse> => {
  const resp = await anonPost<LoginUserResponse>(
    '/user/login',
    loginUserDto,
  );

  if (!resp.isSuccess) {
    throw new Error(resp.errorMsg);
  }

  return resp.data;
};
