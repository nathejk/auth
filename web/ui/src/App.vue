<template>
  <div class="d-flex flex-column h-100">
    <main role="main">
      <router-view></router-view>
    </main>
  </div>
</template>

<style lang="scss">
html, body { height:100%; }
body {
 background: url('/img/bg.jpg');
 background-repeat: no-repeat;
 background-size: cover;
 background-position: center;
}
@font-face {
    font-family: 'Context Reprise Condensed SSi';
    src: url('/ContextRepriseCondensedSSiBoldCondensed.woff2') format('woff2'),
        url('/ContextRepriseCondensedSSiBoldCondensed.woff') format('woff');
    font-weight: bold;
    font-style: normal;
}
</style>

<script>
  export default {
  }
</script>
