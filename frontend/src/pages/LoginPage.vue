<template>
  <div class="row">
    <div class="col-8 window-height hero" :style="{ backgroundImage: `url(${bg})` }" />
    <div class="col window-height q-pa-lg">
      <h4 class="text-weight-medium text-primary">Login to<br />Payment Dashboard</h4>
      <q-btn color="primary" label="Login" @click="login" />
    </div>
  </div>
</template>

<style lang="sass" scoped>
.hero
  background-repeat: no-repeat
  background-size: cover
  background-position: center
</style>

<script setup lang="ts">
import { useRouter } from 'vue-router';

import bg from 'assets/Minimalist-Background-10.jpg';
import { useAuthStore } from 'stores/auth';

const router = useRouter();
const auth = useAuthStore();

async function login() {
  auth.login({ email: 'dummyemail', role: 'dummyrole' });
  return router.push('/');
}
</script>
