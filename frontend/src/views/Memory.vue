<template>
  
    <!--Percentage-->
    <div class="flex flex-col">
      <vc-donut style="color:#8f9aa8" background="#0c123d" has-legend :auto-adjust-text-size="true" :thickness="6" :total="100" foreground="#696c80" :sections="sections">
        <p style="color:#0768fa" class="tracking-wider font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
      </vc-donut>

      <div class="grid grid-cols-3 overflow-auto">
        <vc-donut size="150" v-for="(disk,key) in disks" :key="key" style="color:#8f9aa8" background="#0c123d" has-legend :auto-adjust-text-size="true" :thickness="6" :total="100" foreground="#696c80" :sections="[{value:disk.usedPercent,color:'orange',label:disk.path}]">
          <p style="color:#0768fa" class="tracking-wider font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
        </vc-donut>
      </div>
      
  </div>
</template>

<script>

import * as Wails from '@wailsapp/runtime';

export default{
    data:()=>({
        sections:[0],
        perc:"",
        disks:[]
    }),
    methods:{
        getDisks(){
            window.backend.Memory.GetDisks()
            .then(list=>this.disks=list)
        }
    },
    mounted() {
       
        this.getDisks()
        Wails.Events.On('memory',(perc)=>{
            console.log(perc)
            this.sections=[{value:perc,color:'orange',label:'Memory Usage'}]
            this.perc=perc
        })
    },
}
</script>