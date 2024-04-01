import { browser } from "$app/environment";
import { writable } from "svelte/store";

export const authToken = writable(browser ? window.localStorage.getItem("auth-token") : "");

authToken.subscribe((val: string | null) => {
  if (browser && val) window.localStorage.setItem("auth-token", val);
  else if(browser) window.localStorage.removeItem("auth-token");
})
