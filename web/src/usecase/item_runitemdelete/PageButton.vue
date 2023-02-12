<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunItemDelete">Show</button>

  <ViewModalRunItemDelete ref="modalRunItemDelete" @submit="reload"></ViewModalRunItemDelete>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunItemDelete from "./ModalInput.vue";
const modalRunItemDelete = ref()
const showModalRunItemDelete = () => modalRunItemDelete.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>