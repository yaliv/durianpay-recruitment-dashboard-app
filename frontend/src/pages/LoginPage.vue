<template>
  <div style="padding-left: 4em">
    <h3>Login</h3>
    <q-btn color="primary" label="Login" @click="login" />
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
