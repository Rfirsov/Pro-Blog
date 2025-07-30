
<template>
    <h1 class="text-4xl font-bold mb-6">Latest Posts</h1>
    <div v-if="posts.length === 0" class="text-gray-500">No posts available.</div>
    <div v-else>
        <div v-for="post in posts" :key="post.id" class="mb-4 p-4 border rounded-lg hover:bg-gray-50 transition-colors">
            <h2 class="text-2xl font-semibold mb-2">{{ post.title }}</h2>
            <p class="text-gray-700 mb-2">{{ post.content }}</p>
            <p class="text-sm text-gray-500">By {{ post.author_id }} on {{ new Date(post.created_at).toLocaleDateString() }}</p>
            <div class="mt-2">
                <span v-for="tag in post.tags" :key="tag.id" class="inline-block bg-blue-100 text-blue-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded-full">
                    {{ tag.name }}
                </span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Status {
  id: number
  value: string
  label: string
}

interface Tag {
  id: number
  name: string
}

interface Post {
  id: string
  title: string
  content: string
  author_id: string
  slug: string
  status: Status
  tags: Tag[]
  created_at: string
  updated_at: string
}

interface PostsResponse {
  posts: Post[]
}

const { data } = await useFetch<PostsResponse>(() => `http://localhost:8080/api/v1/posts`)
const posts = computed(() => data.value?.posts || [])

</script>
