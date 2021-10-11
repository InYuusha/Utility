<template>
  
  
  <div >
    <!-- Message -->
    <span v-if="msg" class="bg-purple-300 absolute top-11 p-4 flex flex-row text-gray-800 rounded-3xl w-60">
      {{status}}
      <button class="mx-2" @click="closeMsg()"><UilTimesCircle size="17px" /></button>
    </span>
    <!--Kill port -->
     <span class="absolute top-14 right-14 w-40 ">
      <input class="rounded-lg w-4/6 p-2" v-model="port" type="number">
      <button @click="killConn()" class="rounded-lg w-2/6 text-gray-100 p-2 bg-blue-600">kill</button>
    </span>
     <!--Dropdown-->
    <div class="bg-blue-1000 w-full h-full p-8 text-gray-100 overflow-auto ">
      <t-dropdown :text="protocol" class="mb-8 bg-blue-600 w-16 rounded-lg shadow-2xl absolute">
        <div class="py-1 rounded-md shadow-xs">
          <a
            v-for="(p,key) in protocols" :key="key"
            @click="getConn(p)"
            href="#"
            class="block bg-gray-100 px-4 py-2 text-sm leading-5 text-gray-700 transition duration-150 ease-in-out hover:bg-gray-200 focus:outline-none focus:bg-gray-300"
            role="menuitem"
          >
            {{p}}
          </a>
        
        </div>
      </t-dropdown>
       <h3 class="text-center mb-6 text-xl font-extrabold">{{protocol}} connections</h3>
        <!--Connections-->
      <table class="table-auto w-5/6">
        <thead>
          <tr>
            <th v-for="(attr,key) in attrs" :key="key">{{attr}}</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(conn,key) in conns" :key="key">
            <td v-for="(attr,key) in attrs" :key="key" >
              <p v-if="attr!='localaddr'&& attr!='remoteaddr'">{{conn[attr]}}</p>
              <p v-else>{{conn[attr].ip}}:{{conn[attr].port}}</p>
            </td>
          </tr>
      
        </tbody>
      </table>
    </div>
  </div>
    
</template>

<script>
// @ is an alias to /src
import { UilTimesCircle } from '@iconscout/vue-unicons'

export default {
  name: 'Home',
  components:{
    UilTimesCircle
  },
  data:()=>({
    conns:[],
    attrs:['fd','family','localaddr','remoteaddr','status','pid'],
    protocol:"tcp",
    protocols:['all','tcp','udp','inet'],
    msg:false,
    status:"",
    port:3000
  }),
  methods:{
    getConn(kind="all"){
      window.backend.Connection.GetConnections(kind)
      .then(list=>{
        this.conns=list
        this.protocol=kind
        })
    },
    killConn(){
      window.backend.Connection.killPort(this.port)
      .then(res=>{
        if(res.Status==200)
          this.getConn()
        this.status=res.Msg
        this.msg=true  
      })
    },
    closeMsg(){
      this.msg=false
      this.status=""
    },
  },
  mounted(){
    this.getConn()
  }
}
</script>
<style>

</style>
