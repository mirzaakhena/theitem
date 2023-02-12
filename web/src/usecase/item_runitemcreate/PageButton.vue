<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunItemCreate">Show</button>

  <ViewModalRunItemCreate ref="modalRunItemCreate" @submit="reload"></ViewModalRunItemCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunItemCreate from "./ModalInput.vue";
const modalRunItemCreate = ref()
const showModalRunItemCreate = () => modalRunItemCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>