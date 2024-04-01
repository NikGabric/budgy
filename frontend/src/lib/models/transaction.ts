export type Transaction = {
  transactionId: number,
  title: string,
  description: string,
  amount: number,
  deleted: boolean,
  insertedAt: string,
  updatedAt: string,
  budgyUserId: number,
  transactionTypeId: number,
}