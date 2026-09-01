<script setup>
import { onMounted, ref } from 'vue'
import api from '@/api/api'

const posts = ref([])

const fetchData = async () => {
  try {
    const response = await api.getPosts()
    posts.value = response.data
  } catch (err) {
    err.value = "Ошибка при загрузке данных"
    alert(err)
  }
}

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div class="wall">
    <ul>
      <li v-for="post in posts" :key="post.id">
        {{ post }}
      </li>
    </ul>
  </div>
</template>

<style scoped></style>
