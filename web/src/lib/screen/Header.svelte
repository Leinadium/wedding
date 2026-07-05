<script>
  import Text from "../text/Text.svelte";
</script>

<header>
  <span class="names">
    <Text key="header-names" />
  </span>
</header>

<style>
  header {
    z-index: 1;
    text-align: center;
    margin-bottom: 2rem;
  }

  span {
    text-transform: uppercase;
    letter-spacing: 3px;
    font-size: 1rem;
    color: #fae7b6;
  }
</style>
