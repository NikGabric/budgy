import type { RefSymbol } from "@vue/reactivity";
import { defineStore } from "pinia";

export interface User {
  username: string;
  email: string;
}

const defaultUser: User = {
  username: "Test User",
  email: "test@email.si",
};

export const useUserStore = defineStore("user", () => {
  const user: Ref<User | null | undefined> = ref(useCookie("user"));
  const navigationItems: Ref<NavigationItem[]> = ref(defaultNavigationItems);

  const getUser: ComputedRef<User | null | undefined> = computed(
    () => user.value
  );
  const getNavigationItems: ComputedRef<NavigationItem[]> = computed(
    () => navigationItems.value
  );
  const isLoggedIn: ComputedRef<boolean> = computed(() => !!user.value);

  const login = () => {
    user.value = defaultUser;
    useCookie("user", {
      maxAge: 60 * 24 * 28,
      sameSite: true,
      secure: true,
    }).value = JSON.stringify(user.value);
  };
  const logout = () => {
    user.value = undefined;
    useCookie("user").value = null;
  };

  return {
    user,
    getUser,
    navigationItems,
    getNavigationItems,
    isLoggedIn,
    login,
    logout,
  };
});
