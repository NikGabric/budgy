export type ApiResponse = {
  data: Record<string, unknown>;
  isSuccess: boolean;
  errorMsg?: string;
};
