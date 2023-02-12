<template>

  <div class="d-flex justify-content-between">
    <div>
      <div class="btn-group">
        <button type="button" class="btn btn-success" @click="showModalRunItemCreate">Create New Item</button>
        <button type="button" class="btn btn-primary" @click="reload">Reload</button>
        <button type="button" class="btn btn-dark" @click="showModalFilter">Filter</button>
      </div>
    </div>
    <div>
      <div class="btn-group float-end">
        <button type="button" class="btn btn-success" @click="previousPage">Prev Page</button>
        <button type="button" class="btn btn-primary" @click="showModalPaging">{{state.pagePerTotalRecord}}</button>
        <button type="button" class="btn btn-dark" @click="nextPage">Next Page</button>
      </div>
    </div>
  </div>

  <MirzaTable :fields="fields" :items="state.items">
    <template #action="{item}">
      <div class="btn-group">
        <button type="button" class="btn btn-warning btn-sm" @click="showModalDetail(item)">Detail</button>
        <button type="button" class="btn btn-danger btn-sm" @click="submitRunItemDelete(item)">Delete</button>
        <button type="button" class="btn btn-primary btn-sm" @click="showModalRunItemUpdate(item)">Update</button>
        <button type="button" class="btn btn-info btn-sm" @click="showModalRunItemPurchase(item)">Purchase</button>
      </div>
    </template>
  </MirzaTable>

  <ViewModalDetail ref="modalDetail" @submit="reload"></ViewModalDetail>

  <ViewModalFilter ref="modalFilter" @submit="reload"></ViewModalFilter>

  <ViewModalPaging ref="modalPaging" @submit="reload"></ViewModalPaging>

  <ViewModalRunItemCreate ref="modalRunItemCreate" @submit="reload"></ViewModalRunItemCreate>

  <ViewModalRunItemUpdate ref="modalRunItemUpdate" @submit="reload"></ViewModalRunItemUpdate>

  <ViewModalRunItemPurchase ref="modalRunItemPurchase" @submit="reload"></ViewModalRunItemPurchase>

</template>

<script setup>
import MirzaTable from "../../components/table/MirzaTable.vue";
import {BASE_URL} from "../shared.js";
import {state, getNumberOfPage} from "./state.js";
import {ref} from "vue";
import to from "await-to-js";
import axios from "axios";
import swal from "sweetalert2";

import ViewModalDetail from "./ModalDetail.vue";
const modalDetail = ref()
const showModalDetail = (payload) => modalDetail.value.showModal(payload)

import ViewModalPaging from "./ModalPaging.vue";
const modalPaging = ref()
const showModalPaging = () => modalPaging.value.showModal()

import ViewModalFilter from "./ModalFilter.vue";
const modalFilter = ref()
const showModalFilter = () => modalFilter.value.showModal()

import ViewModalRunItemCreate from "./ModalInput.vue";
const modalRunItemCreate = ref()
const showModalRunItemCreate = () => modalRunItemCreate.value.showModal()

import ViewModalRunItemUpdate from "./ModalUpdate.vue";
const modalRunItemUpdate = ref()
const showModalRunItemUpdate = (item) => modalRunItemUpdate.value.showModal(item)

import ViewModalRunItemPurchase from "./ModalPurchase.vue";
const modalRunItemPurchase = ref()
const showModalRunItemPurchase = (item) => modalRunItemPurchase.value.showModal(item)

const nextPage = () => {
  if (state.filter.page + 1 <= getNumberOfPage()) {
    state.filter.page++
    reload()
  }
}

const previousPage = () => {
  if (state.filter.page - 1 > 0) {
    state.filter.page--
    reload()
  }
}

const reload = async () => {

  const url = `${BASE_URL}/items`

  const requestConfig = { params: { ...state.filter } }

  const [err, res] = await to(axios.get(url, requestConfig).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  state.items = res.data.data.items
  state.totalItems = res.data.data.count
}

const submitRunItemDelete = (item) => {

  swal.fire({
    title: `Are you sure want to delete ${item.name}?`,
    text: "You won't be able to revert this!",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#3085d6',
    cancelButtonColor: '#d33',
    confirmButtonText: 'Yes, delete it!'
  }).then(async (result) => {
    if (result.isConfirmed) {

      const url = `${BASE_URL}/items`

      const [err, res] = await to(axios.delete(`${url}/${item.id}`).catch((err) => Promise.reject(err)))

      if (err) {
        await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
        return
      }

      await reload()

      await swal.fire(
          'Deleted!',
          'Your item has been deleted.',
          'success'
      )

    }
  })


}

const fields = [
  {header: "Action", fieldName: "action",},
  {header: "Name", fieldName: "name",},
  {header: "Quantity", fieldName: "availability",},
]

reload()
</script>

<style scoped>

</style>