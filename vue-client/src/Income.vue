<template>
  <div class="income-dashboard">
    <h2>Отчет о прибылях</h2>
    <p class="subtitle">Квартальные финансовые показатели</p>

    <!-- Карточки-сводки за последний квартал -->
    <div v-if="latest" class="summary-cards">
      <div class="summary-card">
        <span class="label">Последний квартал</span>
        <span class="value">{{ latest.fiscaldateending }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Общая выручка</span>
        <span class="value">{{ formatMoney(latest.totalrevenue) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Чистая прибыль</span>
        <span class="value">{{ formatMoney(latest.netincome) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Операционная прибыль</span>
        <span class="value">{{ formatMoney(latest.operatingincome) }}</span>
      </div>
    </div>

    <!-- График Revenue / Net Income -->
    <div class="chart-section">
      <h3>Динамика выручки и чистой прибыли</h3>
      <div class="bar-chart">
        <div class="chart-row" v-for="(item, idx) in lastEightQuarters" :key="idx">
          <div class="chart-label">{{ shortDate(item.fiscaldateending) }}</div>
          <div class="chart-bars">
            <div class="bar revenue-bar" :style="{ width: revenuePercent(item.totalrevenue) + '%' }" :title="`Выручка: ${formatMoney(item.totalrevenue)}`"></div>
            <div class="bar net-bar" :style="{ width: netPercent(item.netincome) + '%' }" :title="`Чистая прибыль: ${formatMoney(item.netincome)}`"></div>
          </div>
          <div class="chart-values">
            <span class="rev">{{ formatMoneyShort(item.totalrevenue) }}</span>
            <span class="net">{{ formatMoneyShort(item.netincome) }}</span>
          </div>
        </div>
      </div>
      <div class="legend">
        <span><div class="legend-color rev"></div> Выручка</span>
        <span><div class="legend-color net"></div> Чистая прибыль</span>
      </div>
    </div>

    <!-- Таблица всех данных -->
    <div class="table-wrapper">
      <table class="income-table">
        <thead>
          <tr>
            <th>Дата</th><th>Выручка</th><th>Себестоимость выручки</th><th>Валовая прибыль</th>
            <th>Операционная прибыль</th><th>Чистая прибыль</th><th>R&D</th><th>SGA</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in sortedData" :key="item.fiscaldateending">
            <td>{{ item.fiscaldateending }}</td>
            <td>{{ formatMoney(item.totalrevenue) }}</td>
            <td>{{ formatMoney(item.costofrevenue) }}</td>
            <td>{{ formatMoney(item.grossprofit) }}</td>
            <td>{{ formatMoney(item.operatingincome) }}</td>
            <td>{{ formatMoney(item.netincome) }}</td>
            <td>{{ formatMoney(item.researchanddevelopment) }}</td>
            <td>{{ formatMoney(item.sellinggeneralandadministrative) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  data: any // может быть { response: [...] } или напрямую массив
}>()

// Нормализуем данные – извлекаем массив
const rawData = computed(() => {
  if (!props.data) return []
  if (Array.isArray(props.data)) return props.data
  if (props.data.response && Array.isArray(props.data.response)) return props.data.response
  return []
})

// Сортируем по дате (от старых к новым)
const sortedData = computed(() => {
  return [...rawData.value].sort((a, b) => 
    new Date(b.fiscaldateending).getTime() - new Date(a.fiscaldateending).getTime()
  )
})

// Последний квартал
const latest = computed(() => sortedData.value[0])

// Последние 8 кварталов для графика (или все, если меньше 8)
const lastEightQuarters = computed(() => sortedData.value.slice(0, 8))

// Находим максимумы для масштабирования графика
const maxRevenue = computed(() => {
  return Math.max(...sortedData.value.map(d => Number(d.totalrevenue) || 0), 1)
})
// const maxNet = computed(() => {
//   return Math.max(...sortedData.value.map(d => Number(d.netincome) || 0), 1)
// })

function revenuePercent(value: any): number {
  const rev = Number(value) || 0
  return (rev / maxRevenue.value) * 100
}
function netPercent(value: any): number {
  const net = Number(value) || 0
  return (net / maxRevenue.value) * 100
}

// Форматирование денег
function formatMoney(value: any): string {
  const num = Number(value)
  if (isNaN(num)) return '—'
  if (num >= 1e9) return (num / 1e9).toFixed(2) + 'B'
  if (num >= 1e6) return (num / 1e6).toFixed(2) + 'M'
  if (num >= 1e3) return (num / 1e3).toFixed(2) + 'K'
  return num.toFixed(0)
}
function formatMoneyShort(value: any): string {
  const num = Number(value)
  if (isNaN(num)) return '—'
  if (num >= 1e9) return (num / 1e9).toFixed(1) + 'B'
  if (num >= 1e6) return (num / 1e6).toFixed(0) + 'M'
  if (num >= 1e3) return (num / 1e3).toFixed(0) + 'K'
  return num.toFixed(0)
}
function shortDate(dateStr: string): string {
  if (!dateStr) return ''
  const parts = dateStr.split('-')
  if (parts.length >= 2) return `${parts[1]}/${parts[0].slice(2)}`
  return dateStr
}
</script>

<style scoped>
.income-dashboard {
  width: 100%;
}
.subtitle {
  color: #718096;
  margin-bottom: 1.5rem;
}
.summary-cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}
.summary-card {
  background: linear-gradient(135deg, #f7fafc 0%, #edf2f7 100%);
  border-radius: 20px;
  padding: 1rem;
  text-align: center;
  display: flex; flex-direction: column;
}
.summary-card .label {
  font-size: 0.8rem;
  text-transform: uppercase;
  color: #4a5568;
}
.summary-card .value {
  font-size: 1.3rem;
  font-weight: bold;
  color: #2d3748;
}
.chart-section {
  margin-bottom: 2rem;
}
.chart-row {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}
.chart-label {
  width: 60px;
  font-size: 0.75rem;
  color: #4a5568;
}
.chart-bars {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
  background: #edf2f7;
  border-radius: 12px;
  padding: 4px 0;
}
.bar {
  height: 14px;
  border-radius: 7px;
  transition: width 0.2s;
}
.revenue-bar {
  background-color: #42adb9;
}
.net-bar {
  background-color: #658ef5;
}
.chart-values {
  width: 100px;
  display: flex;
  justify-content: space-between;
  font-size: 0.7rem;
  font-weight: 500;
}
.chart-values .rev { color: #42adb9; }
.chart-values .net { color: #658ef5; }
.legend {
  display: flex;
  gap: 1rem;
  margin-top: 0.5rem;
  justify-content: flex-end;
  font-size: 0.8rem;
}
.legend-color {
  display: inline-block;
  width: 12px;
  height: 12px;
  border-radius: 2px;
  margin-right: 4px;
}
.legend-color.rev { background: #42adb9; }
.legend-color.net { background: #658ef5; }
.table-wrapper {
  overflow-x: auto;
  border-radius: 16px;
  border: 1px solid #e2e8f0;
}
.income-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.8rem;
}
.income-table th, .income-table td {
  padding: 0.75rem;
  text-align: right;
  border-bottom: 1px solid #edf2f7;
}
.income-table th:first-child, .income-table td:first-child {
  text-align: left;
  position: sticky;
  left: 0;
  background: white;
}
.income-table th {
  background: #f7fafc;
  font-weight: 600;
  color: #2d3748;
}
.income-table tr:hover td {
  background: #f7fafc;
}
@media (max-width: 768px) {
  .summary-card .value { font-size: 1rem; }
  .chart-values { width: 80px; font-size: 0.6rem; }
  .income-table { font-size: 0.7rem; }
}
</style>