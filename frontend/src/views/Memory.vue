<template>
  
    <!--Percentage-->
    <div class="bg-blue-900  shadow-md  mx-auto p-7 h-5/6">
       <vc-donut has-legend :auto-adjust-text-size="true" background="white" :thickness="10" :total="100" foreground="#4680fa" :sections="sections">
          <p class="tracking-wider text-gray-600 font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
       </vc-donut>
      

  </div>
</template>

<script>

import * as Wails from '@wailsapp/runtime';

export default{
    data:()=>({
        sections:[0],
        perc:""
    }),
    methods:{

    },
    mounted() {
        Wails.Events.On('memory',(perc)=>{
            console.log(perc)
            this.sections=[{value:perc,color:'orange',label:'Memory Usage'}]
            this.perc=perc
        })
    },
}
</script>