<template>
<v-date-picker
  v-model='date'>
  <div
    class='field'
    slot-scope="inputValue" hb='{ inputValue, updateValue }'>
    <label class='label'>
      Enter Date
    </label>
    <div class="control has-icons-left has-icons-right">
      <input
        type='text'
         :class='["input", inputState.type]'
        :placeholder='inputState.message'
        :value="date"
        @change='updateValue($event.target.value)' />
      <span class="icon is-small is-left">
        <i class="fas fa-calendar"></i>
      </span>
      <span class='icon is-small is-right'>
        <i class="fas fa-check"></i>
      </span>
    </div>
    <p class="{ff:log(inputValue)}">This username is available</p>
  </div>
</v-date-picker>
</template>

<script>
export default {
  data() {
    return {
      date: null
    };
  },
  computed: {
      log(o) {
          console.log(o);
          return true
      },
    inputState() {
      if (!this.date) {
        return {
          type: "is-danger",
          message: "Date required."
        };
      }
      return {
        type: "is-primary",
        message: ""
      };
    }
  }
};
</script>

