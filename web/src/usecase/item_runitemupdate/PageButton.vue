<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunItemUpdate">Show</button>

  <ViewModalRunItemUpdate ref="modalRunItemUpdate" @submit="reload"></ViewModalRunItemUpdate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunItemUpdate from "./ModalInput.vue";
const modalRunItemUpdate = ref()
const showModalRunItemUpdate = () => modalRunItemUpdate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>