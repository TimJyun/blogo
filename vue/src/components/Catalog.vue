<script setup lang="ts">
import { ref, Ref } from 'vue'
import { getPostList, PostPreview } from '../apis';

import { useRouter } from 'vue-router'

const router = useRouter()


const postList: Ref<PostPreview[]> = ref([])
getPostList().then((value) => {
  postList.value = value
  console.log(value)
})



const read = (id: number) => {
  router.replace("/post/" + id)
}





</script>

<template>
  <div>
    <RouterLink to="/newPost"> <input type="button" value="创建博文"> </input>    </RouterLink>
  </div>

  博文总数:{{ postList.length }}
  <div v-for="postPreview in postList" style="border: 1px;border-style: dotted;">


    <div> 标题 : {{ postPreview.title }} <input type="button" value="阅读" @click="read(postPreview.id)"></input></div>
    <div> 发布时间：{{ postPreview.lastUpdate }} </div>

  </div>
</template>

<style scoped></style>
