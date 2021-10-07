<template>
  
    <!--Percentage-->
    <div class="bg-blue-900  shadow-md  mx-auto p-7 h-5/6 overflow-auto">
       <vc-donut  has-legend :auto-adjust-text-size="true" background="white" :thickness="10" :total="100" foreground="#4680fa" :sections="sections">
          <p class="tracking-wider text-gray-600 font-extrabold">{{perc}}<span class="text-sm align-top">%</span></p>
       </vc-donut>
      
      <div class="bg-gray-800  full p-8 text-gray-400 mt-4 overflow-auto">
        <h3 class="text-center my-4 text-xl font-extrabold">Host Info</h3>
        <table class="table-auto">
          <thead>
            <tr>
              <th>pid</th>
              <th>cpu usage</th>
              <th>exe</th>
              <th>background</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(p,key) in processList" :key="key">
              <td>{{p.Pid.pid}}</td>
              <td>{{p.Cpu}}%</td>
              <td>{{p.Exe}}</td>
              <td>{{p.Background}}</td>
              <td>
                <button class="bg-gray-500" @click="killP(p.Pid.pid)">Stop x</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div> 

  </div>
</template>

<script>
// @ is an alias to /src
import * as Wails from '@wailsapp/runtime';

export default {
  name: 'Home',

  data:()=>({
    sections:[0],
    perc:"",
    processList:[]
  }),
  methods:{
    killP(p){
      window.backend.Info.KillP(p).then(()=>{
        this.getProcesses()
      })
    },
    getProcesses(){
      window.backend.Info.GetProcesses()
      .then(list=>this.processList = list)
    }
  },
  mounted() {

    this.getProcesses()

    Wails.Events.On("cpu",(per)=>{
      this.sections=[{value:per,color:'orange',label:'CPU Usage'}]
      this.perc=per
    })
  },
}
</script>
<style>

</style>
