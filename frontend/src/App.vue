<template>
  <div id="app" class="h-screen">
   
    <router-view class="bg-blue-1000  shadow-md   p-7 h-5/6 " />

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
