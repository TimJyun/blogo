import axios from "axios";


const host = ""


export async function createPost(request: { title: string, content: string, token: string }): Promise<boolean> {
    try {
        let result: { ok: boolean | undefined | null } = (await axios.post(host + "/api/createPost", request)).data
        return result.ok === true
    } catch (e) {
        console.log(e)
    }
    return false
}
export async function deletePost(request: { id: number, token: string }): Promise<boolean> {
    for (let i = 0; i < 3; i++) {
        try {
            let result: { ok: boolean | undefined | null } = (await axios.post(host + "/api/deletePost", request)).data
            return result.ok === true
        } catch (e) {
            console.log(e)
        }
    }

    return false
}

export type Post = {
    id: number,
    title: string,
    content: string,
    version: number,
    createTime: string,
    lastUpdate: string,
}
export async function getPost(request: { id: number }): Promise<Post | undefined> {
    for (let i = 0; i < 5; i++) {
        try {
            let post: Post = (await axios.post(host + "/api/getPost", request)).data
            return post
        } catch (e) {
            console.log(e)
        }
    }
}

export type PostPreview = { id: number, title: string, lastUpdate: string }

export async function getPostList(): Promise<PostPreview[]> {
    for (let i = 0; i < 5; i++) {
        try {
            let postList: PostPreview[] = (await axios.post(host + "/api/getPostList")).data
            for (let i = 0; i < postList.length; i++) {
                postList[i] .lastUpdate =postList[i].lastUpdate.substring(0,16)
            }



            return postList
        } catch (e) {
            console.log(e)
        }
    }
    return []
}

export async function updatePost(
    request: {
        version: number,
        id: number,
        title: string,
        content: string,
        token: string
    },
): Promise<boolean> {
    try {
        let result: { ok: boolean | undefined | null } = (await axios.post(host + "/api/updatePost", request)).data
        return result.ok === true
    } catch (e) {
        console.log(e)
    }
    return false
}