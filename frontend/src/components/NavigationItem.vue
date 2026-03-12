<template>
  <q-item clickable v-ripple :to="link" :active="active" active-class="bg-accent text-white">
    <q-item-section v-if="icon" avatar>
      <q-icon :name="icon" />
    </q-item-section>

    <q-item-section>
      <q-item-label>{{ title }}</q-item-label>
      <q-item-label caption :class="active ? 'text-amber-2' : ''">{{ caption }}</q-item-label>
    </q-item-section>
  </q-item>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import { ref, watchEffect } from 'vue';

const route = useRoute();
const active = ref(false);

export interface NavigationItemProps {
  title: string;
  caption?: string;
  link?: string;
  icon?: string;
}

const props = withDefaults(defineProps<NavigationItemProps>(), {
  caption: '',
  link: '#',
  icon: '',
});

watchEffect(() => {
  active.value = route.path == props.link;
});
</script>
