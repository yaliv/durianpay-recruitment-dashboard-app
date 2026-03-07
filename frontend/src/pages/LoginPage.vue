<template>
  <div>
    <h2>Login</h2>
    <button @click="login">Login</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const router = useRouter();

async function login() {
  localStorage.setItem('auth', 'true');
  return router.push('/');
}
</script>
