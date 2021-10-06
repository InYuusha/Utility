<template>
  
    <!--Percentage-->
  <div class="bg-blue-900  shadow-md  mx-auto p-7 h-4/6">
    <div class="bg-gray-800 w-3/6 p-10">
      <span v-for="(item,key) in info" :key="key" class="grid grid-cols-2 gap-10 text-lg text-gray-400">
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
