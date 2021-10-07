<template>
  
    <!--Percentage-->
    <div>
     <vc-donut style="color:#8f9aa8" background="#0c123d" has-legend :auto-adjust-text-size="true" :thickness="6" :total="100" foreground="#696c80" :sections="sections">
          <p style="color:#0768fa" class="tracking-wider font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
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