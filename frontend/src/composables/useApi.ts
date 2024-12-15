import { authGet } from '@/api/helpers';
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

export const useGetTransactions = async (
  params?: Record<string, string | number>,
): Promise<Transaction[]> => {
  const userStore = useUserStore();
  const { token } = storeToRefs(userStore);

  const resp = await authGet<Transaction[]>(
    token.value ?? '',
    '/user/transactions',
    params ?? { limit: 5 },
  );

  if (!resp.isSuccess) {
    throw new Error(resp.errorMsg);
  }

  return resp.data;
};
