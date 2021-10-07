<template>
    
   <div class=" bg-blue-900 p-4  grid grid-cols-5" >
      <router-link to="/">
        <UilProcessor :color="icon.color" :size="icon.size" class="logo"></UilProcessor>
      </router-link>

      <router-link to="/host">
        <UilDesktop :color="icon.color" :size="icon.size" class="logo"></UilDesktop>
      </router-link>

      <router-link to="/memory">
        <UilHdd :color="icon.color" :size="icon.size" class="logo"></UilHdd>
      </router-link>

       <router-link to="/docker">
        <UilDocker :color="icon.color" :size="icon.size" class="logo"></UilDocker>
      </router-link>

       <router-link to="/connection">
        <UilRssInterface :color="icon.color" :size="icon.size" class="logo"></UilRssInterface>
      </router-link>
    </div>
</template>

<script>
import { UilProcessor } from '@iconscout/vue-unicons'
import { UilHdd } from '@iconscout/vue-unicons'
import { UilDocker } from '@iconscout/vue-unicons'
import { UilRssInterface } from '@iconscout/vue-unicons'
import { UilDesktop } from '@iconscout/vue-unicons'

export default{
    components: {
    UilProcessor,
    UilHdd,
    UilDocker,
    UilRssInterface,
    UilDesktop

  },
    data:()=>({
        icon:{
          size:'60px',
          color:'#4680fa'
        }
    }),
    methods:{
       
    },
    computed:{
    
    }
}
</script>

<style>



</style>