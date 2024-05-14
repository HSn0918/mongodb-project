<template>
  <div>
    <el-select
        v-model="queryCollegeName"
        :placeholder="queryCollegeName || 'Select'"
        size="large"
        style="width: 240px"
        @focus="fetchCollegeNames"
    >
      <el-option
          v-for="(item, index) in collegeName"
          :key="index"
          :label="item"
          :value="item"
      />
    </el-select>
    <div class="gva-form-box">

      <el-form>
        <el-form-item label="负责人信息">
          <div v-if="collegeDean">
            <p>姓名：{{ collegeDean.name }}</p>
            <p>联系方式：{{ collegeDean.contact }}</p>
          </div>
        </el-form-item>
      </el-form>



    </div>
    <div class="gva-form-box">
      <el-form-item label="教师信息">
        <el-list>
          <el-list-item v-for="(item, index) in teacher" :key="index">
            <div class="teacher-info">
              <span>姓名：{{ item.name }}</span>
              <span>联系方式：{{ item.contact }}</span>
            </div>
          </el-list-item>
        </el-list>
      </el-form-item>
    </div>

  </div>
</template>

<script setup>
import {ElMessage} from 'element-plus'
import {ref, watch} from 'vue'
import {findAllcollegeName, findCollegeDean,findAllTeacher} from "@/api/college/sysCollege"

const collegeName = ref([])
const queryCollegeName = ref('')
const collegeDean = ref('')
const teacher = ref([])
const fetchCollegeNames = async () => {
  try {
    const response = await findAllcollegeName()
    if (response && response.data) {
      collegeName.value = response.data
    }

  } catch (error) {
    ElMessage.error('Failed to fetch college names')
  }
}

watch(queryCollegeName, async (newVal) => {
  if (newVal) {
    try {
      console.log(queryCollegeName.value);
      let data = {
        "name": queryCollegeName.value
      };
      console.log(data);
      let response = await findCollegeDean(data);
      console.log(response);
      if (response && response.data) {
        collegeDean.value = {
          id: response.data.id,
          name: response.data.name,
          contact: response.data.contact
        };
      }
      teacher.value = [];
      response = await findAllTeacher(data);
      if (response && response.data) {
        for(let i=0;i<response.data.length;i++){
          teacher.value.push({
            name: response.data[i].name,
            contact: response.data[i].contact
          })
        }
        console.log(teacher.value);
      }


    } catch (error) {
      ElMessage.error('Failed to fetch college dean');
    }
  }
});
</script>
<style>
.teacher-info {
  display: flex;
  gap: 20px; /* 控制姓名和联系方式之间的间距 */
}
.teacher-info p {
  margin: 0;
}
</style>
