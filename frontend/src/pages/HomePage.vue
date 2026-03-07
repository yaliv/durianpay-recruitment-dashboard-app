<template>
  <q-page padding>
    <h3>Welcome!</h3>
    <table class="text-subtitle1">
      <tbody>
        <tr>
          <td class="text-bold q-pr-sm">Email:</td>
          <td>{{ auth.user?.email }}</td>
        </tr>
        <tr>
          <td class="text-bold q-pr-sm">Role:</td>
          <td>{{ auth.user?.role }}</td>
        </tr>
      </tbody>
    </table>
  </q-page>
</template>

<script setup lang="ts">
import { useAuthStore } from 'stores/auth';

const auth = useAuthStore();
</script>
