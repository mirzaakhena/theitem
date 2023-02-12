<template>
  <MirzaModal id="modalRunItemCreate" ref="modalRunItemCreate" title="RunItemCreate" @submit="submitRunItemCreate">

    <div class="mb-3">
      <label class="form-label">Name</label>
      <input type="text" class="form-control" placeholder="Must longer than 10 char" v-model="payload.data.name">
    </div>

    <div class="mb-3">
      <label class="form-label">Rating</label>
      <input type="number" class="form-control" placeholder="Between 0 to 5" v-model="payload.data.rating">
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
      <label class="form-label">Reputation</label>
      <input type="number" class="form-control" placeholder="Between 0 to 1000" v-model="payload.data.reputation">
    </div>

    <div class="mb-3">
      <label class="form-label">Price</label>
      <input type="number" class="form-control" placeholder="Price" v-model="payload.data.price">
    </div>

    <div class="mb-3">
      <label class="form-label">Availability</label>
      <input type="number" class="form-control" placeholder="Availability" v-model="payload.data.availability">
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

const modalRunItemCreate = ref()

const emit = defineEmits(["submit"])

const payload = reactive({
  data: {
    name: "123 123 123",
    rating: 2,
    category:  "cartoon",
    image: "http://image.com",
    reputation:  34,
    price: 5000,
    availability:  10
  }
})

const submitRunItemCreate = async () => {

  const url = `${BASE_URL}/items`

  const [err, res] = await to(axios.post(`${url}`, payload.data).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  console.log(res.data.data)
  emit("submit")
  hideModal()
}

const showModal = () => {
  modalRunItemCreate.value.showModal()
}

const hideModal = () => {
  modalRunItemCreate.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>