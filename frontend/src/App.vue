<script setup>
import { RouterView, useRoute } from 'vue-router'
import BottomNav from './components/BottomNav.vue'

const route = useRoute()
</script>

<template>
  <div class="bg-[#F2F4F8] flex items-center justify-center min-h-screen">
    <div class="w-full max-w-[375px] h-[812px] bg-gradient-to-br from-[#eff6ff] via-[#fff1f2] to-[#f0f9ff] shadow-2xl overflow-hidden relative sm:rounded-[40px] flex flex-col border-[8px] border-white">
      
      <!-- Background Blobs -->
      <div class="blob w-64 h-64 bg-blue-200 rounded-full -top-10 -left-10 mix-blend-multiply"></div>
      <div class="blob w-64 h-64 bg-pink-200 rounded-full top-20 -right-20 mix-blend-multiply"></div>

      <!-- Main Content Area -->
      <div class="flex-1 overflow-y-auto hide-scrollbar pb-24 relative z-10">
        <RouterView />
      </div>

      <!-- Bottom Navigation -->
      <BottomNav v-if="route.name !== 'add'" />
    </div>
  </div>
</template>
