<template>
  <page-header />
  <main-content />
</template>

<script setup lang="ts">
import PageHeader from '/@/components/Layout/PageHeader/PageHeader.vue'
import MainContent from '/@/components/MainContent.vue'
</script>

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
