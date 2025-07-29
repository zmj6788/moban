<template>
  <ArticleCalendarInner :height="height" :is-title="props.isTitle" :data-list="data.data_list" :theme="store.theme" v-if="show && data.data_list.length"/>
</template>

<script setup>
import ArticleCalendarInner from './article_calendar_inner.vue'
import {reactive, ref, watch} from "vue";
import {useStore} from "@/stores/store";
import {getArticleCalendarApi} from "@/api/article_api";

const props = defineProps({
  isTitle: {
    type: Boolean,
    default: true,
  }
})

const height = ref(190)

const store = useStore()
const show = ref(true)
const data = reactive({
  data_list: []
})
watch(() => store.theme, () => {
  show.value = false
  setTimeout(() => {
    show.value = true
  })
})

async function getData() {
  let res = await getArticleCalendarApi()
  data.data_list = res.data

  if (!props.isTitle){
    height.value = 150
  }

}

getData()
</script>