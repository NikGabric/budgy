import router from '@/router';
import { useUserStore } from '@/stores/user';
import {
  BadgeEuro,
  ChartColumnBig,
  LayoutDashboard,
  LogOut,
  Settings,
} from 'lucide-vue-next';
import type { SidebarItem } from './side-nav.types';

const handleLogout = () => {
  const userStore = useUserStore();
  userStore.logout();
  router.push('/login');
};

export const startItems: SidebarItem[] = [
  {
    title: 'Dashboard',
    icon: LayoutDashboard,
    action: () => router.push('/dashboard'),
  },
  {
    title: 'Transactions',
    icon: BadgeEuro,
    action: () => router.push('/transactions'),
  },
  {
    title: 'Budget',
    icon: ChartColumnBig,
    action: () => router.push('/budget'),
  },
];

export const endItems: SidebarItem[] = [
  {
    title: 'Settings',
    icon: Settings,
    action: () => router.push('/settings'),
  },
  {
    title: 'Logout',
    icon: LogOut,
    action: handleLogout,
    style: 'text-error',
  },
];
