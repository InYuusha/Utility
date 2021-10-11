<template>
  
  
  <div >
     <!--Dropdown-->
    <div class="bg-gray-1000 w-full h-full p-8 text-gray-400 overflow-auto ">
      <t-dropdown :text="protocol" class="mb-8 bg-gray-700 w-16 rounded-lg shadow-2xl absolute">
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


export default {
  name: 'Home',

  data:()=>({
    conns:[],
    attrs:['fd','family','localaddr','remoteaddr','status','pid'],
    protocol:"tcp",
    protocols:['all','tcp','udp','inet'],
  }),
  methods:{
    getConn(kind="all"){
      window.backend.Connection.GetConnections(kind)
      .then(list=>{
        this.conns=list
        this.protocol=kind
        })
    }
  },
  mounted(){
    this.getConn()
  }
}
</script>
<style>

</style>
