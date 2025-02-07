<script setup lang="ts">
import { ref } from 'vue'
import { createPost } from '../apis';
import { useRouter } from 'vue-router'

const router = useRouter()

const title = ref("post title")
const content = ref("new post content")
const token = ref("")

var busying = ref(false)
const submitPost = () => {
    if (busying.value) {
        return
    }
    createPost({ title: title.value, content: content.value, token: token.value },)
        .then((value) => {
            if (value) {
                router.replace("/")
            } else {
                console.log("fail to create post")
            }
        }).catch((e) => {
            console.log(e)
        }).finally(() => {
            busying.value = false
        })
}


</script>

<template>
    <div>
        <RouterLink to="/"> <input type="button" value="回到目录页"> </input></RouterLink>
    </div>

    <div> 标题: <input v-model="title" style="width: 60vw;"></input> </div>
    <div> 正文: <textarea v-model="content" style="width: 60vw;height: 75vh;">    </textarea> </div>
    <div>
        后台密码:<input v-model="token"></input>
        <input type="button" :disabled="busying" value="创建" @click="submitPost"></input>
    </div>





</template>

<style scoped></style>
