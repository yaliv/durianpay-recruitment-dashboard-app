import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('pages/LoginPage.vue'),
  },
  {
    path: '/',
    name: 'Internal',
    meta: { requiresAuth: true },
    component: () => import('layouts/InternalLayout.vue'),
    children: [
      { path: '', name: 'Home', component: () => import('pages/HomePage.vue') },
      { path: 'payments', name: 'Payments', component: () => import('pages/PaymentsPage.vue') },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    name: 'NotFound',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
