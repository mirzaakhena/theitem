import {computed, reactive} from "vue";

export const state = reactive({
    items: [],
    filter: {
        page: 1,
        size: 4,
        rating: -1,
        category: '',
        reputation_badge: '',
        availability_more: 0,
        availability_less: 100,
    },
    totalItems: 0,
    pagePerTotalRecord: computed(()=> `${state.filter.page} / ${getNumberOfPage()}` )
})

export const getNumberOfPage = () => Math.ceil(state.totalItems / state.filter.size)