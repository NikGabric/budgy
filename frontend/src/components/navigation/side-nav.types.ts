import type { LucideIcon } from 'lucide-vue-next';

export interface SidebarItem {
  title: string;
  icon: LucideIcon;
  action: () => void;
  style?: string;
}
