<script setup>
import {onBeforeMount, ref} from "vue";
import AppCard from "../../layout/components/AppCard.vue";
import Client from "./client.vue";
import Tunnel from "./tunnel.vue";
import P2p from "./p2p.vue";
import About from "./about.vue";
import Logger from "./logger.vue";
import {apiConfigQuery} from "../../api/config.js";

const state = ref({
})

</script>

<template>
  <AppCard :show-border="false">
    <n-tabs type="card" animated size="small">
      <n-tab-pane name="client" tab="客户端">
        <Client></Client>
      </n-tab-pane>
      <n-tab-pane name="tunnel" tab="私有隧道">
        <Tunnel></Tunnel>
      </n-tab-pane>
      <n-tab-pane name="p2p" tab="P2P隧道">
        <P2p></P2p>
      </n-tab-pane>
      <n-tab-pane name="logger" tab="日志">
        <Logger></Logger>
      </n-tab-pane>
      <n-tab-pane name="about" tab="关于">
        <About></About>
      </n-tab-pane>
    </n-tabs>
  </AppCard>
</template>