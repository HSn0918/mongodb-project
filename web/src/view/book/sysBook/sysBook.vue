<template>
  <div>
    <div class="gva-search-box">
      <el-input
        v-model="input"
        style="width: 240px"
        placeholder="Please input"
        clearable
      />
      <el-button :icon="Search" circle @click="HandlerSearch" />
    </div>
    <div class="gva-form-box">
      <el-form-item label="出版教材：">
        <el-list>
          <el-list-item v-for="(item, index) in books" :key="index">
            <div class="book-info">
              <span>教材名：{{ item.name }}</span>
              <span>学院：{{ item.college }}</span>
              <span>出版社：{{ item.publisher }}</span>
              <span>ISBN：{{ item.isbn }}</span>
              <span>作者ID：{{ item.authors.join(', ') }}</span>
            </div>
          </el-list-item>
        </el-list>
      </el-form-item>
    </div>
    <div class="gva-form-box">
      <el-input
        v-model="nameReq"
        style="width: 180px"
        placeholder="教材名"
        clearable
      />
      <el-input
        v-model="collegeReq"
        style="width: 180px"
        placeholder="学院"
        clearable
      />
      <el-input
        v-model="publisherReq"
        style="width: 180px"
        placeholder="出版社"
        clearable
      />
      <el-input
        v-model="yearReq"
        style="width: 180px"
        placeholder="year"
        clearable
      />
      <el-input
        v-model="isbnReq"
        style="width: 180px"
        placeholder="isbn"
        clearable
      />
      <el-input
        v-model="idsReq"
        style="width: 180px"
        placeholder="工号使用，分割"
        clearable
      />
      <el-button  type="primary" @click="InsertBook">新增</el-button>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {Search} from '@element-plus/icons-vue'
import {getBooksList, insertBook} from "@/api/book/sysBook"
import {ElMessage} from "element-plus";

const input = ref('')
const books = ref([])
const nameReq = ref('')
const collegeReq = ref('')
const publisherReq = ref('')
const yearReq = ref('')
const isbnReq = ref('')
const idsReq = ref('')
const InsertBook = async () => {
  let data = {
    name: nameReq.value,
    college: collegeReq.value,
    publisher: publisherReq.value,
    year: yearReq.value,
    isbn: isbnReq.value,
    authors: idsReq.value.split(',').map(id => ({id}))
  }

  // 检查所有字段是否不为空
  if (!data.name || !data.college || !data.publisher || !data.year || !data.isbn || !data.authors.length) {
    ElMessage({
      message: '所有字段都是必填项，请检查输入。',
      type: 'warning',
    });
    return;
  }

  let resp = await insertBook(data)
  if (resp.code === 0) {
    console.log(resp)
    if (resp.data.flag === true) {
      ElMessage({
        message: '插入成功',
        type: 'success',
      })
    } else {
      ElMessage({
        message: '插入失败',
        type: 'error',
      })
    }
  } else {
    ElMessage({
      message: '插入失败',
      type: 'error',
    })
  }
}
const HandlerSearch = async () => {
  books.value = []
  const data = {
    name: input.value
  }
  console.log(input.value)
  try {
    const resp = await getBooksList(data)
    console.log(resp)
    if (resp.code === 0) {
      // 详细记录API响应
      console.log('API response:', resp.data)
      // 确保 resp.data.data 是一个数组
      if (resp && resp.data) {
        books.value = resp.data.map(book => ({
          name: book.name,
          college: book.college,
          publisher: book.publisher,
          year: book.year,
          isbn: book.isbn,
          authors: book.authors.map(author => author.id)
        }))
      } else {
        console.error('API response data is not an array')
      }
    }
  } catch (error) {
    console.error('Error fetching books:', error)
  }
}
</script>

<style>
.book-info {
  display: flex;
  gap: 20px; /* 控制姓名和联系方式之间的间距 */
}

.book-info p {
  margin: 0;
}
</style>
