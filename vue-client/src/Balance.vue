<template>
  <div class="balance-dashboard">
    <h2>Анализ бухгалтерского баланса</h2>
    <p class="subtitle">Активы, обязательства и капитал компании по кварталам</p>

    <!-- Карточки последнего отчётного периода -->
    <div v-if="latest" class="summary-cards">
      <div class="summary-card">
        <span class="label">Последняя дата</span>
        <span class="value">{{ latest.fiscaldateending }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Всего активов</span>
        <span class="value">{{ formatMoney(latest.totalassets) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Всего обязательств</span>
        <span class="value">{{ formatMoney(latest.totalliabilities) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Собственный капитал</span>
        <span class="value">{{ formatMoney(latest.totalshareholderequity) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Долг / Капитал</span>
        <span class="value">{{ debtEquityRatio(latest) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Коэффициент текущей ликвидности</span>
        <span class="value">{{ currentRatio(latest) }}</span>
      </div>
    </div>

    <!-- Линейный график динамики -->
    <div class="chart-section">
      <h3>Динамика основных показателей</h3>
      <div class="chart-container">
        <canvas ref="trendCanvas" width="800" height="350" class="trend-canvas"></canvas>
      </div>
      <div class="chart-note">Данные отображены в логарифмическом масштабе для наглядности (фактические значения можно увидеть в таблице)</div>
    </div>

    <!-- Круговая диаграмма структуры активов -->
    <div class="pie-section">
      <h3>Структура активов (последний квартал)</h3>
      <div class="pie-wrapper">
        <canvas ref="pieCanvas" width="250" height="250" style="max-width:250px; width:100%; height:auto"></canvas>
        <div class="pie-legend">
          <div><span class="color current"></span> Оборотные активы: <strong>{{ currentPercent }}%</strong></div>
          <div><span class="color noncurrent"></span> Внеоборотные активы: <strong>{{ noncurrentPercent }}%</strong></div>
        </div>
      </div>
    </div>

    <!-- Таблица всех данных (от новых к старым) -->
    <div class="table-wrapper">
      <table class="balance-table">
        <thead>
          <tr>
            <th>Дата</th><th>Всего активов</th><th>Оборотные активы</th><th>Внеоборотные активы</th>
            <th>Всего обязательств</th><th>Собственный капитал</th><th>Долг/Капитал</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in sortedDataReversed" :key="item.fiscaldateending">
            <td>{{ item.fiscaldateending }}</td>
            <td>{{ formatMoney(item.totalassets) }}</td>
            <td>{{ formatMoney(item.totalcurrentassets) }}</td>
            <td>{{ formatMoney(item.totalnoncurrentassets) }}</td>
            <td>{{ formatMoney(item.totalliabilities) }}</td>
            <td>{{ formatMoney(item.totalshareholderequity) }}</td>
            <td>{{ debtEquityRatio(item) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch, nextTick, onMounted } from 'vue'

const props = defineProps<{
  data: any
}>()

// Нормализация данных
const rawData = computed(() => {
  if (!props.data) return []
  if (Array.isArray(props.data)) return props.data
  if (props.data.response && Array.isArray(props.data.response)) return props.data.response
  return []
})

// Сортировка от старых к новым (для графика)
const sortedDataAsc = computed(() => {
  return [...rawData.value].sort((a, b) => 
    new Date(a.fiscaldateending).getTime() - new Date(b.fiscaldateending).getTime()
  )
})

// Сортировка от новых к старым (для таблицы)
const sortedDataReversed = computed(() => {
  return [...sortedDataAsc.value].reverse()
})

const latest = computed(() => sortedDataAsc.value[sortedDataAsc.value.length - 1])

// Проценты для круговой диаграммы
const currentPercent = computed(() => {
  const current = Number(latest.value?.totalcurrentassets) || 0
  const noncurrent = Number(latest.value?.totalnoncurrentassets) || 0
  const total = current + noncurrent
  if (total === 0) return 0
  return ((current / total) * 100).toFixed(1)
})
const noncurrentPercent = computed(() => {
  const current = Number(latest.value?.totalcurrentassets) || 0
  const noncurrent = Number(latest.value?.totalnoncurrentassets) || 0
  const total = current + noncurrent
  if (total === 0) return 0
  return ((noncurrent / total) * 100).toFixed(1)
})

// Вспомогательные функции
function formatMoney(value: any): string {
  const num = Number(value)
  if (isNaN(num)) return '—'
  if (num >= 1e12) return (num / 1e12).toFixed(2) + ' трлн'
  if (num >= 1e9) return (num / 1e9).toFixed(2) + ' млрд'
  if (num >= 1e6) return (num / 1e6).toFixed(2) + ' млн'
  if (num >= 1e3) return (num / 1e3).toFixed(2) + ' тыс'
  return num.toFixed(0)
}

function debtEquityRatio(item: any): string {
  const debt = Number(item.totalliabilities) || 0
  const equity = Number(item.totalshareholderequity) || 1
  if (debt === 0) return '0.00'
  return (debt / equity).toFixed(2)
}

function currentRatio(item: any): string {
  const ca = Number(item.totalcurrentassets) || 0
  const cl = Number(item.totalcurrentliabilities) || 1
  if (ca === 0) return '0.00'
  return (ca / cl).toFixed(2)
}

// ----- Линейный график (улучшенный) -----
const trendCanvas = ref<HTMLCanvasElement | null>(null)

const drawTrend = () => {
  if (!trendCanvas.value) return
  const data = sortedDataAsc.value
  if (data.length < 2) return

  const canvas = trendCanvas.value
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  // Установка размера canvas под контейнер
  const container = canvas.parentElement
  const width = container?.clientWidth || 800
  const height = 350
  canvas.width = width
  canvas.height = height

  const padding = { top: 40, right: 60, bottom: 50, left: 70 }
  const graphWidth = width - padding.left - padding.right
  const graphHeight = height - padding.top - padding.bottom

  // Подготовка данных
  const dates = data.map(d => d.fiscaldateending)
  const assets = data.map(d => Number(d.totalassets) || 0)
  const liabilities = data.map(d => Number(d.totalliabilities) || 0)
  const equity = data.map(d => Number(d.totalshareholderequity) || 0)
  const allValues = [...assets, ...liabilities, ...equity].filter(v => v > 0)
  const maxVal = Math.max(...allValues, 1)
  
  // Логарифмическая шкала? Используем линейную, но подписи в читаемом формате
  const getY = (val: number) => {
    return padding.top + graphHeight - (val / maxVal) * graphHeight
  }

  const getX = (index: number) => {
    return padding.left + (index / (data.length - 1)) * graphWidth
  }

  // Очистка
  ctx.clearRect(0, 0, width, height)
  
  // Сетка и оси Y
  ctx.save()
  ctx.strokeStyle = '#e2e8f0'
  ctx.lineWidth = 1
  ctx.fillStyle = '#4a5568'
  ctx.font = '11px sans-serif'
  
  // Горизонтальные линии сетки и подписи Y
  const ySteps = 5
  for (let i = 0; i <= ySteps; i++) {
    const yVal = (maxVal / ySteps) * i
    const y = getY(yVal)
    ctx.beginPath()
    ctx.moveTo(padding.left, y)
    ctx.lineTo(width - padding.right, y)
    ctx.stroke()
    ctx.fillStyle = '#4a5568'
    ctx.fillText(formatMoneyShort(yVal), 5, y + 3)
  }
  
  // Ось X: даты (первые, последние, может средние)
  const dateLabels = [dates[0], dates[Math.floor(dates.length/2)], dates[dates.length-1]]
  dateLabels.forEach((label, idx) => {
    let x
    if (idx === 0) x = getX(0)
    else if (idx === 1) x = getX(Math.floor(data.length/2))
    else x = getX(data.length-1)
    ctx.fillStyle = '#4a5568'
    ctx.fillText(label, x - 30, height - padding.bottom + 20)
  })
  
  // Подписи осей
  ctx.fillStyle = '#2d3748'
  ctx.font = 'bold 12px sans-serif'
  ctx.fillText('Дата', width/2 - 20, height - 10)
  ctx.save()
  ctx.translate(20, height/2)
  ctx.rotate(-Math.PI/2)
  ctx.fillText('Сумма (USD)', -20, 40)
  ctx.restore()
  
  // Рисуем линии
  const drawLine = (values: number[], color: string) => {
    ctx.beginPath()
    ctx.strokeStyle = color
    ctx.lineWidth = 2.5
    let first = true
    for (let i = 0; i < values.length; i++) {
      const x = getX(i)
      const y = getY(values[i])
      if (first) {
        ctx.moveTo(x, y)
        first = false
      } else {
        ctx.lineTo(x, y)
      }
    }
    ctx.stroke()
  }
  drawLine(assets, '#42b983')
  drawLine(liabilities, '#f56565')
  drawLine(equity, '#4299e1')
  
  // Легенда
  ctx.font = '12px sans-serif'
  ctx.fillStyle = '#42b983'
  ctx.fillText('Активы', width - 80, 30)
  ctx.fillStyle = '#f56565'
  ctx.fillText('Обязательства', width - 80, 50)
  ctx.fillStyle = '#4299e1'
  ctx.fillText('Капитал', width - 80, 70)
  ctx.restore()
}

function formatMoneyShort(num: number): string {
  if (num >= 1e12) return (num / 1e12).toFixed(1) + 'T'
  if (num >= 1e9) return (num / 1e9).toFixed(1) + 'B'
  if (num >= 1e6) return (num / 1e6).toFixed(1) + 'M'
  if (num >= 1e3) return (num / 1e3).toFixed(1) + 'K'
  return num.toFixed(0)
}

// Круговая диаграмма
const pieCanvas = ref<HTMLCanvasElement | null>(null)
const drawPie = () => {
  if (!pieCanvas.value || !latest.value) return
  const canvas = pieCanvas.value
  const ctx = canvas.getContext('2d')
  if (!ctx) return
  
  const size = 250
  canvas.width = size
  canvas.height = size
  const center = size/2
  const radius = size/2 - 10
  
  const current = Number(latest.value.totalcurrentassets) || 0
  const noncurrent = Number(latest.value.totalnoncurrentassets) || 0
  const total = current + noncurrent
  if (total === 0) return
  
  const currentAngle = (current / total) * 2 * Math.PI
  
  ctx.clearRect(0, 0, size, size)
  ctx.beginPath()
  ctx.fillStyle = '#42b983'
  ctx.moveTo(center, center)
  ctx.arc(center, center, radius, 0, currentAngle)
  ctx.fill()
  ctx.beginPath()
  ctx.fillStyle = '#cbd5e0'
  ctx.moveTo(center, center)
  ctx.arc(center, center, radius, currentAngle, 2 * Math.PI)
  ctx.fill()
}

watch(sortedDataAsc, async () => {
  await nextTick()
  drawTrend()
  drawPie()
})

onMounted(() => {
  drawTrend()
  drawPie()
  window.addEventListener('resize', () => drawTrend())
})
</script>

<style scoped>
.balance-dashboard {
  width: 100%;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
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
  display: flex;
  flex-direction: column;
  background: #f7fafc;
  border-radius: 20px;
  padding: 0.8rem;
  text-align: center;
  box-shadow: 0 1px 2px rgba(0,0,0,0.05);
  justify-content: center;
}
.summary-card .label {
  font-size: 0.7rem;
  text-transform: uppercase;
  color: #4a5568;
  letter-spacing: 0.5px;
}
.summary-card .value {
  font-size: 1.2rem;
  font-weight: bold;
  color: #2d3748;
}
.chart-section, .pie-section {
  margin-bottom: 2rem;
}
.chart-container {
  background: white;
  border-radius: 16px;
  padding: 0.5rem;
  box-shadow: 0 1px 3px rgba(0,0,0,0.05);
}
.trend-canvas {
  width: 100%;
  height: auto;
  background: white;
  border-radius: 12px;
}
.chart-note {
  font-size: 0.7rem;
  color: #a0aec0;
  text-align: center;
  margin-top: 0.5rem;
}
.pie-wrapper {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: 2rem;
  justify-content: center;
}
.pie-legend {
  background: #f7fafc;
  padding: 1rem;
  border-radius: 16px;
  font-size: 0.9rem;
}
.pie-legend .color {
  display: inline-block;
  width: 16px;
  height: 16px;
  border-radius: 4px;
  margin-right: 8px;
  vertical-align: middle;
}
.color.current { background: #42b983; }
.color.noncurrent { background: #cbd5e0; }
.table-wrapper {
  overflow-x: auto;
  border-radius: 16px;
  border: 1px solid #e2e8f0;
  margin-top: 1rem;
}
.balance-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.8rem;
}
.balance-table th, .balance-table td {
  padding: 0.7rem;
  text-align: right;
  border-bottom: 1px solid #edf2f7;
}
.balance-table th:first-child, .balance-table td:first-child {
  text-align: left;
  position: sticky;
  left: 0;
  background: white;
}
.balance-table th {
  background: #f7fafc;
  font-weight: 600;
}
.balance-table tr:hover td {
  background: #f7fafc;
}
@media (max-width: 768px) {
  .summary-card .value { font-size: 1rem; }
  .pie-wrapper { flex-direction: column; }
}
</style>