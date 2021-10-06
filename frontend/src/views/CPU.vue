<template>
  
    <!--Percentage-->
    <div class="bg-blue-900  shadow-md  mx-auto p-7 h-5/6">
       <vc-donut  has-legend :auto-adjust-text-size="true" background="white" :thickness="10" :total="100" foreground="#4680fa" :sections="sections">
          <p class="tracking-wider text-gray-600 font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
       </vc-donut>
      

  </div>
</template>

<script>
// @ is an alias to /src
import * as Wails from '@wailsapp/runtime';

export default {
  name: 'Home',

  data:()=>({
    sections:[0],
    perc:""
  }),
  methods:{

  },
  mounted() {
    Wails.Events.On("cpu",(per)=>{
      this.sections=[{value:per,color:'orange',label:'CPU Usage'}]
      this.perc=per
    })
  },
}
</script>
<style>

</style>
