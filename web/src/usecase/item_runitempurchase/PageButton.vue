<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunItemPurchase">Show</button>

  <ViewModalRunItemPurchase ref="modalRunItemPurchase" @submit="reload"></ViewModalRunItemPurchase>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunItemPurchase from "./ModalInput.vue";
const modalRunItemPurchase = ref()
const showModalRunItemPurchase = () => modalRunItemPurchase.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>