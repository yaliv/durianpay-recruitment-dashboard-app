<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn flat dense round icon="menu" aria-label="Menu" @click="toggleLeftDrawer" />

        <q-toolbar-title> Payment Dashboard </q-toolbar-title>

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
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from 'stores/auth';
import NavigationItem, { type NavigationItemProps } from 'components/NavigationItem.vue';

const router = useRouter();
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
</script>
