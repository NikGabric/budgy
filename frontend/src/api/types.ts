export interface ApiResponse<T = unknown> {
  data: T;
  isSuccess: boolean;
  errorMsg?: string;
}

export type LoginUserDto = {
  username: string;
  password: string;
};