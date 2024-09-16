import { Gauge, HandCoins, LogOut, Settings, Wallet } from "lucide-vue-next";

export interface NavigationItem {
  label: string;
  to: string;
  iconName: string;
  styles?: string;
}

export const defaultNavigationItems: NavigationItem[] = [
  {
    label: "Dashboard",
    to: "/dashboard",
    iconName: "Gauge",
  },
  {
    label: "Transactions",
    to: "/",
    iconName: "HandCoins",
  },
  {
    label: "Wallet",
    to: "/",
    iconName: "Wallet",
  },
];

export const iconMap: Record<string, Component> = {
  Gauge,
  HandCoins,
  Wallet,
  Settings,
  LogOut,
};
