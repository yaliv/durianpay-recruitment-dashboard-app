<template>
  <div>
    <h2>Login</h2>
    <button @click="login">Login</button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useAuthStore } from 'stores/auth';

const router = useRouter();
const auth = useAuthStore();

async function login() {
  auth.login({ email: 'dummyemail', role: 'dummyrole' });
  return router.push('/');
}
</script>
