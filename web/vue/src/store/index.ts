import { createStore } from "vuex";

export interface State {
  upIndex: string;
  activeIndex: string;
}

const store = createStore<State>({
  state: {
    upIndex: "",
    activeIndex: "",
  },
  getters: {
    getActInd: (state) => (isUp: boolean) => {
      return isUp ? state.upIndex + "-" + state.activeIndex : state.activeIndex;
    },
  },
  mutations: {
    changeActInd(state: State, index: string) {
      state.activeIndex = index;
    },
    changeUpInd(state: State, upIndex: string) {
      state.upIndex = upIndex;
    },
  },
});

export default store;
