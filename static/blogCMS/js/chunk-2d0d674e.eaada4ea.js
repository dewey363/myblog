(window.webpackJsonp=window.webpackJsonp||[]).push([["chunk-2d0d674e"],{7340:function(e,t,a){"use strict";a.r(t);var i=a("9bd2"),n={data:function(){return{loading:!0,columns:[{title:"序号",key:"id",width:"50"},{title:"标题",key:"title"},{title:"分类",key:"cate.name"},{title:"标签",key:"labels",component:{name:"my-tag"}},{title:"作者",key:"author"},{title:"摘要",key:"summary"},{title:"阅读量",key:"views",width:"80"}],data:[{id:"1",title:"topic demo",summary:"上海市普陀区金沙江路 1518 弄",cate:{name:"test"},labels:[{id:1,name:"test01"},{id:2,name:"test02"}],author:"sinksmell",views:2}],dataBackUp:[],cates:[],cateForm:{cateId:""},pagination:{currentPage:1,pageSize:10},formOptions:{labelWidth:"80px",labelPosition:"left",saveLoading:!1},rowHandle:{remove:{icon:"el-icon-delete",text:"删除",size:"medium",fixed:"right",confirm:!0},options:{}}}},computed:{currentData:function(){var e=this.pagination.currentPage,t=this.pagination.pageSize;return this.data.slice((e-1)*t,e*t)}},methods:{getTopicsByCate:function(){var t=this;Object(i.a)({method:"get",url:"/topic/cate/"+this.cateForm.cateId}).then(function(e){t.data=e.topics})},resetForm:function(e){this.$refs[e].resetFields(),this.data=this.dataBackUp},handleSizeChange:function(e){this.pagination.pageSize=e},handleCurrentChange:function(e){this.pagination.currentPage=e},handleRowRemove:function(e,t){var a=this,n=(e.index,e.row);Object(i.a)({method:"post",url:"/topic/delete",data:{id:n.id}}).then(function(e){"OK"===e.msg?(a.$message({message:"删除成功",tags:"success"}),t()):a.$message({message:"删除失败",tags:"error"})})},handleDialogOpen:function(e){var t=e.mode;this.$message({message:"打开模态框，模式为："+t,tags:"success"})},handleDialogCancel:function(e){this.$message({message:"取消保存",tags:"warning"}),e()}},mounted:function(){var t=this;i.a.get("/topic/list").then(function(e){t.data=e,t.dataBackUp=e,t.loading=!1}),i.a.get("/category/list").then(function(e){t.cates=e})}},o=a("2877"),s=Object(o.a)(n,function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("d2-container",[a("template",{slot:"header"},[a("el-form",{ref:"cateForm",staticClass:"demo-form-inline",attrs:{inline:!0,model:t.cateForm}},[a("el-form-item",{attrs:{prop:"cateId",label:"文章分类"}},[a("el-select",{attrs:{placeholder:"请选择"},model:{value:t.cateForm.cateId,callback:function(e){t.$set(t.cateForm,"cateId",e)},expression:"cateForm.cateId"}},t._l(t.cates,function(e){return a("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})}),1)],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.getTopicsByCate}},[t._v("查询")]),a("el-button",{attrs:{type:"success"},on:{click:function(e){t.resetForm("cateForm")}}},[t._v("重置")])],1)],1)],1),a("d2-crud",{ref:"d2Crud",attrs:{columns:t.columns,loading:t.loading,data:t.currentData,"form-options":t.formOptions,rowHandle:t.rowHandle},on:{"row-remove":t.handleRowRemove,"dialog-open":t.handleDialogOpen,"dialog-cancel":t.handleDialogCancel}}),a("template",{slot:"footer"},[a("el-pagination",{attrs:{background:"",layout:"total, sizes, prev, pager, next, jumper","current-page":t.pagination.currentPage,"page-sizes":[5,10,20,40],"page-size":t.pagination.pageSize,total:t.data.length},on:{"size-change":t.handleSizeChange,"current-change":t.handleCurrentChange}})],1)],2)},[],!1,null,null,null);s.options.__file="index.vue";t.default=s.exports}}]);