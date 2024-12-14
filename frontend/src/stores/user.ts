import { loginUser } from '@/api/api';
import type { LoginUserDto } from '@/api/types';
import { defineStore } from 'pinia';
import { ref, type Ref } from 'vue';
import { parseJwt } from '../utils/jwt';

type User = {
  username: string;
  email: string;
};

export const useUserStore = defineStore('user', () => {
  const localToken = ref(localStorage.getItem('budgy_token') ?? null);
  const parsedToken = localToken.value ? parseJwt(localToken.value) : null;

  const user: Ref<User | null> = ref(localToken.value ? parsedToken : null);
  const token: Ref<string | null> = ref(
    localToken.value ? localToken.value : null,
  );

  const login = async (loginUserDto: LoginUserDto) => {
    try {
      const resp = await loginUser(loginUserDto);
      token.value = resp.token;
      user.value = parseJwt(resp.token);
      localStorage.setItem('budgy_token', resp.token);
    } catch (err) {
      alert(err);
    }
  };

  const logout = () => {
    user.value = null;
    localToken.value = null;
    localStorage.removeItem('budgy_token');
  };

  return { user, token, login, logout };
});
