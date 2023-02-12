<template>
  <MirzaModal id="modalFilter" ref="modalFilter" title="Filter Items" @submit="submitDataFilter">

    <div class="mb-3">
      <label class="form-label">Rating (-1 for all)</label>
      <input type="number" class="form-control" placeholder="use -1 for all" v-model="state.filter.rating">
    </div>

    <div class="mb-3">
      <label class="form-label">Category</label>
      <select class="form-control" v-model="state.filter.category">
        <option value="">All</option>
        <option value="photo">Photo</option>
        <option value="sketch">Sketch</option>
        <option value="cartoon">Cartoon</option>
        <option value="animation">Animation</option>
      </select>
    </div>

    <div class="mb-3">
      <label class="form-label">Reputation Badge</label>
      <select class="form-control" v-model="state.filter.reputation_badge">
        <option value="">All</option>
        <option value="red">Red</option>
        <option value="yellow">Yellow</option>
        <option value="green">Green</option>
      </select>
    </div>

    <div class="mb-3">
      <label class="form-label">Availability More</label>
      <input type="number" class="form-control" placeholder="More than" v-model="state.filter.availability_more">
    </div>

    <div class="mb-3">
      <label class="form-label">Availability Less</label>
      <input type="number" class="form-control" placeholder="Less than" v-model="state.filter.availability_less">
    </div>

  </MirzaModal>
</template>

<script setup>
import MirzaModal from "../../components/modal/MirzaModal.vue";
import {state} from "./state.js";
import {ref} from "vue";

const modalFilter = ref()

const emit = defineEmits(["submit"])

const submitDataFilter = () => {
  state.filter.page = 1
  emit("submit")
  hideModal()
}

const showModal = () => {
  modalFilter.value.showModal()
}

const hideModal = () => {
  modalFilter.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>