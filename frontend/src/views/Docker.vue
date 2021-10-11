<template>
  
    <!--Percentage-->
  <div>
    <div class="bg-gray-1000 overflow-auto w-full h-full p-8 text-gray-200">
      
      <button class="bg-blue-900 absolute p-3  rounded-lg top-10 right-14" @click="allContainers()"><UilSync :class="loading?'animate-spin':null" size="30px" /></button>

      <h3 class="text-center mb-4 text-xl font-extrabold">Container Info</h3>
      <table class="table-auto w-full">
        <thead>
          <tr class="w-full px-4 bg-gray-700" >
            <th class="text-center">containerID</th>
            <th>name</th>
            <th>image</th>
            <th>status</th>
            <th>running</th>
          </tr>
        </thead>
        <tbody>
            <tr class="leading-10 text-lg" v-for="(cont,key) in containers" :key="key">
              <td>{{cont.containerID}}</td>
              <td>{{cont.name}}</td>
              <td>{{cont.image}}</td>
              <td>{{cont.status}}</td>
              <td>{{cont.running}}</td>
            </tr>
        </tbody>
        
      </table>
    </div> 
  </div>
      

 
</template>

<script>
// @ is an alias to /src
import { UilSync } from '@iconscout/vue-unicons'

export default {
  name: 'Home',
  components:{
    UilSync
  },
  data:()=>({
    containers:[],
    activeID:[],
    loading:false
  }),
  methods:{
 
    allContainers(){
      this.loading=true
      window.backend.Docker.GetAllDocker()
      .then(list=>{
      this.containers = list
      this.loading=false
    })
    }
  
  },

  mounted() {
    this.allContainers()
  },
}
</script>

<style>
table td{
  @apply mx-5 text-center;
};
table tr{
  @apply px-4;
}

</style>
