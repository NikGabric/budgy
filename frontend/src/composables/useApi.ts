import { authGet } from '@/api/helpers';
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
  transaction_type_id?: number;
  from_date?: string;
  to_date?: string;
}

export interface PaginationParams {
  limit?: number;
}

export const useGetUserTransactions = async (
  params?: GetUserTransactionsParams & PaginationParams,
): Promise<Transaction[]> => {
  const userStore = useUserStore();
  const { token } = storeToRefs(userStore);

  const resp = await authGet<Transaction[]>(
    token.value ?? '',
    '/user/transactions',
    { ...params, limit: params?.limit ?? DEFAULT_PAGINATION_LIMIT },
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
