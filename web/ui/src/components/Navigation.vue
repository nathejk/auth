<template>
  <header>
    <div class="navbar navbar-dark bg-dark shadow-sm">
      <div class="container d-flex flex-row-reverse">
        <a class="navbar-brand" href="/" style="background:url(/img/logomoonsmall.png) no-repeat; padding-left:60px;"><strong>{{ title }}</strong></a>
      </div>
    </div>
  </header>
</template>

<style lang="scss">
</style>

<script>
export default {
    props: {
        title: null
    }
}
</script>
