import { createRouter, createWebHistory } from 'vue-router';
import DashboardView from '../views/dashboard/DashboardView.vue';
import { useUserStore } from '@/stores/user';
import LoginView from '@/views/auth/LoginView.vue';
import SettingsView from '@/views/settings/SettingsView.vue';
import BudgetView from '@/views/budget/BudgetView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/login', name: 'login', component: LoginView },
    { path: '/', redirect: 'dashboard' },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
      meta: {
        auth: true,
      },
    },
    {
      path: '/budget',
      name: 'budget',
      component: BudgetView,
      meta: {
        auth: true,
      },
    },
    {
      path: '/settings',
      name: 'settings',
      component: SettingsView,
      meta: {
        auth: true,
      },
    },
  ],
});

router.beforeEach((to) => {
  const { isAuthenticated } = useUserStore();

  if (!isAuthenticated && to.meta.auth) {
    return { name: 'login' };
  }

  return true;
});

export default router;
