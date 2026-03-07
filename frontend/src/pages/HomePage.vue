<template>
  <div>
    <h2>Welcome!</h2>
    <button @click="logout">Logout</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useAuthStore } from 'stores/auth';

const router = useRouter();
const auth = useAuthStore();

async function logout() {
  auth.logout();
  return router.push('/login');
}
</script>
