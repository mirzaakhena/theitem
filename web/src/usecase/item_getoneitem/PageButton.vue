<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalGetOneItem">Show</button>

  <ViewModalGetOneItem ref="modalGetOneItem" @submit="reload"></ViewModalGetOneItem>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalGetOneItem from "./ModalDetail.vue";
const modalGetOneItem = ref()
const showModalGetOneItem = () => modalGetOneItem.value.showModal({name: 'Mirza', address: 'Bandung'})

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>