<template>
  
    <!--Percentage-->
    <div class="flex flex-col text-gray-400">
      <h3 class="text-center my-4 text-xl font-extrabold">Virtual Memory</h3>
      <vc-donut style="color:#8f9aa8" background="#0c123d" has-legend :auto-adjust-text-size="true" :thickness="6" :total="100" foreground="#696c80" :sections="sections">
          <span class="flex flex-col">
            <p style="color:#0768fa" class="tracking-wider font-extrabold">{{virtualMem.UsedPercent}}<span class="text-sm align-top">%</span></p>
            <p style="color:#0768fa" class="tracking-wider font-extrabold">of {{byteToGb(virtualMem.Total)}} GB</p>
          </span>
      </vc-donut>

      <h3 class="text-center my-4 text-xl font-extrabold">Drives</h3>

      <div class="grid grid-cols-3 overflow-auto">
        <vc-donut size="150" v-for="(disk,key) in disks" :key="key" style="color:#8f9aa8" background="#0c123d" has-legend :auto-adjust-text-size="true" :thickness="6" :total="100" foreground="#696c80" :sections="[{value:disk.UsedPercent,color:'#0768fa',label:disk.Path}]">
         <span class="flex flex-col">
            <p style="color:#0768fa" class="tracking-wider font-extrabold">{{disk.UsedPercent}}<span class="text-sm align-top">%</span></p>
            <p style="color:#0768fa" class="tracking-wider font-extrabold"> of{{byteToGb(disk.Total)}} GB</p>
         </span>
        </vc-donut>
      </div>
      
  </div>
</template>

<script>

import * as Wails from '@wailsapp/runtime';

export default{
    data:()=>({
        sections:[0],
        virtualMem:{},
        disks:[]
    }),
    methods:{
        getDisks(){
            window.backend.Memory.GetDisks()
            .then(list=>this.disks=list)
        },
        byteToGb(b){
            let gb =b/(1024*1024*1024)
            gb=Math.round((gb + Number.EPSILON) * 100) / 100;
            return (gb)
        }
    },
    mounted() {
       
        this.getDisks()
        Wails.Events.On('memory',(mem)=>{
            
            this.sections=[{value:mem.UsedPercent,color:'#0768fa',label:'Memory Usage'}]
            this.virtualMem=mem
        })
    },
    computed:{
       
    }
}
</script>

<style >
h3{
    color:
}
    
</style>