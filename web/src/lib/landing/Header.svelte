<script>
  import Text from "../text/Text.svelte";
</script>

<header>
  <span class="date formal">
    <Text key="header-names" />
  </span>
</header>

<style>
  header {
    z-index: 1;
    text-align: center;
    margin-bottom: 2rem;

    display: flex;
    flex-flow: column nowrap;
    justify-content: start;
    align-items: center;
  }

  span {
    letter-spacing: 3px;
    font-size: 1.5rem;
    color: #fae7b6;
  }
</style>
