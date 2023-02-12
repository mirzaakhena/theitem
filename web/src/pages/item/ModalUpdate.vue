<template>
  <MirzaModal id="modalRunItemUpdate" ref="modalRunItemUpdate" title="RunItemUpdate" @submit="submitRunItemUpdate">

    <div class="mb-3">
      <label class="form-label">Name</label>
      <input type="text" class="form-control" placeholder="Must longer than 10 char" v-model="payload.data.name">
    </div>

    <div class="mb-3">
      <label class="form-label">Category</label>
      <select class="form-control" v-model="payload.data.category">
        <option value="photo">Photo</option>
        <option value="sketch">Sketch</option>
        <option value="cartoon">Cartoon</option>
        <option value="animation">Animation</option>
      </select>

    </div>

    <div class="mb-3">
      <label class="form-label">Image</label>
      <input type="text" class="form-control" placeholder="Must be valid url" v-model="payload.data.image">
    </div>

    <div class="mb-3">
      <label class="form-label">Price</label>
      <input type="number" class="form-control" placeholder="Price" v-model="payload.data.price">
    </div>

  </MirzaModal>
</template>

<script setup>
import MirzaModal from "../../components/modal/MirzaModal.vue";
import {BASE_URL} from "../shared.js";
import {ref} from "vue";
import to from "await-to-js";
import axios from "axios";
import swal from "sweetalert2";
import {reactive} from "vue";

const modalRunItemUpdate = ref()

const emit = defineEmits(["submit"])

const payload = reactive({
  itemID: '',
  data: {
    name: "",
    category:"",
    image: "",
    price: 0
  }
})

const submitRunItemUpdate = async () => {

  const url = `${BASE_URL}/items/${payload.itemID}`

  console.log(payload.data)

  const [err, res] = await to(axios.put(`${url}`, payload.data).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  console.log(res.data.data)
  emit("submit")
  hideModal()
}

const showModal = async (item) => {

  payload.itemID = item.id

  const url = `${BASE_URL}/items/${item.id}`

  const requestConfig = { params: {} }

  const [err, res] = await to(axios.get(url, requestConfig).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  payload.data = res.data.data.item
  console.log(res.data.data.item)

  modalRunItemUpdate.value.showModal()
}

const hideModal = () => {
  modalRunItemUpdate.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>