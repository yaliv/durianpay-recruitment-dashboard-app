<template>
  <q-page padding>
    <div class="row items-start q-gutter-xl">
      <div class="col">
        <q-card class="my-card bg-accent text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Total</div>
              <div class="text-h2">{{ summary.total }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="my-card bg-positive text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Completed</div>
              <div class="text-h2">{{ summary.completed }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="my-card bg-info text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Processing</div>
              <div class="text-h2">{{ summary.processing }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="my-card bg-negative text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Failed</div>
              <div class="text-h2">{{ summary.failed }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { api } from 'boot/axios';
import type { Client as PaymentDashboardClient } from '../openapi/openapi';

const client = await api.getClient<PaymentDashboardClient>();

const summaryRes = await client.getPaymentSummary();
const summary = summaryRes.data;
</script>
