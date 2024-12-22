import { authGet } from '@/api/helpers';
import type { PaginationResponse } from '@/api/types';
import { DEFAULT_PAGINATION_LIMIT } from '@/app-constants';
import { useUserStore } from '@/stores/user';
import { storeToRefs } from 'pinia';

export interface Transaction {
  id: number;
  userId: number;
  transactionTypeId: number;
  amount: number;
  description: string;
  transactionDate: string;
  createdAt: string;
  updatedAt: string;
  transactionTypeName: string;
}

export interface GetUserTransactionsParams {
  transactionTypeId?: number;
  fromDate?: string;
  toDate?: string;
}

export interface PaginationParams {
  limit?: number;
  page?: number;
}

export const useGetUserTransactions = async (
  params?: GetUserTransactionsParams & PaginationParams,
): Promise<PaginationResponse<Transaction[]>> => {
  const userStore = useUserStore();
  const { token } = storeToRefs(userStore);

  const limit = params?.limit ?? DEFAULT_PAGINATION_LIMIT;

  const resp = await authGet<PaginationResponse<Transaction[]>>(
    token.value ?? '',
    '/user/transactions',
    {
      ...params,
      limit,
    },
  );

  if (!resp.isSuccess) {
    throw new Error(resp.errorMsg);
  }

  return resp.data;
};

export interface TransactionType {
  id: number;
  userId: number;
  name: string;
  description: string;
  createdAt: string;
  updatedAt: string;
}

export const useGetUserTransactionTypes = async (): Promise<
  TransactionType[]
> => {
  const userStore = useUserStore();
  const { token } = storeToRefs(userStore);

  const resp = await authGet<TransactionType[]>(
    token.value ?? '',
    '/user/transaction-types',
  );

  if (!resp.isSuccess) {
    throw new Error(resp.errorMsg);
  }

  return resp.data;
};
