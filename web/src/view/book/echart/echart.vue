<template>
  <div class="gva-form-box">
    <div class="chart-container">
      <div ref="collegeBarChart" style="width: 600px; height: 400px;" />
      <div ref="collegePieChart" style="width: 600px; height: 400px;"></div>
    </div>
    <div class="chart-container">
      <div ref="publishBarChart" style="width: 600px; height: 400px;" />
      <div ref="publishPieChart" style="width: 600px; height: 400px;"></div>
    </div>
    <div class="chart-container">
      <div ref="yearBarChart" style="width: 600px; height: 400px;"></div>
      <div ref="yearPieChart" style="width: 600px; height: 400px;"></div>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue';
import * as echarts from 'echarts';
import {getEchart} from "@/api/book/sysBook";

const collegeBarChart = ref(null);
const collegePieChart = ref(null);
const publishBarChart = ref(null);
const publishPieChart = ref(null);
const yearBarChart = ref(null);
const yearPieChart = ref(null);

const collegeData = ref([]);
const publishData = ref([]);
const yearData = ref([]);

const fetchData = async () => {
  try {
    const response = await getEchart();
    if (response.code === 0) {
      collegeData.value = response.data.college;
      publishData.value = response.data.publish;
      yearData.value = response.data.year;
      renderCharts();
    } else {
      console.error('Failed to fetch data', response);
    }
  } catch (error) {
    console.error('An error occurred', error);
  }
};

const renderCharts = () => {
  const collegeBarChartInstance = echarts.init(collegeBarChart.value);
  const collegePieChartInstance = echarts.init(collegePieChart.value);
  const publishBarChartInstance = echarts.init(publishBarChart.value);
  const publishPieChartInstance = echarts.init(publishPieChart.value);
  const yearBarChartInstance = echarts.init(yearBarChart.value);
  const yearPieChartInstance = echarts.init(yearPieChart.value);

  const barOption = (title, data) => ({
    title: {text: title},
    tooltip: {},
    xAxis: {type: 'category', data: data.map(item => item.name)},
    yAxis: {type: 'value'},
    series: [{data: data.map(item => item.value), type: 'bar'}]
  });

  const pieOption = (title, data) => ({
    title: {text: title},
    tooltip: {trigger: 'item'},
    series: [{
      type: 'pie',
      radius: '50%',
      data: data.map(item => ({name: item.name, value: item.value}))
    }]
  });

  collegeBarChartInstance.setOption(barOption('学院汇总柱状图', collegeData.value));
  collegePieChartInstance.setOption(pieOption('学院汇总饼图', collegeData.value));
  publishBarChartInstance.setOption(barOption('出版社汇总柱状图', publishData.value));
  publishPieChartInstance.setOption(pieOption('出版社汇总饼图', publishData.value));
  yearBarChartInstance.setOption(barOption('出版年份汇总柱状图', yearData.value));
  yearPieChartInstance.setOption(pieOption('出版年份汇总饼图', yearData.value));
};

onMounted(() => {
  fetchData();
});
</script>

<style scoped>
.gva-form-box {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.chart-container {
  display: flex;
  justify-content: space-between;
}
</style>
