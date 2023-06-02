import {createStore} from "vuex";

export default createStore({
    state() {
        return {
            classes: []
        }
    },
    mutations: {
        setClasses(state, classes) {
            state.classes = classes;
        },
    },
});