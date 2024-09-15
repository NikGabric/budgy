import { Gauge, HandCoins, Wallet } from "lucide-vue-next";

export interface NavigationItem {
  label: string;
  to: string;
  icon?: Component;
  styles?: string;
}

export const navigationItems: NavigationItem[] = [
  {
    label: "Dashboard",
    to: "/dashboard",
    icon: Gauge,
  },
  {
    label: "Transactions",
    to: "/",
    icon: HandCoins,
  },
  {
    label: "Wallet",
    to: "/",
    icon: Wallet,
  },
];
