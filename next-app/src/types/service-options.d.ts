import { AxiosError } from "axios";

export interface IServiceOptions {
  accessToken?: string;
}

export type FetchOptions<T> = {
  isLoading: boolean;
  error: AxiosError | null;
  data: T;
};

export type Result<T, E> = [T | null, E | null];
export type ResultWithoutResponse<E> = [E | null];

export type APIError = AxiosError<{ message: string; title: string }>;
