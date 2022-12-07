<template>
  <div>
    <el-container>
      <el-main>
        <h1>student home</h1>
        <info-card></info-card>
        <el-card style="width: 100%; margin-top: 10px">
          <p><i class="el-icon-s-home" style="margin-right: 18px"></i>项目名：学生选课系统</p>
          <p><i class="el-icon-s-custom" style="margin-right: 18px"></i>开发人员：LZH</p>
          <p><i class="el-icon-location-outline" style="margin-right: 18px"></i>广东工业大学</p>
        </el-card>
        <logout></logout>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Logout from "@/components/logout";
import InfoCard from "@/components/infoCard";
export default {
  name: "home",
  components: {InfoCard, Logout},
}
</script>

<style scoped>

</style>