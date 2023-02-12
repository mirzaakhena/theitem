<template>
  <MirzaModal id="modalRunItemPurchase" ref="modalRunItemPurchase" title="RunItemPurchase" @submit="submitRunItemPurchase">
    
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

const modalRunItemPurchase = ref()

const emit = defineEmits(["submit"])

const payload = reactive({
  data: { 
  }
})

const submitRunItemPurchase = async () => {

  const url = `${BASE_URL}/runitempurchase`

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
  modalRunItemPurchase.value.showModal()
}

const hideModal = () => {
  modalRunItemPurchase.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>