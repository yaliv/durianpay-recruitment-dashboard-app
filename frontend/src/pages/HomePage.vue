<template>
  <div>
    <h2>Welcome!</h2>
    <button @click="logout">Logout</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

async function logout() {
  localStorage.removeItem('auth');
  return router.push('/login');
}
</script>
