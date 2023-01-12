/* eslint-disable */
import { Store } from "vuex";
import { State } from "@/store";
declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $store: Store<State>;
  }
  declare module "*.svg" {
    const content: any;
    export default content;
  }
}
declare module "*.vue" {
  import type { DefineComponent } from "vue";
  const component: DefineComponent<{}, {}, any>;
  export default component;
}

declare module "tsparticles";
