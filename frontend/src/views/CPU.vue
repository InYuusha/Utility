<template>
  
    
    <div >

       <div class="flex flex-row h-2/6 mb-6">
         <!--Cpu Details-->
         <div class="w-3/6 bg-gray-1000 p-8 text-gray-400  overflow-auto">
            <span v-for="(p,key) in cpuDetails" :key="key">
              <h3 class="font-extrabold text-center">Cpu {{p.cpu}}</h3>
              <span class="flex flex-row justify-between" v-for="(field,key) in fields" :key="key">
                <p>{{field}}</p>
                <p>{{p[field]}}</p>
              </span>
        
            </span>
         </div>
          <!--Donut Chart-->
        <div class="w-3/6">
          <vc-donut size="170" style="color:#8f9aa8" background="#0c123d" has-legend  auto-adjust-text-size :thickness="6" :total="100" foreground="#696c80" :sections="sections">
            <p style="color:#0768fa" class="tracking-wider font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
          </vc-donut>
        </div>

       </div>
       <!--Processes-->
       <div class="overflow-auto h-4/6 backdrop-blur-lg">
          <div class="bg-gray-1000  full p-8 text-gray-400 mt-4 overflow-auto">
         <button  class="bg-blue-900 p-3 rounded-lg float-right " @click="getProcesses()"><UilSync :class="loading?'animate-spin':null" size="30px" /></button>
        <h3 class="text-center my-4 text-xl font-extrabold">Running Processes</h3>
        <table class="table-auto w-full">
          <thead>
            <tr>
              <th>pid</th>
              <th>cpu usage</th>
              <th>exe</th>
              <th>background</th>
              <th>status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(p,key) in processList" :key="key">
              <td>{{p.Pid.pid}}</td>
              <td>{{p.Cpu}}%</td>
              <td>{{p.Exe}}</td>
              <td>{{p.Background}}</td>
              <td>
                <button class="bg-gray-700 px-2 mr-4 py-1 rounded-lg" @click="killP(p.Pid.pid)"><UilTimesCircle size="20px"/></button>
                <button class="bg-gray-700 px-3 mr-4 py-1 rounded-lg" @click="stopP(p.Pid.pid)"><UilStopCircle size="20px"/></button>
                <button class="bg-gray-700 px-3 py-1 rounded-lg" @click="contP(p.Pid.pid)"><UilPlay size="20px" /></button>
              </td>
            </tr>
          </tbody>
        </table>
        </div> 
       </div>
     
  </div>
</template>

<script>
// @ is an alias to /src
import * as Wails from '@wailsapp/runtime';
import { UilSync } from '@iconscout/vue-unicons'
import { UilStopCircle } from '@iconscout/vue-unicons'
import { UilPlay } from '@iconscout/vue-unicons'
import { UilTimesCircle } from '@iconscout/vue-unicons'

export default {

  name: 'Home',
  components:{
    UilSync,
    UilStopCircle,
    UilPlay,
    UilTimesCircle
  },
  data:()=>({
    sections:[0],
    perc:"",
    loading:false,
    processList:[],
    cpuDetails:[],
    fields:['vendorId','family','model','stepping','cores','coreId','physicalId','modelName','mhz','cachSize'],
  }),
  methods:{
    killP(p){
      window.backend.Info.KillP(p).then(()=>{
        this.getProcesses()
      })
    },
     stopP(p){
      window.backend.Info.stopP(p).then(()=>{
        this.getProcesses()
      })
    },
     resumeP(p){
      window.backend.Info.contP(p).then(()=>{
        this.getProcesses()
      })
    },
    getProcesses(){
      this.loading=true
      this.list=[]
      window.backend.Info.GetProcesses()
      .then(list=>{
        this.processList = list
        this.loading=false
        })
    },
    getCpuDetails(){
      window.backend.Info.GetCpuDetails()
      .then(list=>this.cpuDetails=list)
    }
  },
  mounted() {

    this.getCpuDetails()
    this.getProcesses()

    Wails.Events.On("cpu",(per)=>{
      this.sections=[{value:per,color:'#0768fa',label:'CPU Usage'}]
      this.perc=per
    })
  },
}
</script>
<style>

</style>
