<template>
  <q-page padding>
    <div class="row items-start q-gutter-xl">
      <div class="col">
        <q-card class="bg-accent text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Total</div>
              <div class="text-h2">{{ summary.total }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="bg-positive text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Completed</div>
              <div class="text-h2">{{ summary.completed }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="bg-info text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Processing</div>
              <div class="text-h2">{{ summary.processing }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col">
        <q-card class="bg-negative text-white">
          <q-card-section>
            <div class="row items-start justify-between">
              <div class="text-h6">Failed</div>
              <div class="text-h2">{{ summary.failed }}</div>
            </div>
          </q-card-section>
        </q-card>
      </div>
    </div>

    <q-card class="q-mt-xl">
      <q-toolbar class="bg-grey-3">
        <q-toolbar-title shrink>Payment List</q-toolbar-title>
        <q-separator vertical inset class="q-mx-sm" />
        <q-btn-dropdown flat label="Status">
          <q-list dense style="min-width: 150px">
            <q-item clickable v-ripple @click="statusSelected('')">
              <q-item-section> All </q-item-section>
              <q-item-section side>
                <q-icon name="check" v-show="filterStatus == ''" />
              </q-item-section>
            </q-item>
            <q-item clickable v-ripple @click="statusSelected('completed')">
              <q-item-section> Completed </q-item-section>
              <q-item-section side>
                <q-icon name="check" v-show="filterStatus == 'completed'" />
              </q-item-section>
            </q-item>
            <q-item clickable v-ripple @click="statusSelected('processing')">
              <q-item-section> Processing </q-item-section>
              <q-item-section side>
                <q-icon name="check" v-show="filterStatus == 'processing'" />
              </q-item-section>
            </q-item>
            <q-item clickable v-ripple @click="statusSelected('failed')">
              <q-item-section> Failed </q-item-section>
              <q-item-section side>
                <q-icon name="check" v-show="filterStatus == 'failed'" />
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
      </q-toolbar>
    </q-card>

    <q-table
      :columns="tblColumns"
      :rows="tblRows || []"
      row-key="merchant"
      v-model:pagination="tblPagination"
      :loading="tblLoading"
    />
  </q-page>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { QTableColumn } from 'quasar';

import { api } from 'boot/axios';
import type {
  Client as PaymentDashboardClient,
  Paths as PaymentDashboardPaths,
  Payment,
} from '../openapi/openapi';

const client = await api.getClient<PaymentDashboardClient>();

// ----- Filter status -----

const filterStatus = ref('');

async function statusSelected(status: string) {
  filterStatus.value = status;
  await fetchPaymentList();
}

// ----- Data table -----

const tblColumns: QTableColumn[] = [
  { name: 'merchant', label: 'Merchant', field: 'merchant', align: 'left', sortable: true },
  { name: 'status', label: 'Status', field: 'status', align: 'left', sortable: true },
  { name: 'amount', label: 'Amount', field: 'amount', sortable: true },
  { name: 'created_at', label: 'Date', field: 'created_at', sortable: true },
];

const tblRows = ref<Payment[]>();
const tblLoading = ref(false);

const tblPagination = ref({
  sortBy: 'created_at',
  descending: true,
  rowsPerPage: 10,
});

// ----- Fetch payment summary, only once -----

const summaryRes = await client.getPaymentSummary();
const summary = summaryRes.data;

// ----- Fetch payment list -----

async function fetchPaymentList() {
  const { sortBy, descending } = tblPagination.value;

  tblLoading.value = true;

  const params = <PaymentDashboardPaths.GetPaymentList.QueryParameters>{};
  if (filterStatus.value) params.status = filterStatus.value;
  if (sortBy) {
    params.sort = descending ? '-' : '';
    params.sort += sortBy;
  }

  const listRes = await client.getPaymentList(params);
  tblRows.value = listRes.data.payments;

  tblLoading.value = false;
}

// ----- Lifecycle hooks -----

onMounted(async () => {
  await fetchPaymentList();
});
</script>
