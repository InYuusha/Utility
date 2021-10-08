<template>
    
   <div class=" bg-blue-900 p-4 flex justify-between px-10" >
      <router-link to="/" :class="$route.path=='/'?'active':null">
        <UilProcessor :color="icon.color" :size="icon.size" class="logo "></UilProcessor>
      </router-link>

      <router-link to="/host" :class="$route.name=='Host'?'active':null">
        <UilDesktop :color="icon.color" :size="icon.size" class="logo"></UilDesktop>
      </router-link>

      <router-link to="/memory" :class="$route.name=='Memory'?'active':null">
        <UilHdd :color="icon.color" :size="icon.size" class="logo"></UilHdd>
      </router-link>

       <router-link to="/docker" :class="$route.name=='Docker'?'active':null">
        <UilDocker :color="icon.color" :size="icon.size" class="logo"></UilDocker>
      </router-link>

       <router-link to="/connection" :class="$route.name=='Connection'?'active':null">
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
.active{
  border-bottom: 10px solid #0c123d; 
}
router-link{
  @apply justify-center;
}

</style>