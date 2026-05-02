<template>
  <div class="cashflow-dashboard">
    <h2>Анализ денежных потоков</h2>
    <p class="subtitle">Операционная, инвестиционная и финансовая деятельность по кварталам</p>

    <!-- Карточки последнего квартала -->
    <div v-if="latest" class="summary-cards">
      <div class="summary-card">
        <span class="label">Последняя дата</span>
        <span class="value">{{ latest.fiscaldateending }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Операционный поток (CFO)</span>
        <span class="value">{{ formatMoney(latest.operatingcashflow) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Инвестиционный поток (CFI)</span>
        <span class="value">{{ formatMoney(latest.cashflowfrominvestment) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Финансовый поток (CFF)</span>
        <span class="value">{{ formatMoney(latest.cashflowfromfinancing) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Свободный денежный поток</span>
        <span class="value">{{ formatMoney(freeCashFlow(latest)) }}</span>
      </div>
      <div class="summary-card">
        <span class="label">Чистая прибыль</span>
        <span class="value">{{ formatMoney(latest.netincome) }}</span>
      </div>
    </div>

    <!-- График динамики денежных потоков -->
    <div class="chart-section">
      <h3>Динамика денежных потоков</h3>
      <div class="chart-container">
        <canvas ref="flowCanvas" width="800" height="350" class="flow-canvas"></canvas>
      </div>
      <div class="chart-note">Положительные значения – приток, отрицательные – отток денежных средств</div>
    </div>

    <!-- Дополнительные метрики: соотношение CFO / Net Income -->
    <div class="ratio-section" v-if="latest">
      <h3>Качество прибыли</h3>
      <div class="ratio-card">
        <div class="ratio-value">{{ cashFlowToNetIncomeRatio(latest) }}</div>
        <div class="ratio-desc">Соотношение операционного потока к чистой прибыли (CFO / Net Income)</div>
        <div class="ratio-note">Значение > 1 указывает на высокое качество прибыли (больше денег, чем учтённой прибыли)</div>
      </div>
    </div>

    <!-- Таблица всех данных (от новых к старым) -->
    <div class="table-wrapper">
      <table class="cashflow-table">
        <thead>
          <tr>
            <th>Дата</th><th>Операционный поток</th><th>Инвестиционный поток</th><th>Финансовый поток</th>
            <th>Свободный поток</th><th>Чистая прибыль</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in sortedDataReversed" :key="item.fiscaldateending">
            <td>{{ item.fiscaldateending }}</td>
            <td>{{ formatMoney(item.operatingcashflow, true) }}</td>
            <td>{{ formatMoney(item.cashflowfrominvestment, true) }}</td>
            <td>{{ formatMoney(item.cashflowfromfinancing, true) }}</td>
            <td>{{ formatMoney(freeCashFlow(item), true) }}</td>
            <td>{{ formatMoney(item.netincome, true) }}</td>
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

// Вспомогательные функции
function formatMoney(value: any, signed: boolean = false): string {
  const num = Number(value)
  if (isNaN(num)) return '—'
  const absNum = Math.abs(num)
  let formatted = ''
  if (absNum >= 1e12) formatted = (absNum / 1e12).toFixed(2) + ' трлн'
  else if (absNum >= 1e9) formatted = (absNum / 1e9).toFixed(2) + ' млрд'
  else if (absNum >= 1e6) formatted = (absNum / 1e6).toFixed(2) + ' млн'
  else if (absNum >= 1e3) formatted = (absNum / 1e3).toFixed(2) + ' тыс'
  else formatted = absNum.toFixed(0)
  if (signed && num < 0) return `-${formatted}`
  return formatted
}

function freeCashFlow(item: any): number {
  const ocf = Number(item.operatingcashflow) || 0
  const capex = Number(item.capitalexpenditures) || 0
  return ocf - capex
}

function cashFlowToNetIncomeRatio(item: any): string {
  const ocf = Number(item.operatingcashflow) || 0
  const net = Number(item.netincome) || 1
  if (net === 0) return '∞'
  return (ocf / net).toFixed(2)
}

// ----- Линейный график денежных потоков -----
const flowCanvas = ref<HTMLCanvasElement | null>(null)

const drawFlowChart = () => {
  if (!flowCanvas.value) return
  const data = sortedDataAsc.value
  if (data.length < 2) return

  const canvas = flowCanvas.value
  const ctx = canvas.getContext('2d')
  if (!ctx) return

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
  const operating = data.map(d => Number(d.operatingcashflow) || 0)
  const investing = data.map(d => Number(d.cashflowfrominvestment) || 0)
  const financing = data.map(d => Number(d.cashflowfromfinancing) || 0)
  const fcf = data.map(d => freeCashFlow(d))

  // Все значения для масштаба (берём абсолютные, но шкала линейная)
  const allValues = [...operating, ...investing, ...financing, ...fcf]
  const maxVal = Math.max(...allValues, 1)
  const minVal = Math.min(...allValues, -1)
  const range = Math.max(Math.abs(maxVal), Math.abs(minVal))
  const yTop = range
  const yBottom = -range

  const getY = (val: number) => {
    // пропорционально от yBottom до yTop
    const t = (val - yBottom) / (yTop - yBottom)
    return padding.top + graphHeight * (1 - t)
  }
  const getX = (index: number) => {
    return padding.left + (index / (data.length - 1)) * graphWidth
  }

  ctx.clearRect(0, 0, width, height)

  // Сетка и оси Y
  ctx.save()
  ctx.strokeStyle = '#e2e8f0'
  ctx.lineWidth = 1
  ctx.fillStyle = '#4a5568'
  ctx.font = '11px sans-serif'

  // Горизонтальные линии для положительных и отрицательных значений
  const ySteps = 5
  for (let i = 0; i <= ySteps; i++) {
    const yVal = yBottom + (range * 2 / ySteps) * i
    const y = getY(yVal)
    ctx.beginPath()
    ctx.moveTo(padding.left, y)
    ctx.lineTo(width - padding.right, y)
    ctx.stroke()
    ctx.fillStyle = '#4a5568'
    ctx.fillText(formatMoneyShort(yVal), 5, y + 3)
  }
  // Нулевая линия
  const yZero = getY(0)
  ctx.beginPath()
  ctx.strokeStyle = '#000'
  ctx.lineWidth = 1.5
  ctx.moveTo(padding.left, yZero)
  ctx.lineTo(width - padding.right, yZero)
  ctx.stroke()

  // Ось X: даты
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
  ctx.fillText('Денежный поток (USD)', -60, 40)
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
  drawLine(operating, '#42b983')
  drawLine(investing, '#f56565')
  drawLine(financing, '#4299e1')
  drawLine(fcf, '#9c27b0')

  // Легенда
  ctx.font = '12px sans-serif'
  ctx.fillStyle = '#42b983'
  ctx.fillText('Операционный поток', width - 130, 30)
  ctx.fillStyle = '#f56565'
  ctx.fillText('Инвестиционный поток', width - 130, 50)
  ctx.fillStyle = '#4299e1'
  ctx.fillText('Финансовый поток', width - 130, 70)
  ctx.fillStyle = '#9c27b0'
  ctx.fillText('Свободный поток', width - 130, 90)
  ctx.restore()
}

function formatMoneyShort(num: number): string {
  const absNum = Math.abs(num)
  let formatted = ''
  if (absNum >= 1e12) formatted = (absNum / 1e12).toFixed(1) + 'T'
  else if (absNum >= 1e9) formatted = (absNum / 1e9).toFixed(1) + 'B'
  else if (absNum >= 1e6) formatted = (absNum / 1e6).toFixed(1) + 'M'
  else if (absNum >= 1e3) formatted = (absNum / 1e3).toFixed(1) + 'K'
  else formatted = absNum.toFixed(0)
  return num < 0 ? `-${formatted}` : formatted
}

watch(sortedDataAsc, async () => {
  await nextTick()
  drawFlowChart()
})

onMounted(() => {
  drawFlowChart()
  window.addEventListener('resize', () => drawFlowChart())
})
</script>

<style scoped>
.cashflow-dashboard {
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
.chart-section {
  margin-bottom: 2rem;
}
.chart-container {
  background: white;
  border-radius: 16px;
  padding: 0.5rem;
  box-shadow: 0 1px 3px rgba(0,0,0,0.05);
}
.flow-canvas {
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
.ratio-section {
  margin-bottom: 2rem;
  text-align: center;
}
.ratio-card {
  background: linear-gradient(135deg, #f7fafc, #edf2f7);
  border-radius: 24px;
  padding: 1.2rem;
  max-width: 400px;
  margin: 0 auto;
}
.ratio-value {
  font-size: 2rem;
  font-weight: bold;
  color: #2d3748;
}
.ratio-desc {
  font-size: 0.85rem;
  color: #4a5568;
  margin-top: 0.3rem;
}
.ratio-note {
  font-size: 0.7rem;
  color: #a0aec0;
  margin-top: 0.5rem;
}
.table-wrapper {
  overflow-x: auto;
  border-radius: 16px;
  border: 1px solid #e2e8f0;
}
.cashflow-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.8rem;
}
.cashflow-table th, .cashflow-table td {
  padding: 0.7rem;
  text-align: right;
  border-bottom: 1px solid #edf2f7;
}
.cashflow-table th:first-child, .cashflow-table td:first-child {
  text-align: left;
  position: sticky;
  left: 0;
  background: white;
}
.cashflow-table th {
  background: #f7fafc;
  font-weight: 600;
}
.cashflow-table tr:hover td {
  background: #f7fafc;
}
@media (max-width: 768px) {
  .summary-card .value { font-size: 1rem; }
  .ratio-value { font-size: 1.5rem; }
}
</style>