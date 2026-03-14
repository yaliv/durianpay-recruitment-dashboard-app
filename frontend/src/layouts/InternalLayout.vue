<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <q-toolbar-title> Payment Dashboard - {{ $route.name }} </q-toolbar-title>

        <q-btn flat dense round icon="logout" aria-label="Logout" @click="logout">
          <q-tooltip> Logout </q-tooltip>
        </q-btn>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header> Navigation </q-item-label>

        <NavigationItem v-for="link in linksList" :key="link.title" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <q-page padding v-if="errMessage">
        <h3 class="text-negative">{{ errMessage }}</h3>
      </q-page>

      <router-view v-else v-slot="{ Component }">
        <suspense timeout="0">
          <component :is="Component" />

          <template #fallback>
            <q-page padding>
              <h3 class="text-info">Loading content...</h3>
            </q-page>
          </template>
        </suspense>
      </router-view>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref, watch, onErrorCaptured } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { isAxiosError } from 'axios';

import { useAuthStore } from 'stores/auth';
import NavigationItem, { type NavigationItemProps } from 'components/NavigationItem.vue';

const router = useRouter();
const route = useRoute();
const auth = useAuthStore();

const linksList: NavigationItemProps[] = [
  {
    title: 'Home',
    caption: '',
    icon: 'home',
    link: '/',
  },
  {
    title: 'Payments',
    caption: 'List of payments',
    icon: 'payments',
    link: '/payments',
  },
];

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}

async function logout() {
  auth.logout();
  return router.push('/login');
}

// ----- Error handling -----

const errMessage = ref('');

watch(
  () => route.path,
  () => (errMessage.value = ''),
);

onErrorCaptured((err) => {
  if (isAxiosError(err)) {
    errMessage.value = err.response ? 'Server Error' : 'Network Error';
  } else {
    errMessage.value = 'Unexpected Error';
  }

  // Returning false stops the error from propagating further up the chain.
  return false;
});
</script>
