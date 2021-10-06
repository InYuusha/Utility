<template>
  <div id="app" class="h-screen">
   
    <router-view />

     <NavBar class="h-1/6" />
  </div>
</template>
<script>

import NavBar from './components/NavigationDrawer.vue'
export default {
  name: 'App',
  components: {
    NavBar
  },
}
</script>
<style>

</style>
