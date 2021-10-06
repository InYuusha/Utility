<template>
  
    <!--Percentage-->
    <div class="bg-blue-900  shadow-md  mx-auto p-7 h-4/6">
       <vc-donut has-legend :auto-adjust-text-size="true" background="white" :thickness="10" :total="100" foreground="#4680fa" :sections="sections">
          <h1 class="text-6xl tracking-wider text-gray-600 font-extrabold">20<span class="text-sm align-top">%</span></h1>
       </vc-donut>
      

  </div>
</template>

<script>
export default{
    data:()=>({
        sections:[{value:40,color:'orange',label:'Memory Usage'}],
    }),
    methods:{

    }
}
</script>