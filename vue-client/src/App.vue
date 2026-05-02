<template>
  <div class="app">
    <header class="app-header">
      <div class="header-container">
        <!-- Гамбургер (показывается только когда сайдбар скрыт) -->
        <div v-if="isSidebarHidden" class="hamburger-wrapper">
          <button class="hamburger-btn" @click="toggleMobileMenu">
            ☰
          </button>
          <div v-if="mobileMenuOpen" class="mobile-dropdown">
            <ul>
              <li
                v-for="company in companies"
                :key="company.symbol"
                @click="selectCompany(company.symbol); mobileMenuOpen = false"
              >
                {{ company.name }}
              </li>
            </ul>
          </div>
        </div>
        <div v-else class="placeholder"></div>

        <nav class="tabs">
          <button
            v-for="tab in tabs"
            :key="tab.key"
            :class="['tab-btn', { active: activeTab === tab.key }]"
            @click="switchTab(tab.key)"
          >
            {{ tab.label }}
          </button>
        </nav>

        <div class="placeholder"></div>
      </div>
    </header>

    <div class="main-layout">
      <aside v-show="!isSidebarHidden" class="company-sidebar">
        <h3>Компании</h3>
        <ul>
          <li
            v-for="company in companies"
            :key="company.symbol"
            :class="{ active: selectedCompany === company.symbol }"
            @click="selectCompany(company.symbol)"
          >
            {{ company.name }}
          </li>
        </ul>
      </aside>

      <main class="content">
        <div v-if="error" class="error">{{ error }}</div>
        <div v-else-if="loading" class="loading">Загрузка...</div>
        <div v-else-if="!selectedCompany" class="info">Выберите компанию из списка</div>

        <!-- Overview Дашборд -->
        <div v-else-if="activeTab === 'overview' && data" class="overview-dashboard">
          <h2>{{ data.name }} ({{ data.symbol }})</h2>
          <p class="sector">{{ data.sector }} · {{ data.industry }}</p>

          <!-- Спидометр финансовой надежности -->
          <div class="gauge-section">
            <h3>Финансовая надёжность</h3>
            <canvas ref="gaugeCanvas" width="400" height="180" class="gauge-canvas"></canvas>
            <div class="gauge-labels">
              <span>Критическая</span><span>Низкая</span><span>Средняя</span><span>Высокая</span>
            </div>
            <div class="reliability-text" :style="{ color: gaugeColor }">
              {{ reliabilityText }}
            </div>
            <div class="reliability-score">Оценка: {{ reliabilityScore }} / 100</div>
          </div>

          <!-- Карточки ключевых метрик -->
          <div class="metrics-grid">
            <div class="metric-card">
              <span class="label">Рыночная капитализация</span>
              <span class="value">{{ formatNumber(data.marketcapitalization) }} $</span>
            </div>
            <div class="metric-card">
              <span class="label">P/E</span>
              <span class="value">{{ data.peratio || '—' }}</span>
            </div>
            <div class="metric-card">
              <span class="label">Дивидендная доходность</span>
              <span class="value">{{ (data.dividendyield * 100).toFixed(2) }}%</span>
            </div>
            <div class="metric-card">
              <span class="label">Рентабельность по чистой прибыли</span>
              <span class="value">{{ (data.profitmargin * 100).toFixed(1) }}%</span>
            </div>
            <div class="metric-card">
              <span class="label">ROE (TTM)</span>
              <span class="value">{{ (data.returnonequityttm * 100).toFixed(1) }}%</span>
            </div>
            <div class="metric-card">
              <span class="label">Выручка (TTM)</span>
              <span class="value">{{ formatNumber(data.revenuettm) }} $</span>
            </div>
          </div>

          <!-- Бар-чарты -->
          <div class="chart-section">
            <h3>Показатели эффективности</h3>
            <div class="bar-item">
              <span>Рентабельность по чистой прибыли</span>
              <div class="bar-bg"><div class="bar-fill" :style="{ width: (data.profitmargin * 100) + '%', background: '#42b983' }"></div></div>
              <span>{{ (data.profitmargin * 100).toFixed(1) }}%</span>
            </div>
            <div class="bar-item">
              <span>Операционная рентабельность</span>
              <div class="bar-bg"><div class="bar-fill" :style="{ width: (data.operatingmarginttm * 100) + '%', background: '#ffc107' }"></div></div>
              <span>{{ (data.operatingmarginttm * 100).toFixed(1) }}%</span>
            </div>
            <div class="bar-item">
              <span>ROE</span>
              <div class="bar-bg"><div class="bar-fill" :style="{ width: (data.returnonequityttm * 100) + '%', background: '#17a2b8' }"></div></div>
              <span>{{ (data.returnonequityttm * 100).toFixed(1) }}%</span>
            </div>
            <div class="bar-item">
              <span>ROA</span>
              <div class="bar-bg"><div class="bar-fill" :style="{ width: (data.returnonassetsttm * 100) + '%', background: '#6f42c1' }"></div></div>
              <span>{{ (data.returnonassetsttm * 100).toFixed(1) }}%</span>
            </div>
          </div>

          <!-- Таблица рейтингов и дополнительных данных -->
          <div class="analyst-grid">
            <div class="info-block">
              <h4>Аналитический рейтинг</h4>
              <ul>
                <li>Сильная покупка: {{ data.analystratingstrongbuy || 0 }}</li>
                <li>Покупка: {{ data.analystratingbuy || 0 }}</li>
                <li>Держать: {{ data.analystratinghold || 0 }}</li>
                <li>Продавать: {{ data.analystratingsell || 0 }}</li>
                <li>Сильная продажа: {{ data.analystratingstrongsell || 0 }}</li>
              </ul>
            </div>
            <div class="info-block">
              <h4>Цена и риск</h4>
              <ul>
                <li>Максимум за 52 недели: {{ data.weekhigh52 || '—' }} $</li>
                <li>Минимум за 52 недели: {{ data.weeklow52 || '—' }} $</li>
                <li>Beta: {{ data.beta || '—' }}</li>
                <li>Целевая цена: {{ data.analysttargetprice || '—' }} $</li>
              </ul>
            </div>
            <div class="info-block">
              <h4>Дивиденды</h4>
              <ul>
                <li>Дивиденд на акцию: {{ data.dividendpershare || '—' }} $</li>
                <li>Дата закрытия реестра: {{ data.exdividenddate || '—' }}</li>
                <li>Дата выплаты дивидендов: {{ data.dividenddate || '—' }}</li>
              </ul>
            </div>
          </div>
        </div>

        <!-- Income Дашборд -->
        <IncomeDashboard v-else-if="activeTab === 'income' && data" :data="data" />
        <BalanceDashboard v-else-if="activeTab === 'balance' && data" :data="data" />
        <CashFlow v-else-if="activeTab === 'cashflow' && data" :data="data" />

        <div v-else-if="data" class="data">
          <h3>Ответ от {{ activeTabEndpoint }}</h3>
          <pre>{{ JSON.stringify(data, null, 2) }}</pre>
        </div>
        <div v-else class="info">Нет данных. Выберите компанию или переключите вкладку.</div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed, watch, nextTick } from 'vue'
import IncomeDashboard from './Income.vue'
import BalanceDashboard from './Balance.vue'
import CashFlow from './CashFlow.vue'

type TabKey = 'overview' | 'income' | 'balance' | 'cashflow'

const API_BASE = "/api"

const companies = [
  { name: 'META', symbol: 'META' },
  { name: 'GOOGL', symbol: 'GOOGL' },
  { name: 'MSFT', symbol: 'MSFT' },
  { name: 'NVDA', symbol: 'NVDA' },
  { name: 'IBM', symbol: 'IBM' },
  { name: 'NFLX', symbol: 'NFLX' },
]

const tabs: { key: TabKey; label: string; endpoint: string }[] = [
  { key: 'overview', label: 'Overview', endpoint: '/overview' },
  { key: 'income', label: 'IncomeStatement', endpoint: '/income' },
  { key: 'balance', label: 'BalanceSheet', endpoint: '/balance' },
  { key: 'cashflow', label: 'CashFlow', endpoint: '/cashflow' },
]

const selectedCompany = ref<string | null>(null)
const activeTab = ref<TabKey>('overview')
const loading = ref(false)
const error = ref<string | null>(null)
const data = ref<any>(null)

// Адаптивное скрытие сайдбара
const isSidebarHidden = ref(false)
const checkSidebarVisibility = () => {
  const { innerWidth: width, innerHeight: height } = window
  isSidebarHidden.value = width < height * 1.381
}

// Мобильное меню (гамбургер)
const mobileMenuOpen = ref(false)
const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

const activeTabEndpoint = computed(() => tabs.find(t => t.key === activeTab.value)?.endpoint || '')

// ----- Логика спидометра -----
const gaugeCanvas = ref<HTMLCanvasElement | null>(null)
const reliabilityScore = ref(0)
const gaugeColor = ref('#aaa')
const reliabilityText = ref('')

const calculateReliability = (companyData: any) => {
  let score = 0
  // Profit margin > 0.2 → 35 баллов
  if (companyData.profitmargin > 0.2) score += 35
  // ROE > 0.2 → 35 баллов
  if (companyData.returnonequityttm > 0.2) score += 35
  // PEG < 1 → 15 баллов
  if (companyData.pegratio < 1) score += 15
  // Beta < 1 → 15 баллов
  if (companyData.beta < 1) score += 15

  let color, text
  if (score <= 25) {
    color = '#e53e3e'
    text = 'Критическая надёжность – высокий риск'
  } else if (score <= 50) {
    color = '#ed8936'
    text = 'Низкая надёжность – действуйте осторожно'
  } else if (score <= 75) {
    color = '#ecc94b'
    text = 'Средняя надёжность – стабильно'
  } else {
    color = '#48bb78'
    text = 'Высокая надёжность – отлично'
  }
  return { score, color, text }
}

// Рисование спидометра (градуированная дуга со стрелкой)
const drawGauge = (score: number) => {
  if (!gaugeCanvas.value) return
  const canvas = gaugeCanvas.value
  const ctx = canvas.getContext('2d')
  if (!ctx) return
  const width = canvas.width, height = canvas.height
  ctx.clearRect(0, 0, width, height)
  
  const centerX = width / 2
  const centerY = height - 20
  const radius = 130
  const startAngle = -Math.PI
  const endAngle = 0
  const angleRange = endAngle - startAngle
  
  // Рисуем 4 цветных сектора
  const sections = [
    { color: '#e53e3e', start: 0, end: 0.25 },
    { color: '#ed8936', start: 0.25, end: 0.5 },
    { color: '#ecc94b', start: 0.5, end: 0.75 },
    { color: '#48bb78', start: 0.75, end: 1.0 }
  ]
  sections.forEach(s => {
    const angleStart = startAngle + angleRange * s.start
    const angleEnd = startAngle + angleRange * s.end
    ctx.beginPath();
    ctx.arc(centerX, centerY, radius, angleStart, angleEnd);
    ctx.arc(centerX, centerY, radius-50, angleEnd, angleStart, true);
    ctx.closePath();
    ctx.fillStyle = s.color;
    ctx.fill();
  })
  
  // Стрелка: угол пропорционально score
  const arrowAngle = startAngle + angleRange * (score / 100)
  const arrowLength = radius - 15
  const arrowX = centerX + Math.cos(arrowAngle) * arrowLength
  const arrowY = centerY + Math.sin(arrowAngle) * arrowLength
  ctx.beginPath()
  ctx.moveTo(centerX, centerY)
  ctx.lineTo(arrowX, arrowY)
  ctx.lineWidth = 4
  ctx.strokeStyle = '#2d3748'
  ctx.stroke()
  // Круг в центре
  ctx.beginPath()
  ctx.arc(centerX, centerY, 8, 0, 2 * Math.PI)
  ctx.fillStyle = '#2d3748'
  ctx.fill()
}

watch(data, async (newData) => {
  if (newData && activeTab.value === 'overview') {
    const { score, color, text } = calculateReliability(newData)
    reliabilityScore.value = score
    gaugeColor.value = color
    reliabilityText.value = text
    await nextTick()
    drawGauge(score)
  }
})

// ----- Вспомогательные функции -----
function formatNumber(num: string | number | undefined): string {
  if (!num) return '—'
  const n = Number(num)
  if (n >= 1e12) return (n / 1e12).toFixed(2) + 'T'
  if (n >= 1e9) return (n / 1e9).toFixed(2) + 'B'
  if (n >= 1e6) return (n / 1e6).toFixed(2) + 'M'
  return n.toFixed(2)
}

async function fetchData() {
  if (!selectedCompany.value) return
  loading.value = true
  error.value = null
  data.value = null

  const endpoint = activeTabEndpoint.value
  let body: any = { symbol: selectedCompany.value }
  if (activeTab.value !== 'overview') {
    body.range = {
      start: { year: 2005, quarter: 4 },
      end: { year: 2026, quarter: 3 },
    }
  }

  try {
    const response = await fetch(`${API_BASE}${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body),
    })
    if (!response.ok) throw new Error(`HTTP ${response.status}: ${response.statusText}`)
    data.value = await response.json()
  } catch (err: any) {
    console.error(err)
    error.value = err.message || 'Не удалось загрузить данные'
  } finally {
    loading.value = false
  }
}

function selectCompany(symbol: string) {
  if (selectedCompany.value !== symbol) {
    selectedCompany.value = symbol
    fetchData()
  }
}

function switchTab(tabKey: TabKey) {
  if (tabKey === activeTab.value) return
  activeTab.value = tabKey
  if (selectedCompany.value) fetchData()
}

onMounted(() => {
  checkSidebarVisibility()
  window.addEventListener('resize', checkSidebarVisibility)
  document.addEventListener('click', (e) => {
    const target = e.target as HTMLElement
    if (!target.closest('.hamburger-wrapper')) mobileMenuOpen.value = false
  })
})

onUnmounted(() => {
  window.removeEventListener('resize', checkSidebarVisibility)
})
</script>