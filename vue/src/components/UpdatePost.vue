<script setup lang="ts">
import { Ref, ref } from 'vue'
import { getPost, Post, updatePost } from '../apis';
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const route = useRoute()
const token = ref("")

var busying = ref(true)
const post: Ref<Post> = ref({
  id: -1,
  title: "unknown",
  content: "unknown",
  version: -1,
  createTime: "unknown",
  lastUpdate: "unknown",
})

getPost({ id: Number(route.params.id) })
  .then((value) => {
    if (value) {
      post.value = value
    }
  }).finally(() => {
    busying.value = false
  })


const submitPost = () => {
  if (busying.value) {
    return
  }

  updatePost(
    {
      title: post.value.title,
      content: post.value.content,
      version: Number(post.value.version),
      id: Number(route.params.id),
      token: token.value
    },
  )
    .then((value) => {
      if (value) {
        router.replace("/")
      } else {
        console.log("failed to update post")
      }
    })
    .finally(() => {
      busying.value = false
    })
}

</script>

<template>
  <div>    <RouterLink to="/"> <input type="button" value="回到目录页"> </input></RouterLink>  </div>

  <div> 标题: <input v-model="post.title" style="width: 60vw;"></input> </div>
  <div> 正文: <textarea v-model="post.content" style="width: 60vw;height: 75vh;">    </textarea> </div>
  <div>
    后台密码:<input v-model="token"></input>
    <input type="button" value="更新" :disabled="busying" @click="submitPost"></input>
  </div>



</template>

<style scoped></style>
