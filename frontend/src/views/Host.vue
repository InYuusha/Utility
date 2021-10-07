<template>
  
    <!--Percentage-->
  <div >
    <div class="bg-gray-1000 overflow-auto full p-8 text-gray-400">
      <h3 class="text-center my-4 text-xl font-extrabold">Host Info</h3>
      <span v-for="(item,key) in info" :key="key" class="grid grid-cols-2 gap-10 text-lg ">
        <p>{{item.K}}</p>
        <p>{{item.V}}</p>
      </span>
    </div> 
  </div>
      
</template>

<script>
// @ is an alias to /src


export default {
  name: 'Home',

  data:()=>({
    info:[],
  }),
  methods:{

  },
  mounted(){
    window.backend.GetHost()
    .then(host=>{
      Object.keys(host).map(item=>{
        let obj={
          K:item,
          V:host[item]
        }
        this.info.push(obj)
      })
    })
  }
}
</script>
<style>

</style>
