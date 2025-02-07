<script setup lang="ts">
import { Ref, ref } from 'vue'
import { deletePost, getPost, Post } from '../apis';
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router';



const post: Ref<Post> = ref({
  id: -1,
  title: "unknown",
  content: "unknown",
  version: -1,
  createTime: "unknown",
  lastUpdate: "unknown",
})
const token = ref("")

const route = useRoute()
const router = useRouter()


var busying = ref(true)

const id = Number(route.params.id)
console.log("id:" + id)

getPost({ id: id }).then((value) => {
  if (value) {
    post.value = value
  }
  console.log(value)
}).finally(() => {
  busying.value = false
})

const delPost = () => {
  if (busying.value) {
    return
  }
  deletePost({ id: id, token: token.value })
    .then((success) => {
      if (success) {
        router.replace("/")
      } else {
        console.log("delete post fail")
        console.log("id:" + id)
        console.log("token:" + token.value)
      }
    })
    .finally(() => { busying.value = false })
}

const editePost = () => {
  router.replace("/updatePost/" + id)
}

</script>

<template>
  <div>
    <RouterLink to="/"> <input type="button" value="回到目录页"> </input></RouterLink>
  </div>

  <div>
    <h3>{{ post.title }}</h3>
  </div>
  <div style="min-height: 75vh;"> {{ post.content }}</div>



  <div>
    <input type="button" value="编辑" @click="editePost"> </input>
    <input type="button" value="删除" :disabled="busying" @click="delPost"></input>
    后台密码:<input v-model="token"></input>
  </div>



</template>

<style scoped></style>
