<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import PageHeader from './components/Layout/PageHeader/PageHeader.vue'
import MainContent from './components/MainContent.vue'
</script>

<template>
  <page-header />
  <main-content />
</template>

<style>
html,
body {
  margin: 0;
  padding: 0;
  border: 0;
  height: 100%;
}

#app {
  height: 100%;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
