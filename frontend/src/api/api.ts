import type { Transaction } from '@/composables/useApi';
import { anonPost, authGet, authPost } from './helpers';
import type { LoginUserDto } from './types';
import { useUserStore } from '@/stores/user';

export interface LoginUserResponse {
  token: string;
}

export const loginUser = async (
  loginUserDto: LoginUserDto,
): Promise<LoginUserResponse> => {
  const resp = await anonPost<LoginUserResponse>('/user/login', loginUserDto);

  if (!resp.isSuccess) {
    throw new Error(resp.errorMsg);
  }

  return resp.data;
};

export interface TransactionDto {
  description: string;
  amount: number;
  transactionTypeId: number;
  transactionDate: string | Date;
}

export const postTransaction = async (
  transactionDto: TransactionDto,
): Promise<Transaction> => {
  const userStore = useUserStore();
  const d = new Date(transactionDto.transactionDate);
  transactionDto.transactionDate = d.toISOString();

  const resp = await authPost<Transaction>(
    userStore.token ?? '',
    '/transaction',
    transactionDto,
  );

  return resp.data;
};

export const getTransaction = async (id: number): Promise<Transaction> => {
  const userStore = useUserStore();

  const resp = await authGet<Transaction>(
    userStore.token ?? '',
    `/transaction/${id}`,
  );

  return resp.data;
};
