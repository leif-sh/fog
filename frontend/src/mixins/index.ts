// https://github.com/vuejs/vue-class-component#using-mixins
import { timestampToTime } from "../utils/utils";

const mixin = {
  methods: {
    formatTime(value: number): string {
      return timestampToTime(value, true);
    }
  }
};
export default mixin;
