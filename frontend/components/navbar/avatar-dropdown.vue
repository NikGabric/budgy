<script setup lang="ts">
import { LogOut, Settings, User } from "lucide-vue-next";
</script>

<template>
  <div class="dropdown dropdown-end">
    <button
      class="btn btn-circle avatar placeholder"
      tabindex="0"
      role="button"
    >
      <div class="bg-neutral text-neutral-content w-12 rounded-full">
        <span>SY</span>
      </div>
    </button>
    <ul
      tabindex="0"
      class="dropdown-content menu bg-base-300 z-[1] w-56 p-2 shadow rounded-xl gap-1"
    >
      <li>
        <button class="rounded-lg"><User />Profile</button>
      </li>
      <li>
        <button class="rounded-lg"><Settings /> Settings</button>
      </li>
      <div class="divider my-0" />
      <li>
        <button class="rounded-lg text-error"><LogOut /> Logout</button>
      </li>
    </ul>
  </div>
</template>
